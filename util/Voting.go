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
	// Parse the private key
	key, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Failed to parse private key: %v", err)
	}

	// Create a transactor with the chain ID
	auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(11155111)) // Replace with your chain ID
	if err != nil {

		log.Fatalf("Failed to create transactor: %v", err)
	}

	return auth
}

func Vote(candidateIndex uint64, privateKey string) error {
    // Create Ethereum client
    client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to connect to Ethereum client: %v", err)
    }
    defer client.Close()

    // Contract address
    contractAddress := common.HexToAddress("0x357d43845E59c6A140E57827b95405199120dF4e")

    // Create contract instance
    instance, err := bindings.NewBindings(contractAddress, client)
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to instantiate contract: %v", err)
    }

    // Parse private key
    key, err := crypto.HexToECDSA(privateKey)
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to parse private key: %v", err)
    }

    // Create transactor
    auth, err := bind.NewKeyedTransactorWithChainID(key, big.NewInt(11155111))
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to create transactor: %v", err)
    }

    // Set gas limit and price
    auth.GasLimit = uint64(300000)  // Adjust gas limit as needed
    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to get gas price: %v", err)
    }
    auth.GasPrice = gasPrice

    // Call the vote function
    tx, err := instance.Vote(auth, big.NewInt(int64(candidateIndex)))
    if err != nil {
		fmt.Println(err)
        return fmt.Errorf("failed to cast vote: %v", err)
    }

    receipt, err := bind.WaitMined(context.Background(), client, tx)
    if err != nil {fmt.Println(err)
        return fmt.Errorf("error waiting for transaction to be mined: %v", err)
    }

    if receipt.Status != 1 {
		fmt.Println(err)
        return fmt.Errorf("transaction failed: status %d", receipt.Status)
    }

    fmt.Printf("Vote transaction successful. Hash: %s\n", tx.Hash().Hex())
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

// Get all candidates
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

// Add a new candidate (owner only)
func AddCandidate(name string, privateKey string) {
	instance,client := ConnectVotingSystem()
	fmt.Println(client)
	auth := GetAuth(privateKey)

	// Call the addCandidate function
	tx, err := instance.AddCandidate(auth, name)
	if err != nil {
		log.Fatalf("Failed to add candidate: %v", err)
	}

	fmt.Printf("Add candidate transaction submitted: %s\n", tx.Hash().Hex())
}