package util

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"strings"

	"pollsbackend/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	//"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ValidateWallet(){
	userID:="1804072310"
	client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
		fmt.Println(err)
        log.Fatal(err)
    }

	contractAddress := common.HexToAddress("0x84Eb5C50Fcd8d6F2eeBDb929381af5AC4e80321c")
    instance, err := bindings.NewWallet(contractAddress, client)
    if err != nil {
		fmt.Println(err)
        log.Fatal(err)
    }

	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}
	walletAddress, err := instance.UserWallets(callOpts, userID)
	if err != nil {
		fmt.Println(err)
		log.Fatalf("Failed to query userWallets: %v", err)
	}
	if walletAddress == (common.Address{}) {
		fmt.Printf("No wallet found for user: %s\n", userID)
	} else {
		fmt.Printf("Wallet address for user %s: %s\n", userID, walletAddress.Hex())
	}

} 

func InitializeUserWallet(ownerAddress string, nftContractAddress string) error {

    client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to connect to the Ethereum client: %v", err)
    }

    privateKey, err := crypto.HexToECDSA("526938daf3a62f82fc13d7abe8d063104160bfd869ddbc25e3feb6a2f8a8042e")
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to convert hex to ECDSA: %v", err)
    }


    contractAddress := common.HexToAddress("0x858581A5c619bA15f21C23598aB74e1e317ABECc")

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to create authorized transactor: %v", err)
    }


    walletInstance, err := bindings.NewWalletImplementation1(contractAddress, client)
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to instantiate contract: %v", err)
    }


    owner := common.HexToAddress(ownerAddress)
    nftContract := common.HexToAddress(nftContractAddress)


    tx, err := walletInstance.Initialize(auth, owner, nftContract)
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to initialize wallet: %v", err)
    }


    fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())

    return nil
}

func EstimateMintNFTGas(walletInstance *bindings.WalletImplementation1, auth *bind.TransactOpts) (uint64, error) {
    gasEstimate, err := walletInstance.MintNFT(&bind.TransactOpts{
        From:     auth.From,
        Nonce:    auth.Nonce,
        Signer:   auth.Signer,
        Value:    big.NewInt(1e16),
        GasLimit: 0,
        GasPrice: auth.GasPrice,
    })
    if err != nil {
        return 0, fmt.Errorf("failed to estimate gas: %v", err)
    }
    return gasEstimate.Gas(), nil
}

func MintNFT() error {
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to connect to Ethereum client: %v", err)
    }
    defer client.Close()

    privateKey, err := crypto.HexToECDSA("526938daf3a62f82fc13d7abe8d063104160bfd869ddbc25e3feb6a2f8a8042e")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to convert hex to ECDSA: %v", err)
    }

  
    contractWalletAddress := common.HexToAddress("0x0F73feCF10C6E777A4213e611a8ED5df3185A489")

    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111)) // Sepolia's chain ID
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to create authorized transactor: %v", err)
    }

    auth.From = contractWalletAddress 
    auth.Value = big.NewInt(10000000000000000) 
    auth.GasLimit = uint64(300000) 


    contractAddress := common.HexToAddress("0x858581A5c619bA15f21C23598aB74e1e317ABECc")
    walletInstance, err := bindings.NewBindings(contractAddress, client)
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to instantiate contract: %v", err)
    }

    tx, err := walletInstance.MintNFT(auth)
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to mint NFT: %v", err)
    }

    fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
    return nil
}

func MintNFTWithExecute(walletaddress,nftContractAddress string) error {
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to connect to Ethereum client: %v", err)
    }
    defer client.Close()
    privateKey, err := crypto.HexToECDSA("526938daf3a62f82fc13d7abe8d063104160bfd869ddbc25e3feb6a2f8a8042e")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to convert hex to ECDSA: %v", err)
    }
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(11155111))
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to create transactor: %v", err)
    }

    nftABI, err := abi.JSON(strings.NewReader(`[{"inputs":[],"name":"mint","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"payable","type":"function"}]`))
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to parse NFT contract ABI: %v", err)
    }

    data, err := nftABI.Pack("mint")
    if err != nil {
        fmt.Println(err)
        return fmt.Errorf("failed to encode mint function: %v", err)
    }

    messagehash := crypto.Keccak256Hash([]byte(walletaddress + nftContractAddress + string(data)))

    if err != nil {
        return fmt.Errorf("failed to sign message: %v", err)
    } else{
        fmt.Println(err)
        fmt.Println("messagehash",messagehash)
    }
/*
    signature,err := crypto.Sign(messagehash.Bytes(),privateKey)
    if err != nil{
        fmt.Println(err)
        return fmt.Errorf("failed to load wallet instance")
    }
*/
    walletinstance,err := bindings.NewWalletImplementation1(common.HexToAddress(walletaddress), client) 
    if err != nil {
        fmt.Println("wallet instance error",err)
        return fmt.Errorf("failed to load wallet instance: %v", err)
    }
    tx, err := walletinstance.Execute(auth, common.HexToAddress(nftContractAddress), big.NewInt(0), data)
    if err != nil {
        fmt.Println("execute error",err)
        return fmt.Errorf("failed to execute mint: %v", err)
    }


    fmt.Printf("Transaction sent: %s\n", tx.Hash().Hex())
    return nil
}