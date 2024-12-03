package util

import (
	"context"
	//"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"pollsbackend/bindings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)


func ConnectVotingSystem() (*bindings.Bindings, *ethclient.Client) {
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
        log.Fatalf("Failed to connect to Ethereum client: %v", err)
    }

    contractAddress := common.HexToAddress("0x357d43845E59c6A140E57827b95405199120dF4e")
    instance, err := bindings.NewBindings(contractAddress, client)
    if err != nil {
        log.Fatalf("Failed to instantiate contract: %v", err)
    }

    return instance, client
}


func GetAuth(privateKey string) *bind.TransactOpts {

	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}


	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(11155111))
	if err != nil {

		log.Fatalf("Failed to create transactor: %v", err)
	}

	return auth
}

func Vote(candidateIndex uint64, privateKey string) error {

    instance, client := ConnectVotingSystem()
    defer client.Close()

    key, err := crypto.HexToECDSA(privateKey)
    if err != nil {
        return fmt.Errorf("failed to parse private key: %w", err)
    }


    auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(11155111))
    if err != nil {
        return fmt.Errorf("failed to create transactor: %w", err)
    }

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		return fmt.Errorf("failed to get gas price: %w", err)
	}

gasPriceWithPremium := new(big.Int).Mul(gasPrice, big.NewInt(120))
gasPriceWithPremium.Div(gasPriceWithPremium, big.NewInt(100))
fmt.Printf("Network Gas Price: %v Gwei", new(big.Float).Quo(
    new(big.Float).SetInt(gasPrice), 
    big.NewFloat(1e9),
))
auth.GasPrice = gasPriceWithPremium
fmt.Println(auth.GasPrice)

    tx, err := instance.Vote(auth, big.NewInt(int64(candidateIndex)))

	fmt.Println(tx.Hash().String())
    if err != nil {
        return fmt.Errorf("failed to cast vote: %w", err)
    }

    receipt, err := bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        return fmt.Errorf("error waiting for transaction to be mined: %w", err)
    }

    if receipt.Status != 1 {
        return fmt.Errorf("transaction failed: status %d", receipt.Status)
    }

    log.Printf("Vote transaction successful. Hash: %s\n", tx.Hash().Hex())
    return nil
}



func GetLeader() {
	instance, client := ConnectVotingSystem()
	fmt.Println(client)

	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	leaderName, voteCount, err := instance.GetLeader(callOpts)
	if err != nil {
		log.Fatalf("Failed to get leader: %v", err)
	}

	fmt.Printf("Current leader: %s with %d votes\n", leaderName, voteCount.Uint64())
}


func GetAllCandidates() {
	instance, client := ConnectVotingSystem()
fmt.Println(client)
	callOpts := &bind.CallOpts{
		Context: context.Background(),
	}

	candidates, err := instance.GetAllCandidates(callOpts)
	if err != nil {
		log.Fatalf("Failed to get candidates: %v", err)
	}

	fmt.Println("Candidates:")
	for i, candidate := range candidates {
		fmt.Printf("Index: %d, Name: %s, Votes: %d\n", i, candidate.Name, candidate.VoteCount.Uint64())
	}
}


func AddCandidate(name string, privateKey string) {
	instance,client := ConnectVotingSystem()
	fmt.Println(client)
	auth := GetAuth(privateKey)

	tx, err := instance.AddCandidate(auth, name)
	if err != nil {
		log.Fatalf("Failed to add candidate: %v", err)
	}

	fmt.Printf("Add candidate transaction submitted: %s\n", tx.Hash().Hex())
}