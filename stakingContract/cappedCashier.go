// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package CappedCashier

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

// CappedCashierMetaData contains all meta data concerning the CappedCashier contract.
var CappedCashierMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakingToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rewardToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minStakePerTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxStakePerAddress\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldCap\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCap\",\"type\":\"uint256\"}],\"name\":\"CapacityChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MaxStakePerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"}],\"name\":\"MinStakePerTimeChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stopTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"StakingEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stopTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"roundNum\",\"type\":\"uint256\"}],\"name\":\"StakingPreparing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rewardAmount\",\"type\":\"uint256\"}],\"name\":\"Withdrawal\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"capacity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"duration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getRewardTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_stakeholder\",\"type\":\"address\"}],\"name\":\"getStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingTokenBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"loadReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxStakePerAddress\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minStakePerTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stopTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rewardAmount\",\"type\":\"uint256\"}],\"name\":\"newRound\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_token\",\"type\":\"address\"}],\"name\":\"recoverTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardLoaded\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"roundNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newCap\",\"type\":\"uint256\"}],\"name\":\"setCapcity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMaxStakePerAddress\",\"type\":\"uint256\"}],\"name\":\"setMaxStakePerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newMinStakePerTime\",\"type\":\"uint256\"}],\"name\":\"setMinStakePerTime\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakeholders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"staked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stopTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalStaked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"withdrawRewardPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawalAllowed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// CappedCashierABI is the input ABI used to generate the binding from.
// Deprecated: Use CappedCashierMetaData.ABI instead.
var CappedCashierABI = CappedCashierMetaData.ABI

// CappedCashier is an auto generated Go binding around an Ethereum contract.
type CappedCashier struct {
	CappedCashierCaller     // Read-only binding to the contract
	CappedCashierTransactor // Write-only binding to the contract
	CappedCashierFilterer   // Log filterer for contract events
}

// CappedCashierCaller is an auto generated read-only Go binding around an Ethereum contract.
type CappedCashierCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCashierTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CappedCashierTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCashierFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CappedCashierFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CappedCashierSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CappedCashierSession struct {
	Contract     *CappedCashier    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CappedCashierCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CappedCashierCallerSession struct {
	Contract *CappedCashierCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// CappedCashierTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CappedCashierTransactorSession struct {
	Contract     *CappedCashierTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// CappedCashierRaw is an auto generated low-level Go binding around an Ethereum contract.
type CappedCashierRaw struct {
	Contract *CappedCashier // Generic contract binding to access the raw methods on
}

// CappedCashierCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CappedCashierCallerRaw struct {
	Contract *CappedCashierCaller // Generic read-only contract binding to access the raw methods on
}

// CappedCashierTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CappedCashierTransactorRaw struct {
	Contract *CappedCashierTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCappedCashier creates a new instance of CappedCashier, bound to a specific deployed contract.
func NewCappedCashier(address common.Address, backend bind.ContractBackend) (*CappedCashier, error) {
	contract, err := bindCappedCashier(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CappedCashier{CappedCashierCaller: CappedCashierCaller{contract: contract}, CappedCashierTransactor: CappedCashierTransactor{contract: contract}, CappedCashierFilterer: CappedCashierFilterer{contract: contract}}, nil
}

// NewCappedCashierCaller creates a new read-only instance of CappedCashier, bound to a specific deployed contract.
func NewCappedCashierCaller(address common.Address, caller bind.ContractCaller) (*CappedCashierCaller, error) {
	contract, err := bindCappedCashier(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCashierCaller{contract: contract}, nil
}

// NewCappedCashierTransactor creates a new write-only instance of CappedCashier, bound to a specific deployed contract.
func NewCappedCashierTransactor(address common.Address, transactor bind.ContractTransactor) (*CappedCashierTransactor, error) {
	contract, err := bindCappedCashier(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CappedCashierTransactor{contract: contract}, nil
}

// NewCappedCashierFilterer creates a new log filterer instance of CappedCashier, bound to a specific deployed contract.
func NewCappedCashierFilterer(address common.Address, filterer bind.ContractFilterer) (*CappedCashierFilterer, error) {
	contract, err := bindCappedCashier(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CappedCashierFilterer{contract: contract}, nil
}

// bindCappedCashier binds a generic wrapper to an already deployed contract.
func bindCappedCashier(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CappedCashierABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCashier *CappedCashierRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedCashier.Contract.CappedCashierCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCashier *CappedCashierRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.Contract.CappedCashierTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCashier *CappedCashierRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCashier.Contract.CappedCashierTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CappedCashier *CappedCashierCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CappedCashier.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CappedCashier *CappedCashierTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CappedCashier *CappedCashierTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CappedCashier.Contract.contract.Transact(opts, method, params...)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) Capacity(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "capacity")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedCashier *CappedCashierSession) Capacity() (*big.Int, error) {
	return _CappedCashier.Contract.Capacity(&_CappedCashier.CallOpts)
}

// Capacity is a free data retrieval call binding the contract method 0x5cfc1a51.
//
// Solidity: function capacity() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) Capacity() (*big.Int, error) {
	return _CappedCashier.Contract.Capacity(&_CappedCashier.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) Duration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "duration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedCashier *CappedCashierSession) Duration() (*big.Int, error) {
	return _CappedCashier.Contract.Duration(&_CappedCashier.CallOpts)
}

// Duration is a free data retrieval call binding the contract method 0x0fb5a6b4.
//
// Solidity: function duration() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) Duration() (*big.Int, error) {
	return _CappedCashier.Contract.Duration(&_CappedCashier.CallOpts)
}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) GetRewardTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "getRewardTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierSession) GetRewardTokenBalance() (*big.Int, error) {
	return _CappedCashier.Contract.GetRewardTokenBalance(&_CappedCashier.CallOpts)
}

// GetRewardTokenBalance is a free data retrieval call binding the contract method 0x93ce5343.
//
// Solidity: function getRewardTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) GetRewardTokenBalance() (*big.Int, error) {
	return _CappedCashier.Contract.GetRewardTokenBalance(&_CappedCashier.CallOpts)
}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _stakeholder) view returns(uint256)
func (_CappedCashier *CappedCashierCaller) GetStaked(opts *bind.CallOpts, _stakeholder common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "getStaked", _stakeholder)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _stakeholder) view returns(uint256)
func (_CappedCashier *CappedCashierSession) GetStaked(_stakeholder common.Address) (*big.Int, error) {
	return _CappedCashier.Contract.GetStaked(&_CappedCashier.CallOpts, _stakeholder)
}

// GetStaked is a free data retrieval call binding the contract method 0x399080ec.
//
// Solidity: function getStaked(address _stakeholder) view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) GetStaked(_stakeholder common.Address) (*big.Int, error) {
	return _CappedCashier.Contract.GetStaked(&_CappedCashier.CallOpts, _stakeholder)
}

// GetStakingTokenBalance is a free data retrieval call binding the contract method 0x20a1318c.
//
// Solidity: function getStakingTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) GetStakingTokenBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "getStakingTokenBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingTokenBalance is a free data retrieval call binding the contract method 0x20a1318c.
//
// Solidity: function getStakingTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierSession) GetStakingTokenBalance() (*big.Int, error) {
	return _CappedCashier.Contract.GetStakingTokenBalance(&_CappedCashier.CallOpts)
}

// GetStakingTokenBalance is a free data retrieval call binding the contract method 0x20a1318c.
//
// Solidity: function getStakingTokenBalance() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) GetStakingTokenBalance() (*big.Int, error) {
	return _CappedCashier.Contract.GetStakingTokenBalance(&_CappedCashier.CallOpts)
}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) MaxStakePerAddress(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "maxStakePerAddress")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedCashier *CappedCashierSession) MaxStakePerAddress() (*big.Int, error) {
	return _CappedCashier.Contract.MaxStakePerAddress(&_CappedCashier.CallOpts)
}

// MaxStakePerAddress is a free data retrieval call binding the contract method 0x57559cf2.
//
// Solidity: function maxStakePerAddress() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) MaxStakePerAddress() (*big.Int, error) {
	return _CappedCashier.Contract.MaxStakePerAddress(&_CappedCashier.CallOpts)
}

// MinStakePerTime is a free data retrieval call binding the contract method 0x6620dca1.
//
// Solidity: function minStakePerTime() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) MinStakePerTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "minStakePerTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinStakePerTime is a free data retrieval call binding the contract method 0x6620dca1.
//
// Solidity: function minStakePerTime() view returns(uint256)
func (_CappedCashier *CappedCashierSession) MinStakePerTime() (*big.Int, error) {
	return _CappedCashier.Contract.MinStakePerTime(&_CappedCashier.CallOpts)
}

// MinStakePerTime is a free data retrieval call binding the contract method 0x6620dca1.
//
// Solidity: function minStakePerTime() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) MinStakePerTime() (*big.Int, error) {
	return _CappedCashier.Contract.MinStakePerTime(&_CappedCashier.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address staker) view returns(uint256)
func (_CappedCashier *CappedCashierCaller) Nonces(opts *bind.CallOpts, staker common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "nonces", staker)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address staker) view returns(uint256)
func (_CappedCashier *CappedCashierSession) Nonces(staker common.Address) (*big.Int, error) {
	return _CappedCashier.Contract.Nonces(&_CappedCashier.CallOpts, staker)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address staker) view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) Nonces(staker common.Address) (*big.Int, error) {
	return _CappedCashier.Contract.Nonces(&_CappedCashier.CallOpts, staker)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCashier *CappedCashierCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCashier *CappedCashierSession) Owner() (common.Address, error) {
	return _CappedCashier.Contract.Owner(&_CappedCashier.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_CappedCashier *CappedCashierCallerSession) Owner() (common.Address, error) {
	return _CappedCashier.Contract.Owner(&_CappedCashier.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCashier *CappedCashierCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCashier *CappedCashierSession) Paused() (bool, error) {
	return _CappedCashier.Contract.Paused(&_CappedCashier.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_CappedCashier *CappedCashierCallerSession) Paused() (bool, error) {
	return _CappedCashier.Contract.Paused(&_CappedCashier.CallOpts)
}

// RewardAmount is a free data retrieval call binding the contract method 0xf7b2a7be.
//
// Solidity: function rewardAmount() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) RewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "rewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardAmount is a free data retrieval call binding the contract method 0xf7b2a7be.
//
// Solidity: function rewardAmount() view returns(uint256)
func (_CappedCashier *CappedCashierSession) RewardAmount() (*big.Int, error) {
	return _CappedCashier.Contract.RewardAmount(&_CappedCashier.CallOpts)
}

// RewardAmount is a free data retrieval call binding the contract method 0xf7b2a7be.
//
// Solidity: function rewardAmount() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) RewardAmount() (*big.Int, error) {
	return _CappedCashier.Contract.RewardAmount(&_CappedCashier.CallOpts)
}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedCashier *CappedCashierCaller) RewardLoaded(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "rewardLoaded")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedCashier *CappedCashierSession) RewardLoaded() (bool, error) {
	return _CappedCashier.Contract.RewardLoaded(&_CappedCashier.CallOpts)
}

// RewardLoaded is a free data retrieval call binding the contract method 0x2a6da428.
//
// Solidity: function rewardLoaded() view returns(bool)
func (_CappedCashier *CappedCashierCallerSession) RewardLoaded() (bool, error) {
	return _CappedCashier.Contract.RewardLoaded(&_CappedCashier.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedCashier *CappedCashierCaller) RewardToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "rewardToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedCashier *CappedCashierSession) RewardToken() (common.Address, error) {
	return _CappedCashier.Contract.RewardToken(&_CappedCashier.CallOpts)
}

// RewardToken is a free data retrieval call binding the contract method 0xf7c618c1.
//
// Solidity: function rewardToken() view returns(address)
func (_CappedCashier *CappedCashierCallerSession) RewardToken() (common.Address, error) {
	return _CappedCashier.Contract.RewardToken(&_CappedCashier.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) RoundNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "roundNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CappedCashier *CappedCashierSession) RoundNum() (*big.Int, error) {
	return _CappedCashier.Contract.RoundNum(&_CappedCashier.CallOpts)
}

// RoundNum is a free data retrieval call binding the contract method 0x119b22b3.
//
// Solidity: function roundNum() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) RoundNum() (*big.Int, error) {
	return _CappedCashier.Contract.RoundNum(&_CappedCashier.CallOpts)
}

// Stakeholders is a free data retrieval call binding the contract method 0xc07cfca9.
//
// Solidity: function stakeholders(address ) view returns(uint256 staked, uint256 timestamp)
func (_CappedCashier *CappedCashierCaller) Stakeholders(opts *bind.CallOpts, arg0 common.Address) (struct {
	Staked    *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "stakeholders", arg0)

	outstruct := new(struct {
		Staked    *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Staked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Stakeholders is a free data retrieval call binding the contract method 0xc07cfca9.
//
// Solidity: function stakeholders(address ) view returns(uint256 staked, uint256 timestamp)
func (_CappedCashier *CappedCashierSession) Stakeholders(arg0 common.Address) (struct {
	Staked    *big.Int
	Timestamp *big.Int
}, error) {
	return _CappedCashier.Contract.Stakeholders(&_CappedCashier.CallOpts, arg0)
}

// Stakeholders is a free data retrieval call binding the contract method 0xc07cfca9.
//
// Solidity: function stakeholders(address ) view returns(uint256 staked, uint256 timestamp)
func (_CappedCashier *CappedCashierCallerSession) Stakeholders(arg0 common.Address) (struct {
	Staked    *big.Int
	Timestamp *big.Int
}, error) {
	return _CappedCashier.Contract.Stakeholders(&_CappedCashier.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CappedCashier *CappedCashierCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CappedCashier *CappedCashierSession) StakingToken() (common.Address, error) {
	return _CappedCashier.Contract.StakingToken(&_CappedCashier.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_CappedCashier *CappedCashierCallerSession) StakingToken() (common.Address, error) {
	return _CappedCashier.Contract.StakingToken(&_CappedCashier.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) StartTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "startTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedCashier *CappedCashierSession) StartTime() (*big.Int, error) {
	return _CappedCashier.Contract.StartTime(&_CappedCashier.CallOpts)
}

// StartTime is a free data retrieval call binding the contract method 0x78e97925.
//
// Solidity: function startTime() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) StartTime() (*big.Int, error) {
	return _CappedCashier.Contract.StartTime(&_CappedCashier.CallOpts)
}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) StopTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "stopTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedCashier *CappedCashierSession) StopTime() (*big.Int, error) {
	return _CappedCashier.Contract.StopTime(&_CappedCashier.CallOpts)
}

// StopTime is a free data retrieval call binding the contract method 0x03ff5e73.
//
// Solidity: function stopTime() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) StopTime() (*big.Int, error) {
	return _CappedCashier.Contract.StopTime(&_CappedCashier.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedCashier *CappedCashierCaller) TotalStaked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "totalStaked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedCashier *CappedCashierSession) TotalStaked() (*big.Int, error) {
	return _CappedCashier.Contract.TotalStaked(&_CappedCashier.CallOpts)
}

// TotalStaked is a free data retrieval call binding the contract method 0x817b1cd2.
//
// Solidity: function totalStaked() view returns(uint256)
func (_CappedCashier *CappedCashierCallerSession) TotalStaked() (*big.Int, error) {
	return _CappedCashier.Contract.TotalStaked(&_CappedCashier.CallOpts)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0xa1d19361.
//
// Solidity: function withdrawalAllowed() view returns(bool)
func (_CappedCashier *CappedCashierCaller) WithdrawalAllowed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CappedCashier.contract.Call(opts, &out, "withdrawalAllowed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0xa1d19361.
//
// Solidity: function withdrawalAllowed() view returns(bool)
func (_CappedCashier *CappedCashierSession) WithdrawalAllowed() (bool, error) {
	return _CappedCashier.Contract.WithdrawalAllowed(&_CappedCashier.CallOpts)
}

// WithdrawalAllowed is a free data retrieval call binding the contract method 0xa1d19361.
//
// Solidity: function withdrawalAllowed() view returns(bool)
func (_CappedCashier *CappedCashierCallerSession) WithdrawalAllowed() (bool, error) {
	return _CappedCashier.Contract.WithdrawalAllowed(&_CappedCashier.CallOpts)
}

// LoadReward is a paid mutator transaction binding the contract method 0x0ecd9bf8.
//
// Solidity: function loadReward() returns()
func (_CappedCashier *CappedCashierTransactor) LoadReward(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "loadReward")
}

// LoadReward is a paid mutator transaction binding the contract method 0x0ecd9bf8.
//
// Solidity: function loadReward() returns()
func (_CappedCashier *CappedCashierSession) LoadReward() (*types.Transaction, error) {
	return _CappedCashier.Contract.LoadReward(&_CappedCashier.TransactOpts)
}

// LoadReward is a paid mutator transaction binding the contract method 0x0ecd9bf8.
//
// Solidity: function loadReward() returns()
func (_CappedCashier *CappedCashierTransactorSession) LoadReward() (*types.Transaction, error) {
	return _CappedCashier.Contract.LoadReward(&_CappedCashier.TransactOpts)
}

// NewRound is a paid mutator transaction binding the contract method 0xc71bf91a.
//
// Solidity: function newRound(uint256 _startTime, uint256 _stopTime, uint256 _rewardAmount) returns()
func (_CappedCashier *CappedCashierTransactor) NewRound(opts *bind.TransactOpts, _startTime *big.Int, _stopTime *big.Int, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "newRound", _startTime, _stopTime, _rewardAmount)
}

// NewRound is a paid mutator transaction binding the contract method 0xc71bf91a.
//
// Solidity: function newRound(uint256 _startTime, uint256 _stopTime, uint256 _rewardAmount) returns()
func (_CappedCashier *CappedCashierSession) NewRound(_startTime *big.Int, _stopTime *big.Int, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.NewRound(&_CappedCashier.TransactOpts, _startTime, _stopTime, _rewardAmount)
}

// NewRound is a paid mutator transaction binding the contract method 0xc71bf91a.
//
// Solidity: function newRound(uint256 _startTime, uint256 _stopTime, uint256 _rewardAmount) returns()
func (_CappedCashier *CappedCashierTransactorSession) NewRound(_startTime *big.Int, _stopTime *big.Int, _rewardAmount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.NewRound(&_CappedCashier.TransactOpts, _startTime, _stopTime, _rewardAmount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCashier *CappedCashierTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCashier *CappedCashierSession) Pause() (*types.Transaction, error) {
	return _CappedCashier.Contract.Pause(&_CappedCashier.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_CappedCashier *CappedCashierTransactorSession) Pause() (*types.Transaction, error) {
	return _CappedCashier.Contract.Pause(&_CappedCashier.TransactOpts)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address _token) returns()
func (_CappedCashier *CappedCashierTransactor) RecoverTokens(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "recoverTokens", _token)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address _token) returns()
func (_CappedCashier *CappedCashierSession) RecoverTokens(_token common.Address) (*types.Transaction, error) {
	return _CappedCashier.Contract.RecoverTokens(&_CappedCashier.TransactOpts, _token)
}

// RecoverTokens is a paid mutator transaction binding the contract method 0x16114acd.
//
// Solidity: function recoverTokens(address _token) returns()
func (_CappedCashier *CappedCashierTransactorSession) RecoverTokens(_token common.Address) (*types.Transaction, error) {
	return _CappedCashier.Contract.RecoverTokens(&_CappedCashier.TransactOpts, _token)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCashier *CappedCashierTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCashier *CappedCashierSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedCashier.Contract.RenounceOwnership(&_CappedCashier.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_CappedCashier *CappedCashierTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _CappedCashier.Contract.RenounceOwnership(&_CappedCashier.TransactOpts)
}

// SetCapcity is a paid mutator transaction binding the contract method 0xb9a6a484.
//
// Solidity: function setCapcity(uint256 _newCap) returns()
func (_CappedCashier *CappedCashierTransactor) SetCapcity(opts *bind.TransactOpts, _newCap *big.Int) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "setCapcity", _newCap)
}

// SetCapcity is a paid mutator transaction binding the contract method 0xb9a6a484.
//
// Solidity: function setCapcity(uint256 _newCap) returns()
func (_CappedCashier *CappedCashierSession) SetCapcity(_newCap *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetCapcity(&_CappedCashier.TransactOpts, _newCap)
}

// SetCapcity is a paid mutator transaction binding the contract method 0xb9a6a484.
//
// Solidity: function setCapcity(uint256 _newCap) returns()
func (_CappedCashier *CappedCashierTransactorSession) SetCapcity(_newCap *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetCapcity(&_CappedCashier.TransactOpts, _newCap)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedCashier *CappedCashierTransactor) SetMaxStakePerAddress(opts *bind.TransactOpts, _newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "setMaxStakePerAddress", _newMaxStakePerAddress)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedCashier *CappedCashierSession) SetMaxStakePerAddress(_newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetMaxStakePerAddress(&_CappedCashier.TransactOpts, _newMaxStakePerAddress)
}

// SetMaxStakePerAddress is a paid mutator transaction binding the contract method 0x36a42425.
//
// Solidity: function setMaxStakePerAddress(uint256 _newMaxStakePerAddress) returns()
func (_CappedCashier *CappedCashierTransactorSession) SetMaxStakePerAddress(_newMaxStakePerAddress *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetMaxStakePerAddress(&_CappedCashier.TransactOpts, _newMaxStakePerAddress)
}

// SetMinStakePerTime is a paid mutator transaction binding the contract method 0x2f6f0a6a.
//
// Solidity: function setMinStakePerTime(uint256 _newMinStakePerTime) returns()
func (_CappedCashier *CappedCashierTransactor) SetMinStakePerTime(opts *bind.TransactOpts, _newMinStakePerTime *big.Int) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "setMinStakePerTime", _newMinStakePerTime)
}

// SetMinStakePerTime is a paid mutator transaction binding the contract method 0x2f6f0a6a.
//
// Solidity: function setMinStakePerTime(uint256 _newMinStakePerTime) returns()
func (_CappedCashier *CappedCashierSession) SetMinStakePerTime(_newMinStakePerTime *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetMinStakePerTime(&_CappedCashier.TransactOpts, _newMinStakePerTime)
}

// SetMinStakePerTime is a paid mutator transaction binding the contract method 0x2f6f0a6a.
//
// Solidity: function setMinStakePerTime(uint256 _newMinStakePerTime) returns()
func (_CappedCashier *CappedCashierTransactorSession) SetMinStakePerTime(_newMinStakePerTime *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.SetMinStakePerTime(&_CappedCashier.TransactOpts, _newMinStakePerTime)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CappedCashier *CappedCashierTransactor) Stake(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "stake", _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CappedCashier *CappedCashierSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.Stake(&_CappedCashier.TransactOpts, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 _amount) returns()
func (_CappedCashier *CappedCashierTransactorSession) Stake(_amount *big.Int) (*types.Transaction, error) {
	return _CappedCashier.Contract.Stake(&_CappedCashier.TransactOpts, _amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCashier *CappedCashierTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCashier *CappedCashierSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedCashier.Contract.TransferOwnership(&_CappedCashier.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_CappedCashier *CappedCashierTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _CappedCashier.Contract.TransferOwnership(&_CappedCashier.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCashier *CappedCashierTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCashier *CappedCashierSession) Unpause() (*types.Transaction, error) {
	return _CappedCashier.Contract.Unpause(&_CappedCashier.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_CappedCashier *CappedCashierTransactorSession) Unpause() (*types.Transaction, error) {
	return _CappedCashier.Contract.Unpause(&_CappedCashier.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedCashier *CappedCashierTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedCashier *CappedCashierSession) Withdraw() (*types.Transaction, error) {
	return _CappedCashier.Contract.Withdraw(&_CappedCashier.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CappedCashier *CappedCashierTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CappedCashier.Contract.Withdraw(&_CappedCashier.TransactOpts)
}

// WithdrawRewardPermit is a paid mutator transaction binding the contract method 0x14952f94.
//
// Solidity: function withdrawRewardPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CappedCashier *CappedCashierTransactor) WithdrawRewardPermit(opts *bind.TransactOpts, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CappedCashier.contract.Transact(opts, "withdrawRewardPermit", value, deadline, v, r, s)
}

// WithdrawRewardPermit is a paid mutator transaction binding the contract method 0x14952f94.
//
// Solidity: function withdrawRewardPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CappedCashier *CappedCashierSession) WithdrawRewardPermit(value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CappedCashier.Contract.WithdrawRewardPermit(&_CappedCashier.TransactOpts, value, deadline, v, r, s)
}

// WithdrawRewardPermit is a paid mutator transaction binding the contract method 0x14952f94.
//
// Solidity: function withdrawRewardPermit(uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_CappedCashier *CappedCashierTransactorSession) WithdrawRewardPermit(value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _CappedCashier.Contract.WithdrawRewardPermit(&_CappedCashier.TransactOpts, value, deadline, v, r, s)
}

// CappedCashierCapacityChangedIterator is returned from FilterCapacityChanged and is used to iterate over the raw logs and unpacked data for CapacityChanged events raised by the CappedCashier contract.
type CappedCashierCapacityChangedIterator struct {
	Event *CappedCashierCapacityChanged // Event containing the contract specifics and raw log

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
func (it *CappedCashierCapacityChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierCapacityChanged)
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
		it.Event = new(CappedCashierCapacityChanged)
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
func (it *CappedCashierCapacityChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierCapacityChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierCapacityChanged represents a CapacityChanged event raised by the CappedCashier contract.
type CappedCashierCapacityChanged struct {
	OldCap *big.Int
	NewCap *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCapacityChanged is a free log retrieval operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedCashier *CappedCashierFilterer) FilterCapacityChanged(opts *bind.FilterOpts) (*CappedCashierCapacityChangedIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "CapacityChanged")
	if err != nil {
		return nil, err
	}
	return &CappedCashierCapacityChangedIterator{contract: _CappedCashier.contract, event: "CapacityChanged", logs: logs, sub: sub}, nil
}

// WatchCapacityChanged is a free log subscription operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedCashier *CappedCashierFilterer) WatchCapacityChanged(opts *bind.WatchOpts, sink chan<- *CappedCashierCapacityChanged) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "CapacityChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierCapacityChanged)
				if err := _CappedCashier.contract.UnpackLog(event, "CapacityChanged", log); err != nil {
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

// ParseCapacityChanged is a log parse operation binding the contract event 0xc017b6ded710bda1c186f1a47c7fd57340569e95d6818b718bf85585e2f9e4b7.
//
// Solidity: event CapacityChanged(uint256 oldCap, uint256 newCap)
func (_CappedCashier *CappedCashierFilterer) ParseCapacityChanged(log types.Log) (*CappedCashierCapacityChanged, error) {
	event := new(CappedCashierCapacityChanged)
	if err := _CappedCashier.contract.UnpackLog(event, "CapacityChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierMaxStakePerAddressChangedIterator is returned from FilterMaxStakePerAddressChanged and is used to iterate over the raw logs and unpacked data for MaxStakePerAddressChanged events raised by the CappedCashier contract.
type CappedCashierMaxStakePerAddressChangedIterator struct {
	Event *CappedCashierMaxStakePerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CappedCashierMaxStakePerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierMaxStakePerAddressChanged)
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
		it.Event = new(CappedCashierMaxStakePerAddressChanged)
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
func (it *CappedCashierMaxStakePerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierMaxStakePerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierMaxStakePerAddressChanged represents a MaxStakePerAddressChanged event raised by the CappedCashier contract.
type CappedCashierMaxStakePerAddressChanged struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMaxStakePerAddressChanged is a free log retrieval operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) FilterMaxStakePerAddressChanged(opts *bind.FilterOpts) (*CappedCashierMaxStakePerAddressChangedIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "MaxStakePerAddressChanged")
	if err != nil {
		return nil, err
	}
	return &CappedCashierMaxStakePerAddressChangedIterator{contract: _CappedCashier.contract, event: "MaxStakePerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchMaxStakePerAddressChanged is a free log subscription operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) WatchMaxStakePerAddressChanged(opts *bind.WatchOpts, sink chan<- *CappedCashierMaxStakePerAddressChanged) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "MaxStakePerAddressChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierMaxStakePerAddressChanged)
				if err := _CappedCashier.contract.UnpackLog(event, "MaxStakePerAddressChanged", log); err != nil {
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

// ParseMaxStakePerAddressChanged is a log parse operation binding the contract event 0xcd601ef0c358383f04ece11d4261f0ecdd1333fe0620b2ad6f4624d963d0977f.
//
// Solidity: event MaxStakePerAddressChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) ParseMaxStakePerAddressChanged(log types.Log) (*CappedCashierMaxStakePerAddressChanged, error) {
	event := new(CappedCashierMaxStakePerAddressChanged)
	if err := _CappedCashier.contract.UnpackLog(event, "MaxStakePerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierMinStakePerTimeChangedIterator is returned from FilterMinStakePerTimeChanged and is used to iterate over the raw logs and unpacked data for MinStakePerTimeChanged events raised by the CappedCashier contract.
type CappedCashierMinStakePerTimeChangedIterator struct {
	Event *CappedCashierMinStakePerTimeChanged // Event containing the contract specifics and raw log

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
func (it *CappedCashierMinStakePerTimeChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierMinStakePerTimeChanged)
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
		it.Event = new(CappedCashierMinStakePerTimeChanged)
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
func (it *CappedCashierMinStakePerTimeChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierMinStakePerTimeChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierMinStakePerTimeChanged represents a MinStakePerTimeChanged event raised by the CappedCashier contract.
type CappedCashierMinStakePerTimeChanged struct {
	OldAmount *big.Int
	NewAmount *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinStakePerTimeChanged is a free log retrieval operation binding the contract event 0xf109c4476c1c4414002a0f7e10de46f87055ecdb5b3f55a76a064dc26a882a40.
//
// Solidity: event MinStakePerTimeChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) FilterMinStakePerTimeChanged(opts *bind.FilterOpts) (*CappedCashierMinStakePerTimeChangedIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "MinStakePerTimeChanged")
	if err != nil {
		return nil, err
	}
	return &CappedCashierMinStakePerTimeChangedIterator{contract: _CappedCashier.contract, event: "MinStakePerTimeChanged", logs: logs, sub: sub}, nil
}

// WatchMinStakePerTimeChanged is a free log subscription operation binding the contract event 0xf109c4476c1c4414002a0f7e10de46f87055ecdb5b3f55a76a064dc26a882a40.
//
// Solidity: event MinStakePerTimeChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) WatchMinStakePerTimeChanged(opts *bind.WatchOpts, sink chan<- *CappedCashierMinStakePerTimeChanged) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "MinStakePerTimeChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierMinStakePerTimeChanged)
				if err := _CappedCashier.contract.UnpackLog(event, "MinStakePerTimeChanged", log); err != nil {
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

// ParseMinStakePerTimeChanged is a log parse operation binding the contract event 0xf109c4476c1c4414002a0f7e10de46f87055ecdb5b3f55a76a064dc26a882a40.
//
// Solidity: event MinStakePerTimeChanged(uint256 oldAmount, uint256 newAmount)
func (_CappedCashier *CappedCashierFilterer) ParseMinStakePerTimeChanged(log types.Log) (*CappedCashierMinStakePerTimeChanged, error) {
	event := new(CappedCashierMinStakePerTimeChanged)
	if err := _CappedCashier.contract.UnpackLog(event, "MinStakePerTimeChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the CappedCashier contract.
type CappedCashierOwnershipTransferredIterator struct {
	Event *CappedCashierOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *CappedCashierOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierOwnershipTransferred)
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
		it.Event = new(CappedCashierOwnershipTransferred)
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
func (it *CappedCashierOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierOwnershipTransferred represents a OwnershipTransferred event raised by the CappedCashier contract.
type CappedCashierOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCashier *CappedCashierFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*CappedCashierOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &CappedCashierOwnershipTransferredIterator{contract: _CappedCashier.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCashier *CappedCashierFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *CappedCashierOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierOwnershipTransferred)
				if err := _CappedCashier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_CappedCashier *CappedCashierFilterer) ParseOwnershipTransferred(log types.Log) (*CappedCashierOwnershipTransferred, error) {
	event := new(CappedCashierOwnershipTransferred)
	if err := _CappedCashier.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CappedCashier contract.
type CappedCashierPausedIterator struct {
	Event *CappedCashierPaused // Event containing the contract specifics and raw log

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
func (it *CappedCashierPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierPaused)
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
		it.Event = new(CappedCashierPaused)
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
func (it *CappedCashierPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierPaused represents a Paused event raised by the CappedCashier contract.
type CappedCashierPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCashier *CappedCashierFilterer) FilterPaused(opts *bind.FilterOpts) (*CappedCashierPausedIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CappedCashierPausedIterator{contract: _CappedCashier.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCashier *CappedCashierFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CappedCashierPaused) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierPaused)
				if err := _CappedCashier.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_CappedCashier *CappedCashierFilterer) ParsePaused(log types.Log) (*CappedCashierPaused, error) {
	event := new(CappedCashierPaused)
	if err := _CappedCashier.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierRecoveredIterator is returned from FilterRecovered and is used to iterate over the raw logs and unpacked data for Recovered events raised by the CappedCashier contract.
type CappedCashierRecoveredIterator struct {
	Event *CappedCashierRecovered // Event containing the contract specifics and raw log

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
func (it *CappedCashierRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierRecovered)
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
		it.Event = new(CappedCashierRecovered)
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
func (it *CappedCashierRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierRecovered represents a Recovered event raised by the CappedCashier contract.
type CappedCashierRecovered struct {
	Token  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRecovered is a free log retrieval operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) FilterRecovered(opts *bind.FilterOpts, token []common.Address) (*CappedCashierRecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Recovered", tokenRule)
	if err != nil {
		return nil, err
	}
	return &CappedCashierRecoveredIterator{contract: _CappedCashier.contract, event: "Recovered", logs: logs, sub: sub}, nil
}

// WatchRecovered is a free log subscription operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) WatchRecovered(opts *bind.WatchOpts, sink chan<- *CappedCashierRecovered, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Recovered", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierRecovered)
				if err := _CappedCashier.contract.UnpackLog(event, "Recovered", log); err != nil {
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

// ParseRecovered is a log parse operation binding the contract event 0x8c1256b8896378cd5044f80c202f9772b9d77dc85c8a6eb51967210b09bfaa28.
//
// Solidity: event Recovered(address indexed token, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) ParseRecovered(log types.Log) (*CappedCashierRecovered, error) {
	event := new(CappedCashierRecovered)
	if err := _CappedCashier.contract.UnpackLog(event, "Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the CappedCashier contract.
type CappedCashierRewardIterator struct {
	Event *CappedCashierReward // Event containing the contract specifics and raw log

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
func (it *CappedCashierRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierReward)
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
		it.Event = new(CappedCashierReward)
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
func (it *CappedCashierRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierReward represents a Reward event raised by the CappedCashier contract.
type CappedCashierReward struct {
	Staker       common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x619caafabdd75649b302ba8419e48cccf64f37f1983ac4727cfb38b57703ffc9.
//
// Solidity: event Reward(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) FilterReward(opts *bind.FilterOpts, staker []common.Address) (*CappedCashierRewardIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Reward", stakerRule)
	if err != nil {
		return nil, err
	}
	return &CappedCashierRewardIterator{contract: _CappedCashier.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x619caafabdd75649b302ba8419e48cccf64f37f1983ac4727cfb38b57703ffc9.
//
// Solidity: event Reward(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *CappedCashierReward, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Reward", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierReward)
				if err := _CappedCashier.contract.UnpackLog(event, "Reward", log); err != nil {
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

// ParseReward is a log parse operation binding the contract event 0x619caafabdd75649b302ba8419e48cccf64f37f1983ac4727cfb38b57703ffc9.
//
// Solidity: event Reward(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) ParseReward(log types.Log) (*CappedCashierReward, error) {
	event := new(CappedCashierReward)
	if err := _CappedCashier.contract.UnpackLog(event, "Reward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the CappedCashier contract.
type CappedCashierStakedIterator struct {
	Event *CappedCashierStaked // Event containing the contract specifics and raw log

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
func (it *CappedCashierStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierStaked)
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
		it.Event = new(CappedCashierStaked)
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
func (it *CappedCashierStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierStaked represents a Staked event raised by the CappedCashier contract.
type CappedCashierStaked struct {
	Staker common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address) (*CappedCashierStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return &CappedCashierStakedIterator{contract: _CappedCashier.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *CappedCashierStaked, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Staked", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierStaked)
				if err := _CappedCashier.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x9e71bc8eea02a63969f509818f2dafb9254532904319f9dbda79b67bd34a5f3d.
//
// Solidity: event Staked(address indexed staker, uint256 amount)
func (_CappedCashier *CappedCashierFilterer) ParseStaked(log types.Log) (*CappedCashierStaked, error) {
	event := new(CappedCashierStaked)
	if err := _CappedCashier.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierStakingEnabledIterator is returned from FilterStakingEnabled and is used to iterate over the raw logs and unpacked data for StakingEnabled events raised by the CappedCashier contract.
type CappedCashierStakingEnabledIterator struct {
	Event *CappedCashierStakingEnabled // Event containing the contract specifics and raw log

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
func (it *CappedCashierStakingEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierStakingEnabled)
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
		it.Event = new(CappedCashierStakingEnabled)
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
func (it *CappedCashierStakingEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierStakingEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierStakingEnabled represents a StakingEnabled event raised by the CappedCashier contract.
type CappedCashierStakingEnabled struct {
	StartTime    *big.Int
	StopTime     *big.Int
	RewardAmount *big.Int
	RoundNum     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingEnabled is a free log retrieval operation binding the contract event 0xb5129b4da23f5262ac3e5f0c04ece22914c4498b07fb4ae13a5ac29723c6028f.
//
// Solidity: event StakingEnabled(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) FilterStakingEnabled(opts *bind.FilterOpts) (*CappedCashierStakingEnabledIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "StakingEnabled")
	if err != nil {
		return nil, err
	}
	return &CappedCashierStakingEnabledIterator{contract: _CappedCashier.contract, event: "StakingEnabled", logs: logs, sub: sub}, nil
}

// WatchStakingEnabled is a free log subscription operation binding the contract event 0xb5129b4da23f5262ac3e5f0c04ece22914c4498b07fb4ae13a5ac29723c6028f.
//
// Solidity: event StakingEnabled(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) WatchStakingEnabled(opts *bind.WatchOpts, sink chan<- *CappedCashierStakingEnabled) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "StakingEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierStakingEnabled)
				if err := _CappedCashier.contract.UnpackLog(event, "StakingEnabled", log); err != nil {
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

// ParseStakingEnabled is a log parse operation binding the contract event 0xb5129b4da23f5262ac3e5f0c04ece22914c4498b07fb4ae13a5ac29723c6028f.
//
// Solidity: event StakingEnabled(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) ParseStakingEnabled(log types.Log) (*CappedCashierStakingEnabled, error) {
	event := new(CappedCashierStakingEnabled)
	if err := _CappedCashier.contract.UnpackLog(event, "StakingEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierStakingPreparingIterator is returned from FilterStakingPreparing and is used to iterate over the raw logs and unpacked data for StakingPreparing events raised by the CappedCashier contract.
type CappedCashierStakingPreparingIterator struct {
	Event *CappedCashierStakingPreparing // Event containing the contract specifics and raw log

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
func (it *CappedCashierStakingPreparingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierStakingPreparing)
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
		it.Event = new(CappedCashierStakingPreparing)
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
func (it *CappedCashierStakingPreparingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierStakingPreparingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierStakingPreparing represents a StakingPreparing event raised by the CappedCashier contract.
type CappedCashierStakingPreparing struct {
	StartTime    *big.Int
	StopTime     *big.Int
	RewardAmount *big.Int
	RoundNum     *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterStakingPreparing is a free log retrieval operation binding the contract event 0x2bd1d6574f534f2f8253d40b169e293b43d6bfc8731679569e291bb9f4779cb2.
//
// Solidity: event StakingPreparing(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) FilterStakingPreparing(opts *bind.FilterOpts) (*CappedCashierStakingPreparingIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "StakingPreparing")
	if err != nil {
		return nil, err
	}
	return &CappedCashierStakingPreparingIterator{contract: _CappedCashier.contract, event: "StakingPreparing", logs: logs, sub: sub}, nil
}

// WatchStakingPreparing is a free log subscription operation binding the contract event 0x2bd1d6574f534f2f8253d40b169e293b43d6bfc8731679569e291bb9f4779cb2.
//
// Solidity: event StakingPreparing(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) WatchStakingPreparing(opts *bind.WatchOpts, sink chan<- *CappedCashierStakingPreparing) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "StakingPreparing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierStakingPreparing)
				if err := _CappedCashier.contract.UnpackLog(event, "StakingPreparing", log); err != nil {
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

// ParseStakingPreparing is a log parse operation binding the contract event 0x2bd1d6574f534f2f8253d40b169e293b43d6bfc8731679569e291bb9f4779cb2.
//
// Solidity: event StakingPreparing(uint256 startTime, uint256 stopTime, uint256 rewardAmount, uint256 roundNum)
func (_CappedCashier *CappedCashierFilterer) ParseStakingPreparing(log types.Log) (*CappedCashierStakingPreparing, error) {
	event := new(CappedCashierStakingPreparing)
	if err := _CappedCashier.contract.UnpackLog(event, "StakingPreparing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the CappedCashier contract.
type CappedCashierUnpausedIterator struct {
	Event *CappedCashierUnpaused // Event containing the contract specifics and raw log

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
func (it *CappedCashierUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierUnpaused)
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
		it.Event = new(CappedCashierUnpaused)
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
func (it *CappedCashierUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierUnpaused represents a Unpaused event raised by the CappedCashier contract.
type CappedCashierUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCashier *CappedCashierFilterer) FilterUnpaused(opts *bind.FilterOpts) (*CappedCashierUnpausedIterator, error) {

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &CappedCashierUnpausedIterator{contract: _CappedCashier.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCashier *CappedCashierFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *CappedCashierUnpaused) (event.Subscription, error) {

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierUnpaused)
				if err := _CappedCashier.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_CappedCashier *CappedCashierFilterer) ParseUnpaused(log types.Log) (*CappedCashierUnpaused, error) {
	event := new(CappedCashierUnpaused)
	if err := _CappedCashier.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CappedCashierWithdrawalIterator is returned from FilterWithdrawal and is used to iterate over the raw logs and unpacked data for Withdrawal events raised by the CappedCashier contract.
type CappedCashierWithdrawalIterator struct {
	Event *CappedCashierWithdrawal // Event containing the contract specifics and raw log

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
func (it *CappedCashierWithdrawalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CappedCashierWithdrawal)
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
		it.Event = new(CappedCashierWithdrawal)
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
func (it *CappedCashierWithdrawalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CappedCashierWithdrawalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CappedCashierWithdrawal represents a Withdrawal event raised by the CappedCashier contract.
type CappedCashierWithdrawal struct {
	Staker       common.Address
	RewardAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterWithdrawal is a free log retrieval operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) FilterWithdrawal(opts *bind.FilterOpts, staker []common.Address) (*CappedCashierWithdrawalIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.FilterLogs(opts, "Withdrawal", stakerRule)
	if err != nil {
		return nil, err
	}
	return &CappedCashierWithdrawalIterator{contract: _CappedCashier.contract, event: "Withdrawal", logs: logs, sub: sub}, nil
}

// WatchWithdrawal is a free log subscription operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) WatchWithdrawal(opts *bind.WatchOpts, sink chan<- *CappedCashierWithdrawal, staker []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _CappedCashier.contract.WatchLogs(opts, "Withdrawal", stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CappedCashierWithdrawal)
				if err := _CappedCashier.contract.UnpackLog(event, "Withdrawal", log); err != nil {
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

// ParseWithdrawal is a log parse operation binding the contract event 0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65.
//
// Solidity: event Withdrawal(address indexed staker, uint256 rewardAmount)
func (_CappedCashier *CappedCashierFilterer) ParseWithdrawal(log types.Log) (*CappedCashierWithdrawal, error) {
	event := new(CappedCashierWithdrawal)
	if err := _CappedCashier.contract.UnpackLog(event, "Withdrawal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
