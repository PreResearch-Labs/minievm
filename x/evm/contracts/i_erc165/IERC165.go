// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package i_erc165

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

// IErc165MetaData contains all meta data concerning the IErc165 contract.
var IErc165MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IErc165ABI is the input ABI used to generate the binding from.
// Deprecated: Use IErc165MetaData.ABI instead.
var IErc165ABI = IErc165MetaData.ABI

// IErc165 is an auto generated Go binding around an Ethereum contract.
type IErc165 struct {
	IErc165Caller     // Read-only binding to the contract
	IErc165Transactor // Write-only binding to the contract
	IErc165Filterer   // Log filterer for contract events
}

// IErc165Caller is an auto generated read-only Go binding around an Ethereum contract.
type IErc165Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc165Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IErc165Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc165Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IErc165Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IErc165Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IErc165Session struct {
	Contract     *IErc165          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IErc165CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IErc165CallerSession struct {
	Contract *IErc165Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// IErc165TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IErc165TransactorSession struct {
	Contract     *IErc165Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IErc165Raw is an auto generated low-level Go binding around an Ethereum contract.
type IErc165Raw struct {
	Contract *IErc165 // Generic contract binding to access the raw methods on
}

// IErc165CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IErc165CallerRaw struct {
	Contract *IErc165Caller // Generic read-only contract binding to access the raw methods on
}

// IErc165TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IErc165TransactorRaw struct {
	Contract *IErc165Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIErc165 creates a new instance of IErc165, bound to a specific deployed contract.
func NewIErc165(address common.Address, backend bind.ContractBackend) (*IErc165, error) {
	contract, err := bindIErc165(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IErc165{IErc165Caller: IErc165Caller{contract: contract}, IErc165Transactor: IErc165Transactor{contract: contract}, IErc165Filterer: IErc165Filterer{contract: contract}}, nil
}

// NewIErc165Caller creates a new read-only instance of IErc165, bound to a specific deployed contract.
func NewIErc165Caller(address common.Address, caller bind.ContractCaller) (*IErc165Caller, error) {
	contract, err := bindIErc165(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IErc165Caller{contract: contract}, nil
}

// NewIErc165Transactor creates a new write-only instance of IErc165, bound to a specific deployed contract.
func NewIErc165Transactor(address common.Address, transactor bind.ContractTransactor) (*IErc165Transactor, error) {
	contract, err := bindIErc165(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IErc165Transactor{contract: contract}, nil
}

// NewIErc165Filterer creates a new log filterer instance of IErc165, bound to a specific deployed contract.
func NewIErc165Filterer(address common.Address, filterer bind.ContractFilterer) (*IErc165Filterer, error) {
	contract, err := bindIErc165(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IErc165Filterer{contract: contract}, nil
}

// bindIErc165 binds a generic wrapper to an already deployed contract.
func bindIErc165(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IErc165MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc165 *IErc165Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc165.Contract.IErc165Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc165 *IErc165Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc165.Contract.IErc165Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc165 *IErc165Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc165.Contract.IErc165Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IErc165 *IErc165CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IErc165.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IErc165 *IErc165TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IErc165.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IErc165 *IErc165TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IErc165.Contract.contract.Transact(opts, method, params...)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IErc165 *IErc165Caller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _IErc165.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IErc165 *IErc165Session) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IErc165.Contract.SupportsInterface(&_IErc165.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_IErc165 *IErc165CallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _IErc165.Contract.SupportsInterface(&_IErc165.CallOpts, interfaceId)
}
