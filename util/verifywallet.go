package util

import (
	//"context"
	"fmt"
	"log"
	//"math/big"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	//"github.com/ethereum/go-ethereum/core/types"
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

	// Aquí puedes añadir la lógica para interactuar con el contrato utilizando el ABI parseado
	// Ejemplo: llamar métodos del contrato usando client.CallContract o client.TransactionSender
//0xF554f9646581F79aF174F144Ea4De42AE13AF9c1
contract = common.HexToAddress(contractAddress)
	fmt.Printf("Interacting with contract at: %s\n", contract.Hex())

	data,err :=  parsedABI.Pack("initialize",owneraddress,nftcontractaddress)
	if err != nil { 
		log.Fatalf("Error al empaquetar la llamada a la función initialize: %v", err)
	}
	fmt.Println(data)
}