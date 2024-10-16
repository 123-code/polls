package util

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

type WalletFactory struct{
    abi abi.ABI
    address common.Address
    backend bind.ContractBackend
}




func loadABI(filepath string) (abi.ABI, error) {
    abiFile, err := ioutil.ReadFile(filepath)
    if err != nil {
        fmt.Println("error reading ABI file:", err)
        return abi.ABI{}, fmt.Errorf("failed to read ABI file: %v", err)
    }

    // Unmarshal directly into a []map[string]interface{} to match the expected ABI structure
    var jsonABI []interface{}
    err = json.Unmarshal(abiFile, &jsonABI)
    if err != nil {
        fmt.Println("error parsing ABI:", err)
        return abi.ABI{}, fmt.Errorf("failed to parse ABI JSON: %v", err)
    }

    // Re-marshal the JSON ABI array back to a string for parsing with abi.JSON
    abiJSON, err := json.Marshal(jsonABI)
    if err != nil {
        fmt.Println("error re-marshaling ABI:", err)
        return abi.ABI{}, fmt.Errorf("failed to marshal ABI: %v", err)
    }

    // Parse the ABI using the go-ethereum abi.JSON method
    parsedABI, err := abi.JSON(strings.NewReader(string(abiJSON)))
    if err != nil {
        fmt.Println("error parsing ABI:", err)
        return abi.ABI{}, fmt.Errorf("failed to parse ABI: %v", err)
    }

    return parsedABI, nil
}



 func NewWalletFactory(address common.Address, backend bind.ContractBackend) (*WalletFactory, error) {
    parsedABI, err := loadABI("/Users/joseignacionaranjo/Polls_backend/util/walletfactory.json")
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    return &WalletFactory{abi: parsedABI, address: address, backend: backend}, nil
}
func (w *WalletFactory) CreateWallet(opts *bind.TransactOpts, userId string) (*types.Transaction, error) {
    if opts == nil {
        fmt.Println("transaction options are nil")
        return nil, errors.New("transaction options are nil")
    }
    
    if opts.Nonce == nil {
        nonce, err := w.backend.PendingNonceAt(context.Background(), opts.From)
        if err != nil {
            fmt.Println(err)
            return nil, fmt.Errorf("failed to retrieve nonce: %v", err)
        }
        opts.Nonce = big.NewInt(int64(nonce))
    }

    data, err := w.abi.Pack("createWallet", userId)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    
    tx := types.NewTransaction(opts.Nonce.Uint64(), w.address, opts.Value, opts.GasLimit, opts.GasPrice, data)
    signedTx, err := opts.Signer(opts.From, tx)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    
    err = w.backend.SendTransaction(context.Background(), signedTx)
    if err != nil {
        fmt.Println(err)
        return nil, err
    }
    
    
    return signedTx, nil
}

func (w *WalletFactory) ParseWalletCreated(log types.Log) (struct {
    UserId string
    WalletAddress common.Address
}, error) {
    event := struct {
        UserId string
        WalletAddress common.Address
    }{}
    err := w.abi.UnpackIntoInterface(&event, "WalletCreated", log.Data)
    if err != nil {
        fmt.Println(err)
        return event, err
    }
    return event, nil
}

//auth

func CreateWallet(userID string) (string, error) {
    // Create a custom HTTP client
    httpClient := &http.Client{}
    fmt.Println("llego")
    // Create a custom HTTP RoundTripper
    customTransport := &customTransport{
        underlying: http.DefaultTransport,
        apiKey:     "t-66f6f6ef6be651758a55d255-01527e40b9bd45c1845d4a4b",
    }
    httpClient.Transport = customTransport

    // Create an rpc.Client with the custom HTTP client
    rpcClient, err := rpc.DialHTTPWithClient("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca", httpClient)
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
    }
//0x141920D050A904B6e3feDd732200fbF9B1f0b65A
//6476049119d25379b9cd234d390c757db268cac2729d54c19c6d1b884a9da5e9
    client := ethclient.NewClient(rpcClient)


    privateKey, err := crypto.HexToECDSA("526938daf3a62f82fc13d7abe8d063104160bfd869ddbc25e3feb6a2f8a8042e")
    if err != nil {
        fmt.Println("error loading private key",err)
        return "", fmt.Errorf("failed to load private key: %v", err)
    }

    // Create an authorized transactor
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // 11155111 is the chain ID for Sepolia
    if err != nil {
        fmt.Println("error transactor",err)
        return "", fmt.Errorf("failed to create authorized transactor: %v", err)
    }

    // Set gas price and limit
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        fmt.Println("error suggesting gas price",err)
        return "", fmt.Errorf("failed to suggest gas price: %v", err)
    }
    auth.GasPrice = gasPrice
    auth.GasLimit = uint64(3000000) // Adjust as needed

    // Load the WalletFactory contract
    factoryAddress := common.HexToAddress("0x84Eb5C50Fcd8d6F2eeBDb929381af5AC4e80321c")
    factory, err := NewWalletFactory(factoryAddress, client)
    if err != nil {
        fmt.Println("error with wallet contract",err)
        return "", fmt.Errorf("failed to load WalletFactory contract: %v", err)
    }

    // Call the createWallet function
    tx, err := factory.CreateWallet(auth, userID)
    if err != nil {
        fmt.Println("error creating wallet",err)
        return "", fmt.Errorf("failed to create wallet: %v", err)
    }

    // Wait for the transaction to be mined
    receipt, err := bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        fmt.Println("error mining transaction",err)
        return "", fmt.Errorf("failed to wait for transaction to be mined: %v", err)
    }


    // Get the wallet address from the event logs
    event, err := factory.ParseWalletCreated(*receipt.Logs[0])
    if err != nil {
        fmt.Println("error parsing",err)
        return "", fmt.Errorf("failed to parse WalletCreated event: %v", err)
    }

    if receipt.Status == 0 {
        fmt.Println("transaction failed")
        return "", fmt.Errorf("transaction failed")
    }
    fmt.Printf("Transaction receipt: %+v\n", receipt)

    return event.WalletAddress.Hex(), nil
}

// customTransport is a custom http.RoundTripper that adds the API key to each request
type customTransport struct {
    underlying http.RoundTripper
    apiKey     string
}

func (t *customTransport) RoundTrip(req *http.Request) (*http.Response, error) {
    req.Header.Add("x-api-key", t.apiKey)
    return t.underlying.RoundTrip(req)
}


