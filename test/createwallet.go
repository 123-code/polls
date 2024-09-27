package test

import (
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
)

func TestCreateWallet(t *testing.T) {
    // Replace with a test user ID
    userID := "testUser123"

    // Call the createWallet function
    walletAddress, err := createWallet(userID)

    // Assert that there's no error
    assert.NoError(t, err)

    // Assert that the wallet address is not empty
    assert.NotEmpty(t, walletAddress)

    // Optional: Check if the wallet address is a valid Ethereum address
    assert.True(t, common.IsHexAddress(walletAddress))

    // You might want to add more assertions here, such as checking if the wallet
    // is actually created on the blockchain, but this would require interacting
    // with the blockchain again
}