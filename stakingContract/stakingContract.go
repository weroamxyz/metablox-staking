// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingContract

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
)

// StakingContractMetaData contains all meta data concerning the StakingContract contract.
var StakingContractMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"receiveTokens\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingContractABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingContractMetaData.ABI instead.
var StakingContractABI = StakingContractMetaData.ABI

// StakingContract is an auto generated Go binding around an Ethereum contract.
type StakingContract struct {
	StakingContractCaller     // Read-only binding to the contract
	StakingContractTransactor // Write-only binding to the contract
	StakingContractFilterer   // Log filterer for contract events
}

// StakingContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingContractSession struct {
	Contract     *StakingContract  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingContractCallerSession struct {
	Contract *StakingContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StakingContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingContractTransactorSession struct {
	Contract     *StakingContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StakingContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingContractRaw struct {
	Contract *StakingContract // Generic contract binding to access the raw methods on
}

// StakingContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingContractCallerRaw struct {
	Contract *StakingContractCaller // Generic read-only contract binding to access the raw methods on
}

// StakingContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingContractTransactorRaw struct {
	Contract *StakingContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingContract creates a new instance of StakingContract, bound to a specific deployed contract.
func NewStakingContract(address common.Address, backend bind.ContractBackend) (*StakingContract, error) {
	contract, err := bindStakingContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakingContract{StakingContractCaller: StakingContractCaller{contract: contract}, StakingContractTransactor: StakingContractTransactor{contract: contract}, StakingContractFilterer: StakingContractFilterer{contract: contract}}, nil
}

// NewStakingContractCaller creates a new read-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractCaller(address common.Address, caller bind.ContractCaller) (*StakingContractCaller, error) {
	contract, err := bindStakingContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractCaller{contract: contract}, nil
}

// NewStakingContractTransactor creates a new write-only instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingContractTransactor, error) {
	contract, err := bindStakingContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingContractTransactor{contract: contract}, nil
}

// NewStakingContractFilterer creates a new log filterer instance of StakingContract, bound to a specific deployed contract.
func NewStakingContractFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingContractFilterer, error) {
	contract, err := bindStakingContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingContractFilterer{contract: contract}, nil
}

// bindStakingContract binds a generic wrapper to an already deployed contract.
func bindStakingContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.StakingContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.StakingContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakingContract *StakingContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakingContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakingContract *StakingContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakingContract *StakingContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakingContract.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractSession) Owner() (common.Address, error) {
	return _StakingContract.Contract.Owner(&_StakingContract.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_StakingContract *StakingContractCallerSession) Owner() (common.Address, error) {
	return _StakingContract.Contract.Owner(&_StakingContract.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0x9e1a4d19.
//
// Solidity: function tokenBalance() view returns(uint256)
func (_StakingContract *StakingContractCaller) TokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakingContract.contract.Call(opts, &out, "tokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBalance is a free data retrieval call binding the contract method 0x9e1a4d19.
//
// Solidity: function tokenBalance() view returns(uint256)
func (_StakingContract *StakingContractSession) TokenBalance() (*big.Int, error) {
	return _StakingContract.Contract.TokenBalance(&_StakingContract.CallOpts)
}

// TokenBalance is a free data retrieval call binding the contract method 0x9e1a4d19.
//
// Solidity: function tokenBalance() view returns(uint256)
func (_StakingContract *StakingContractCallerSession) TokenBalance() (*big.Int, error) {
	return _StakingContract.Contract.TokenBalance(&_StakingContract.CallOpts)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) returns()
func (_StakingContract *StakingContractTransactor) Mint(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "mint", _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) returns()
func (_StakingContract *StakingContractSession) Mint(_amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Mint(&_StakingContract.TransactOpts, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 _amount) returns()
func (_StakingContract *StakingContractTransactorSession) Mint(_amount *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Mint(&_StakingContract.TransactOpts, _amount)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x35729130.
//
// Solidity: function receiveTokens(address _from, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractTransactor) ReceiveTokens(opts *bind.TransactOpts, _from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "receiveTokens", _from, _value)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x35729130.
//
// Solidity: function receiveTokens(address _from, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractSession) ReceiveTokens(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.ReceiveTokens(&_StakingContract.TransactOpts, _from, _value)
}

// ReceiveTokens is a paid mutator transaction binding the contract method 0x35729130.
//
// Solidity: function receiveTokens(address _from, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractTransactorSession) ReceiveTokens(_from common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.ReceiveTokens(&_StakingContract.TransactOpts, _from, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Transfer(&_StakingContract.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool success)
func (_StakingContract *StakingContractTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _StakingContract.Contract.Transfer(&_StakingContract.TransactOpts, _to, _value)
}

// StakingContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakingContract contract.
type StakingContractTransferIterator struct {
	Event *StakingContractTransfer // Event containing the contract specifics and raw log

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
func (it *StakingContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingContractTransfer)
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
		it.Event = new(StakingContractTransfer)
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
func (it *StakingContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingContractTransfer represents a Transfer event raised by the StakingContract contract.
type StakingContractTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_StakingContract *StakingContractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address) (*StakingContractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StakingContract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StakingContractTransferIterator{contract: _StakingContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_StakingContract *StakingContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakingContractTransfer, _from []common.Address, _to []common.Address) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _StakingContract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingContractTransfer)
				if err := _StakingContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed _from, address indexed _to, uint256 _value)
func (_StakingContract *StakingContractFilterer) ParseTransfer(log types.Log) (*StakingContractTransfer, error) {
	event := new(StakingContractTransfer)
	if err := _StakingContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
