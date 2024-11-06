// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// WalletContractMetaData contains all meta data concerning the WalletContract contract.
var WalletContractMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nftContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintNFT\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"nftContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMyToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
}

// WalletContractABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletContractMetaData.ABI instead.
var WalletContractABI = WalletContractMetaData.ABI

// WalletContract is an auto generated Go binding around an Ethereum contract.
type WalletContract struct {
	WalletContractCaller     // Read-only binding to the contract
	WalletContractTransactor // Write-only binding to the contract
	WalletContractFilterer   // Log filterer for contract events
}

// WalletContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type WalletContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletContractSession struct {
	Contract     *WalletContract   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WalletContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletContractCallerSession struct {
	Contract *WalletContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// WalletContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletContractTransactorSession struct {
	Contract     *WalletContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// WalletContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type WalletContractRaw struct {
	Contract *WalletContract // Generic contract binding to access the raw methods on
}

// WalletContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletContractCallerRaw struct {
	Contract *WalletContractCaller // Generic read-only contract binding to access the raw methods on
}

// WalletContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletContractTransactorRaw struct {
	Contract *WalletContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletContract creates a new instance of WalletContract, bound to a specific deployed contract.
func NewWalletContract(address common.Address, backend bind.ContractBackend) (*WalletContract, error) {
	contract, err := bindWalletContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletContract{WalletContractCaller: WalletContractCaller{contract: contract}, WalletContractTransactor: WalletContractTransactor{contract: contract}, WalletContractFilterer: WalletContractFilterer{contract: contract}}, nil
}

// NewWalletContractCaller creates a new read-only instance of WalletContract, bound to a specific deployed contract.
func NewWalletContractCaller(address common.Address, caller bind.ContractCaller) (*WalletContractCaller, error) {
	contract, err := bindWalletContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletContractCaller{contract: contract}, nil
}

// NewWalletContractTransactor creates a new write-only instance of WalletContract, bound to a specific deployed contract.
func NewWalletContractTransactor(address common.Address, transactor bind.ContractTransactor) (*WalletContractTransactor, error) {
	contract, err := bindWalletContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletContractTransactor{contract: contract}, nil
}

// NewWalletContractFilterer creates a new log filterer instance of WalletContract, bound to a specific deployed contract.
func NewWalletContractFilterer(address common.Address, filterer bind.ContractFilterer) (*WalletContractFilterer, error) {
	contract, err := bindWalletContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletContractFilterer{contract: contract}, nil
}

// bindWalletContract binds a generic wrapper to an already deployed contract.
func bindWalletContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletContract *WalletContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletContract.Contract.WalletContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletContract *WalletContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletContract.Contract.WalletContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletContract *WalletContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletContract.Contract.WalletContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletContract *WalletContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletContract *WalletContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletContract *WalletContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletContract.Contract.contract.Transact(opts, method, params...)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletContract *WalletContractCaller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletContract.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletContract *WalletContractSession) NftContract() (common.Address, error) {
	return _WalletContract.Contract.NftContract(&_WalletContract.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletContract *WalletContractCallerSession) NftContract() (common.Address, error) {
	return _WalletContract.Contract.NftContract(&_WalletContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletContract *WalletContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletContract *WalletContractSession) Owner() (common.Address, error) {
	return _WalletContract.Contract.Owner(&_WalletContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletContract *WalletContractCallerSession) Owner() (common.Address, error) {
	return _WalletContract.Contract.Owner(&_WalletContract.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xda0980c7.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes signature) returns()
func (_WalletContract *WalletContractTransactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _WalletContract.contract.Transact(opts, "execute", target, value, data, signature)
}

// Execute is a paid mutator transaction binding the contract method 0xda0980c7.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes signature) returns()
func (_WalletContract *WalletContractSession) Execute(target common.Address, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _WalletContract.Contract.Execute(&_WalletContract.TransactOpts, target, value, data, signature)
}

// Execute is a paid mutator transaction binding the contract method 0xda0980c7.
//
// Solidity: function execute(address target, uint256 value, bytes data, bytes signature) returns()
func (_WalletContract *WalletContractTransactorSession) Execute(target common.Address, value *big.Int, data []byte, signature []byte) (*types.Transaction, error) {
	return _WalletContract.Contract.Execute(&_WalletContract.TransactOpts, target, value, data, signature)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletContract *WalletContractTransactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletContract.contract.Transact(opts, "initialize", _owner, _nftContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletContract *WalletContractSession) Initialize(_owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletContract.Contract.Initialize(&_WalletContract.TransactOpts, _owner, _nftContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletContract *WalletContractTransactorSession) Initialize(_owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletContract.Contract.Initialize(&_WalletContract.TransactOpts, _owner, _nftContractAddress)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletContract *WalletContractTransactor) MintNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletContract.contract.Transact(opts, "mintNFT")
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletContract *WalletContractSession) MintNFT() (*types.Transaction, error) {
	return _WalletContract.Contract.MintNFT(&_WalletContract.TransactOpts)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletContract *WalletContractTransactorSession) MintNFT() (*types.Transaction, error) {
	return _WalletContract.Contract.MintNFT(&_WalletContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletContract *WalletContractTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletContract.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletContract *WalletContractSession) Withdraw() (*types.Transaction, error) {
	return _WalletContract.Contract.Withdraw(&_WalletContract.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletContract *WalletContractTransactorSession) Withdraw() (*types.Transaction, error) {
	return _WalletContract.Contract.Withdraw(&_WalletContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletContract *WalletContractTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletContract.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletContract *WalletContractSession) Receive() (*types.Transaction, error) {
	return _WalletContract.Contract.Receive(&_WalletContract.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletContract *WalletContractTransactorSession) Receive() (*types.Transaction, error) {
	return _WalletContract.Contract.Receive(&_WalletContract.TransactOpts)
}
