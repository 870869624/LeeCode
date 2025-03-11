// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package staking

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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"_stakingToken\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"EnforcedPause\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExpectedPause\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"OwnableInvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"OwnableUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"completedPeriod\",\"type\":\"uint256\"}],\"name\":\"PeriodSwitchCompleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processedUpToSpaceId\",\"type\":\"uint256\"}],\"name\":\"PeriodSwitchContinued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newPeriod\",\"type\":\"uint256\"}],\"name\":\"PeriodSwitchStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"SpaceCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"SpaceDestroyed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"SpaceRateApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"SpaceRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pendingAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockedAmount\",\"type\":\"uint256\"}],\"name\":\"StakerStatusProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"spaceId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"status\",\"type\":\"bool\"}],\"name\":\"WhitelistUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"addToWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"}],\"name\":\"continueSingleSwitchPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"continueSwitchPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_rate\",\"type\":\"uint256\"}],\"name\":\"createSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"findSpaceByOwner\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"}],\"name\":\"getTotalStakeInfo\",\"outputs\":[{\"internalType\":\"uint256[4]\",\"name\":\"\",\"type\":\"uint256[4]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_staker\",\"type\":\"address\"}],\"name\":\"getUserStakeInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"userPending\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userLocked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userUnlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"userWithdraw\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSwitching\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextSpaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"ownerToSpaceIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pendingRateUpdates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"Period\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"removeFromWhitelist\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"spaceCreated\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"spaces\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStaked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"stakingToken\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"startSwitchPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchingStartSpaceId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"}],\"name\":\"toDestroy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_newRate\",\"type\":\"uint256\"}],\"name\":\"updateSpaceRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userLockedAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"enumStakingContract.StakeStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"userStakeAmounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"userStakes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"isUnstaking\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"whitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_spaceId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Staking *StakingCaller) CurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "currentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Staking *StakingSession) CurrentPeriod() (*big.Int, error) {
	return _Staking.Contract.CurrentPeriod(&_Staking.CallOpts)
}

// CurrentPeriod is a free data retrieval call binding the contract method 0x06040618.
//
// Solidity: function currentPeriod() view returns(uint256)
func (_Staking *StakingCallerSession) CurrentPeriod() (*big.Int, error) {
	return _Staking.Contract.CurrentPeriod(&_Staking.CallOpts)
}

// FindSpaceByOwner is a free data retrieval call binding the contract method 0xf895c925.
//
// Solidity: function findSpaceByOwner(address _owner) view returns(uint256[])
func (_Staking *StakingCaller) FindSpaceByOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "findSpaceByOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// FindSpaceByOwner is a free data retrieval call binding the contract method 0xf895c925.
//
// Solidity: function findSpaceByOwner(address _owner) view returns(uint256[])
func (_Staking *StakingSession) FindSpaceByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Staking.Contract.FindSpaceByOwner(&_Staking.CallOpts, _owner)
}

// FindSpaceByOwner is a free data retrieval call binding the contract method 0xf895c925.
//
// Solidity: function findSpaceByOwner(address _owner) view returns(uint256[])
func (_Staking *StakingCallerSession) FindSpaceByOwner(_owner common.Address) ([]*big.Int, error) {
	return _Staking.Contract.FindSpaceByOwner(&_Staking.CallOpts, _owner)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_Staking *StakingCaller) GetCurrentPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getCurrentPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_Staking *StakingSession) GetCurrentPeriod() (*big.Int, error) {
	return _Staking.Contract.GetCurrentPeriod(&_Staking.CallOpts)
}

// GetCurrentPeriod is a free data retrieval call binding the contract method 0x086146d2.
//
// Solidity: function getCurrentPeriod() view returns(uint256)
func (_Staking *StakingCallerSession) GetCurrentPeriod() (*big.Int, error) {
	return _Staking.Contract.GetCurrentPeriod(&_Staking.CallOpts)
}

// GetStakers is a free data retrieval call binding the contract method 0x72366d98.
//
// Solidity: function getStakers(uint256 _spaceId) view returns(address[])
func (_Staking *StakingCaller) GetStakers(opts *bind.CallOpts, _spaceId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getStakers", _spaceId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0x72366d98.
//
// Solidity: function getStakers(uint256 _spaceId) view returns(address[])
func (_Staking *StakingSession) GetStakers(_spaceId *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetStakers(&_Staking.CallOpts, _spaceId)
}

// GetStakers is a free data retrieval call binding the contract method 0x72366d98.
//
// Solidity: function getStakers(uint256 _spaceId) view returns(address[])
func (_Staking *StakingCallerSession) GetStakers(_spaceId *big.Int) ([]common.Address, error) {
	return _Staking.Contract.GetStakers(&_Staking.CallOpts, _spaceId)
}

// GetTotalStakeInfo is a free data retrieval call binding the contract method 0xbbc3c128.
//
// Solidity: function getTotalStakeInfo(uint256 _spaceId) view returns(uint256[4])
func (_Staking *StakingCaller) GetTotalStakeInfo(opts *bind.CallOpts, _spaceId *big.Int) ([4]*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getTotalStakeInfo", _spaceId)

	if err != nil {
		return *new([4]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([4]*big.Int)).(*[4]*big.Int)

	return out0, err

}

// GetTotalStakeInfo is a free data retrieval call binding the contract method 0xbbc3c128.
//
// Solidity: function getTotalStakeInfo(uint256 _spaceId) view returns(uint256[4])
func (_Staking *StakingSession) GetTotalStakeInfo(_spaceId *big.Int) ([4]*big.Int, error) {
	return _Staking.Contract.GetTotalStakeInfo(&_Staking.CallOpts, _spaceId)
}

// GetTotalStakeInfo is a free data retrieval call binding the contract method 0xbbc3c128.
//
// Solidity: function getTotalStakeInfo(uint256 _spaceId) view returns(uint256[4])
func (_Staking *StakingCallerSession) GetTotalStakeInfo(_spaceId *big.Int) ([4]*big.Int, error) {
	return _Staking.Contract.GetTotalStakeInfo(&_Staking.CallOpts, _spaceId)
}

// GetUserStakeInfo is a free data retrieval call binding the contract method 0x3300a2d3.
//
// Solidity: function getUserStakeInfo(uint256 _spaceId, address _staker) view returns(uint256 userPending, uint256 userLocked, uint256 userUnlock, uint256 userWithdraw)
func (_Staking *StakingCaller) GetUserStakeInfo(opts *bind.CallOpts, _spaceId *big.Int, _staker common.Address) (struct {
	UserPending  *big.Int
	UserLocked   *big.Int
	UserUnlock   *big.Int
	UserWithdraw *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getUserStakeInfo", _spaceId, _staker)

	outstruct := new(struct {
		UserPending  *big.Int
		UserLocked   *big.Int
		UserUnlock   *big.Int
		UserWithdraw *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.UserPending = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.UserLocked = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UserUnlock = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UserWithdraw = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserStakeInfo is a free data retrieval call binding the contract method 0x3300a2d3.
//
// Solidity: function getUserStakeInfo(uint256 _spaceId, address _staker) view returns(uint256 userPending, uint256 userLocked, uint256 userUnlock, uint256 userWithdraw)
func (_Staking *StakingSession) GetUserStakeInfo(_spaceId *big.Int, _staker common.Address) (struct {
	UserPending  *big.Int
	UserLocked   *big.Int
	UserUnlock   *big.Int
	UserWithdraw *big.Int
}, error) {
	return _Staking.Contract.GetUserStakeInfo(&_Staking.CallOpts, _spaceId, _staker)
}

// GetUserStakeInfo is a free data retrieval call binding the contract method 0x3300a2d3.
//
// Solidity: function getUserStakeInfo(uint256 _spaceId, address _staker) view returns(uint256 userPending, uint256 userLocked, uint256 userUnlock, uint256 userWithdraw)
func (_Staking *StakingCallerSession) GetUserStakeInfo(_spaceId *big.Int, _staker common.Address) (struct {
	UserPending  *big.Int
	UserLocked   *big.Int
	UserUnlock   *big.Int
	UserWithdraw *big.Int
}, error) {
	return _Staking.Contract.GetUserStakeInfo(&_Staking.CallOpts, _spaceId, _staker)
}

// IsSwitching is a free data retrieval call binding the contract method 0x8194fd89.
//
// Solidity: function isSwitching() view returns(bool)
func (_Staking *StakingCaller) IsSwitching(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isSwitching")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSwitching is a free data retrieval call binding the contract method 0x8194fd89.
//
// Solidity: function isSwitching() view returns(bool)
func (_Staking *StakingSession) IsSwitching() (bool, error) {
	return _Staking.Contract.IsSwitching(&_Staking.CallOpts)
}

// IsSwitching is a free data retrieval call binding the contract method 0x8194fd89.
//
// Solidity: function isSwitching() view returns(bool)
func (_Staking *StakingCallerSession) IsSwitching() (bool, error) {
	return _Staking.Contract.IsSwitching(&_Staking.CallOpts)
}

// NextSpaceId is a free data retrieval call binding the contract method 0xe565e661.
//
// Solidity: function nextSpaceId() view returns(uint256)
func (_Staking *StakingCaller) NextSpaceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "nextSpaceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextSpaceId is a free data retrieval call binding the contract method 0xe565e661.
//
// Solidity: function nextSpaceId() view returns(uint256)
func (_Staking *StakingSession) NextSpaceId() (*big.Int, error) {
	return _Staking.Contract.NextSpaceId(&_Staking.CallOpts)
}

// NextSpaceId is a free data retrieval call binding the contract method 0xe565e661.
//
// Solidity: function nextSpaceId() view returns(uint256)
func (_Staking *StakingCallerSession) NextSpaceId() (*big.Int, error) {
	return _Staking.Contract.NextSpaceId(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// OwnerToSpaceIds is a free data retrieval call binding the contract method 0x99c09f07.
//
// Solidity: function ownerToSpaceIds(address , uint256 ) view returns(uint256)
func (_Staking *StakingCaller) OwnerToSpaceIds(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "ownerToSpaceIds", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerToSpaceIds is a free data retrieval call binding the contract method 0x99c09f07.
//
// Solidity: function ownerToSpaceIds(address , uint256 ) view returns(uint256)
func (_Staking *StakingSession) OwnerToSpaceIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Staking.Contract.OwnerToSpaceIds(&_Staking.CallOpts, arg0, arg1)
}

// OwnerToSpaceIds is a free data retrieval call binding the contract method 0x99c09f07.
//
// Solidity: function ownerToSpaceIds(address , uint256 ) view returns(uint256)
func (_Staking *StakingCallerSession) OwnerToSpaceIds(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Staking.Contract.OwnerToSpaceIds(&_Staking.CallOpts, arg0, arg1)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Staking *StakingCallerSession) Paused() (bool, error) {
	return _Staking.Contract.Paused(&_Staking.CallOpts)
}

// PendingRateUpdates is a free data retrieval call binding the contract method 0xa0966217.
//
// Solidity: function pendingRateUpdates(uint256 ) view returns(uint256 rate, uint256 Period)
func (_Staking *StakingCaller) PendingRateUpdates(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Rate   *big.Int
	Period *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "pendingRateUpdates", arg0)

	outstruct := new(struct {
		Rate   *big.Int
		Period *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Rate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Period = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PendingRateUpdates is a free data retrieval call binding the contract method 0xa0966217.
//
// Solidity: function pendingRateUpdates(uint256 ) view returns(uint256 rate, uint256 Period)
func (_Staking *StakingSession) PendingRateUpdates(arg0 *big.Int) (struct {
	Rate   *big.Int
	Period *big.Int
}, error) {
	return _Staking.Contract.PendingRateUpdates(&_Staking.CallOpts, arg0)
}

// PendingRateUpdates is a free data retrieval call binding the contract method 0xa0966217.
//
// Solidity: function pendingRateUpdates(uint256 ) view returns(uint256 rate, uint256 Period)
func (_Staking *StakingCallerSession) PendingRateUpdates(arg0 *big.Int) (struct {
	Rate   *big.Int
	Period *big.Int
}, error) {
	return _Staking.Contract.PendingRateUpdates(&_Staking.CallOpts, arg0)
}

// SpaceCreated is a free data retrieval call binding the contract method 0x0ceca1f7.
//
// Solidity: function spaceCreated(address ) view returns(bool)
func (_Staking *StakingCaller) SpaceCreated(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "spaceCreated", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SpaceCreated is a free data retrieval call binding the contract method 0x0ceca1f7.
//
// Solidity: function spaceCreated(address ) view returns(bool)
func (_Staking *StakingSession) SpaceCreated(arg0 common.Address) (bool, error) {
	return _Staking.Contract.SpaceCreated(&_Staking.CallOpts, arg0)
}

// SpaceCreated is a free data retrieval call binding the contract method 0x0ceca1f7.
//
// Solidity: function spaceCreated(address ) view returns(bool)
func (_Staking *StakingCallerSession) SpaceCreated(arg0 common.Address) (bool, error) {
	return _Staking.Contract.SpaceCreated(&_Staking.CallOpts, arg0)
}

// Spaces is a free data retrieval call binding the contract method 0x7586353b.
//
// Solidity: function spaces(uint256 ) view returns(address owner, uint256 rate, uint256 totalStaked, uint256 status)
func (_Staking *StakingCaller) Spaces(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Owner       common.Address
	Rate        *big.Int
	TotalStaked *big.Int
	Status      *big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "spaces", arg0)

	outstruct := new(struct {
		Owner       common.Address
		Rate        *big.Int
		TotalStaked *big.Int
		Status      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Rate = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalStaked = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Spaces is a free data retrieval call binding the contract method 0x7586353b.
//
// Solidity: function spaces(uint256 ) view returns(address owner, uint256 rate, uint256 totalStaked, uint256 status)
func (_Staking *StakingSession) Spaces(arg0 *big.Int) (struct {
	Owner       common.Address
	Rate        *big.Int
	TotalStaked *big.Int
	Status      *big.Int
}, error) {
	return _Staking.Contract.Spaces(&_Staking.CallOpts, arg0)
}

// Spaces is a free data retrieval call binding the contract method 0x7586353b.
//
// Solidity: function spaces(uint256 ) view returns(address owner, uint256 rate, uint256 totalStaked, uint256 status)
func (_Staking *StakingCallerSession) Spaces(arg0 *big.Int) (struct {
	Owner       common.Address
	Rate        *big.Int
	TotalStaked *big.Int
	Status      *big.Int
}, error) {
	return _Staking.Contract.Spaces(&_Staking.CallOpts, arg0)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingCaller) StakingToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "stakingToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingSession) StakingToken() (common.Address, error) {
	return _Staking.Contract.StakingToken(&_Staking.CallOpts)
}

// StakingToken is a free data retrieval call binding the contract method 0x72f702f3.
//
// Solidity: function stakingToken() view returns(address)
func (_Staking *StakingCallerSession) StakingToken() (common.Address, error) {
	return _Staking.Contract.StakingToken(&_Staking.CallOpts)
}

// SwitchingStartSpaceId is a free data retrieval call binding the contract method 0x1de22364.
//
// Solidity: function switchingStartSpaceId() view returns(uint256)
func (_Staking *StakingCaller) SwitchingStartSpaceId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "switchingStartSpaceId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwitchingStartSpaceId is a free data retrieval call binding the contract method 0x1de22364.
//
// Solidity: function switchingStartSpaceId() view returns(uint256)
func (_Staking *StakingSession) SwitchingStartSpaceId() (*big.Int, error) {
	return _Staking.Contract.SwitchingStartSpaceId(&_Staking.CallOpts)
}

// SwitchingStartSpaceId is a free data retrieval call binding the contract method 0x1de22364.
//
// Solidity: function switchingStartSpaceId() view returns(uint256)
func (_Staking *StakingCallerSession) SwitchingStartSpaceId() (*big.Int, error) {
	return _Staking.Contract.SwitchingStartSpaceId(&_Staking.CallOpts)
}

// UserLockedAmounts is a free data retrieval call binding the contract method 0xa4d5f97e.
//
// Solidity: function userLockedAmounts(uint256 , uint256 , address ) view returns(uint256)
func (_Staking *StakingCaller) UserLockedAmounts(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userLockedAmounts", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserLockedAmounts is a free data retrieval call binding the contract method 0xa4d5f97e.
//
// Solidity: function userLockedAmounts(uint256 , uint256 , address ) view returns(uint256)
func (_Staking *StakingSession) UserLockedAmounts(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserLockedAmounts(&_Staking.CallOpts, arg0, arg1, arg2)
}

// UserLockedAmounts is a free data retrieval call binding the contract method 0xa4d5f97e.
//
// Solidity: function userLockedAmounts(uint256 , uint256 , address ) view returns(uint256)
func (_Staking *StakingCallerSession) UserLockedAmounts(arg0 *big.Int, arg1 *big.Int, arg2 common.Address) (*big.Int, error) {
	return _Staking.Contract.UserLockedAmounts(&_Staking.CallOpts, arg0, arg1, arg2)
}

// UserStakeAmounts is a free data retrieval call binding the contract method 0xef842c57.
//
// Solidity: function userStakeAmounts(uint256 , address , uint8 ) view returns(uint256)
func (_Staking *StakingCaller) UserStakeAmounts(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address, arg2 uint8) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStakeAmounts", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UserStakeAmounts is a free data retrieval call binding the contract method 0xef842c57.
//
// Solidity: function userStakeAmounts(uint256 , address , uint8 ) view returns(uint256)
func (_Staking *StakingSession) UserStakeAmounts(arg0 *big.Int, arg1 common.Address, arg2 uint8) (*big.Int, error) {
	return _Staking.Contract.UserStakeAmounts(&_Staking.CallOpts, arg0, arg1, arg2)
}

// UserStakeAmounts is a free data retrieval call binding the contract method 0xef842c57.
//
// Solidity: function userStakeAmounts(uint256 , address , uint8 ) view returns(uint256)
func (_Staking *StakingCallerSession) UserStakeAmounts(arg0 *big.Int, arg1 common.Address, arg2 uint8) (*big.Int, error) {
	return _Staking.Contract.UserStakeAmounts(&_Staking.CallOpts, arg0, arg1, arg2)
}

// UserStakes is a free data retrieval call binding the contract method 0xe8990524.
//
// Solidity: function userStakes(uint256 , address ) view returns(uint256 amount, uint256 timestamp, uint8 isUnstaking)
func (_Staking *StakingCaller) UserStakes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Amount      *big.Int
	Timestamp   *big.Int
	IsUnstaking uint8
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "userStakes", arg0, arg1)

	outstruct := new(struct {
		Amount      *big.Int
		Timestamp   *big.Int
		IsUnstaking uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Amount = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IsUnstaking = *abi.ConvertType(out[2], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserStakes is a free data retrieval call binding the contract method 0xe8990524.
//
// Solidity: function userStakes(uint256 , address ) view returns(uint256 amount, uint256 timestamp, uint8 isUnstaking)
func (_Staking *StakingSession) UserStakes(arg0 *big.Int, arg1 common.Address) (struct {
	Amount      *big.Int
	Timestamp   *big.Int
	IsUnstaking uint8
}, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0, arg1)
}

// UserStakes is a free data retrieval call binding the contract method 0xe8990524.
//
// Solidity: function userStakes(uint256 , address ) view returns(uint256 amount, uint256 timestamp, uint8 isUnstaking)
func (_Staking *StakingCallerSession) UserStakes(arg0 *big.Int, arg1 common.Address) (struct {
	Amount      *big.Int
	Timestamp   *big.Int
	IsUnstaking uint8
}, error) {
	return _Staking.Contract.UserStakes(&_Staking.CallOpts, arg0, arg1)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCaller) Whitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "whitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// Whitelist is a free data retrieval call binding the contract method 0x9b19251a.
//
// Solidity: function whitelist(address ) view returns(bool)
func (_Staking *StakingCallerSession) Whitelist(arg0 common.Address) (bool, error) {
	return _Staking.Contract.Whitelist(&_Staking.CallOpts, arg0)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _user) returns()
func (_Staking *StakingTransactor) AddToWhitelist(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "addToWhitelist", _user)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _user) returns()
func (_Staking *StakingSession) AddToWhitelist(_user common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddToWhitelist(&_Staking.TransactOpts, _user)
}

// AddToWhitelist is a paid mutator transaction binding the contract method 0xe43252d7.
//
// Solidity: function addToWhitelist(address _user) returns()
func (_Staking *StakingTransactorSession) AddToWhitelist(_user common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AddToWhitelist(&_Staking.TransactOpts, _user)
}

// ContinueSingleSwitchPeriod is a paid mutator transaction binding the contract method 0x9fff9d24.
//
// Solidity: function continueSingleSwitchPeriod(uint256 _spaceId) returns()
func (_Staking *StakingTransactor) ContinueSingleSwitchPeriod(opts *bind.TransactOpts, _spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "continueSingleSwitchPeriod", _spaceId)
}

// ContinueSingleSwitchPeriod is a paid mutator transaction binding the contract method 0x9fff9d24.
//
// Solidity: function continueSingleSwitchPeriod(uint256 _spaceId) returns()
func (_Staking *StakingSession) ContinueSingleSwitchPeriod(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ContinueSingleSwitchPeriod(&_Staking.TransactOpts, _spaceId)
}

// ContinueSingleSwitchPeriod is a paid mutator transaction binding the contract method 0x9fff9d24.
//
// Solidity: function continueSingleSwitchPeriod(uint256 _spaceId) returns()
func (_Staking *StakingTransactorSession) ContinueSingleSwitchPeriod(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ContinueSingleSwitchPeriod(&_Staking.TransactOpts, _spaceId)
}

// ContinueSwitchPeriod is a paid mutator transaction binding the contract method 0xa2296a66.
//
// Solidity: function continueSwitchPeriod() returns()
func (_Staking *StakingTransactor) ContinueSwitchPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "continueSwitchPeriod")
}

// ContinueSwitchPeriod is a paid mutator transaction binding the contract method 0xa2296a66.
//
// Solidity: function continueSwitchPeriod() returns()
func (_Staking *StakingSession) ContinueSwitchPeriod() (*types.Transaction, error) {
	return _Staking.Contract.ContinueSwitchPeriod(&_Staking.TransactOpts)
}

// ContinueSwitchPeriod is a paid mutator transaction binding the contract method 0xa2296a66.
//
// Solidity: function continueSwitchPeriod() returns()
func (_Staking *StakingTransactorSession) ContinueSwitchPeriod() (*types.Transaction, error) {
	return _Staking.Contract.ContinueSwitchPeriod(&_Staking.TransactOpts)
}

// CreateSpace is a paid mutator transaction binding the contract method 0x8a8f5f05.
//
// Solidity: function createSpace(uint256 _rate) returns()
func (_Staking *StakingTransactor) CreateSpace(opts *bind.TransactOpts, _rate *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createSpace", _rate)
}

// CreateSpace is a paid mutator transaction binding the contract method 0x8a8f5f05.
//
// Solidity: function createSpace(uint256 _rate) returns()
func (_Staking *StakingSession) CreateSpace(_rate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateSpace(&_Staking.TransactOpts, _rate)
}

// CreateSpace is a paid mutator transaction binding the contract method 0x8a8f5f05.
//
// Solidity: function createSpace(uint256 _rate) returns()
func (_Staking *StakingTransactorSession) CreateSpace(_rate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.CreateSpace(&_Staking.TransactOpts, _rate)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Staking *StakingTransactorSession) Pause() (*types.Transaction, error) {
	return _Staking.Contract.Pause(&_Staking.TransactOpts)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _user) returns()
func (_Staking *StakingTransactor) RemoveFromWhitelist(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "removeFromWhitelist", _user)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _user) returns()
func (_Staking *StakingSession) RemoveFromWhitelist(_user common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveFromWhitelist(&_Staking.TransactOpts, _user)
}

// RemoveFromWhitelist is a paid mutator transaction binding the contract method 0x8ab1d681.
//
// Solidity: function removeFromWhitelist(address _user) returns()
func (_Staking *StakingTransactorSession) RemoveFromWhitelist(_user common.Address) (*types.Transaction, error) {
	return _Staking.Contract.RemoveFromWhitelist(&_Staking.TransactOpts, _user)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, _spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", _spaceId, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingSession) Stake(_spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _spaceId, _amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7b0472f0.
//
// Solidity: function stake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingTransactorSession) Stake(_spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, _spaceId, _amount)
}

// StartSwitchPeriod is a paid mutator transaction binding the contract method 0x0b73612d.
//
// Solidity: function startSwitchPeriod() returns()
func (_Staking *StakingTransactor) StartSwitchPeriod(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "startSwitchPeriod")
}

// StartSwitchPeriod is a paid mutator transaction binding the contract method 0x0b73612d.
//
// Solidity: function startSwitchPeriod() returns()
func (_Staking *StakingSession) StartSwitchPeriod() (*types.Transaction, error) {
	return _Staking.Contract.StartSwitchPeriod(&_Staking.TransactOpts)
}

// StartSwitchPeriod is a paid mutator transaction binding the contract method 0x0b73612d.
//
// Solidity: function startSwitchPeriod() returns()
func (_Staking *StakingTransactorSession) StartSwitchPeriod() (*types.Transaction, error) {
	return _Staking.Contract.StartSwitchPeriod(&_Staking.TransactOpts)
}

// ToDestroy is a paid mutator transaction binding the contract method 0x284268b9.
//
// Solidity: function toDestroy(uint256 _spaceId) returns()
func (_Staking *StakingTransactor) ToDestroy(opts *bind.TransactOpts, _spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "toDestroy", _spaceId)
}

// ToDestroy is a paid mutator transaction binding the contract method 0x284268b9.
//
// Solidity: function toDestroy(uint256 _spaceId) returns()
func (_Staking *StakingSession) ToDestroy(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ToDestroy(&_Staking.TransactOpts, _spaceId)
}

// ToDestroy is a paid mutator transaction binding the contract method 0x284268b9.
//
// Solidity: function toDestroy(uint256 _spaceId) returns()
func (_Staking *StakingTransactorSession) ToDestroy(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.ToDestroy(&_Staking.TransactOpts, _spaceId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Staking *StakingTransactorSession) Unpause() (*types.Transaction, error) {
	return _Staking.Contract.Unpause(&_Staking.TransactOpts)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingTransactor) Unstake(opts *bind.TransactOpts, _spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "unstake", _spaceId, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingSession) Unstake(_spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, _spaceId, _amount)
}

// Unstake is a paid mutator transaction binding the contract method 0x9e2c8a5b.
//
// Solidity: function unstake(uint256 _spaceId, uint256 _amount) returns()
func (_Staking *StakingTransactorSession) Unstake(_spaceId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Unstake(&_Staking.TransactOpts, _spaceId, _amount)
}

// UpdateSpaceRate is a paid mutator transaction binding the contract method 0xc7f4ccd2.
//
// Solidity: function updateSpaceRate(uint256 _spaceId, uint256 _newRate) returns()
func (_Staking *StakingTransactor) UpdateSpaceRate(opts *bind.TransactOpts, _spaceId *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "updateSpaceRate", _spaceId, _newRate)
}

// UpdateSpaceRate is a paid mutator transaction binding the contract method 0xc7f4ccd2.
//
// Solidity: function updateSpaceRate(uint256 _spaceId, uint256 _newRate) returns()
func (_Staking *StakingSession) UpdateSpaceRate(_spaceId *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateSpaceRate(&_Staking.TransactOpts, _spaceId, _newRate)
}

// UpdateSpaceRate is a paid mutator transaction binding the contract method 0xc7f4ccd2.
//
// Solidity: function updateSpaceRate(uint256 _spaceId, uint256 _newRate) returns()
func (_Staking *StakingTransactorSession) UpdateSpaceRate(_spaceId *big.Int, _newRate *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.UpdateSpaceRate(&_Staking.TransactOpts, _spaceId, _newRate)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _spaceId) returns()
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, _spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", _spaceId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _spaceId) returns()
func (_Staking *StakingSession) Withdraw(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, _spaceId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _spaceId) returns()
func (_Staking *StakingTransactorSession) Withdraw(_spaceId *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, _spaceId)
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Staking contract.
type StakingPausedIterator struct {
	Event *StakingPaused // Event containing the contract specifics and raw log

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
func (it *StakingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPaused)
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
		it.Event = new(StakingPaused)
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
func (it *StakingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPaused represents a Paused event raised by the Staking contract.
type StakingPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) FilterPaused(opts *bind.FilterOpts) (*StakingPausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StakingPausedIterator{contract: _Staking.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Staking *StakingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StakingPaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPaused)
				if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Staking *StakingFilterer) ParsePaused(log types.Log) (*StakingPaused, error) {
	event := new(StakingPaused)
	if err := _Staking.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPeriodSwitchCompletedIterator is returned from FilterPeriodSwitchCompleted and is used to iterate over the raw logs and unpacked data for PeriodSwitchCompleted events raised by the Staking contract.
type StakingPeriodSwitchCompletedIterator struct {
	Event *StakingPeriodSwitchCompleted // Event containing the contract specifics and raw log

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
func (it *StakingPeriodSwitchCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPeriodSwitchCompleted)
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
		it.Event = new(StakingPeriodSwitchCompleted)
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
func (it *StakingPeriodSwitchCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPeriodSwitchCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPeriodSwitchCompleted represents a PeriodSwitchCompleted event raised by the Staking contract.
type StakingPeriodSwitchCompleted struct {
	CompletedPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterPeriodSwitchCompleted is a free log retrieval operation binding the contract event 0x5b9a2abda3d7e4d8f67994cc3b3b9a930d2d7c0bf0b6719ef557457a8d7e6bec.
//
// Solidity: event PeriodSwitchCompleted(uint256 completedPeriod)
func (_Staking *StakingFilterer) FilterPeriodSwitchCompleted(opts *bind.FilterOpts) (*StakingPeriodSwitchCompletedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PeriodSwitchCompleted")
	if err != nil {
		return nil, err
	}
	return &StakingPeriodSwitchCompletedIterator{contract: _Staking.contract, event: "PeriodSwitchCompleted", logs: logs, sub: sub}, nil
}

// WatchPeriodSwitchCompleted is a free log subscription operation binding the contract event 0x5b9a2abda3d7e4d8f67994cc3b3b9a930d2d7c0bf0b6719ef557457a8d7e6bec.
//
// Solidity: event PeriodSwitchCompleted(uint256 completedPeriod)
func (_Staking *StakingFilterer) WatchPeriodSwitchCompleted(opts *bind.WatchOpts, sink chan<- *StakingPeriodSwitchCompleted) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PeriodSwitchCompleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPeriodSwitchCompleted)
				if err := _Staking.contract.UnpackLog(event, "PeriodSwitchCompleted", log); err != nil {
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

// ParsePeriodSwitchCompleted is a log parse operation binding the contract event 0x5b9a2abda3d7e4d8f67994cc3b3b9a930d2d7c0bf0b6719ef557457a8d7e6bec.
//
// Solidity: event PeriodSwitchCompleted(uint256 completedPeriod)
func (_Staking *StakingFilterer) ParsePeriodSwitchCompleted(log types.Log) (*StakingPeriodSwitchCompleted, error) {
	event := new(StakingPeriodSwitchCompleted)
	if err := _Staking.contract.UnpackLog(event, "PeriodSwitchCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPeriodSwitchContinuedIterator is returned from FilterPeriodSwitchContinued and is used to iterate over the raw logs and unpacked data for PeriodSwitchContinued events raised by the Staking contract.
type StakingPeriodSwitchContinuedIterator struct {
	Event *StakingPeriodSwitchContinued // Event containing the contract specifics and raw log

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
func (it *StakingPeriodSwitchContinuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPeriodSwitchContinued)
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
		it.Event = new(StakingPeriodSwitchContinued)
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
func (it *StakingPeriodSwitchContinuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPeriodSwitchContinuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPeriodSwitchContinued represents a PeriodSwitchContinued event raised by the Staking contract.
type StakingPeriodSwitchContinued struct {
	ProcessedUpToSpaceId *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterPeriodSwitchContinued is a free log retrieval operation binding the contract event 0x9238c092a52f5f32af190f40d8d30f18c1ef44483baa3ebe31eb2121286f4633.
//
// Solidity: event PeriodSwitchContinued(uint256 processedUpToSpaceId)
func (_Staking *StakingFilterer) FilterPeriodSwitchContinued(opts *bind.FilterOpts) (*StakingPeriodSwitchContinuedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PeriodSwitchContinued")
	if err != nil {
		return nil, err
	}
	return &StakingPeriodSwitchContinuedIterator{contract: _Staking.contract, event: "PeriodSwitchContinued", logs: logs, sub: sub}, nil
}

// WatchPeriodSwitchContinued is a free log subscription operation binding the contract event 0x9238c092a52f5f32af190f40d8d30f18c1ef44483baa3ebe31eb2121286f4633.
//
// Solidity: event PeriodSwitchContinued(uint256 processedUpToSpaceId)
func (_Staking *StakingFilterer) WatchPeriodSwitchContinued(opts *bind.WatchOpts, sink chan<- *StakingPeriodSwitchContinued) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PeriodSwitchContinued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPeriodSwitchContinued)
				if err := _Staking.contract.UnpackLog(event, "PeriodSwitchContinued", log); err != nil {
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

// ParsePeriodSwitchContinued is a log parse operation binding the contract event 0x9238c092a52f5f32af190f40d8d30f18c1ef44483baa3ebe31eb2121286f4633.
//
// Solidity: event PeriodSwitchContinued(uint256 processedUpToSpaceId)
func (_Staking *StakingFilterer) ParsePeriodSwitchContinued(log types.Log) (*StakingPeriodSwitchContinued, error) {
	event := new(StakingPeriodSwitchContinued)
	if err := _Staking.contract.UnpackLog(event, "PeriodSwitchContinued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingPeriodSwitchStartedIterator is returned from FilterPeriodSwitchStarted and is used to iterate over the raw logs and unpacked data for PeriodSwitchStarted events raised by the Staking contract.
type StakingPeriodSwitchStartedIterator struct {
	Event *StakingPeriodSwitchStarted // Event containing the contract specifics and raw log

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
func (it *StakingPeriodSwitchStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingPeriodSwitchStarted)
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
		it.Event = new(StakingPeriodSwitchStarted)
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
func (it *StakingPeriodSwitchStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingPeriodSwitchStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingPeriodSwitchStarted represents a PeriodSwitchStarted event raised by the Staking contract.
type StakingPeriodSwitchStarted struct {
	NewPeriod *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPeriodSwitchStarted is a free log retrieval operation binding the contract event 0xf4407b200bd8af965ee06a960f67118073146e5686605ddf1c487f08693530e1.
//
// Solidity: event PeriodSwitchStarted(uint256 newPeriod)
func (_Staking *StakingFilterer) FilterPeriodSwitchStarted(opts *bind.FilterOpts) (*StakingPeriodSwitchStartedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "PeriodSwitchStarted")
	if err != nil {
		return nil, err
	}
	return &StakingPeriodSwitchStartedIterator{contract: _Staking.contract, event: "PeriodSwitchStarted", logs: logs, sub: sub}, nil
}

// WatchPeriodSwitchStarted is a free log subscription operation binding the contract event 0xf4407b200bd8af965ee06a960f67118073146e5686605ddf1c487f08693530e1.
//
// Solidity: event PeriodSwitchStarted(uint256 newPeriod)
func (_Staking *StakingFilterer) WatchPeriodSwitchStarted(opts *bind.WatchOpts, sink chan<- *StakingPeriodSwitchStarted) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "PeriodSwitchStarted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingPeriodSwitchStarted)
				if err := _Staking.contract.UnpackLog(event, "PeriodSwitchStarted", log); err != nil {
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

// ParsePeriodSwitchStarted is a log parse operation binding the contract event 0xf4407b200bd8af965ee06a960f67118073146e5686605ddf1c487f08693530e1.
//
// Solidity: event PeriodSwitchStarted(uint256 newPeriod)
func (_Staking *StakingFilterer) ParsePeriodSwitchStarted(log types.Log) (*StakingPeriodSwitchStarted, error) {
	event := new(StakingPeriodSwitchStarted)
	if err := _Staking.contract.UnpackLog(event, "PeriodSwitchStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSpaceCreatedIterator is returned from FilterSpaceCreated and is used to iterate over the raw logs and unpacked data for SpaceCreated events raised by the Staking contract.
type StakingSpaceCreatedIterator struct {
	Event *StakingSpaceCreated // Event containing the contract specifics and raw log

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
func (it *StakingSpaceCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSpaceCreated)
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
		it.Event = new(StakingSpaceCreated)
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
func (it *StakingSpaceCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSpaceCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSpaceCreated represents a SpaceCreated event raised by the Staking contract.
type StakingSpaceCreated struct {
	SpaceId *big.Int
	Owner   common.Address
	Rate    *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpaceCreated is a free log retrieval operation binding the contract event 0x8dfa835b1141f12cb096cf2cad07edd3b4d731d7854921443f4578a93eef2c0d.
//
// Solidity: event SpaceCreated(uint256 indexed spaceId, address owner, uint256 rate)
func (_Staking *StakingFilterer) FilterSpaceCreated(opts *bind.FilterOpts, spaceId []*big.Int) (*StakingSpaceCreatedIterator, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SpaceCreated", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingSpaceCreatedIterator{contract: _Staking.contract, event: "SpaceCreated", logs: logs, sub: sub}, nil
}

// WatchSpaceCreated is a free log subscription operation binding the contract event 0x8dfa835b1141f12cb096cf2cad07edd3b4d731d7854921443f4578a93eef2c0d.
//
// Solidity: event SpaceCreated(uint256 indexed spaceId, address owner, uint256 rate)
func (_Staking *StakingFilterer) WatchSpaceCreated(opts *bind.WatchOpts, sink chan<- *StakingSpaceCreated, spaceId []*big.Int) (event.Subscription, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SpaceCreated", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSpaceCreated)
				if err := _Staking.contract.UnpackLog(event, "SpaceCreated", log); err != nil {
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

// ParseSpaceCreated is a log parse operation binding the contract event 0x8dfa835b1141f12cb096cf2cad07edd3b4d731d7854921443f4578a93eef2c0d.
//
// Solidity: event SpaceCreated(uint256 indexed spaceId, address owner, uint256 rate)
func (_Staking *StakingFilterer) ParseSpaceCreated(log types.Log) (*StakingSpaceCreated, error) {
	event := new(StakingSpaceCreated)
	if err := _Staking.contract.UnpackLog(event, "SpaceCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSpaceDestroyedIterator is returned from FilterSpaceDestroyed and is used to iterate over the raw logs and unpacked data for SpaceDestroyed events raised by the Staking contract.
type StakingSpaceDestroyedIterator struct {
	Event *StakingSpaceDestroyed // Event containing the contract specifics and raw log

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
func (it *StakingSpaceDestroyedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSpaceDestroyed)
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
		it.Event = new(StakingSpaceDestroyed)
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
func (it *StakingSpaceDestroyedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSpaceDestroyedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSpaceDestroyed represents a SpaceDestroyed event raised by the Staking contract.
type StakingSpaceDestroyed struct {
	SpaceId *big.Int
	Owner   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpaceDestroyed is a free log retrieval operation binding the contract event 0xad698dc4e5d3f1331c4aad60a215e76dcaa52435144b6c83d5742f4c1fda3785.
//
// Solidity: event SpaceDestroyed(uint256 indexed spaceId, address indexed owner)
func (_Staking *StakingFilterer) FilterSpaceDestroyed(opts *bind.FilterOpts, spaceId []*big.Int, owner []common.Address) (*StakingSpaceDestroyedIterator, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SpaceDestroyed", spaceIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &StakingSpaceDestroyedIterator{contract: _Staking.contract, event: "SpaceDestroyed", logs: logs, sub: sub}, nil
}

// WatchSpaceDestroyed is a free log subscription operation binding the contract event 0xad698dc4e5d3f1331c4aad60a215e76dcaa52435144b6c83d5742f4c1fda3785.
//
// Solidity: event SpaceDestroyed(uint256 indexed spaceId, address indexed owner)
func (_Staking *StakingFilterer) WatchSpaceDestroyed(opts *bind.WatchOpts, sink chan<- *StakingSpaceDestroyed, spaceId []*big.Int, owner []common.Address) (event.Subscription, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SpaceDestroyed", spaceIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSpaceDestroyed)
				if err := _Staking.contract.UnpackLog(event, "SpaceDestroyed", log); err != nil {
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

// ParseSpaceDestroyed is a log parse operation binding the contract event 0xad698dc4e5d3f1331c4aad60a215e76dcaa52435144b6c83d5742f4c1fda3785.
//
// Solidity: event SpaceDestroyed(uint256 indexed spaceId, address indexed owner)
func (_Staking *StakingFilterer) ParseSpaceDestroyed(log types.Log) (*StakingSpaceDestroyed, error) {
	event := new(StakingSpaceDestroyed)
	if err := _Staking.contract.UnpackLog(event, "SpaceDestroyed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSpaceRateAppliedIterator is returned from FilterSpaceRateApplied and is used to iterate over the raw logs and unpacked data for SpaceRateApplied events raised by the Staking contract.
type StakingSpaceRateAppliedIterator struct {
	Event *StakingSpaceRateApplied // Event containing the contract specifics and raw log

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
func (it *StakingSpaceRateAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSpaceRateApplied)
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
		it.Event = new(StakingSpaceRateApplied)
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
func (it *StakingSpaceRateAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSpaceRateAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSpaceRateApplied represents a SpaceRateApplied event raised by the Staking contract.
type StakingSpaceRateApplied struct {
	SpaceId *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpaceRateApplied is a free log retrieval operation binding the contract event 0xc6f2348b69178bcfce68053e52724b30040efb25ce6b6e87ba92c6a6b1e1a989.
//
// Solidity: event SpaceRateApplied(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) FilterSpaceRateApplied(opts *bind.FilterOpts, spaceId []*big.Int) (*StakingSpaceRateAppliedIterator, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SpaceRateApplied", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingSpaceRateAppliedIterator{contract: _Staking.contract, event: "SpaceRateApplied", logs: logs, sub: sub}, nil
}

// WatchSpaceRateApplied is a free log subscription operation binding the contract event 0xc6f2348b69178bcfce68053e52724b30040efb25ce6b6e87ba92c6a6b1e1a989.
//
// Solidity: event SpaceRateApplied(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) WatchSpaceRateApplied(opts *bind.WatchOpts, sink chan<- *StakingSpaceRateApplied, spaceId []*big.Int) (event.Subscription, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SpaceRateApplied", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSpaceRateApplied)
				if err := _Staking.contract.UnpackLog(event, "SpaceRateApplied", log); err != nil {
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

// ParseSpaceRateApplied is a log parse operation binding the contract event 0xc6f2348b69178bcfce68053e52724b30040efb25ce6b6e87ba92c6a6b1e1a989.
//
// Solidity: event SpaceRateApplied(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) ParseSpaceRateApplied(log types.Log) (*StakingSpaceRateApplied, error) {
	event := new(StakingSpaceRateApplied)
	if err := _Staking.contract.UnpackLog(event, "SpaceRateApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingSpaceRateUpdatedIterator is returned from FilterSpaceRateUpdated and is used to iterate over the raw logs and unpacked data for SpaceRateUpdated events raised by the Staking contract.
type StakingSpaceRateUpdatedIterator struct {
	Event *StakingSpaceRateUpdated // Event containing the contract specifics and raw log

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
func (it *StakingSpaceRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingSpaceRateUpdated)
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
		it.Event = new(StakingSpaceRateUpdated)
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
func (it *StakingSpaceRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingSpaceRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingSpaceRateUpdated represents a SpaceRateUpdated event raised by the Staking contract.
type StakingSpaceRateUpdated struct {
	SpaceId *big.Int
	NewRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterSpaceRateUpdated is a free log retrieval operation binding the contract event 0x73884a6861f5816802d387e2f05365572beaf532208dc5e446eaf5790a11a02b.
//
// Solidity: event SpaceRateUpdated(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) FilterSpaceRateUpdated(opts *bind.FilterOpts, spaceId []*big.Int) (*StakingSpaceRateUpdatedIterator, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "SpaceRateUpdated", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingSpaceRateUpdatedIterator{contract: _Staking.contract, event: "SpaceRateUpdated", logs: logs, sub: sub}, nil
}

// WatchSpaceRateUpdated is a free log subscription operation binding the contract event 0x73884a6861f5816802d387e2f05365572beaf532208dc5e446eaf5790a11a02b.
//
// Solidity: event SpaceRateUpdated(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) WatchSpaceRateUpdated(opts *bind.WatchOpts, sink chan<- *StakingSpaceRateUpdated, spaceId []*big.Int) (event.Subscription, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "SpaceRateUpdated", spaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingSpaceRateUpdated)
				if err := _Staking.contract.UnpackLog(event, "SpaceRateUpdated", log); err != nil {
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

// ParseSpaceRateUpdated is a log parse operation binding the contract event 0x73884a6861f5816802d387e2f05365572beaf532208dc5e446eaf5790a11a02b.
//
// Solidity: event SpaceRateUpdated(uint256 indexed spaceId, uint256 newRate)
func (_Staking *StakingFilterer) ParseSpaceRateUpdated(log types.Log) (*StakingSpaceRateUpdated, error) {
	event := new(StakingSpaceRateUpdated)
	if err := _Staking.contract.UnpackLog(event, "SpaceRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Staking contract.
type StakingStakedIterator struct {
	Event *StakingStaked // Event containing the contract specifics and raw log

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
func (it *StakingStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStaked)
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
		it.Event = new(StakingStaked)
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
func (it *StakingStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStaked represents a Staked event raised by the Staking contract.
type StakingStaked struct {
	User    common.Address
	SpaceId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) FilterStaked(opts *bind.FilterOpts, user []common.Address, spaceId []*big.Int) (*StakingStakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Staked", userRule, spaceIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakedIterator{contract: _Staking.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakingStaked, user []common.Address, spaceId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Staked", userRule, spaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStaked)
				if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
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

// ParseStaked is a log parse operation binding the contract event 0x1449c6dd7851abc30abf37f57715f492010519147cc2652fbc38202c18a6ee90.
//
// Solidity: event Staked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) ParseStaked(log types.Log) (*StakingStaked, error) {
	event := new(StakingStaked)
	if err := _Staking.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingStakerStatusProcessedIterator is returned from FilterStakerStatusProcessed and is used to iterate over the raw logs and unpacked data for StakerStatusProcessed events raised by the Staking contract.
type StakingStakerStatusProcessedIterator struct {
	Event *StakingStakerStatusProcessed // Event containing the contract specifics and raw log

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
func (it *StakingStakerStatusProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingStakerStatusProcessed)
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
		it.Event = new(StakingStakerStatusProcessed)
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
func (it *StakingStakerStatusProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingStakerStatusProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingStakerStatusProcessed represents a StakerStatusProcessed event raised by the Staking contract.
type StakingStakerStatusProcessed struct {
	SpaceId       *big.Int
	Staker        common.Address
	PendingAmount *big.Int
	LockedAmount  *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterStakerStatusProcessed is a free log retrieval operation binding the contract event 0xe39b4653f5ad7df59807b1daf78efa6e764e3467e6f2768c6ad31a1384e80b0d.
//
// Solidity: event StakerStatusProcessed(uint256 indexed spaceId, address indexed staker, uint256 pendingAmount, uint256 lockedAmount)
func (_Staking *StakingFilterer) FilterStakerStatusProcessed(opts *bind.FilterOpts, spaceId []*big.Int, staker []common.Address) (*StakingStakerStatusProcessedIterator, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "StakerStatusProcessed", spaceIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return &StakingStakerStatusProcessedIterator{contract: _Staking.contract, event: "StakerStatusProcessed", logs: logs, sub: sub}, nil
}

// WatchStakerStatusProcessed is a free log subscription operation binding the contract event 0xe39b4653f5ad7df59807b1daf78efa6e764e3467e6f2768c6ad31a1384e80b0d.
//
// Solidity: event StakerStatusProcessed(uint256 indexed spaceId, address indexed staker, uint256 pendingAmount, uint256 lockedAmount)
func (_Staking *StakingFilterer) WatchStakerStatusProcessed(opts *bind.WatchOpts, sink chan<- *StakingStakerStatusProcessed, spaceId []*big.Int, staker []common.Address) (event.Subscription, error) {

	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}
	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "StakerStatusProcessed", spaceIdRule, stakerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingStakerStatusProcessed)
				if err := _Staking.contract.UnpackLog(event, "StakerStatusProcessed", log); err != nil {
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

// ParseStakerStatusProcessed is a log parse operation binding the contract event 0xe39b4653f5ad7df59807b1daf78efa6e764e3467e6f2768c6ad31a1384e80b0d.
//
// Solidity: event StakerStatusProcessed(uint256 indexed spaceId, address indexed staker, uint256 pendingAmount, uint256 lockedAmount)
func (_Staking *StakingFilterer) ParseStakerStatusProcessed(log types.Log) (*StakingStakerStatusProcessed, error) {
	event := new(StakingStakerStatusProcessed)
	if err := _Staking.contract.UnpackLog(event, "StakerStatusProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Staking contract.
type StakingUnpausedIterator struct {
	Event *StakingUnpaused // Event containing the contract specifics and raw log

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
func (it *StakingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnpaused)
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
		it.Event = new(StakingUnpaused)
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
func (it *StakingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnpaused represents a Unpaused event raised by the Staking contract.
type StakingUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StakingUnpausedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StakingUnpausedIterator{contract: _Staking.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Staking *StakingFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StakingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnpaused)
				if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Staking *StakingFilterer) ParseUnpaused(log types.Log) (*StakingUnpaused, error) {
	event := new(StakingUnpaused)
	if err := _Staking.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Staking contract.
type StakingUnstakedIterator struct {
	Event *StakingUnstaked // Event containing the contract specifics and raw log

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
func (it *StakingUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingUnstaked)
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
		it.Event = new(StakingUnstaked)
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
func (it *StakingUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingUnstaked represents a Unstaked event raised by the Staking contract.
type StakingUnstaked struct {
	User    common.Address
	SpaceId *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) FilterUnstaked(opts *bind.FilterOpts, user []common.Address, spaceId []*big.Int) (*StakingUnstakedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Unstaked", userRule, spaceIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingUnstakedIterator{contract: _Staking.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakingUnstaked, user []common.Address, spaceId []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var spaceIdRule []interface{}
	for _, spaceIdItem := range spaceId {
		spaceIdRule = append(spaceIdRule, spaceIdItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Unstaked", userRule, spaceIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingUnstaked)
				if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
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

// ParseUnstaked is a log parse operation binding the contract event 0x7fc4727e062e336010f2c282598ef5f14facb3de68cf8195c2f23e1454b2b74e.
//
// Solidity: event Unstaked(address indexed user, uint256 indexed spaceId, uint256 amount)
func (_Staking *StakingFilterer) ParseUnstaked(log types.Log) (*StakingUnstaked, error) {
	event := new(StakingUnstaked)
	if err := _Staking.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWhitelistUpdatedIterator is returned from FilterWhitelistUpdated and is used to iterate over the raw logs and unpacked data for WhitelistUpdated events raised by the Staking contract.
type StakingWhitelistUpdatedIterator struct {
	Event *StakingWhitelistUpdated // Event containing the contract specifics and raw log

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
func (it *StakingWhitelistUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWhitelistUpdated)
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
		it.Event = new(StakingWhitelistUpdated)
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
func (it *StakingWhitelistUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWhitelistUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWhitelistUpdated represents a WhitelistUpdated event raised by the Staking contract.
type StakingWhitelistUpdated struct {
	User   common.Address
	Status bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWhitelistUpdated is a free log retrieval operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address user, bool status)
func (_Staking *StakingFilterer) FilterWhitelistUpdated(opts *bind.FilterOpts) (*StakingWhitelistUpdatedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "WhitelistUpdated")
	if err != nil {
		return nil, err
	}
	return &StakingWhitelistUpdatedIterator{contract: _Staking.contract, event: "WhitelistUpdated", logs: logs, sub: sub}, nil
}

// WatchWhitelistUpdated is a free log subscription operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address user, bool status)
func (_Staking *StakingFilterer) WatchWhitelistUpdated(opts *bind.WatchOpts, sink chan<- *StakingWhitelistUpdated) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "WhitelistUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWhitelistUpdated)
				if err := _Staking.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
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

// ParseWhitelistUpdated is a log parse operation binding the contract event 0xf93f9a76c1bf3444d22400a00cb9fe990e6abe9dbb333fda48859cfee864543d.
//
// Solidity: event WhitelistUpdated(address user, bool status)
func (_Staking *StakingFilterer) ParseWhitelistUpdated(log types.Log) (*StakingWhitelistUpdated, error) {
	event := new(StakingWhitelistUpdated)
	if err := _Staking.contract.UnpackLog(event, "WhitelistUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Staking contract.
type StakingWithdrawnIterator struct {
	Event *StakingWithdrawn // Event containing the contract specifics and raw log

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
func (it *StakingWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingWithdrawn)
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
		it.Event = new(StakingWithdrawn)
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
func (it *StakingWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingWithdrawn represents a Withdrawn event raised by the Staking contract.
type StakingWithdrawn struct {
	User   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Staking *StakingFilterer) FilterWithdrawn(opts *bind.FilterOpts, user []common.Address) (*StakingWithdrawnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return &StakingWithdrawnIterator{contract: _Staking.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Staking *StakingFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *StakingWithdrawn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Withdrawn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingWithdrawn)
				if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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

// ParseWithdrawn is a log parse operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed user, uint256 amount)
func (_Staking *StakingFilterer) ParseWithdrawn(log types.Log) (*StakingWithdrawn, error) {
	event := new(StakingWithdrawn)
	if err := _Staking.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
