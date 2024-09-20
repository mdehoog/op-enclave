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

// ResourceMeteringResourceConfig is an auto generated low-level Go binding around an user-defined struct.
type ResourceMeteringResourceConfig struct {
	MaxResourceLimit            uint32
	ElasticityMultiplier        uint8
	BaseFeeMaxChangeDenominator uint8
	MinimumBaseFee              uint32
	SystemTxMaxGas              uint32
	MaximumBaseFee              *big.Int
}

// SystemConfigAddresses is an auto generated low-level Go binding around an user-defined struct.
type SystemConfigAddresses struct {
	L1CrossDomainMessenger       common.Address
	L1ERC721Bridge               common.Address
	L1StandardBridge             common.Address
	DisputeGameFactory           common.Address
	OptimismPortal               common.Address
	OptimismMintableERC20Factory common.Address
	GasPayingToken               common.Address
}

// SystemConfigOwnableMetaData contains all meta data concerning the SystemConfigOwnable contract.
var SystemConfigOwnableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BATCH_INBOX_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DISPUTE_GAME_FACTORY_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_CROSS_DOMAIN_MESSENGER_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_ERC_721_BRIDGE_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"L1_STANDARD_BRIDGE_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OPTIMISM_PORTAL_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"START_BLOCK_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNSAFE_BLOCK_SIGNER_SLOT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"basefeeScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batchInbox\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"batcherHash\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"blobbasefeeScalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"disputeGameFactory\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPayingToken\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"decimals_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPayingTokenName\",\"inputs\":[],\"outputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPayingTokenSymbol\",\"inputs\":[],\"outputs\":[{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_basefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_blobbasefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_batcherHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_unsafeBlockSigner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_config\",\"type\":\"tuple\",\"internalType\":\"structResourceMetering.ResourceConfig\",\"components\":[{\"name\":\"maxResourceLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minimumBaseFee\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"systemTxMaxGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maximumBaseFee\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]},{\"name\":\"_batchInbox\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_addresses\",\"type\":\"tuple\",\"internalType\":\"structSystemConfig.Addresses\",\"components\":[{\"name\":\"l1CrossDomainMessenger\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1ERC721Bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1StandardBridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"disputeGameFactory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismPortal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"optimismMintableERC20Factory\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"gasPayingToken\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isCustomGasToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1CrossDomainMessenger\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ERC721Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1StandardBridge\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"maximumGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"minimumGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismMintableERC20Factory\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"optimismPortal\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"overhead\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resourceConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structResourceMetering.ResourceConfig\",\"components\":[{\"name\":\"maxResourceLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"elasticityMultiplier\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseFeeMaxChangeDenominator\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"minimumBaseFee\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"systemTxMaxGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maximumBaseFee\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"scalar\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setBatcherHash\",\"inputs\":[{\"name\":\"_batcherHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasConfig\",\"inputs\":[{\"name\":\"_overhead\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_scalar\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasConfigEcotone\",\"inputs\":[{\"name\":\"_basefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"_blobbasefeeScalar\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasLimit\",\"inputs\":[{\"name\":\"_gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUnsafeBlockSigner\",\"inputs\":[{\"name\":\"_unsafeBlockSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"startBlock\",\"inputs\":[],\"outputs\":[{\"name\":\"startBlock_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unsafeBlockSigner\",\"inputs\":[],\"outputs\":[{\"name\":\"addr_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ConfigUpdate\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"updateType\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumSystemConfig.UpdateType\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false}]",
}

// SystemConfigOwnableABI is the input ABI used to generate the binding from.
// Deprecated: Use SystemConfigOwnableMetaData.ABI instead.
var SystemConfigOwnableABI = SystemConfigOwnableMetaData.ABI

// SystemConfigOwnable is an auto generated Go binding around an Ethereum contract.
type SystemConfigOwnable struct {
	SystemConfigOwnableCaller     // Read-only binding to the contract
	SystemConfigOwnableTransactor // Write-only binding to the contract
	SystemConfigOwnableFilterer   // Log filterer for contract events
}

// SystemConfigOwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type SystemConfigOwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigOwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SystemConfigOwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigOwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SystemConfigOwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SystemConfigOwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SystemConfigOwnableSession struct {
	Contract     *SystemConfigOwnable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SystemConfigOwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SystemConfigOwnableCallerSession struct {
	Contract *SystemConfigOwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// SystemConfigOwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SystemConfigOwnableTransactorSession struct {
	Contract     *SystemConfigOwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// SystemConfigOwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type SystemConfigOwnableRaw struct {
	Contract *SystemConfigOwnable // Generic contract binding to access the raw methods on
}

// SystemConfigOwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SystemConfigOwnableCallerRaw struct {
	Contract *SystemConfigOwnableCaller // Generic read-only contract binding to access the raw methods on
}

// SystemConfigOwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SystemConfigOwnableTransactorRaw struct {
	Contract *SystemConfigOwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSystemConfigOwnable creates a new instance of SystemConfigOwnable, bound to a specific deployed contract.
func NewSystemConfigOwnable(address common.Address, backend bind.ContractBackend) (*SystemConfigOwnable, error) {
	contract, err := bindSystemConfigOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnable{SystemConfigOwnableCaller: SystemConfigOwnableCaller{contract: contract}, SystemConfigOwnableTransactor: SystemConfigOwnableTransactor{contract: contract}, SystemConfigOwnableFilterer: SystemConfigOwnableFilterer{contract: contract}}, nil
}

// NewSystemConfigOwnableCaller creates a new read-only instance of SystemConfigOwnable, bound to a specific deployed contract.
func NewSystemConfigOwnableCaller(address common.Address, caller bind.ContractCaller) (*SystemConfigOwnableCaller, error) {
	contract, err := bindSystemConfigOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnableCaller{contract: contract}, nil
}

// NewSystemConfigOwnableTransactor creates a new write-only instance of SystemConfigOwnable, bound to a specific deployed contract.
func NewSystemConfigOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*SystemConfigOwnableTransactor, error) {
	contract, err := bindSystemConfigOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnableTransactor{contract: contract}, nil
}

// NewSystemConfigOwnableFilterer creates a new log filterer instance of SystemConfigOwnable, bound to a specific deployed contract.
func NewSystemConfigOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*SystemConfigOwnableFilterer, error) {
	contract, err := bindSystemConfigOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnableFilterer{contract: contract}, nil
}

// bindSystemConfigOwnable binds a generic wrapper to an already deployed contract.
func bindSystemConfigOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SystemConfigOwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemConfigOwnable *SystemConfigOwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemConfigOwnable.Contract.SystemConfigOwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemConfigOwnable *SystemConfigOwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SystemConfigOwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemConfigOwnable *SystemConfigOwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SystemConfigOwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SystemConfigOwnable *SystemConfigOwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SystemConfigOwnable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SystemConfigOwnable *SystemConfigOwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SystemConfigOwnable *SystemConfigOwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.contract.Transact(opts, method, params...)
}

// BATCHINBOXSLOT is a free data retrieval call binding the contract method 0xbc49ce5f.
//
// Solidity: function BATCH_INBOX_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) BATCHINBOXSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "BATCH_INBOX_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BATCHINBOXSLOT is a free data retrieval call binding the contract method 0xbc49ce5f.
//
// Solidity: function BATCH_INBOX_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) BATCHINBOXSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.BATCHINBOXSLOT(&_SystemConfigOwnable.CallOpts)
}

// BATCHINBOXSLOT is a free data retrieval call binding the contract method 0xbc49ce5f.
//
// Solidity: function BATCH_INBOX_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) BATCHINBOXSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.BATCHINBOXSLOT(&_SystemConfigOwnable.CallOpts)
}

// DISPUTEGAMEFACTORYSLOT is a free data retrieval call binding the contract method 0xe2a3285c.
//
// Solidity: function DISPUTE_GAME_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) DISPUTEGAMEFACTORYSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "DISPUTE_GAME_FACTORY_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DISPUTEGAMEFACTORYSLOT is a free data retrieval call binding the contract method 0xe2a3285c.
//
// Solidity: function DISPUTE_GAME_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) DISPUTEGAMEFACTORYSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.DISPUTEGAMEFACTORYSLOT(&_SystemConfigOwnable.CallOpts)
}

// DISPUTEGAMEFACTORYSLOT is a free data retrieval call binding the contract method 0xe2a3285c.
//
// Solidity: function DISPUTE_GAME_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) DISPUTEGAMEFACTORYSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.DISPUTEGAMEFACTORYSLOT(&_SystemConfigOwnable.CallOpts)
}

// L1CROSSDOMAINMESSENGERSLOT is a free data retrieval call binding the contract method 0x5d73369c.
//
// Solidity: function L1_CROSS_DOMAIN_MESSENGER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1CROSSDOMAINMESSENGERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "L1_CROSS_DOMAIN_MESSENGER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1CROSSDOMAINMESSENGERSLOT is a free data retrieval call binding the contract method 0x5d73369c.
//
// Solidity: function L1_CROSS_DOMAIN_MESSENGER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1CROSSDOMAINMESSENGERSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1CROSSDOMAINMESSENGERSLOT(&_SystemConfigOwnable.CallOpts)
}

// L1CROSSDOMAINMESSENGERSLOT is a free data retrieval call binding the contract method 0x5d73369c.
//
// Solidity: function L1_CROSS_DOMAIN_MESSENGER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1CROSSDOMAINMESSENGERSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1CROSSDOMAINMESSENGERSLOT(&_SystemConfigOwnable.CallOpts)
}

// L1ERC721BRIDGESLOT is a free data retrieval call binding the contract method 0x19f5cea8.
//
// Solidity: function L1_ERC_721_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1ERC721BRIDGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "L1_ERC_721_BRIDGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1ERC721BRIDGESLOT is a free data retrieval call binding the contract method 0x19f5cea8.
//
// Solidity: function L1_ERC_721_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1ERC721BRIDGESLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1ERC721BRIDGESLOT(&_SystemConfigOwnable.CallOpts)
}

// L1ERC721BRIDGESLOT is a free data retrieval call binding the contract method 0x19f5cea8.
//
// Solidity: function L1_ERC_721_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1ERC721BRIDGESLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1ERC721BRIDGESLOT(&_SystemConfigOwnable.CallOpts)
}

// L1STANDARDBRIDGESLOT is a free data retrieval call binding the contract method 0xf8c68de0.
//
// Solidity: function L1_STANDARD_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1STANDARDBRIDGESLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "L1_STANDARD_BRIDGE_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// L1STANDARDBRIDGESLOT is a free data retrieval call binding the contract method 0xf8c68de0.
//
// Solidity: function L1_STANDARD_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1STANDARDBRIDGESLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1STANDARDBRIDGESLOT(&_SystemConfigOwnable.CallOpts)
}

// L1STANDARDBRIDGESLOT is a free data retrieval call binding the contract method 0xf8c68de0.
//
// Solidity: function L1_STANDARD_BRIDGE_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1STANDARDBRIDGESLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.L1STANDARDBRIDGESLOT(&_SystemConfigOwnable.CallOpts)
}

// OPTIMISMMINTABLEERC20FACTORYSLOT is a free data retrieval call binding the contract method 0x06c92657.
//
// Solidity: function OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) OPTIMISMMINTABLEERC20FACTORYSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPTIMISMMINTABLEERC20FACTORYSLOT is a free data retrieval call binding the contract method 0x06c92657.
//
// Solidity: function OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) OPTIMISMMINTABLEERC20FACTORYSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.OPTIMISMMINTABLEERC20FACTORYSLOT(&_SystemConfigOwnable.CallOpts)
}

// OPTIMISMMINTABLEERC20FACTORYSLOT is a free data retrieval call binding the contract method 0x06c92657.
//
// Solidity: function OPTIMISM_MINTABLE_ERC20_FACTORY_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) OPTIMISMMINTABLEERC20FACTORYSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.OPTIMISMMINTABLEERC20FACTORYSLOT(&_SystemConfigOwnable.CallOpts)
}

// OPTIMISMPORTALSLOT is a free data retrieval call binding the contract method 0xfd32aa0f.
//
// Solidity: function OPTIMISM_PORTAL_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) OPTIMISMPORTALSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "OPTIMISM_PORTAL_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// OPTIMISMPORTALSLOT is a free data retrieval call binding the contract method 0xfd32aa0f.
//
// Solidity: function OPTIMISM_PORTAL_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) OPTIMISMPORTALSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.OPTIMISMPORTALSLOT(&_SystemConfigOwnable.CallOpts)
}

// OPTIMISMPORTALSLOT is a free data retrieval call binding the contract method 0xfd32aa0f.
//
// Solidity: function OPTIMISM_PORTAL_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) OPTIMISMPORTALSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.OPTIMISMPORTALSLOT(&_SystemConfigOwnable.CallOpts)
}

// STARTBLOCKSLOT is a free data retrieval call binding the contract method 0xe0e2016d.
//
// Solidity: function START_BLOCK_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) STARTBLOCKSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "START_BLOCK_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STARTBLOCKSLOT is a free data retrieval call binding the contract method 0xe0e2016d.
//
// Solidity: function START_BLOCK_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) STARTBLOCKSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.STARTBLOCKSLOT(&_SystemConfigOwnable.CallOpts)
}

// STARTBLOCKSLOT is a free data retrieval call binding the contract method 0xe0e2016d.
//
// Solidity: function START_BLOCK_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) STARTBLOCKSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.STARTBLOCKSLOT(&_SystemConfigOwnable.CallOpts)
}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) UNSAFEBLOCKSIGNERSLOT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "UNSAFE_BLOCK_SIGNER_SLOT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) UNSAFEBLOCKSIGNERSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.UNSAFEBLOCKSIGNERSLOT(&_SystemConfigOwnable.CallOpts)
}

// UNSAFEBLOCKSIGNERSLOT is a free data retrieval call binding the contract method 0x4f16540b.
//
// Solidity: function UNSAFE_BLOCK_SIGNER_SLOT() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) UNSAFEBLOCKSIGNERSLOT() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.UNSAFEBLOCKSIGNERSLOT(&_SystemConfigOwnable.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) VERSION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableSession) VERSION() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.VERSION(&_SystemConfigOwnable.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) VERSION() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.VERSION(&_SystemConfigOwnable.CallOpts)
}

// BasefeeScalar is a free data retrieval call binding the contract method 0xbfb14fb7.
//
// Solidity: function basefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) BasefeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "basefeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BasefeeScalar is a free data retrieval call binding the contract method 0xbfb14fb7.
//
// Solidity: function basefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) BasefeeScalar() (uint32, error) {
	return _SystemConfigOwnable.Contract.BasefeeScalar(&_SystemConfigOwnable.CallOpts)
}

// BasefeeScalar is a free data retrieval call binding the contract method 0xbfb14fb7.
//
// Solidity: function basefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) BasefeeScalar() (uint32, error) {
	return _SystemConfigOwnable.Contract.BasefeeScalar(&_SystemConfigOwnable.CallOpts)
}

// BatchInbox is a free data retrieval call binding the contract method 0xdac6e63a.
//
// Solidity: function batchInbox() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) BatchInbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "batchInbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BatchInbox is a free data retrieval call binding the contract method 0xdac6e63a.
//
// Solidity: function batchInbox() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) BatchInbox() (common.Address, error) {
	return _SystemConfigOwnable.Contract.BatchInbox(&_SystemConfigOwnable.CallOpts)
}

// BatchInbox is a free data retrieval call binding the contract method 0xdac6e63a.
//
// Solidity: function batchInbox() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) BatchInbox() (common.Address, error) {
	return _SystemConfigOwnable.Contract.BatchInbox(&_SystemConfigOwnable.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) BatcherHash(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "batcherHash")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) BatcherHash() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.BatcherHash(&_SystemConfigOwnable.CallOpts)
}

// BatcherHash is a free data retrieval call binding the contract method 0xe81b2c6d.
//
// Solidity: function batcherHash() view returns(bytes32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) BatcherHash() ([32]byte, error) {
	return _SystemConfigOwnable.Contract.BatcherHash(&_SystemConfigOwnable.CallOpts)
}

// BlobbasefeeScalar is a free data retrieval call binding the contract method 0xec707517.
//
// Solidity: function blobbasefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) BlobbasefeeScalar(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "blobbasefeeScalar")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BlobbasefeeScalar is a free data retrieval call binding the contract method 0xec707517.
//
// Solidity: function blobbasefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableSession) BlobbasefeeScalar() (uint32, error) {
	return _SystemConfigOwnable.Contract.BlobbasefeeScalar(&_SystemConfigOwnable.CallOpts)
}

// BlobbasefeeScalar is a free data retrieval call binding the contract method 0xec707517.
//
// Solidity: function blobbasefeeScalar() view returns(uint32)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) BlobbasefeeScalar() (uint32, error) {
	return _SystemConfigOwnable.Contract.BlobbasefeeScalar(&_SystemConfigOwnable.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) DisputeGameFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "disputeGameFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) DisputeGameFactory() (common.Address, error) {
	return _SystemConfigOwnable.Contract.DisputeGameFactory(&_SystemConfigOwnable.CallOpts)
}

// DisputeGameFactory is a free data retrieval call binding the contract method 0xf2b4e617.
//
// Solidity: function disputeGameFactory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) DisputeGameFactory() (common.Address, error) {
	return _SystemConfigOwnable.Contract.DisputeGameFactory(&_SystemConfigOwnable.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) GasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "gasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableSession) GasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.GasLimit(&_SystemConfigOwnable.CallOpts)
}

// GasLimit is a free data retrieval call binding the contract method 0xf68016b7.
//
// Solidity: function gasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) GasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.GasLimit(&_SystemConfigOwnable.CallOpts)
}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) GasPayingToken(opts *bind.CallOpts) (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "gasPayingToken")

	outstruct := new(struct {
		Addr     common.Address
		Decimals uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Decimals = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) GasPayingToken() (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	return _SystemConfigOwnable.Contract.GasPayingToken(&_SystemConfigOwnable.CallOpts)
}

// GasPayingToken is a free data retrieval call binding the contract method 0x4397dfef.
//
// Solidity: function gasPayingToken() view returns(address addr_, uint8 decimals_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) GasPayingToken() (struct {
	Addr     common.Address
	Decimals uint8
}, error) {
	return _SystemConfigOwnable.Contract.GasPayingToken(&_SystemConfigOwnable.CallOpts)
}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) GasPayingTokenName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "gasPayingTokenName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) GasPayingTokenName() (string, error) {
	return _SystemConfigOwnable.Contract.GasPayingTokenName(&_SystemConfigOwnable.CallOpts)
}

// GasPayingTokenName is a free data retrieval call binding the contract method 0xd8444715.
//
// Solidity: function gasPayingTokenName() view returns(string name_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) GasPayingTokenName() (string, error) {
	return _SystemConfigOwnable.Contract.GasPayingTokenName(&_SystemConfigOwnable.CallOpts)
}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) GasPayingTokenSymbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "gasPayingTokenSymbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) GasPayingTokenSymbol() (string, error) {
	return _SystemConfigOwnable.Contract.GasPayingTokenSymbol(&_SystemConfigOwnable.CallOpts)
}

// GasPayingTokenSymbol is a free data retrieval call binding the contract method 0x550fcdc9.
//
// Solidity: function gasPayingTokenSymbol() view returns(string symbol_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) GasPayingTokenSymbol() (string, error) {
	return _SystemConfigOwnable.Contract.GasPayingTokenSymbol(&_SystemConfigOwnable.CallOpts)
}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) IsCustomGasToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "isCustomGasToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_SystemConfigOwnable *SystemConfigOwnableSession) IsCustomGasToken() (bool, error) {
	return _SystemConfigOwnable.Contract.IsCustomGasToken(&_SystemConfigOwnable.CallOpts)
}

// IsCustomGasToken is a free data retrieval call binding the contract method 0x21326849.
//
// Solidity: function isCustomGasToken() view returns(bool)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) IsCustomGasToken() (bool, error) {
	return _SystemConfigOwnable.Contract.IsCustomGasToken(&_SystemConfigOwnable.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1CrossDomainMessenger(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "l1CrossDomainMessenger")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1CrossDomainMessenger() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1CrossDomainMessenger(&_SystemConfigOwnable.CallOpts)
}

// L1CrossDomainMessenger is a free data retrieval call binding the contract method 0xa7119869.
//
// Solidity: function l1CrossDomainMessenger() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1CrossDomainMessenger() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1CrossDomainMessenger(&_SystemConfigOwnable.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1ERC721Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "l1ERC721Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1ERC721Bridge() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1ERC721Bridge(&_SystemConfigOwnable.CallOpts)
}

// L1ERC721Bridge is a free data retrieval call binding the contract method 0xc4e8ddfa.
//
// Solidity: function l1ERC721Bridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1ERC721Bridge() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1ERC721Bridge(&_SystemConfigOwnable.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) L1StandardBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "l1StandardBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) L1StandardBridge() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1StandardBridge(&_SystemConfigOwnable.CallOpts)
}

// L1StandardBridge is a free data retrieval call binding the contract method 0x078f29cf.
//
// Solidity: function l1StandardBridge() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) L1StandardBridge() (common.Address, error) {
	return _SystemConfigOwnable.Contract.L1StandardBridge(&_SystemConfigOwnable.CallOpts)
}

// MaximumGasLimit is a free data retrieval call binding the contract method 0x0ae14b1b.
//
// Solidity: function maximumGasLimit() pure returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) MaximumGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "maximumGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MaximumGasLimit is a free data retrieval call binding the contract method 0x0ae14b1b.
//
// Solidity: function maximumGasLimit() pure returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableSession) MaximumGasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.MaximumGasLimit(&_SystemConfigOwnable.CallOpts)
}

// MaximumGasLimit is a free data retrieval call binding the contract method 0x0ae14b1b.
//
// Solidity: function maximumGasLimit() pure returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) MaximumGasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.MaximumGasLimit(&_SystemConfigOwnable.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) MinimumGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "minimumGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableSession) MinimumGasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.MinimumGasLimit(&_SystemConfigOwnable.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0x4add321d.
//
// Solidity: function minimumGasLimit() view returns(uint64)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) MinimumGasLimit() (uint64, error) {
	return _SystemConfigOwnable.Contract.MinimumGasLimit(&_SystemConfigOwnable.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) OptimismMintableERC20Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "optimismMintableERC20Factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _SystemConfigOwnable.Contract.OptimismMintableERC20Factory(&_SystemConfigOwnable.CallOpts)
}

// OptimismMintableERC20Factory is a free data retrieval call binding the contract method 0x9b7d7f0a.
//
// Solidity: function optimismMintableERC20Factory() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) OptimismMintableERC20Factory() (common.Address, error) {
	return _SystemConfigOwnable.Contract.OptimismMintableERC20Factory(&_SystemConfigOwnable.CallOpts)
}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) OptimismPortal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "optimismPortal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) OptimismPortal() (common.Address, error) {
	return _SystemConfigOwnable.Contract.OptimismPortal(&_SystemConfigOwnable.CallOpts)
}

// OptimismPortal is a free data retrieval call binding the contract method 0x0a49cb03.
//
// Solidity: function optimismPortal() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) OptimismPortal() (common.Address, error) {
	return _SystemConfigOwnable.Contract.OptimismPortal(&_SystemConfigOwnable.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) Overhead(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "overhead")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableSession) Overhead() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.Overhead(&_SystemConfigOwnable.CallOpts)
}

// Overhead is a free data retrieval call binding the contract method 0x0c18c162.
//
// Solidity: function overhead() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) Overhead() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.Overhead(&_SystemConfigOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfigOwnable *SystemConfigOwnableSession) Owner() (common.Address, error) {
	return _SystemConfigOwnable.Contract.Owner(&_SystemConfigOwnable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) Owner() (common.Address, error) {
	return _SystemConfigOwnable.Contract.Owner(&_SystemConfigOwnable.CallOpts)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfigOwnable *SystemConfigOwnableCaller) ResourceConfig(opts *bind.CallOpts) (ResourceMeteringResourceConfig, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "resourceConfig")

	if err != nil {
		return *new(ResourceMeteringResourceConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ResourceMeteringResourceConfig)).(*ResourceMeteringResourceConfig)

	return out0, err

}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfigOwnable *SystemConfigOwnableSession) ResourceConfig() (ResourceMeteringResourceConfig, error) {
	return _SystemConfigOwnable.Contract.ResourceConfig(&_SystemConfigOwnable.CallOpts)
}

// ResourceConfig is a free data retrieval call binding the contract method 0xcc731b02.
//
// Solidity: function resourceConfig() view returns((uint32,uint8,uint8,uint32,uint32,uint128))
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) ResourceConfig() (ResourceMeteringResourceConfig, error) {
	return _SystemConfigOwnable.Contract.ResourceConfig(&_SystemConfigOwnable.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) Scalar(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "scalar")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableSession) Scalar() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.Scalar(&_SystemConfigOwnable.CallOpts)
}

// Scalar is a free data retrieval call binding the contract method 0xf45e65d8.
//
// Solidity: function scalar() view returns(uint256)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) Scalar() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.Scalar(&_SystemConfigOwnable.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) StartBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "startBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) StartBlock() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.StartBlock(&_SystemConfigOwnable.CallOpts)
}

// StartBlock is a free data retrieval call binding the contract method 0x48cd4cb1.
//
// Solidity: function startBlock() view returns(uint256 startBlock_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) StartBlock() (*big.Int, error) {
	return _SystemConfigOwnable.Contract.StartBlock(&_SystemConfigOwnable.CallOpts)
}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) UnsafeBlockSigner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "unsafeBlockSigner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableSession) UnsafeBlockSigner() (common.Address, error) {
	return _SystemConfigOwnable.Contract.UnsafeBlockSigner(&_SystemConfigOwnable.CallOpts)
}

// UnsafeBlockSigner is a free data retrieval call binding the contract method 0x1fd19ee1.
//
// Solidity: function unsafeBlockSigner() view returns(address addr_)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) UnsafeBlockSigner() (common.Address, error) {
	return _SystemConfigOwnable.Contract.UnsafeBlockSigner(&_SystemConfigOwnable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemConfigOwnable *SystemConfigOwnableCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SystemConfigOwnable.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemConfigOwnable *SystemConfigOwnableSession) Version() (string, error) {
	return _SystemConfigOwnable.Contract.Version(&_SystemConfigOwnable.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_SystemConfigOwnable *SystemConfigOwnableCallerSession) Version() (string, error) {
	return _SystemConfigOwnable.Contract.Version(&_SystemConfigOwnable.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x43951d62.
//
// Solidity: function initialize(uint32 _basefeeScalar, uint32 _blobbasefeeScalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, (uint32,uint8,uint8,uint32,uint32,uint128) _config, address _batchInbox, (address,address,address,address,address,address,address) _addresses) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) Initialize(opts *bind.TransactOpts, _basefeeScalar uint32, _blobbasefeeScalar uint32, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _config ResourceMeteringResourceConfig, _batchInbox common.Address, _addresses SystemConfigAddresses) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "initialize", _basefeeScalar, _blobbasefeeScalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _config, _batchInbox, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x43951d62.
//
// Solidity: function initialize(uint32 _basefeeScalar, uint32 _blobbasefeeScalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, (uint32,uint8,uint8,uint32,uint32,uint128) _config, address _batchInbox, (address,address,address,address,address,address,address) _addresses) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) Initialize(_basefeeScalar uint32, _blobbasefeeScalar uint32, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _config ResourceMeteringResourceConfig, _batchInbox common.Address, _addresses SystemConfigAddresses) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.Initialize(&_SystemConfigOwnable.TransactOpts, _basefeeScalar, _blobbasefeeScalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _config, _batchInbox, _addresses)
}

// Initialize is a paid mutator transaction binding the contract method 0x43951d62.
//
// Solidity: function initialize(uint32 _basefeeScalar, uint32 _blobbasefeeScalar, bytes32 _batcherHash, uint64 _gasLimit, address _unsafeBlockSigner, (uint32,uint8,uint8,uint32,uint32,uint128) _config, address _batchInbox, (address,address,address,address,address,address,address) _addresses) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) Initialize(_basefeeScalar uint32, _blobbasefeeScalar uint32, _batcherHash [32]byte, _gasLimit uint64, _unsafeBlockSigner common.Address, _config ResourceMeteringResourceConfig, _batchInbox common.Address, _addresses SystemConfigAddresses) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.Initialize(&_SystemConfigOwnable.TransactOpts, _basefeeScalar, _blobbasefeeScalar, _batcherHash, _gasLimit, _unsafeBlockSigner, _config, _batchInbox, _addresses)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) SetBatcherHash(opts *bind.TransactOpts, _batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "setBatcherHash", _batcherHash)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) SetBatcherHash(_batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetBatcherHash(&_SystemConfigOwnable.TransactOpts, _batcherHash)
}

// SetBatcherHash is a paid mutator transaction binding the contract method 0xc9b26f61.
//
// Solidity: function setBatcherHash(bytes32 _batcherHash) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) SetBatcherHash(_batcherHash [32]byte) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetBatcherHash(&_SystemConfigOwnable.TransactOpts, _batcherHash)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) SetGasConfig(opts *bind.TransactOpts, _overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "setGasConfig", _overhead, _scalar)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) SetGasConfig(_overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasConfig(&_SystemConfigOwnable.TransactOpts, _overhead, _scalar)
}

// SetGasConfig is a paid mutator transaction binding the contract method 0x935f029e.
//
// Solidity: function setGasConfig(uint256 _overhead, uint256 _scalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) SetGasConfig(_overhead *big.Int, _scalar *big.Int) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasConfig(&_SystemConfigOwnable.TransactOpts, _overhead, _scalar)
}

// SetGasConfigEcotone is a paid mutator transaction binding the contract method 0x21d7fde5.
//
// Solidity: function setGasConfigEcotone(uint32 _basefeeScalar, uint32 _blobbasefeeScalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) SetGasConfigEcotone(opts *bind.TransactOpts, _basefeeScalar uint32, _blobbasefeeScalar uint32) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "setGasConfigEcotone", _basefeeScalar, _blobbasefeeScalar)
}

// SetGasConfigEcotone is a paid mutator transaction binding the contract method 0x21d7fde5.
//
// Solidity: function setGasConfigEcotone(uint32 _basefeeScalar, uint32 _blobbasefeeScalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) SetGasConfigEcotone(_basefeeScalar uint32, _blobbasefeeScalar uint32) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasConfigEcotone(&_SystemConfigOwnable.TransactOpts, _basefeeScalar, _blobbasefeeScalar)
}

// SetGasConfigEcotone is a paid mutator transaction binding the contract method 0x21d7fde5.
//
// Solidity: function setGasConfigEcotone(uint32 _basefeeScalar, uint32 _blobbasefeeScalar) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) SetGasConfigEcotone(_basefeeScalar uint32, _blobbasefeeScalar uint32) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasConfigEcotone(&_SystemConfigOwnable.TransactOpts, _basefeeScalar, _blobbasefeeScalar)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) SetGasLimit(opts *bind.TransactOpts, _gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "setGasLimit", _gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) SetGasLimit(_gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasLimit(&_SystemConfigOwnable.TransactOpts, _gasLimit)
}

// SetGasLimit is a paid mutator transaction binding the contract method 0xb40a817c.
//
// Solidity: function setGasLimit(uint64 _gasLimit) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) SetGasLimit(_gasLimit uint64) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetGasLimit(&_SystemConfigOwnable.TransactOpts, _gasLimit)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactor) SetUnsafeBlockSigner(opts *bind.TransactOpts, _unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfigOwnable.contract.Transact(opts, "setUnsafeBlockSigner", _unsafeBlockSigner)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfigOwnable *SystemConfigOwnableSession) SetUnsafeBlockSigner(_unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetUnsafeBlockSigner(&_SystemConfigOwnable.TransactOpts, _unsafeBlockSigner)
}

// SetUnsafeBlockSigner is a paid mutator transaction binding the contract method 0x18d13918.
//
// Solidity: function setUnsafeBlockSigner(address _unsafeBlockSigner) returns()
func (_SystemConfigOwnable *SystemConfigOwnableTransactorSession) SetUnsafeBlockSigner(_unsafeBlockSigner common.Address) (*types.Transaction, error) {
	return _SystemConfigOwnable.Contract.SetUnsafeBlockSigner(&_SystemConfigOwnable.TransactOpts, _unsafeBlockSigner)
}

// SystemConfigOwnableConfigUpdateIterator is returned from FilterConfigUpdate and is used to iterate over the raw logs and unpacked data for ConfigUpdate events raised by the SystemConfigOwnable contract.
type SystemConfigOwnableConfigUpdateIterator struct {
	Event *SystemConfigOwnableConfigUpdate // Event containing the contract specifics and raw log

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
func (it *SystemConfigOwnableConfigUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemConfigOwnableConfigUpdate)
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
		it.Event = new(SystemConfigOwnableConfigUpdate)
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
func (it *SystemConfigOwnableConfigUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemConfigOwnableConfigUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemConfigOwnableConfigUpdate represents a ConfigUpdate event raised by the SystemConfigOwnable contract.
type SystemConfigOwnableConfigUpdate struct {
	Version    *big.Int
	UpdateType uint8
	Data       []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigUpdate is a free log retrieval operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) FilterConfigUpdate(opts *bind.FilterOpts, version []*big.Int, updateType []uint8) (*SystemConfigOwnableConfigUpdateIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _SystemConfigOwnable.contract.FilterLogs(opts, "ConfigUpdate", versionRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnableConfigUpdateIterator{contract: _SystemConfigOwnable.contract, event: "ConfigUpdate", logs: logs, sub: sub}, nil
}

// WatchConfigUpdate is a free log subscription operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) WatchConfigUpdate(opts *bind.WatchOpts, sink chan<- *SystemConfigOwnableConfigUpdate, version []*big.Int, updateType []uint8) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var updateTypeRule []interface{}
	for _, updateTypeItem := range updateType {
		updateTypeRule = append(updateTypeRule, updateTypeItem)
	}

	logs, sub, err := _SystemConfigOwnable.contract.WatchLogs(opts, "ConfigUpdate", versionRule, updateTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemConfigOwnableConfigUpdate)
				if err := _SystemConfigOwnable.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
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

// ParseConfigUpdate is a log parse operation binding the contract event 0x1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be.
//
// Solidity: event ConfigUpdate(uint256 indexed version, uint8 indexed updateType, bytes data)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) ParseConfigUpdate(log types.Log) (*SystemConfigOwnableConfigUpdate, error) {
	event := new(SystemConfigOwnableConfigUpdate)
	if err := _SystemConfigOwnable.contract.UnpackLog(event, "ConfigUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SystemConfigOwnableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SystemConfigOwnable contract.
type SystemConfigOwnableInitializedIterator struct {
	Event *SystemConfigOwnableInitialized // Event containing the contract specifics and raw log

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
func (it *SystemConfigOwnableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SystemConfigOwnableInitialized)
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
		it.Event = new(SystemConfigOwnableInitialized)
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
func (it *SystemConfigOwnableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SystemConfigOwnableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SystemConfigOwnableInitialized represents a Initialized event raised by the SystemConfigOwnable contract.
type SystemConfigOwnableInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) FilterInitialized(opts *bind.FilterOpts) (*SystemConfigOwnableInitializedIterator, error) {

	logs, sub, err := _SystemConfigOwnable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SystemConfigOwnableInitializedIterator{contract: _SystemConfigOwnable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SystemConfigOwnableInitialized) (event.Subscription, error) {

	logs, sub, err := _SystemConfigOwnable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SystemConfigOwnableInitialized)
				if err := _SystemConfigOwnable.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_SystemConfigOwnable *SystemConfigOwnableFilterer) ParseInitialized(log types.Log) (*SystemConfigOwnableInitialized, error) {
	event := new(SystemConfigOwnableInitialized)
	if err := _SystemConfigOwnable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
