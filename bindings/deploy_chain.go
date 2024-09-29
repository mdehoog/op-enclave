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
)

// DeployChainDeploy is an auto generated low-level Go binding around an user-defined struct.
type DeployChainDeploy struct {
	L2OutputOracle               common.Address
	SystemConfig                 common.Address
	OptimismPortal               common.Address
	L1CrossDomainMessenger       common.Address
	L1StandardBridge             common.Address
	L1ERC721Bridge               common.Address
	OptimismMintableERC20Factory common.Address
}

// DeployChainMetaData contains all meta data concerning the DeployChain contract.
var DeployChainMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_proxyAdmin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_outputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_superchainConfig\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deploy\",\"inputs\":[{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainID\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"genesisHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structDeployChain.Deploy\",\"components\":[{\"name\":\"l2OutputOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"systemConfig\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"l1CrossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ERC721Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1StandardBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismMintableERC20Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outputOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxyAdmin\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"superchainConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"}]",
}

// DeployChainABI is the input ABI used to generate the binding from.
// Deprecated: Use DeployChainMetaData.ABI instead.
var DeployChainABI = DeployChainMetaData.ABI

// DeployChain is an auto generated Go binding around an Ethereum contract.
type DeployChain struct {
	DeployChainCaller     // Read-only binding to the contract
	DeployChainTransactor // Write-only binding to the contract
	DeployChainFilterer   // Log filterer for contract events
}

// DeployChainCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeployChainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeployChainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeployChainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeployChainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeployChainSession struct {
	Contract     *DeployChain      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeployChainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeployChainCallerSession struct {
	Contract *DeployChainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// DeployChainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeployChainTransactorSession struct {
	Contract     *DeployChainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// DeployChainRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeployChainRaw struct {
	Contract *DeployChain // Generic contract binding to access the raw methods on
}

// DeployChainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeployChainCallerRaw struct {
	Contract *DeployChainCaller // Generic read-only contract binding to access the raw methods on
}

// DeployChainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeployChainTransactorRaw struct {
	Contract *DeployChainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeployChain creates a new instance of DeployChain, bound to a specific deployed contract.
func NewDeployChain(address common.Address, backend bind.ContractBackend) (*DeployChain, error) {
	contract, err := bindDeployChain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DeployChain{DeployChainCaller: DeployChainCaller{contract: contract}, DeployChainTransactor: DeployChainTransactor{contract: contract}, DeployChainFilterer: DeployChainFilterer{contract: contract}}, nil
}

// NewDeployChainCaller creates a new read-only instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainCaller(address common.Address, caller bind.ContractCaller) (*DeployChainCaller, error) {
	contract, err := bindDeployChain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeployChainCaller{contract: contract}, nil
}

// NewDeployChainTransactor creates a new write-only instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainTransactor(address common.Address, transactor bind.ContractTransactor) (*DeployChainTransactor, error) {
	contract, err := bindDeployChain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeployChainTransactor{contract: contract}, nil
}

// NewDeployChainFilterer creates a new log filterer instance of DeployChain, bound to a specific deployed contract.
func NewDeployChainFilterer(address common.Address, filterer bind.ContractFilterer) (*DeployChainFilterer, error) {
	contract, err := bindDeployChain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeployChainFilterer{contract: contract}, nil
}

// bindDeployChain binds a generic wrapper to an already deployed contract.
func bindDeployChain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeployChainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployChain *DeployChainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployChain.Contract.DeployChainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployChain *DeployChainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployChainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployChain *DeployChainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployChain.Contract.DeployChainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DeployChain *DeployChainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DeployChain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DeployChain *DeployChainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DeployChain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DeployChain *DeployChainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DeployChain.Contract.contract.Transact(opts, method, params...)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainCaller) L1CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1CrossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainSession) L1CrossDomainMessenger() (common.Address, error) {
	return _DeployChain.Contract.L1CrossDomainMessenger(&_DeployChain.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _DeployChain.Contract.L1CrossDomainMessenger(&_DeployChain.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainCaller) L1ERC721Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1ERC721Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainSession) L1ERC721Bridge() (common.Address, error) {
	return _DeployChain.Contract.L1ERC721Bridge(&_DeployChain.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1ERC721Bridge() (common.Address, error) {
	return _DeployChain.Contract.L1ERC721Bridge(&_DeployChain.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainCaller) L1StandardBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "l1StandardBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainSession) L1StandardBridge() (common.Address, error) {
	return _DeployChain.Contract.L1StandardBridge(&_DeployChain.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address)
func (_DeployChain *DeployChainCallerSession) L1StandardBridge() (common.Address, error) {
	return _DeployChain.Contract.L1StandardBridge(&_DeployChain.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainCaller) OptimismMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "optimismMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _DeployChain.Contract.OptimismMintableERC20Factory(&_DeployChain.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address)
func (_DeployChain *DeployChainCallerSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _DeployChain.Contract.OptimismMintableERC20Factory(&_DeployChain.CallOpts)
}

// OutputOracle is a free data retrieval call binding the contract method 0xe02b4c4a.
//
// Solidity: function outputOracle() view returns(address)
func (_DeployChain *DeployChainCaller) OutputOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "outputOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutputOracle is a free data retrieval call binding the contract method 0xe02b4c4a.
//
// Solidity: function outputOracle() view returns(address)
func (_DeployChain *DeployChainSession) OutputOracle() (common.Address, error) {
	return _DeployChain.Contract.OutputOracle(&_DeployChain.CallOpts)
}

// OutputOracle is a free data retrieval call binding the contract method 0xe02b4c4a.
//
// Solidity: function outputOracle() view returns(address)
func (_DeployChain *DeployChainCallerSession) OutputOracle() (common.Address, error) {
	return _DeployChain.Contract.OutputOracle(&_DeployChain.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_DeployChain *DeployChainCaller) Portal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "portal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_DeployChain *DeployChainSession) Portal() (common.Address, error) {
	return _DeployChain.Contract.Portal(&_DeployChain.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_DeployChain *DeployChainCallerSession) Portal() (common.Address, error) {
	return _DeployChain.Contract.Portal(&_DeployChain.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainCaller) ProxyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "proxyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainSession) ProxyAdmin() (common.Address, error) {
	return _DeployChain.Contract.ProxyAdmin(&_DeployChain.CallOpts)
}

// ProxyAdmin is a free data retrieval call binding the contract method 0x3e47158c.
//
// Solidity: function proxyAdmin() view returns(address)
func (_DeployChain *DeployChainCallerSession) ProxyAdmin() (common.Address, error) {
	return _DeployChain.Contract.ProxyAdmin(&_DeployChain.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainCaller) SuperchainConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "superchainConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainSession) SuperchainConfig() (common.Address, error) {
	return _DeployChain.Contract.SuperchainConfig(&_DeployChain.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_DeployChain *DeployChainCallerSession) SuperchainConfig() (common.Address, error) {
	return _DeployChain.Contract.SuperchainConfig(&_DeployChain.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainCaller) SystemConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DeployChain.contract.Call(opts, &out, "systemConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainSession) SystemConfig() (common.Address, error) {
	return _DeployChain.Contract.SystemConfig(&_DeployChain.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_DeployChain *DeployChainCallerSession) SystemConfig() (common.Address, error) {
	return _DeployChain.Contract.SystemConfig(&_DeployChain.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0xb6a944da.
//
// Solidity: function deploy(bytes32 salt, uint256 chainID, bytes32 genesisHash) returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainTransactor) Deploy(opts *bind.TransactOpts, salt [32]byte, chainID *big.Int, genesisHash [32]byte) (*types.Transaction, error) {
	return _DeployChain.contract.Transact(opts, "deploy", salt, chainID, genesisHash)
}

// Deploy is a paid mutator transaction binding the contract method 0xb6a944da.
//
// Solidity: function deploy(bytes32 salt, uint256 chainID, bytes32 genesisHash) returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainSession) Deploy(salt [32]byte, chainID *big.Int, genesisHash [32]byte) (*types.Transaction, error) {
	return _DeployChain.Contract.Deploy(&_DeployChain.TransactOpts, salt, chainID, genesisHash)
}

// Deploy is a paid mutator transaction binding the contract method 0xb6a944da.
//
// Solidity: function deploy(bytes32 salt, uint256 chainID, bytes32 genesisHash) returns((address,address,address,address,address,address,address))
func (_DeployChain *DeployChainTransactorSession) Deploy(salt [32]byte, chainID *big.Int, genesisHash [32]byte) (*types.Transaction, error) {
	return _DeployChain.Contract.Deploy(&_DeployChain.TransactOpts, salt, chainID, genesisHash)
}
