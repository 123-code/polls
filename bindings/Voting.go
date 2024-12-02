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

// VotingSystemCandidate is an auto generated low-level Go binding around an user-defined struct.
type VotingSystemCandidate struct {
	Name      string
	VoteCount *big.Int
}

// BindingsMetaData contains all meta data concerning the Bindings contract.
var BindingsMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_candidateNames\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addCandidate\",\"inputs\":[{\"name\":\"_name\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"candidateIndex\",\"inputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"candidates\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"voteCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getAllCandidates\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structVotingSystem.Candidate[]\",\"components\":[{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"voteCount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLeader\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasVoted\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"vote\",\"inputs\":[{\"name\":\"_candidateIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"VoteCast\",\"inputs\":[{\"name\":\"voter\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"candidateIndex\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
}

// BindingsABI is the input ABI used to generate the binding from.
// Deprecated: Use BindingsMetaData.ABI instead.
var BindingsABI = BindingsMetaData.ABI

// Bindings is an auto generated Go binding around an Ethereum contract.
type Bindings struct {
	BindingsCaller     // Read-only binding to the contract
	BindingsTransactor // Write-only binding to the contract
	BindingsFilterer   // Log filterer for contract events
}

// BindingsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BindingsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BindingsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BindingsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BindingsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BindingsSession struct {
	Contract     *Bindings         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BindingsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BindingsCallerSession struct {
	Contract *BindingsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BindingsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BindingsTransactorSession struct {
	Contract     *BindingsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BindingsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BindingsRaw struct {
	Contract *Bindings // Generic contract binding to access the raw methods on
}

// BindingsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BindingsCallerRaw struct {
	Contract *BindingsCaller // Generic read-only contract binding to access the raw methods on
}

// BindingsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BindingsTransactorRaw struct {
	Contract *BindingsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBindings creates a new instance of Bindings, bound to a specific deployed contract.
func NewBindings(address common.Address, backend bind.ContractBackend) (*Bindings, error) {
	contract, err := bindBindings(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bindings{BindingsCaller: BindingsCaller{contract: contract}, BindingsTransactor: BindingsTransactor{contract: contract}, BindingsFilterer: BindingsFilterer{contract: contract}}, nil
}

// NewBindingsCaller creates a new read-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsCaller(address common.Address, caller bind.ContractCaller) (*BindingsCaller, error) {
	contract, err := bindBindings(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsCaller{contract: contract}, nil
}

// NewBindingsTransactor creates a new write-only instance of Bindings, bound to a specific deployed contract.
func NewBindingsTransactor(address common.Address, transactor bind.ContractTransactor) (*BindingsTransactor, error) {
	contract, err := bindBindings(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BindingsTransactor{contract: contract}, nil
}

// NewBindingsFilterer creates a new log filterer instance of Bindings, bound to a specific deployed contract.
func NewBindingsFilterer(address common.Address, filterer bind.ContractFilterer) (*BindingsFilterer, error) {
	contract, err := bindBindings(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BindingsFilterer{contract: contract}, nil
}

// bindBindings binds a generic wrapper to an already deployed contract.
func bindBindings(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BindingsMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.BindingsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.BindingsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bindings *BindingsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bindings.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bindings *BindingsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bindings *BindingsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bindings.Contract.contract.Transact(opts, method, params...)
}

// CandidateIndex is a free data retrieval call binding the contract method 0x9df9ae0c.
//
// Solidity: function candidateIndex(string ) view returns(uint256)
func (_Bindings *BindingsCaller) CandidateIndex(opts *bind.CallOpts, arg0 string) (*big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "candidateIndex", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CandidateIndex is a free data retrieval call binding the contract method 0x9df9ae0c.
//
// Solidity: function candidateIndex(string ) view returns(uint256)
func (_Bindings *BindingsSession) CandidateIndex(arg0 string) (*big.Int, error) {
	return _Bindings.Contract.CandidateIndex(&_Bindings.CallOpts, arg0)
}

// CandidateIndex is a free data retrieval call binding the contract method 0x9df9ae0c.
//
// Solidity: function candidateIndex(string ) view returns(uint256)
func (_Bindings *BindingsCallerSession) CandidateIndex(arg0 string) (*big.Int, error) {
	return _Bindings.Contract.CandidateIndex(&_Bindings.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Bindings *BindingsCaller) Candidates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "candidates", arg0)

	outstruct := new(struct {
		Name      string
		VoteCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.VoteCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Bindings *BindingsSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Bindings.Contract.Candidates(&_Bindings.CallOpts, arg0)
}

// Candidates is a free data retrieval call binding the contract method 0x3477ee2e.
//
// Solidity: function candidates(uint256 ) view returns(string name, uint256 voteCount)
func (_Bindings *BindingsCallerSession) Candidates(arg0 *big.Int) (struct {
	Name      string
	VoteCount *big.Int
}, error) {
	return _Bindings.Contract.Candidates(&_Bindings.CallOpts, arg0)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256)[])
func (_Bindings *BindingsCaller) GetAllCandidates(opts *bind.CallOpts) ([]VotingSystemCandidate, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getAllCandidates")

	if err != nil {
		return *new([]VotingSystemCandidate), err
	}

	out0 := *abi.ConvertType(out[0], new([]VotingSystemCandidate)).(*[]VotingSystemCandidate)

	return out0, err

}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256)[])
func (_Bindings *BindingsSession) GetAllCandidates() ([]VotingSystemCandidate, error) {
	return _Bindings.Contract.GetAllCandidates(&_Bindings.CallOpts)
}

// GetAllCandidates is a free data retrieval call binding the contract method 0x2e6997fe.
//
// Solidity: function getAllCandidates() view returns((string,uint256)[])
func (_Bindings *BindingsCallerSession) GetAllCandidates() ([]VotingSystemCandidate, error) {
	return _Bindings.Contract.GetAllCandidates(&_Bindings.CallOpts)
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns(string, uint256)
func (_Bindings *BindingsCaller) GetLeader(opts *bind.CallOpts) (string, *big.Int, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "getLeader")

	if err != nil {
		return *new(string), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns(string, uint256)
func (_Bindings *BindingsSession) GetLeader() (string, *big.Int, error) {
	return _Bindings.Contract.GetLeader(&_Bindings.CallOpts)
}

// GetLeader is a free data retrieval call binding the contract method 0x4c051f14.
//
// Solidity: function getLeader() view returns(string, uint256)
func (_Bindings *BindingsCallerSession) GetLeader() (string, *big.Int, error) {
	return _Bindings.Contract.GetLeader(&_Bindings.CallOpts)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Bindings *BindingsCaller) HasVoted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "hasVoted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Bindings *BindingsSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.HasVoted(&_Bindings.CallOpts, arg0)
}

// HasVoted is a free data retrieval call binding the contract method 0x09eef43e.
//
// Solidity: function hasVoted(address ) view returns(bool)
func (_Bindings *BindingsCallerSession) HasVoted(arg0 common.Address) (bool, error) {
	return _Bindings.Contract.HasVoted(&_Bindings.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bindings.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Bindings *BindingsCallerSession) Owner() (common.Address, error) {
	return _Bindings.Contract.Owner(&_Bindings.CallOpts)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Bindings *BindingsTransactor) AddCandidate(opts *bind.TransactOpts, _name string) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "addCandidate", _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Bindings *BindingsSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Bindings.Contract.AddCandidate(&_Bindings.TransactOpts, _name)
}

// AddCandidate is a paid mutator transaction binding the contract method 0x462e91ec.
//
// Solidity: function addCandidate(string _name) returns()
func (_Bindings *BindingsTransactorSession) AddCandidate(_name string) (*types.Transaction, error) {
	return _Bindings.Contract.AddCandidate(&_Bindings.TransactOpts, _name)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Bindings *BindingsTransactor) Vote(opts *bind.TransactOpts, _candidateIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.contract.Transact(opts, "vote", _candidateIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Bindings *BindingsSession) Vote(_candidateIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Vote(&_Bindings.TransactOpts, _candidateIndex)
}

// Vote is a paid mutator transaction binding the contract method 0x0121b93f.
//
// Solidity: function vote(uint256 _candidateIndex) returns()
func (_Bindings *BindingsTransactorSession) Vote(_candidateIndex *big.Int) (*types.Transaction, error) {
	return _Bindings.Contract.Vote(&_Bindings.TransactOpts, _candidateIndex)
}

// BindingsVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the Bindings contract.
type BindingsVoteCastIterator struct {
	Event *BindingsVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BindingsVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BindingsVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BindingsVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BindingsVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BindingsVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BindingsVoteCast represents a VoteCast event raised by the Bindings contract.
type BindingsVoteCast struct {
	Voter          common.Address
	CandidateIndex *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 candidateIndex)
func (_Bindings *BindingsFilterer) FilterVoteCast(opts *bind.FilterOpts) (*BindingsVoteCastIterator, error) {

	logs, sub, err := _Bindings.contract.FilterLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return &BindingsVoteCastIterator{contract: _Bindings.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 candidateIndex)
func (_Bindings *BindingsFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *BindingsVoteCast) (event.Subscription, error) {

	logs, sub, err := _Bindings.contract.WatchLogs(opts, "VoteCast")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BindingsVoteCast)
				if err := _Bindings.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0xa36cc2bebb74db33e9f88110a07ef56e1b31b24b4c4f51b54b1664266e29f45b.
//
// Solidity: event VoteCast(address voter, uint256 candidateIndex)
func (_Bindings *BindingsFilterer) ParseVoteCast(log types.Log) (*BindingsVoteCast, error) {
	event := new(BindingsVoteCast)
	if err := _Bindings.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
