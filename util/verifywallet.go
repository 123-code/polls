package util

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func VerifyContract() {
	const contractAddress = "0x09cE5cC3737C7a6F2745FE39a75e9397cB68C666"
	const owneraddress = "0x141920D050A904B6e3feDd732200fbF9B1f0b65A"
	const nftcontractaddress = "0xF554f9646581F79aF174F144Ea4De42AE13AF9c1"

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/682c39bac1294baeb74ae767786db1ca")
	if err != nil {
		log.Fatalf("Error al conectar a la red: %v", err)
	}
	defer client.Close()

	fmt.Println("Conectado a la red Sepolia")

	abiFile, err := os.ReadFile("/Users/joseignacionaranjo/Polls_backend/util/wallet.json")
	if err != nil {
		log.Fatalf("Error al leer el archivo ABI: %v", err)
	}

	parsedABI, err := abi.JSON(strings.NewReader(string(abiFile)))
	if err != nil {
		log.Fatalf("Error al parsear el ABI: %v", err)
	}
	fmt.Println("ABI parseado:", parsedABI)

	contract := common.HexToAddress(contractAddress)
	fmt.Printf("Interacción con contrato en: %s\n", contract.Hex())

	// Convertimos las direcciones a common.Address
	owner := common.HexToAddress(owneraddress)
	nftContract := common.HexToAddress(nftcontractaddress)

	
	data, err := parsedABI.Pack("initialize", owner, nftContract)
	if err != nil {
		log.Fatalf("Error al empaquetar la llamada a la función initialize: %v", err)
	}
	fmt.Println("Datos empaquetados:", data)

	nftdata, err := parsedABI.Pack("mintNFT")
	if err != nil {
		log.Fatalf("Error al empaquetar la llamada a la función initialize: %v", err)
	}
	fmt.Println("Datos empaquetados:", nftdata)

	privateKey, err := crypto.HexToECDSA("6476049119d25379b9cd234d390c757db268cac2729d54c19c6d1b884a9da5e9")
    if err != nil {
        log.Fatalf("Error al convertir la clave privada: %v", err)
    }
	publicKey := privateKey.Public()
    publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
    if !ok {
        log.Fatal("Error al convertir la clave pública")
    }

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
    nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
    if err != nil {
        log.Fatalf("Error al obtener el nonce: %v", err)
    }

	gasPrice := big.NewInt(1000000000) 

	tx := types.NewTransaction(nonce, contract, big.NewInt(0), uint64(21788), gasPrice, nftdata)

    chainID, err := client.NetworkID(context.Background())
    if err != nil {
        log.Fatalf("Error al obtener el ID de la red: %v", err)
    }

    signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
    if err != nil {
        log.Fatalf("Error al firmar la transacción: %v", err)
    }

    err = client.SendTransaction(context.Background(), signedTx)
    if err != nil {
        log.Fatalf("Error al enviar la transacción: %v", err)
    }

	fmt.Printf("Transacción enviada: %s\n", signedTx.Hash().Hex())


}