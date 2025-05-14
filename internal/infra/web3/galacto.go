// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package web3

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

// GalactoMetaData contains all meta data concerning the Galacto contract.
var GalactoMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalanceBreakdown\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fromStaff\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromTransfer\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"join\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"publishAction\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"userWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"username\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actionId\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actionName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"performedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"staff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromStaff\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActionPublished\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"userWallet\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"userID\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"username\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"actionId\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"actionName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rewardAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"performedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Joined\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFromStaff\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMember\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NotMember\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NotStaff\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// GalactoABI is the input ABI used to generate the binding from.
// Deprecated: Use GalactoMetaData.ABI instead.
var GalactoABI = GalactoMetaData.ABI

// Galacto is an auto generated Go binding around an Ethereum contract.
type Galacto struct {
	GalactoCaller     // Read-only binding to the contract
	GalactoTransactor // Write-only binding to the contract
	GalactoFilterer   // Log filterer for contract events
}

// GalactoCaller is an auto generated read-only Go binding around an Ethereum contract.
type GalactoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GalactoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GalactoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GalactoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GalactoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GalactoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GalactoSession struct {
	Contract     *Galacto          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GalactoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GalactoCallerSession struct {
	Contract *GalactoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GalactoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GalactoTransactorSession struct {
	Contract     *GalactoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GalactoRaw is an auto generated low-level Go binding around an Ethereum contract.
type GalactoRaw struct {
	Contract *Galacto // Generic contract binding to access the raw methods on
}

// GalactoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GalactoCallerRaw struct {
	Contract *GalactoCaller // Generic read-only contract binding to access the raw methods on
}

// GalactoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GalactoTransactorRaw struct {
	Contract *GalactoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGalacto creates a new instance of Galacto, bound to a specific deployed contract.
func NewGalacto(address common.Address, backend bind.ContractBackend) (*Galacto, error) {
	contract, err := bindGalacto(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Galacto{GalactoCaller: GalactoCaller{contract: contract}, GalactoTransactor: GalactoTransactor{contract: contract}, GalactoFilterer: GalactoFilterer{contract: contract}}, nil
}

// NewGalactoCaller creates a new read-only instance of Galacto, bound to a specific deployed contract.
func NewGalactoCaller(address common.Address, caller bind.ContractCaller) (*GalactoCaller, error) {
	contract, err := bindGalacto(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GalactoCaller{contract: contract}, nil
}

// NewGalactoTransactor creates a new write-only instance of Galacto, bound to a specific deployed contract.
func NewGalactoTransactor(address common.Address, transactor bind.ContractTransactor) (*GalactoTransactor, error) {
	contract, err := bindGalacto(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GalactoTransactor{contract: contract}, nil
}

// NewGalactoFilterer creates a new log filterer instance of Galacto, bound to a specific deployed contract.
func NewGalactoFilterer(address common.Address, filterer bind.ContractFilterer) (*GalactoFilterer, error) {
	contract, err := bindGalacto(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GalactoFilterer{contract: contract}, nil
}

// bindGalacto binds a generic wrapper to an already deployed contract.
func bindGalacto(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GalactoMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Galacto *GalactoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Galacto.Contract.GalactoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Galacto *GalactoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Galacto.Contract.GalactoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Galacto *GalactoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Galacto.Contract.GalactoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Galacto *GalactoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Galacto.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Galacto *GalactoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Galacto.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Galacto *GalactoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Galacto.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Galacto *GalactoCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Galacto *GalactoSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Galacto.Contract.Allowance(&_Galacto.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Galacto *GalactoCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Galacto.Contract.Allowance(&_Galacto.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Galacto *GalactoCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Galacto *GalactoSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Galacto.Contract.BalanceOf(&_Galacto.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Galacto *GalactoCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Galacto.Contract.BalanceOf(&_Galacto.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Galacto *GalactoCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Galacto *GalactoSession) Decimals() (uint8, error) {
	return _Galacto.Contract.Decimals(&_Galacto.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Galacto *GalactoCallerSession) Decimals() (uint8, error) {
	return _Galacto.Contract.Decimals(&_Galacto.CallOpts)
}

// GetBalanceBreakdown is a free data retrieval call binding the contract method 0x15b5709a.
//
// Solidity: function getBalanceBreakdown(address wallet) view returns(uint256 fromStaff, uint256 fromTransfer)
func (_Galacto *GalactoCaller) GetBalanceBreakdown(opts *bind.CallOpts, wallet common.Address) (struct {
	FromStaff    *big.Int
	FromTransfer *big.Int
}, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "getBalanceBreakdown", wallet)

	outstruct := new(struct {
		FromStaff    *big.Int
		FromTransfer *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FromStaff = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FromTransfer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBalanceBreakdown is a free data retrieval call binding the contract method 0x15b5709a.
//
// Solidity: function getBalanceBreakdown(address wallet) view returns(uint256 fromStaff, uint256 fromTransfer)
func (_Galacto *GalactoSession) GetBalanceBreakdown(wallet common.Address) (struct {
	FromStaff    *big.Int
	FromTransfer *big.Int
}, error) {
	return _Galacto.Contract.GetBalanceBreakdown(&_Galacto.CallOpts, wallet)
}

// GetBalanceBreakdown is a free data retrieval call binding the contract method 0x15b5709a.
//
// Solidity: function getBalanceBreakdown(address wallet) view returns(uint256 fromStaff, uint256 fromTransfer)
func (_Galacto *GalactoCallerSession) GetBalanceBreakdown(wallet common.Address) (struct {
	FromStaff    *big.Int
	FromTransfer *big.Int
}, error) {
	return _Galacto.Contract.GetBalanceBreakdown(&_Galacto.CallOpts, wallet)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address ) view returns(bool)
func (_Galacto *GalactoCaller) IsMember(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "isMember", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address ) view returns(bool)
func (_Galacto *GalactoSession) IsMember(arg0 common.Address) (bool, error) {
	return _Galacto.Contract.IsMember(&_Galacto.CallOpts, arg0)
}

// IsMember is a free data retrieval call binding the contract method 0xa230c524.
//
// Solidity: function isMember(address ) view returns(bool)
func (_Galacto *GalactoCallerSession) IsMember(arg0 common.Address) (bool, error) {
	return _Galacto.Contract.IsMember(&_Galacto.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Galacto *GalactoCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Galacto *GalactoSession) Name() (string, error) {
	return _Galacto.Contract.Name(&_Galacto.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Galacto *GalactoCallerSession) Name() (string, error) {
	return _Galacto.Contract.Name(&_Galacto.CallOpts)
}

// Staff is a free data retrieval call binding the contract method 0xd11345e0.
//
// Solidity: function staff() view returns(address)
func (_Galacto *GalactoCaller) Staff(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "staff")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Staff is a free data retrieval call binding the contract method 0xd11345e0.
//
// Solidity: function staff() view returns(address)
func (_Galacto *GalactoSession) Staff() (common.Address, error) {
	return _Galacto.Contract.Staff(&_Galacto.CallOpts)
}

// Staff is a free data retrieval call binding the contract method 0xd11345e0.
//
// Solidity: function staff() view returns(address)
func (_Galacto *GalactoCallerSession) Staff() (common.Address, error) {
	return _Galacto.Contract.Staff(&_Galacto.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Galacto *GalactoCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Galacto *GalactoSession) Symbol() (string, error) {
	return _Galacto.Contract.Symbol(&_Galacto.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Galacto *GalactoCallerSession) Symbol() (string, error) {
	return _Galacto.Contract.Symbol(&_Galacto.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Galacto *GalactoCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Galacto.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Galacto *GalactoSession) TotalSupply() (*big.Int, error) {
	return _Galacto.Contract.TotalSupply(&_Galacto.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Galacto *GalactoCallerSession) TotalSupply() (*big.Int, error) {
	return _Galacto.Contract.TotalSupply(&_Galacto.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Galacto *GalactoTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Galacto *GalactoSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.Approve(&_Galacto.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Galacto *GalactoTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.Approve(&_Galacto.TransactOpts, spender, value)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Galacto *GalactoTransactor) Join(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "join")
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Galacto *GalactoSession) Join() (*types.Transaction, error) {
	return _Galacto.Contract.Join(&_Galacto.TransactOpts)
}

// Join is a paid mutator transaction binding the contract method 0xb688a363.
//
// Solidity: function join() returns()
func (_Galacto *GalactoTransactorSession) Join() (*types.Transaction, error) {
	return _Galacto.Contract.Join(&_Galacto.TransactOpts)
}

// PublishAction is a paid mutator transaction binding the contract method 0xccf52af8.
//
// Solidity: function publishAction(string id, address userWallet, string userId, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoTransactor) PublishAction(opts *bind.TransactOpts, id string, userWallet common.Address, userId string, username string, actionId string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "publishAction", id, userWallet, userId, username, actionId, actionName, rewardAmount, performedAt)
}

// PublishAction is a paid mutator transaction binding the contract method 0xccf52af8.
//
// Solidity: function publishAction(string id, address userWallet, string userId, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoSession) PublishAction(id string, userWallet common.Address, userId string, username string, actionId string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.PublishAction(&_Galacto.TransactOpts, id, userWallet, userId, username, actionId, actionName, rewardAmount, performedAt)
}

// PublishAction is a paid mutator transaction binding the contract method 0xccf52af8.
//
// Solidity: function publishAction(string id, address userWallet, string userId, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoTransactorSession) PublishAction(id string, userWallet common.Address, userId string, username string, actionId string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.PublishAction(&_Galacto.TransactOpts, id, userWallet, userId, username, actionId, actionName, rewardAmount, performedAt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Galacto *GalactoTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Galacto *GalactoSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.Transfer(&_Galacto.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Galacto *GalactoTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.Transfer(&_Galacto.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Galacto *GalactoTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Galacto *GalactoSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.TransferFrom(&_Galacto.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Galacto *GalactoTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.TransferFrom(&_Galacto.TransactOpts, from, to, value)
}

// TransferFromStaff is a paid mutator transaction binding the contract method 0x0bbae1b3.
//
// Solidity: function transferFromStaff(address to, uint256 amount) returns()
func (_Galacto *GalactoTransactor) TransferFromStaff(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "transferFromStaff", to, amount)
}

// TransferFromStaff is a paid mutator transaction binding the contract method 0x0bbae1b3.
//
// Solidity: function transferFromStaff(address to, uint256 amount) returns()
func (_Galacto *GalactoSession) TransferFromStaff(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.TransferFromStaff(&_Galacto.TransactOpts, to, amount)
}

// TransferFromStaff is a paid mutator transaction binding the contract method 0x0bbae1b3.
//
// Solidity: function transferFromStaff(address to, uint256 amount) returns()
func (_Galacto *GalactoTransactorSession) TransferFromStaff(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.TransferFromStaff(&_Galacto.TransactOpts, to, amount)
}

// GalactoActionPublishedIterator is returned from FilterActionPublished and is used to iterate over the raw logs and unpacked data for ActionPublished events raised by the Galacto contract.
type GalactoActionPublishedIterator struct {
	Event *GalactoActionPublished // Event containing the contract specifics and raw log

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
func (it *GalactoActionPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GalactoActionPublished)
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
		it.Event = new(GalactoActionPublished)
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
func (it *GalactoActionPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GalactoActionPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GalactoActionPublished represents a ActionPublished event raised by the Galacto contract.
type GalactoActionPublished struct {
	Id           string
	UserWallet   common.Address
	UserID       string
	Username     string
	ActionId     string
	ActionName   string
	RewardAmount *big.Int
	PerformedAt  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterActionPublished is a free log retrieval operation binding the contract event 0x747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c4.
//
// Solidity: event ActionPublished(string id, address userWallet, string userID, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt)
func (_Galacto *GalactoFilterer) FilterActionPublished(opts *bind.FilterOpts) (*GalactoActionPublishedIterator, error) {

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "ActionPublished")
	if err != nil {
		return nil, err
	}
	return &GalactoActionPublishedIterator{contract: _Galacto.contract, event: "ActionPublished", logs: logs, sub: sub}, nil
}

// WatchActionPublished is a free log subscription operation binding the contract event 0x747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c4.
//
// Solidity: event ActionPublished(string id, address userWallet, string userID, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt)
func (_Galacto *GalactoFilterer) WatchActionPublished(opts *bind.WatchOpts, sink chan<- *GalactoActionPublished) (event.Subscription, error) {

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "ActionPublished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GalactoActionPublished)
				if err := _Galacto.contract.UnpackLog(event, "ActionPublished", log); err != nil {
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

// ParseActionPublished is a log parse operation binding the contract event 0x747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c4.
//
// Solidity: event ActionPublished(string id, address userWallet, string userID, string username, string actionId, string actionName, uint256 rewardAmount, uint256 performedAt)
func (_Galacto *GalactoFilterer) ParseActionPublished(log types.Log) (*GalactoActionPublished, error) {
	event := new(GalactoActionPublished)
	if err := _Galacto.contract.UnpackLog(event, "ActionPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GalactoApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Galacto contract.
type GalactoApprovalIterator struct {
	Event *GalactoApproval // Event containing the contract specifics and raw log

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
func (it *GalactoApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GalactoApproval)
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
		it.Event = new(GalactoApproval)
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
func (it *GalactoApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GalactoApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GalactoApproval represents a Approval event raised by the Galacto contract.
type GalactoApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Galacto *GalactoFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*GalactoApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &GalactoApprovalIterator{contract: _Galacto.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Galacto *GalactoFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *GalactoApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GalactoApproval)
				if err := _Galacto.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Galacto *GalactoFilterer) ParseApproval(log types.Log) (*GalactoApproval, error) {
	event := new(GalactoApproval)
	if err := _Galacto.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GalactoJoinedIterator is returned from FilterJoined and is used to iterate over the raw logs and unpacked data for Joined events raised by the Galacto contract.
type GalactoJoinedIterator struct {
	Event *GalactoJoined // Event containing the contract specifics and raw log

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
func (it *GalactoJoinedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GalactoJoined)
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
		it.Event = new(GalactoJoined)
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
func (it *GalactoJoinedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GalactoJoinedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GalactoJoined represents a Joined event raised by the Galacto contract.
type GalactoJoined struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterJoined is a free log retrieval operation binding the contract event 0x7073afa60b48e833a1a66ebb442bc8e0d19e9f3cc05f34bdcd1da22c8d87f275.
//
// Solidity: event Joined(address indexed user)
func (_Galacto *GalactoFilterer) FilterJoined(opts *bind.FilterOpts, user []common.Address) (*GalactoJoinedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "Joined", userRule)
	if err != nil {
		return nil, err
	}
	return &GalactoJoinedIterator{contract: _Galacto.contract, event: "Joined", logs: logs, sub: sub}, nil
}

// WatchJoined is a free log subscription operation binding the contract event 0x7073afa60b48e833a1a66ebb442bc8e0d19e9f3cc05f34bdcd1da22c8d87f275.
//
// Solidity: event Joined(address indexed user)
func (_Galacto *GalactoFilterer) WatchJoined(opts *bind.WatchOpts, sink chan<- *GalactoJoined, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "Joined", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GalactoJoined)
				if err := _Galacto.contract.UnpackLog(event, "Joined", log); err != nil {
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

// ParseJoined is a log parse operation binding the contract event 0x7073afa60b48e833a1a66ebb442bc8e0d19e9f3cc05f34bdcd1da22c8d87f275.
//
// Solidity: event Joined(address indexed user)
func (_Galacto *GalactoFilterer) ParseJoined(log types.Log) (*GalactoJoined, error) {
	event := new(GalactoJoined)
	if err := _Galacto.contract.UnpackLog(event, "Joined", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GalactoTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Galacto contract.
type GalactoTransferIterator struct {
	Event *GalactoTransfer // Event containing the contract specifics and raw log

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
func (it *GalactoTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GalactoTransfer)
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
		it.Event = new(GalactoTransfer)
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
func (it *GalactoTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GalactoTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GalactoTransfer represents a Transfer event raised by the Galacto contract.
type GalactoTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Galacto *GalactoFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*GalactoTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &GalactoTransferIterator{contract: _Galacto.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Galacto *GalactoFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *GalactoTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GalactoTransfer)
				if err := _Galacto.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Galacto *GalactoFilterer) ParseTransfer(log types.Log) (*GalactoTransfer, error) {
	event := new(GalactoTransfer)
	if err := _Galacto.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GalactoTransferFromStaffIterator is returned from FilterTransferFromStaff and is used to iterate over the raw logs and unpacked data for TransferFromStaff events raised by the Galacto contract.
type GalactoTransferFromStaffIterator struct {
	Event *GalactoTransferFromStaff // Event containing the contract specifics and raw log

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
func (it *GalactoTransferFromStaffIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GalactoTransferFromStaff)
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
		it.Event = new(GalactoTransferFromStaff)
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
func (it *GalactoTransferFromStaffIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GalactoTransferFromStaffIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GalactoTransferFromStaff represents a TransferFromStaff event raised by the Galacto contract.
type GalactoTransferFromStaff struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransferFromStaff is a free log retrieval operation binding the contract event 0x8091ad0f6e4ac8d4debf0ec435f6321d71ce46af83248ee4e3860328e24faf1b.
//
// Solidity: event TransferFromStaff(address indexed to, uint256 amount)
func (_Galacto *GalactoFilterer) FilterTransferFromStaff(opts *bind.FilterOpts, to []common.Address) (*GalactoTransferFromStaffIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "TransferFromStaff", toRule)
	if err != nil {
		return nil, err
	}
	return &GalactoTransferFromStaffIterator{contract: _Galacto.contract, event: "TransferFromStaff", logs: logs, sub: sub}, nil
}

// WatchTransferFromStaff is a free log subscription operation binding the contract event 0x8091ad0f6e4ac8d4debf0ec435f6321d71ce46af83248ee4e3860328e24faf1b.
//
// Solidity: event TransferFromStaff(address indexed to, uint256 amount)
func (_Galacto *GalactoFilterer) WatchTransferFromStaff(opts *bind.WatchOpts, sink chan<- *GalactoTransferFromStaff, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "TransferFromStaff", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GalactoTransferFromStaff)
				if err := _Galacto.contract.UnpackLog(event, "TransferFromStaff", log); err != nil {
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

// ParseTransferFromStaff is a log parse operation binding the contract event 0x8091ad0f6e4ac8d4debf0ec435f6321d71ce46af83248ee4e3860328e24faf1b.
//
// Solidity: event TransferFromStaff(address indexed to, uint256 amount)
func (_Galacto *GalactoFilterer) ParseTransferFromStaff(log types.Log) (*GalactoTransferFromStaff, error) {
	event := new(GalactoTransferFromStaff)
	if err := _Galacto.contract.UnpackLog(event, "TransferFromStaff", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
