package util

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"pollsbackend/bindings"

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