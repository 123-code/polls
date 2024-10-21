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

// WalletImplementation1MetaData contains all meta data concerning the WalletImplementation1 contract.
var WalletImplementation1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"execute\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_nftContractAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mintNFT\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"nftContract\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIMyToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
}

// WalletImplementation1ABI is the input ABI used to generate the binding from.
// Deprecated: Use WalletImplementation1MetaData.ABI instead.
var WalletImplementation1ABI = WalletImplementation1MetaData.ABI

// WalletImplementation1 is an auto generated Go binding around an Ethereum contract.
type WalletImplementation1 struct {
	WalletImplementation1Caller     // Read-only binding to the contract
	WalletImplementation1Transactor // Write-only binding to the contract
	WalletImplementation1Filterer   // Log filterer for contract events
}

// WalletImplementation1Caller is an auto generated read-only Go binding around an Ethereum contract.
type WalletImplementation1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletImplementation1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type WalletImplementation1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletImplementation1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WalletImplementation1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WalletImplementation1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WalletImplementation1Session struct {
	Contract     *WalletImplementation1 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WalletImplementation1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WalletImplementation1CallerSession struct {
	Contract *WalletImplementation1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// WalletImplementation1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WalletImplementation1TransactorSession struct {
	Contract     *WalletImplementation1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// WalletImplementation1Raw is an auto generated low-level Go binding around an Ethereum contract.
type WalletImplementation1Raw struct {
	Contract *WalletImplementation1 // Generic contract binding to access the raw methods on
}

// WalletImplementation1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WalletImplementation1CallerRaw struct {
	Contract *WalletImplementation1Caller // Generic read-only contract binding to access the raw methods on
}

// WalletImplementation1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WalletImplementation1TransactorRaw struct {
	Contract *WalletImplementation1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewWalletImplementation1 creates a new instance of WalletImplementation1, bound to a specific deployed contract.
func NewWalletImplementation1(address common.Address, backend bind.ContractBackend) (*WalletImplementation1, error) {
	contract, err := bindWalletImplementation1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WalletImplementation1{WalletImplementation1Caller: WalletImplementation1Caller{contract: contract}, WalletImplementation1Transactor: WalletImplementation1Transactor{contract: contract}, WalletImplementation1Filterer: WalletImplementation1Filterer{contract: contract}}, nil
}

// NewWalletImplementation1Caller creates a new read-only instance of WalletImplementation1, bound to a specific deployed contract.
func NewWalletImplementation1Caller(address common.Address, caller bind.ContractCaller) (*WalletImplementation1Caller, error) {
	contract, err := bindWalletImplementation1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WalletImplementation1Caller{contract: contract}, nil
}

// NewWalletImplementation1Transactor creates a new write-only instance of WalletImplementation1, bound to a specific deployed contract.
func NewWalletImplementation1Transactor(address common.Address, transactor bind.ContractTransactor) (*WalletImplementation1Transactor, error) {
	contract, err := bindWalletImplementation1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WalletImplementation1Transactor{contract: contract}, nil
}

// NewWalletImplementation1Filterer creates a new log filterer instance of WalletImplementation1, bound to a specific deployed contract.
func NewWalletImplementation1Filterer(address common.Address, filterer bind.ContractFilterer) (*WalletImplementation1Filterer, error) {
	contract, err := bindWalletImplementation1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WalletImplementation1Filterer{contract: contract}, nil
}

// bindWalletImplementation1 binds a generic wrapper to an already deployed contract.
func bindWalletImplementation1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WalletImplementation1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletImplementation1 *WalletImplementation1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletImplementation1.Contract.WalletImplementation1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletImplementation1 *WalletImplementation1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.WalletImplementation1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletImplementation1 *WalletImplementation1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.WalletImplementation1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WalletImplementation1 *WalletImplementation1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WalletImplementation1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WalletImplementation1 *WalletImplementation1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WalletImplementation1 *WalletImplementation1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.contract.Transact(opts, method, params...)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletImplementation1 *WalletImplementation1Caller) NftContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletImplementation1.contract.Call(opts, &out, "nftContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletImplementation1 *WalletImplementation1Session) NftContract() (common.Address, error) {
	return _WalletImplementation1.Contract.NftContract(&_WalletImplementation1.CallOpts)
}

// NftContract is a free data retrieval call binding the contract method 0xd56d229d.
//
// Solidity: function nftContract() view returns(address)
func (_WalletImplementation1 *WalletImplementation1CallerSession) NftContract() (common.Address, error) {
	return _WalletImplementation1.Contract.NftContract(&_WalletImplementation1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletImplementation1 *WalletImplementation1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WalletImplementation1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletImplementation1 *WalletImplementation1Session) Owner() (common.Address, error) {
	return _WalletImplementation1.Contract.Owner(&_WalletImplementation1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WalletImplementation1 *WalletImplementation1CallerSession) Owner() (common.Address, error) {
	return _WalletImplementation1.Contract.Owner(&_WalletImplementation1.CallOpts)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_WalletImplementation1 *WalletImplementation1Transactor) Execute(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletImplementation1.contract.Transact(opts, "execute", target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_WalletImplementation1 *WalletImplementation1Session) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Execute(&_WalletImplementation1.TransactOpts, target, value, data)
}

// Execute is a paid mutator transaction binding the contract method 0xb61d27f6.
//
// Solidity: function execute(address target, uint256 value, bytes data) returns()
func (_WalletImplementation1 *WalletImplementation1TransactorSession) Execute(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Execute(&_WalletImplementation1.TransactOpts, target, value, data)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletImplementation1 *WalletImplementation1Transactor) Initialize(opts *bind.TransactOpts, _owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletImplementation1.contract.Transact(opts, "initialize", _owner, _nftContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletImplementation1 *WalletImplementation1Session) Initialize(_owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Initialize(&_WalletImplementation1.TransactOpts, _owner, _nftContractAddress)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _owner, address _nftContractAddress) returns()
func (_WalletImplementation1 *WalletImplementation1TransactorSession) Initialize(_owner common.Address, _nftContractAddress common.Address) (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Initialize(&_WalletImplementation1.TransactOpts, _owner, _nftContractAddress)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletImplementation1 *WalletImplementation1Transactor) MintNFT(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletImplementation1.contract.Transact(opts, "mintNFT")
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletImplementation1 *WalletImplementation1Session) MintNFT() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.MintNFT(&_WalletImplementation1.TransactOpts)
}

// MintNFT is a paid mutator transaction binding the contract method 0x14f710fe.
//
// Solidity: function mintNFT() payable returns()
func (_WalletImplementation1 *WalletImplementation1TransactorSession) MintNFT() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.MintNFT(&_WalletImplementation1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletImplementation1 *WalletImplementation1Transactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletImplementation1.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletImplementation1 *WalletImplementation1Session) Withdraw() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Withdraw(&_WalletImplementation1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_WalletImplementation1 *WalletImplementation1TransactorSession) Withdraw() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Withdraw(&_WalletImplementation1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletImplementation1 *WalletImplementation1Transactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WalletImplementation1.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletImplementation1 *WalletImplementation1Session) Receive() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Receive(&_WalletImplementation1.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WalletImplementation1 *WalletImplementation1TransactorSession) Receive() (*types.Transaction, error) {
	return _WalletImplementation1.Contract.Receive(&_WalletImplementation1.TransactOpts)
}
