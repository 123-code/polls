package util
 
import (
    "context"
    "fmt"
    "math/big"
    "github.com/ethereum/go-ethereum/accounts/abi/bind"
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/ethereum/go-ethereum/crypto"
    
    //"your_project/walletfactory"
)

func createWallet(userID string) (string, error) {

    client, err := ethclient.Dial("https://optimism-goerli.infura.io/v3/YOUR_INFURA_PROJECT_ID")
    if err != nil {
        return "", fmt.Errorf("failed to connect to the Ethereum client: %v", err)
    }

    // Load your private key
    privateKey, err := crypto.HexToECDSA("your_private_key_hex")
    if err != nil {
        return "", fmt.Errorf("failed to load private key: %v", err)
    }

    // Create an authorized transactor
    auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(420)) // 420 is the chain ID for Optimism Goerli
    if err != nil {
        return "", fmt.Errorf("failed to create authorized transactor: %v", err)
    }


    gasPrice, err := client.SuggestGasPrice(context.Background())
    if err != nil {
        return "", fmt.Errorf("failed to suggest gas price: %v", err)
    }
    auth.GasPrice = gasPrice
    auth.GasLimit = uint64(3000000)

    factoryAddress := common.HexToAddress("YOUR_WALLET_FACTORY_CONTRACT_ADDRESS")
    factory, err := walletfactory.NewWalletFactory(factoryAddress, client)
    if err != nil {
        return "", fmt.Errorf("failed to load WalletFactory contract: %v", err)
    }

    tx, err := factory.CreateWallet(auth, userID)
    if err != nil {
        return "", fmt.Errorf("failed to create wallet: %v", err)
    }


    receipt, err := bind.WaitMined(context.Background(), client, tx)
    if err != nil {
        return "", fmt.Errorf("failed to wait for transaction to be mined: %v", err)
    }

    event, err := factory.ParseWalletCreated(*receipt.Logs[0])
    if err != nil {
        return "", fmt.Errorf("failed to parse WalletCreated event: %v", err)
    }

    return event.WalletAddress.Hex(), nil
}