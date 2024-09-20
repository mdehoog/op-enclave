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

// TypesOutputRootProof is an auto generated low-level Go binding around an user-defined struct.
type TypesOutputRootProof struct {
	Version                  [32]byte
	StateRoot                [32]byte
	MessagePasserStorageRoot [32]byte
	LatestBlockhash          [32]byte
}

// TypesWithdrawalTransaction is an auto generated low-level Go binding around an user-defined struct.
type TypesWithdrawalTransaction struct {
	Nonce    *big.Int
	Sender   common.Address
	Target   common.Address
	Value    *big.Int
	GasLimit *big.Int
	Data     []byte
}

// PortalMetaData contains all meta data concerning the Portal contract.
var PortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"balance\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositERC20Transaction\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_mint\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_isCreation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"depositTransaction\",\"inputs\":[{\"name\":\"_to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"_isCreation\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"donateETH\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"finalizeWithdrawalTransaction\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structTypes.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizedWithdrawals\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"guardian\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"_l2Oracle\",\"type\":\"address\",\"internalType\":\"contractOutputOracle\"},{\"name\":\"_systemConfig\",\"type\":\"address\",\"internalType\":\"contractSystemConfig\"},{\"name\":\"_superchainConfig\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isOutputFinalized\",\"inputs\":[{\"name\":\"_l2OutputIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2Oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutputOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l2Sender\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minimumGasLimit\",\"inputs\":[{\"name\":\"_byteCount\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"params\",\"inputs\":[],\"outputs\":[{\"name\":\"prevBaseFee\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"prevBoughtGas\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"prevBlockNum\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"paused_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proveAndFinalizeWithdrawalTransaction\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structTypes.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_l2OutputIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proveWithdrawalTransaction\",\"inputs\":[{\"name\":\"_tx\",\"type\":\"tuple\",\"internalType\":\"structTypes.WithdrawalTransaction\",\"components\":[{\"name\":\"nonce\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"_l2OutputIndex\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_outputRootProof\",\"type\":\"tuple\",\"internalType\":\"structTypes.OutputRootProof\",\"components\":[{\"name\":\"version\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"stateRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"messagePasserStorageRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"latestBlockhash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"_withdrawalProof\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPayingToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_decimals\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"_name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_symbol\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"superchainConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSuperchainConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"systemConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractSystemConfig\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransactionDeposited\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"version\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"opaqueData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalFinalized\",\"inputs\":[{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WithdrawalProven\",\"inputs\":[{\"name\":\"withdrawalHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BadTarget\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallPaused\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ContentLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EmptyItem\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"GasEstimation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDataRemainder\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidHeader\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"LargeCalldata\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoValue\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NonReentrant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OnlyCustomGasToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OutOfGas\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SmallGasLimit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedList\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UnexpectedString\",\"inputs\":[]}]",
}

// PortalABI is the input ABI used to generate the binding from.
// Deprecated: Use PortalMetaData.ABI instead.
var PortalABI = PortalMetaData.ABI

// Portal is an auto generated Go binding around an Ethereum contract.
type Portal struct {
	PortalCaller     // Read-only binding to the contract
	PortalTransactor // Write-only binding to the contract
	PortalFilterer   // Log filterer for contract events
}

// PortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortalSession struct {
	Contract     *Portal           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortalCallerSession struct {
	Contract *PortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortalTransactorSession struct {
	Contract     *PortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortalRaw struct {
	Contract *Portal // Generic contract binding to access the raw methods on
}

// PortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortalCallerRaw struct {
	Contract *PortalCaller // Generic read-only contract binding to access the raw methods on
}

// PortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortalTransactorRaw struct {
	Contract *PortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortal creates a new instance of Portal, bound to a specific deployed contract.
func NewPortal(address common.Address, backend bind.ContractBackend) (*Portal, error) {
	contract, err := bindPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Portal{PortalCaller: PortalCaller{contract: contract}, PortalTransactor: PortalTransactor{contract: contract}, PortalFilterer: PortalFilterer{contract: contract}}, nil
}

// NewPortalCaller creates a new read-only instance of Portal, bound to a specific deployed contract.
func NewPortalCaller(address common.Address, caller bind.ContractCaller) (*PortalCaller, error) {
	contract, err := bindPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortalCaller{contract: contract}, nil
}

// NewPortalTransactor creates a new write-only instance of Portal, bound to a specific deployed contract.
func NewPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*PortalTransactor, error) {
	contract, err := bindPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortalTransactor{contract: contract}, nil
}

// NewPortalFilterer creates a new log filterer instance of Portal, bound to a specific deployed contract.
func NewPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*PortalFilterer, error) {
	contract, err := bindPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortalFilterer{contract: contract}, nil
}

// bindPortal binds a generic wrapper to an already deployed contract.
func bindPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PortalABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.PortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.PortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Portal *PortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Portal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Portal *PortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Portal *PortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Portal.Contract.contract.Transact(opts, method, params...)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Portal *PortalCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Portal *PortalSession) Balance() (*big.Int, error) {
	return _Portal.Contract.Balance(&_Portal.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Portal *PortalCallerSession) Balance() (*big.Int, error) {
	return _Portal.Contract.Balance(&_Portal.CallOpts)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Portal *PortalCaller) FinalizedWithdrawals(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "finalizedWithdrawals", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Portal *PortalSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _Portal.Contract.FinalizedWithdrawals(&_Portal.CallOpts, arg0)
}

// FinalizedWithdrawals is a free data retrieval call binding the contract method 0xa14238e7.
//
// Solidity: function finalizedWithdrawals(bytes32 ) view returns(bool)
func (_Portal *PortalCallerSession) FinalizedWithdrawals(arg0 [32]byte) (bool, error) {
	return _Portal.Contract.FinalizedWithdrawals(&_Portal.CallOpts, arg0)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Portal *PortalCaller) Guardian(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "guardian")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Portal *PortalSession) Guardian() (common.Address, error) {
	return _Portal.Contract.Guardian(&_Portal.CallOpts)
}

// Guardian is a free data retrieval call binding the contract method 0x452a9320.
//
// Solidity: function guardian() view returns(address)
func (_Portal *PortalCallerSession) Guardian() (common.Address, error) {
	return _Portal.Contract.Guardian(&_Portal.CallOpts)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_Portal *PortalCaller) IsOutputFinalized(opts *bind.CallOpts, _l2OutputIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "isOutputFinalized", _l2OutputIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_Portal *PortalSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _Portal.Contract.IsOutputFinalized(&_Portal.CallOpts, _l2OutputIndex)
}

// IsOutputFinalized is a free data retrieval call binding the contract method 0x6dbffb78.
//
// Solidity: function isOutputFinalized(uint256 _l2OutputIndex) view returns(bool)
func (_Portal *PortalCallerSession) IsOutputFinalized(_l2OutputIndex *big.Int) (bool, error) {
	return _Portal.Contract.IsOutputFinalized(&_Portal.CallOpts, _l2OutputIndex)
}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_Portal *PortalCaller) L2Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "l2Oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_Portal *PortalSession) L2Oracle() (common.Address, error) {
	return _Portal.Contract.L2Oracle(&_Portal.CallOpts)
}

// L2Oracle is a free data retrieval call binding the contract method 0x9b5f694a.
//
// Solidity: function l2Oracle() view returns(address)
func (_Portal *PortalCallerSession) L2Oracle() (common.Address, error) {
	return _Portal.Contract.L2Oracle(&_Portal.CallOpts)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_Portal *PortalCaller) L2Sender(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "l2Sender")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_Portal *PortalSession) L2Sender() (common.Address, error) {
	return _Portal.Contract.L2Sender(&_Portal.CallOpts)
}

// L2Sender is a free data retrieval call binding the contract method 0x9bf62d82.
//
// Solidity: function l2Sender() view returns(address)
func (_Portal *PortalCallerSession) L2Sender() (common.Address, error) {
	return _Portal.Contract.L2Sender(&_Portal.CallOpts)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_Portal *PortalCaller) MinimumGasLimit(opts *bind.CallOpts, _byteCount uint64) (uint64, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "minimumGasLimit", _byteCount)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_Portal *PortalSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _Portal.Contract.MinimumGasLimit(&_Portal.CallOpts, _byteCount)
}

// MinimumGasLimit is a free data retrieval call binding the contract method 0xa35d99df.
//
// Solidity: function minimumGasLimit(uint64 _byteCount) pure returns(uint64)
func (_Portal *PortalCallerSession) MinimumGasLimit(_byteCount uint64) (uint64, error) {
	return _Portal.Contract.MinimumGasLimit(&_Portal.CallOpts, _byteCount)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_Portal *PortalCaller) Params(opts *bind.CallOpts) (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "params")

	outstruct := new(struct {
		PrevBaseFee   *big.Int
		PrevBoughtGas uint64
		PrevBlockNum  uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PrevBaseFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PrevBoughtGas = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.PrevBlockNum = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_Portal *PortalSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _Portal.Contract.Params(&_Portal.CallOpts)
}

// Params is a free data retrieval call binding the contract method 0xcff0ab96.
//
// Solidity: function params() view returns(uint128 prevBaseFee, uint64 prevBoughtGas, uint64 prevBlockNum)
func (_Portal *PortalCallerSession) Params() (struct {
	PrevBaseFee   *big.Int
	PrevBoughtGas uint64
	PrevBlockNum  uint64
}, error) {
	return _Portal.Contract.Params(&_Portal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_Portal *PortalCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_Portal *PortalSession) Paused() (bool, error) {
	return _Portal.Contract.Paused(&_Portal.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool paused_)
func (_Portal *PortalCallerSession) Paused() (bool, error) {
	return _Portal.Contract.Paused(&_Portal.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_Portal *PortalCaller) SuperchainConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "superchainConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_Portal *PortalSession) SuperchainConfig() (common.Address, error) {
	return _Portal.Contract.SuperchainConfig(&_Portal.CallOpts)
}

// SuperchainConfig is a free data retrieval call binding the contract method 0x35e80ab3.
//
// Solidity: function superchainConfig() view returns(address)
func (_Portal *PortalCallerSession) SuperchainConfig() (common.Address, error) {
	return _Portal.Contract.SuperchainConfig(&_Portal.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_Portal *PortalCaller) SystemConfig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "systemConfig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_Portal *PortalSession) SystemConfig() (common.Address, error) {
	return _Portal.Contract.SystemConfig(&_Portal.CallOpts)
}

// SystemConfig is a free data retrieval call binding the contract method 0x33d7e2bd.
//
// Solidity: function systemConfig() view returns(address)
func (_Portal *PortalCallerSession) SystemConfig() (common.Address, error) {
	return _Portal.Contract.SystemConfig(&_Portal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Portal *PortalCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Portal.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Portal *PortalSession) Version() (string, error) {
	return _Portal.Contract.Version(&_Portal.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(string)
func (_Portal *PortalCallerSession) Version() (string, error) {
	return _Portal.Contract.Version(&_Portal.CallOpts)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_Portal *PortalTransactor) DepositERC20Transaction(opts *bind.TransactOpts, _to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "depositERC20Transaction", _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_Portal *PortalSession) DepositERC20Transaction(_to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.Contract.DepositERC20Transaction(&_Portal.TransactOpts, _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositERC20Transaction is a paid mutator transaction binding the contract method 0x149f2f22.
//
// Solidity: function depositERC20Transaction(address _to, uint256 _mint, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) returns()
func (_Portal *PortalTransactorSession) DepositERC20Transaction(_to common.Address, _mint *big.Int, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.Contract.DepositERC20Transaction(&_Portal.TransactOpts, _to, _mint, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_Portal *PortalTransactor) DepositTransaction(opts *bind.TransactOpts, _to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "depositTransaction", _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_Portal *PortalSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.Contract.DepositTransaction(&_Portal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DepositTransaction is a paid mutator transaction binding the contract method 0xe9e05c42.
//
// Solidity: function depositTransaction(address _to, uint256 _value, uint64 _gasLimit, bool _isCreation, bytes _data) payable returns()
func (_Portal *PortalTransactorSession) DepositTransaction(_to common.Address, _value *big.Int, _gasLimit uint64, _isCreation bool, _data []byte) (*types.Transaction, error) {
	return _Portal.Contract.DepositTransaction(&_Portal.TransactOpts, _to, _value, _gasLimit, _isCreation, _data)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_Portal *PortalTransactor) DonateETH(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "donateETH")
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_Portal *PortalSession) DonateETH() (*types.Transaction, error) {
	return _Portal.Contract.DonateETH(&_Portal.TransactOpts)
}

// DonateETH is a paid mutator transaction binding the contract method 0x8b4c40b0.
//
// Solidity: function donateETH() payable returns()
func (_Portal *PortalTransactorSession) DonateETH() (*types.Transaction, error) {
	return _Portal.Contract.DonateETH(&_Portal.TransactOpts)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_Portal *PortalTransactor) FinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "finalizeWithdrawalTransaction", _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_Portal *PortalSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _Portal.Contract.FinalizeWithdrawalTransaction(&_Portal.TransactOpts, _tx)
}

// FinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x8c3152e9.
//
// Solidity: function finalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx) returns()
func (_Portal *PortalTransactorSession) FinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction) (*types.Transaction, error) {
	return _Portal.Contract.FinalizeWithdrawalTransaction(&_Portal.TransactOpts, _tx)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l2Oracle, address _systemConfig, address _superchainConfig) returns()
func (_Portal *PortalTransactor) Initialize(opts *bind.TransactOpts, _l2Oracle common.Address, _systemConfig common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "initialize", _l2Oracle, _systemConfig, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l2Oracle, address _systemConfig, address _superchainConfig) returns()
func (_Portal *PortalSession) Initialize(_l2Oracle common.Address, _systemConfig common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Initialize(&_Portal.TransactOpts, _l2Oracle, _systemConfig, _superchainConfig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc0c53b8b.
//
// Solidity: function initialize(address _l2Oracle, address _systemConfig, address _superchainConfig) returns()
func (_Portal *PortalTransactorSession) Initialize(_l2Oracle common.Address, _systemConfig common.Address, _superchainConfig common.Address) (*types.Transaction, error) {
	return _Portal.Contract.Initialize(&_Portal.TransactOpts, _l2Oracle, _systemConfig, _superchainConfig)
}

// ProveAndFinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x47f55db5.
//
// Solidity: function proveAndFinalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalTransactor) ProveAndFinalizeWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "proveAndFinalizeWithdrawalTransaction", _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveAndFinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x47f55db5.
//
// Solidity: function proveAndFinalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalSession) ProveAndFinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.Contract.ProveAndFinalizeWithdrawalTransaction(&_Portal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveAndFinalizeWithdrawalTransaction is a paid mutator transaction binding the contract method 0x47f55db5.
//
// Solidity: function proveAndFinalizeWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalTransactorSession) ProveAndFinalizeWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.Contract.ProveAndFinalizeWithdrawalTransaction(&_Portal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalTransactor) ProveWithdrawalTransaction(opts *bind.TransactOpts, _tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "proveWithdrawalTransaction", _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.Contract.ProveWithdrawalTransaction(&_Portal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// ProveWithdrawalTransaction is a paid mutator transaction binding the contract method 0x4870496f.
//
// Solidity: function proveWithdrawalTransaction((uint256,address,address,uint256,uint256,bytes) _tx, uint256 _l2OutputIndex, (bytes32,bytes32,bytes32,bytes32) _outputRootProof, bytes[] _withdrawalProof) returns()
func (_Portal *PortalTransactorSession) ProveWithdrawalTransaction(_tx TypesWithdrawalTransaction, _l2OutputIndex *big.Int, _outputRootProof TypesOutputRootProof, _withdrawalProof [][]byte) (*types.Transaction, error) {
	return _Portal.Contract.ProveWithdrawalTransaction(&_Portal.TransactOpts, _tx, _l2OutputIndex, _outputRootProof, _withdrawalProof)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_Portal *PortalTransactor) SetGasPayingToken(opts *bind.TransactOpts, _token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _Portal.contract.Transact(opts, "setGasPayingToken", _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_Portal *PortalSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.SetGasPayingToken(&_Portal.TransactOpts, _token, _decimals, _name, _symbol)
}

// SetGasPayingToken is a paid mutator transaction binding the contract method 0x71cfaa3f.
//
// Solidity: function setGasPayingToken(address _token, uint8 _decimals, bytes32 _name, bytes32 _symbol) returns()
func (_Portal *PortalTransactorSession) SetGasPayingToken(_token common.Address, _decimals uint8, _name [32]byte, _symbol [32]byte) (*types.Transaction, error) {
	return _Portal.Contract.SetGasPayingToken(&_Portal.TransactOpts, _token, _decimals, _name, _symbol)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portal *PortalTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Portal.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portal *PortalSession) Receive() (*types.Transaction, error) {
	return _Portal.Contract.Receive(&_Portal.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Portal *PortalTransactorSession) Receive() (*types.Transaction, error) {
	return _Portal.Contract.Receive(&_Portal.TransactOpts)
}

// PortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Portal contract.
type PortalInitializedIterator struct {
	Event *PortalInitialized // Event containing the contract specifics and raw log

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
func (it *PortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalInitialized)
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
		it.Event = new(PortalInitialized)
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
func (it *PortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalInitialized represents a Initialized event raised by the Portal contract.
type PortalInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Portal *PortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*PortalInitializedIterator, error) {

	logs, sub, err := _Portal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PortalInitializedIterator{contract: _Portal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Portal *PortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PortalInitialized) (event.Subscription, error) {

	logs, sub, err := _Portal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalInitialized)
				if err := _Portal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Portal *PortalFilterer) ParseInitialized(log types.Log) (*PortalInitialized, error) {
	event := new(PortalInitialized)
	if err := _Portal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalTransactionDepositedIterator is returned from FilterTransactionDeposited and is used to iterate over the raw logs and unpacked data for TransactionDeposited events raised by the Portal contract.
type PortalTransactionDepositedIterator struct {
	Event *PortalTransactionDeposited // Event containing the contract specifics and raw log

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
func (it *PortalTransactionDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalTransactionDeposited)
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
		it.Event = new(PortalTransactionDeposited)
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
func (it *PortalTransactionDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalTransactionDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalTransactionDeposited represents a TransactionDeposited event raised by the Portal contract.
type PortalTransactionDeposited struct {
	From       common.Address
	To         common.Address
	Version    *big.Int
	OpaqueData []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTransactionDeposited is a free log retrieval operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_Portal *PortalFilterer) FilterTransactionDeposited(opts *bind.FilterOpts, from []common.Address, to []common.Address, version []*big.Int) (*PortalTransactionDepositedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return &PortalTransactionDepositedIterator{contract: _Portal.contract, event: "TransactionDeposited", logs: logs, sub: sub}, nil
}

// WatchTransactionDeposited is a free log subscription operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_Portal *PortalFilterer) WatchTransactionDeposited(opts *bind.WatchOpts, sink chan<- *PortalTransactionDeposited, from []common.Address, to []common.Address, version []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "TransactionDeposited", fromRule, toRule, versionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalTransactionDeposited)
				if err := _Portal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
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

// ParseTransactionDeposited is a log parse operation binding the contract event 0xb3813568d9991fc951961fcb4c784893574240a28925604d09fc577c55bb7c32.
//
// Solidity: event TransactionDeposited(address indexed from, address indexed to, uint256 indexed version, bytes opaqueData)
func (_Portal *PortalFilterer) ParseTransactionDeposited(log types.Log) (*PortalTransactionDeposited, error) {
	event := new(PortalTransactionDeposited)
	if err := _Portal.contract.UnpackLog(event, "TransactionDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalWithdrawalFinalizedIterator is returned from FilterWithdrawalFinalized and is used to iterate over the raw logs and unpacked data for WithdrawalFinalized events raised by the Portal contract.
type PortalWithdrawalFinalizedIterator struct {
	Event *PortalWithdrawalFinalized // Event containing the contract specifics and raw log

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
func (it *PortalWithdrawalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalWithdrawalFinalized)
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
		it.Event = new(PortalWithdrawalFinalized)
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
func (it *PortalWithdrawalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalWithdrawalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalWithdrawalFinalized represents a WithdrawalFinalized event raised by the Portal contract.
type PortalWithdrawalFinalized struct {
	WithdrawalHash [32]byte
	Success        bool
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalFinalized is a free log retrieval operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_Portal *PortalFilterer) FilterWithdrawalFinalized(opts *bind.FilterOpts, withdrawalHash [][32]byte) (*PortalWithdrawalFinalizedIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return &PortalWithdrawalFinalizedIterator{contract: _Portal.contract, event: "WithdrawalFinalized", logs: logs, sub: sub}, nil
}

// WatchWithdrawalFinalized is a free log subscription operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_Portal *PortalFilterer) WatchWithdrawalFinalized(opts *bind.WatchOpts, sink chan<- *PortalWithdrawalFinalized, withdrawalHash [][32]byte) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "WithdrawalFinalized", withdrawalHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalWithdrawalFinalized)
				if err := _Portal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
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

// ParseWithdrawalFinalized is a log parse operation binding the contract event 0xdb5c7652857aa163daadd670e116628fb42e869d8ac4251ef8971d9e5727df1b.
//
// Solidity: event WithdrawalFinalized(bytes32 indexed withdrawalHash, bool success)
func (_Portal *PortalFilterer) ParseWithdrawalFinalized(log types.Log) (*PortalWithdrawalFinalized, error) {
	event := new(PortalWithdrawalFinalized)
	if err := _Portal.contract.UnpackLog(event, "WithdrawalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalWithdrawalProvenIterator is returned from FilterWithdrawalProven and is used to iterate over the raw logs and unpacked data for WithdrawalProven events raised by the Portal contract.
type PortalWithdrawalProvenIterator struct {
	Event *PortalWithdrawalProven // Event containing the contract specifics and raw log

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
func (it *PortalWithdrawalProvenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalWithdrawalProven)
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
		it.Event = new(PortalWithdrawalProven)
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
func (it *PortalWithdrawalProvenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalWithdrawalProvenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalWithdrawalProven represents a WithdrawalProven event raised by the Portal contract.
type PortalWithdrawalProven struct {
	WithdrawalHash [32]byte
	From           common.Address
	To             common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalProven is a free log retrieval operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_Portal *PortalFilterer) FilterWithdrawalProven(opts *bind.FilterOpts, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (*PortalWithdrawalProvenIterator, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Portal.contract.FilterLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PortalWithdrawalProvenIterator{contract: _Portal.contract, event: "WithdrawalProven", logs: logs, sub: sub}, nil
}

// WatchWithdrawalProven is a free log subscription operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_Portal *PortalFilterer) WatchWithdrawalProven(opts *bind.WatchOpts, sink chan<- *PortalWithdrawalProven, withdrawalHash [][32]byte, from []common.Address, to []common.Address) (event.Subscription, error) {

	var withdrawalHashRule []interface{}
	for _, withdrawalHashItem := range withdrawalHash {
		withdrawalHashRule = append(withdrawalHashRule, withdrawalHashItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Portal.contract.WatchLogs(opts, "WithdrawalProven", withdrawalHashRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalWithdrawalProven)
				if err := _Portal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
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

// ParseWithdrawalProven is a log parse operation binding the contract event 0x67a6208cfcc0801d50f6cbe764733f4fddf66ac0b04442061a8a8c0cb6b63f62.
//
// Solidity: event WithdrawalProven(bytes32 indexed withdrawalHash, address indexed from, address indexed to)
func (_Portal *PortalFilterer) ParseWithdrawalProven(log types.Log) (*PortalWithdrawalProven, error) {
	event := new(PortalWithdrawalProven)
	if err := _Portal.contract.UnpackLog(event, "WithdrawalProven", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
