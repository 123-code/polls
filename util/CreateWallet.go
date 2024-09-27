package util

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"

	"github.com/ethereum/go-ethereum/common"

)

type WalletFactory struct{
    abi abi.ABI
    address common.Address
    backend bind.ContractBackend
}




 func loadABI(filepath string)(abi.ABI,error){
abiFile,err := ioutil.ReadFile(filepath)
if err != nil{
    fmt.Println("error",err);
    return abi.ABI{}, fmt.Errorf("failed to read ABI file: %v", err)
} 
var jsonABI map[string]interface{}

err = json.Unmarshal(abiFile,&jsonABI)
if err != nil {
    return abi.ABI{}, fmt.Errorf("failed to parse ABI JSON: %v", err)
}

abiJSON,err := json.Marshal(jsonABI["abi"])
if err != nil{
    fmt.Println(err)
    return abi.ABI{},fmt.Errorf("error")
}
parsedABI,err := abi.JSON(strings.NewReader(string(abiJSON)))
if err != nil {
    return abi.ABI{}, fmt.Errorf("failed to parse ABI: %v", err)
}
return parsedABI,nil
 }



 func NewWalletFactory(address common.Address, backend bind.ContractBackend) (*WalletFactory, error) {
    parsedABI, err := loadABI("constants/walletfactoryABI.json")
    if err != nil {
        return nil, err
    }
    return &WalletFactory{abi: parsedABI, address: address, backend: backend}, nil
}
func (w *WalletFactory) CreateWallet(opts *bind.TransactOpts, userId string) (*types.Transaction, error) {
    data, err := w.abi.Pack("createWallet", userId)
    if err != nil {
        return nil, err
    }
    tx := types.NewTransaction(opts.Nonce.Uint64(), w.address, opts.Value, opts.GasLimit, opts.GasPrice, data)
    signedTx, err := opts.Signer(opts.From, tx)
    if err != nil {
        return nil, err
    }
    err = w.backend.SendTransaction(context.Background(), signedTx)
    if err != nil {
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
        return event, err
    }
    return event, nil
}



func CreateWallet(userID string) (string, error) {
    // Create a custom HTTP client
    httpClient := &http.Client{}
    
    // Create a custom HTTP RoundTripper
    customTransport := &customTransport{
        underlying: http.DefaultTransport,
        apiKey:     "t-66f6f6ef6be651758a55d255-01527e40b9bd45c1845d4a4b",
    }
    httpClient.Transport = customTransport

    // Create an rpc.Client with the custom HTTP client
    rpcClient, err := rpc.DialHTTPWithClient("https://ethereum-sepolia.gateway.tatum.io/", httpClient)
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
    }

    // Create an ethclient.Client with the rpc.Client
    client := ethclient.NewClient(rpcClient)

    // Load your private key
    privateKey, err := crypto.HexToECDSA("your_private_key_hex")
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to load private key: %v", err)
    }

    // Create an authorized transactor
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // 11155111 is the chain ID for Sepolia
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to create authorized transactor: %v", err)
    }

    // Set gas price and limit
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to suggest gas price: %v", err)
    }
    auth.GasPrice = gasPrice
    auth.GasLimit = uint64(3000000) // Adjust as needed

    // Load the WalletFactory contract
    factoryAddress := common.HexToAddress("YOUR_WALLET_FACTORY_CONTRACT_ADDRESS")
    factory, err := NewWalletFactory(factoryAddress, client)
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to load WalletFactory contract: %v", err)
    }

    // Call the createWallet function
    tx, err := factory.CreateWallet(auth, userID)
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to create wallet: %v", err)
    }

    // Wait for the transaction to be mined
    receipt, err := bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to wait for transaction to be mined: %v", err)
    }
    // Get the wallet address from the event logs
    event, err := factory.ParseWalletCreated(*receipt.Logs[0])
    if err != nil {
        fmt.Println("error",err)
        return "", fmt.Errorf("failed to parse WalletCreated event: %v", err)
    }

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


