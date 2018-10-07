// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// ContractABI is the input ABI used to generate the binding from.
const ContractABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pixelByCoordinate\",\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"price\",\"type\":\"uint256\"},{\"name\":\"auctionId\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InterfaceId_ERC165\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes4\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"lastPaidTaxes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"auctionById\",\"outputs\":[{\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"name\":\"blockId\",\"type\":\"bytes32\"},{\"name\":\"x\",\"type\":\"uint256\"},{\"name\":\"y\",\"type\":\"uint256\"},{\"name\":\"currentPrice\",\"type\":\"uint256\"},{\"name\":\"currentLeader\",\"type\":\"address\"},{\"name\":\"endTime\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"userBalanceAtLastPaid\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"valueHeld\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"xMax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"addFunds\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxPercentage\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"taxCollector\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"transferTaxes\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"user\",\"type\":\"address\"}],\"name\":\"userHasPositveBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"yMax\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_xMax\",\"type\":\"uint256\"},{\"name\":\"_yMax\",\"type\":\"uint256\"},{\"name\":\"_taxPercentage\",\"type\":\"uint256\"},{\"name\":\"_taxCollector\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"buyer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"BuyPixel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"seller\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"SetPixelPrice\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pixelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"BeginDutchAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pixelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"auctionId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amountBet\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"timeBet\",\"type\":\"uint256\"}],\"name\":\"UpdateAuctionBid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"pixelId\",\"type\":\"bytes32\"},{\"indexed\":true,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"claimer\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"x\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"y\",\"type\":\"uint256\"}],\"name\":\"EndDutchAuction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_approved\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_operator\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"buyUninitializedPixelBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256[]\"},{\"name\":\"_y\",\"type\":\"uint256[]\"},{\"name\":\"_price\",\"type\":\"uint256[]\"}],\"name\":\"buyUninitializedPixelBlocks\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"buyPixelBlock\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256[]\"},{\"name\":\"_y\",\"type\":\"uint256[]\"},{\"name\":\"_price\",\"type\":\"uint256[]\"}],\"name\":\"buyPixelBlocks\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"setPixelBlockPrice\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256[]\"},{\"name\":\"_y\",\"type\":\"uint256[]\"},{\"name\":\"_price\",\"type\":\"uint256[]\"}],\"name\":\"setPixelBlockPrices\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"beginDutchAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"},{\"name\":\"_bid\",\"type\":\"uint256\"}],\"name\":\"bidInAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"endDutchAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_x\",\"type\":\"uint256\"},{\"name\":\"_y\",\"type\":\"uint256\"}],\"name\":\"encodeTokenId\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Contract *ContractCaller) InterfaceIdERC165(opts *bind.CallOpts) ([4]byte, error) {
	var (
		ret0 = new([4]byte)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "InterfaceId_ERC165")
	return *ret0, err
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Contract *ContractSession) InterfaceIdERC165() ([4]byte, error) {
	return _Contract.Contract.InterfaceIdERC165(&_Contract.CallOpts)
}

// InterfaceIdERC165 is a free data retrieval call binding the contract method 0x19fa8f50.
//
// Solidity: function InterfaceId_ERC165() constant returns(bytes4)
func (_Contract *ContractCallerSession) InterfaceIdERC165() ([4]byte, error) {
	return _Contract.Contract.InterfaceIdERC165(&_Contract.CallOpts)
}

// AuctionById is a free data retrieval call binding the contract method 0x2353f37b.
//
// Solidity: function auctionById( bytes32) constant returns(auctionId bytes32, blockId bytes32, x uint256, y uint256, currentPrice uint256, currentLeader address, endTime uint256)
func (_Contract *ContractCaller) AuctionById(opts *bind.CallOpts, arg0 [32]byte) (struct {
	AuctionId     [32]byte
	BlockId       [32]byte
	X             *big.Int
	Y             *big.Int
	CurrentPrice  *big.Int
	CurrentLeader common.Address
	EndTime       *big.Int
}, error) {
	ret := new(struct {
		AuctionId     [32]byte
		BlockId       [32]byte
		X             *big.Int
		Y             *big.Int
		CurrentPrice  *big.Int
		CurrentLeader common.Address
		EndTime       *big.Int
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "auctionById", arg0)
	return *ret, err
}

// AuctionById is a free data retrieval call binding the contract method 0x2353f37b.
//
// Solidity: function auctionById( bytes32) constant returns(auctionId bytes32, blockId bytes32, x uint256, y uint256, currentPrice uint256, currentLeader address, endTime uint256)
func (_Contract *ContractSession) AuctionById(arg0 [32]byte) (struct {
	AuctionId     [32]byte
	BlockId       [32]byte
	X             *big.Int
	Y             *big.Int
	CurrentPrice  *big.Int
	CurrentLeader common.Address
	EndTime       *big.Int
}, error) {
	return _Contract.Contract.AuctionById(&_Contract.CallOpts, arg0)
}

// AuctionById is a free data retrieval call binding the contract method 0x2353f37b.
//
// Solidity: function auctionById( bytes32) constant returns(auctionId bytes32, blockId bytes32, x uint256, y uint256, currentPrice uint256, currentLeader address, endTime uint256)
func (_Contract *ContractCallerSession) AuctionById(arg0 [32]byte) (struct {
	AuctionId     [32]byte
	BlockId       [32]byte
	X             *big.Int
	Y             *big.Int
	CurrentPrice  *big.Int
	CurrentLeader common.Address
	EndTime       *big.Int
}, error) {
	return _Contract.Contract.AuctionById(&_Contract.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contract *ContractCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contract *ContractSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(uint256)
func (_Contract *ContractCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _Contract.Contract.BalanceOf(&_Contract.CallOpts, _owner)
}

// EncodeTokenId is a free data retrieval call binding the contract method 0xebd46d64.
//
// Solidity: function encodeTokenId(_x uint256, _y uint256) constant returns(uint256)
func (_Contract *ContractCaller) EncodeTokenId(opts *bind.CallOpts, _x *big.Int, _y *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "encodeTokenId", _x, _y)
	return *ret0, err
}

// EncodeTokenId is a free data retrieval call binding the contract method 0xebd46d64.
//
// Solidity: function encodeTokenId(_x uint256, _y uint256) constant returns(uint256)
func (_Contract *ContractSession) EncodeTokenId(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Contract.Contract.EncodeTokenId(&_Contract.CallOpts, _x, _y)
}

// EncodeTokenId is a free data retrieval call binding the contract method 0xebd46d64.
//
// Solidity: function encodeTokenId(_x uint256, _y uint256) constant returns(uint256)
func (_Contract *ContractCallerSession) EncodeTokenId(_x *big.Int, _y *big.Int) (*big.Int, error) {
	return _Contract.Contract.EncodeTokenId(&_Contract.CallOpts, _x, _y)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_Contract *ContractCaller) Exists(opts *bind.CallOpts, _tokenId *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "exists", _tokenId)
	return *ret0, err
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_Contract *ContractSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, _tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(_tokenId uint256) constant returns(bool)
func (_Contract *ContractCallerSession) Exists(_tokenId *big.Int) (bool, error) {
	return _Contract.Contract.Exists(&_Contract.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Contract *ContractCaller) GetApproved(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "getApproved", _tokenId)
	return *ret0, err
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Contract *ContractSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, _tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(_tokenId uint256) constant returns(address)
func (_Contract *ContractCallerSession) GetApproved(_tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.GetApproved(&_Contract.CallOpts, _tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Contract *ContractCaller) IsApprovedForAll(opts *bind.CallOpts, _owner common.Address, _operator common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "isApprovedForAll", _owner, _operator)
	return *ret0, err
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Contract *ContractSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, _owner, _operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(_owner address, _operator address) constant returns(bool)
func (_Contract *ContractCallerSession) IsApprovedForAll(_owner common.Address, _operator common.Address) (bool, error) {
	return _Contract.Contract.IsApprovedForAll(&_Contract.CallOpts, _owner, _operator)
}

// LastPaidTaxes is a free data retrieval call binding the contract method 0x1f153836.
//
// Solidity: function lastPaidTaxes( address) constant returns(uint256)
func (_Contract *ContractCaller) LastPaidTaxes(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "lastPaidTaxes", arg0)
	return *ret0, err
}

// LastPaidTaxes is a free data retrieval call binding the contract method 0x1f153836.
//
// Solidity: function lastPaidTaxes( address) constant returns(uint256)
func (_Contract *ContractSession) LastPaidTaxes(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.LastPaidTaxes(&_Contract.CallOpts, arg0)
}

// LastPaidTaxes is a free data retrieval call binding the contract method 0x1f153836.
//
// Solidity: function lastPaidTaxes( address) constant returns(uint256)
func (_Contract *ContractCallerSession) LastPaidTaxes(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.LastPaidTaxes(&_Contract.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contract *ContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contract *ContractSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_Contract *ContractCallerSession) Name() (string, error) {
	return _Contract.Contract.Name(&_Contract.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Contract *ContractCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "ownerOf", _tokenId)
	return *ret0, err
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Contract *ContractSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(_tokenId uint256) constant returns(address)
func (_Contract *ContractCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _Contract.Contract.OwnerOf(&_Contract.CallOpts, _tokenId)
}

// PixelByCoordinate is a free data retrieval call binding the contract method 0x0855b6d4.
//
// Solidity: function pixelByCoordinate( uint256,  uint256) constant returns(id bytes32, seller address, x uint256, y uint256, price uint256, auctionId bytes32)
func (_Contract *ContractCaller) PixelByCoordinate(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Id        [32]byte
	Seller    common.Address
	X         *big.Int
	Y         *big.Int
	Price     *big.Int
	AuctionId [32]byte
}, error) {
	ret := new(struct {
		Id        [32]byte
		Seller    common.Address
		X         *big.Int
		Y         *big.Int
		Price     *big.Int
		AuctionId [32]byte
	})
	out := ret
	err := _Contract.contract.Call(opts, out, "pixelByCoordinate", arg0, arg1)
	return *ret, err
}

// PixelByCoordinate is a free data retrieval call binding the contract method 0x0855b6d4.
//
// Solidity: function pixelByCoordinate( uint256,  uint256) constant returns(id bytes32, seller address, x uint256, y uint256, price uint256, auctionId bytes32)
func (_Contract *ContractSession) PixelByCoordinate(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        [32]byte
	Seller    common.Address
	X         *big.Int
	Y         *big.Int
	Price     *big.Int
	AuctionId [32]byte
}, error) {
	return _Contract.Contract.PixelByCoordinate(&_Contract.CallOpts, arg0, arg1)
}

// PixelByCoordinate is a free data retrieval call binding the contract method 0x0855b6d4.
//
// Solidity: function pixelByCoordinate( uint256,  uint256) constant returns(id bytes32, seller address, x uint256, y uint256, price uint256, auctionId bytes32)
func (_Contract *ContractCallerSession) PixelByCoordinate(arg0 *big.Int, arg1 *big.Int) (struct {
	Id        [32]byte
	Seller    common.Address
	X         *big.Int
	Y         *big.Int
	Price     *big.Int
	AuctionId [32]byte
}, error) {
	return _Contract.Contract.PixelByCoordinate(&_Contract.CallOpts, arg0, arg1)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_Contract *ContractCaller) SupportsInterface(opts *bind.CallOpts, _interfaceId [4]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "supportsInterface", _interfaceId)
	return *ret0, err
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_Contract *ContractSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, _interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(_interfaceId bytes4) constant returns(bool)
func (_Contract *ContractCallerSession) SupportsInterface(_interfaceId [4]byte) (bool, error) {
	return _Contract.Contract.SupportsInterface(&_Contract.CallOpts, _interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contract *ContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contract *ContractSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_Contract *ContractCallerSession) Symbol() (string, error) {
	return _Contract.Contract.Symbol(&_Contract.CallOpts)
}

// TaxCollector is a free data retrieval call binding the contract method 0xbea1dcf8.
//
// Solidity: function taxCollector() constant returns(address)
func (_Contract *ContractCaller) TaxCollector(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "taxCollector")
	return *ret0, err
}

// TaxCollector is a free data retrieval call binding the contract method 0xbea1dcf8.
//
// Solidity: function taxCollector() constant returns(address)
func (_Contract *ContractSession) TaxCollector() (common.Address, error) {
	return _Contract.Contract.TaxCollector(&_Contract.CallOpts)
}

// TaxCollector is a free data retrieval call binding the contract method 0xbea1dcf8.
//
// Solidity: function taxCollector() constant returns(address)
func (_Contract *ContractCallerSession) TaxCollector() (common.Address, error) {
	return _Contract.Contract.TaxCollector(&_Contract.CallOpts)
}

// TaxPercentage is a free data retrieval call binding the contract method 0xae7b6d16.
//
// Solidity: function taxPercentage() constant returns(uint256)
func (_Contract *ContractCaller) TaxPercentage(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "taxPercentage")
	return *ret0, err
}

// TaxPercentage is a free data retrieval call binding the contract method 0xae7b6d16.
//
// Solidity: function taxPercentage() constant returns(uint256)
func (_Contract *ContractSession) TaxPercentage() (*big.Int, error) {
	return _Contract.Contract.TaxPercentage(&_Contract.CallOpts)
}

// TaxPercentage is a free data retrieval call binding the contract method 0xae7b6d16.
//
// Solidity: function taxPercentage() constant returns(uint256)
func (_Contract *ContractCallerSession) TaxPercentage() (*big.Int, error) {
	return _Contract.Contract.TaxPercentage(&_Contract.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_Contract *ContractCaller) TokenByIndex(opts *bind.CallOpts, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "tokenByIndex", _index)
	return *ret0, err
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_Contract *ContractSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, _index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(_index uint256) constant returns(uint256)
func (_Contract *ContractCallerSession) TokenByIndex(_index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenByIndex(&_Contract.CallOpts, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_Contract *ContractCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, _owner common.Address, _index *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "tokenOfOwnerByIndex", _owner, _index)
	return *ret0, err
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_Contract *ContractSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, _owner, _index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(_owner address, _index uint256) constant returns(uint256)
func (_Contract *ContractCallerSession) TokenOfOwnerByIndex(_owner common.Address, _index *big.Int) (*big.Int, error) {
	return _Contract.Contract.TokenOfOwnerByIndex(&_Contract.CallOpts, _owner, _index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Contract *ContractCaller) TokenURI(opts *bind.CallOpts, _tokenId *big.Int) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "tokenURI", _tokenId)
	return *ret0, err
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Contract *ContractSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, _tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(_tokenId uint256) constant returns(string)
func (_Contract *ContractCallerSession) TokenURI(_tokenId *big.Int) (string, error) {
	return _Contract.Contract.TokenURI(&_Contract.CallOpts, _tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contract *ContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contract *ContractSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_Contract *ContractCallerSession) TotalSupply() (*big.Int, error) {
	return _Contract.Contract.TotalSupply(&_Contract.CallOpts)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(user address) constant returns(uint256)
func (_Contract *ContractCaller) UserBalance(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "userBalance", user)
	return *ret0, err
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(user address) constant returns(uint256)
func (_Contract *ContractSession) UserBalance(user common.Address) (*big.Int, error) {
	return _Contract.Contract.UserBalance(&_Contract.CallOpts, user)
}

// UserBalance is a free data retrieval call binding the contract method 0x0103c92b.
//
// Solidity: function userBalance(user address) constant returns(uint256)
func (_Contract *ContractCallerSession) UserBalance(user common.Address) (*big.Int, error) {
	return _Contract.Contract.UserBalance(&_Contract.CallOpts, user)
}

// UserBalanceAtLastPaid is a free data retrieval call binding the contract method 0x2fd53674.
//
// Solidity: function userBalanceAtLastPaid( address) constant returns(uint256)
func (_Contract *ContractCaller) UserBalanceAtLastPaid(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "userBalanceAtLastPaid", arg0)
	return *ret0, err
}

// UserBalanceAtLastPaid is a free data retrieval call binding the contract method 0x2fd53674.
//
// Solidity: function userBalanceAtLastPaid( address) constant returns(uint256)
func (_Contract *ContractSession) UserBalanceAtLastPaid(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserBalanceAtLastPaid(&_Contract.CallOpts, arg0)
}

// UserBalanceAtLastPaid is a free data retrieval call binding the contract method 0x2fd53674.
//
// Solidity: function userBalanceAtLastPaid( address) constant returns(uint256)
func (_Contract *ContractCallerSession) UserBalanceAtLastPaid(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.UserBalanceAtLastPaid(&_Contract.CallOpts, arg0)
}

// UserHasPositveBalance is a free data retrieval call binding the contract method 0xdad6b482.
//
// Solidity: function userHasPositveBalance(user address) constant returns(bool)
func (_Contract *ContractCaller) UserHasPositveBalance(opts *bind.CallOpts, user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "userHasPositveBalance", user)
	return *ret0, err
}

// UserHasPositveBalance is a free data retrieval call binding the contract method 0xdad6b482.
//
// Solidity: function userHasPositveBalance(user address) constant returns(bool)
func (_Contract *ContractSession) UserHasPositveBalance(user common.Address) (bool, error) {
	return _Contract.Contract.UserHasPositveBalance(&_Contract.CallOpts, user)
}

// UserHasPositveBalance is a free data retrieval call binding the contract method 0xdad6b482.
//
// Solidity: function userHasPositveBalance(user address) constant returns(bool)
func (_Contract *ContractCallerSession) UserHasPositveBalance(user common.Address) (bool, error) {
	return _Contract.Contract.UserHasPositveBalance(&_Contract.CallOpts, user)
}

// ValueHeld is a free data retrieval call binding the contract method 0x3caa2ea1.
//
// Solidity: function valueHeld( address) constant returns(uint256)
func (_Contract *ContractCaller) ValueHeld(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "valueHeld", arg0)
	return *ret0, err
}

// ValueHeld is a free data retrieval call binding the contract method 0x3caa2ea1.
//
// Solidity: function valueHeld( address) constant returns(uint256)
func (_Contract *ContractSession) ValueHeld(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ValueHeld(&_Contract.CallOpts, arg0)
}

// ValueHeld is a free data retrieval call binding the contract method 0x3caa2ea1.
//
// Solidity: function valueHeld( address) constant returns(uint256)
func (_Contract *ContractCallerSession) ValueHeld(arg0 common.Address) (*big.Int, error) {
	return _Contract.Contract.ValueHeld(&_Contract.CallOpts, arg0)
}

// XMax is a free data retrieval call binding the contract method 0x9d8a5eaf.
//
// Solidity: function xMax() constant returns(uint256)
func (_Contract *ContractCaller) XMax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "xMax")
	return *ret0, err
}

// XMax is a free data retrieval call binding the contract method 0x9d8a5eaf.
//
// Solidity: function xMax() constant returns(uint256)
func (_Contract *ContractSession) XMax() (*big.Int, error) {
	return _Contract.Contract.XMax(&_Contract.CallOpts)
}

// XMax is a free data retrieval call binding the contract method 0x9d8a5eaf.
//
// Solidity: function xMax() constant returns(uint256)
func (_Contract *ContractCallerSession) XMax() (*big.Int, error) {
	return _Contract.Contract.XMax(&_Contract.CallOpts)
}

// YMax is a free data retrieval call binding the contract method 0xfc6ff9e1.
//
// Solidity: function yMax() constant returns(uint256)
func (_Contract *ContractCaller) YMax(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Contract.contract.Call(opts, out, "yMax")
	return *ret0, err
}

// YMax is a free data retrieval call binding the contract method 0xfc6ff9e1.
//
// Solidity: function yMax() constant returns(uint256)
func (_Contract *ContractSession) YMax() (*big.Int, error) {
	return _Contract.Contract.YMax(&_Contract.CallOpts)
}

// YMax is a free data retrieval call binding the contract method 0xfc6ff9e1.
//
// Solidity: function yMax() constant returns(uint256)
func (_Contract *ContractCallerSession) YMax() (*big.Int, error) {
	return _Contract.Contract.YMax(&_Contract.CallOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Contract *ContractTransactor) AddFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "addFunds")
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Contract *ContractSession) AddFunds() (*types.Transaction, error) {
	return _Contract.Contract.AddFunds(&_Contract.TransactOpts)
}

// AddFunds is a paid mutator transaction binding the contract method 0xa26759cb.
//
// Solidity: function addFunds() returns()
func (_Contract *ContractTransactorSession) AddFunds() (*types.Transaction, error) {
	return _Contract.Contract.AddFunds(&_Contract.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Contract *ContractTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Contract *ContractSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_to address, _tokenId uint256) returns()
func (_Contract *ContractTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Approve(&_Contract.TransactOpts, _to, _tokenId)
}

// BeginDutchAuction is a paid mutator transaction binding the contract method 0xfe13b118.
//
// Solidity: function beginDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractTransactor) BeginDutchAuction(opts *bind.TransactOpts, _x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "beginDutchAuction", _x, _y)
}

// BeginDutchAuction is a paid mutator transaction binding the contract method 0xfe13b118.
//
// Solidity: function beginDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractSession) BeginDutchAuction(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BeginDutchAuction(&_Contract.TransactOpts, _x, _y)
}

// BeginDutchAuction is a paid mutator transaction binding the contract method 0xfe13b118.
//
// Solidity: function beginDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractTransactorSession) BeginDutchAuction(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BeginDutchAuction(&_Contract.TransactOpts, _x, _y)
}

// BidInAuction is a paid mutator transaction binding the contract method 0xc1f573ca.
//
// Solidity: function bidInAuction(_x uint256, _y uint256, _bid uint256) returns()
func (_Contract *ContractTransactor) BidInAuction(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "bidInAuction", _x, _y, _bid)
}

// BidInAuction is a paid mutator transaction binding the contract method 0xc1f573ca.
//
// Solidity: function bidInAuction(_x uint256, _y uint256, _bid uint256) returns()
func (_Contract *ContractSession) BidInAuction(_x *big.Int, _y *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BidInAuction(&_Contract.TransactOpts, _x, _y, _bid)
}

// BidInAuction is a paid mutator transaction binding the contract method 0xc1f573ca.
//
// Solidity: function bidInAuction(_x uint256, _y uint256, _bid uint256) returns()
func (_Contract *ContractTransactorSession) BidInAuction(_x *big.Int, _y *big.Int, _bid *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BidInAuction(&_Contract.TransactOpts, _x, _y, _bid)
}

// BuyPixelBlock is a paid mutator transaction binding the contract method 0xfb9c19bf.
//
// Solidity: function buyPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactor) BuyPixelBlock(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyPixelBlock", _x, _y, _price)
}

// BuyPixelBlock is a paid mutator transaction binding the contract method 0xfb9c19bf.
//
// Solidity: function buyPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractSession) BuyPixelBlock(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyPixelBlock(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyPixelBlock is a paid mutator transaction binding the contract method 0xfb9c19bf.
//
// Solidity: function buyPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactorSession) BuyPixelBlock(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyPixelBlock(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyPixelBlocks is a paid mutator transaction binding the contract method 0x2e0ca840.
//
// Solidity: function buyPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactor) BuyPixelBlocks(opts *bind.TransactOpts, _x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyPixelBlocks", _x, _y, _price)
}

// BuyPixelBlocks is a paid mutator transaction binding the contract method 0x2e0ca840.
//
// Solidity: function buyPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractSession) BuyPixelBlocks(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyPixelBlocks(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyPixelBlocks is a paid mutator transaction binding the contract method 0x2e0ca840.
//
// Solidity: function buyPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactorSession) BuyPixelBlocks(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyPixelBlocks(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyUninitializedPixelBlock is a paid mutator transaction binding the contract method 0xe3696aa5.
//
// Solidity: function buyUninitializedPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactor) BuyUninitializedPixelBlock(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyUninitializedPixelBlock", _x, _y, _price)
}

// BuyUninitializedPixelBlock is a paid mutator transaction binding the contract method 0xe3696aa5.
//
// Solidity: function buyUninitializedPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractSession) BuyUninitializedPixelBlock(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyUninitializedPixelBlock(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyUninitializedPixelBlock is a paid mutator transaction binding the contract method 0xe3696aa5.
//
// Solidity: function buyUninitializedPixelBlock(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactorSession) BuyUninitializedPixelBlock(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyUninitializedPixelBlock(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyUninitializedPixelBlocks is a paid mutator transaction binding the contract method 0xa70753b3.
//
// Solidity: function buyUninitializedPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactor) BuyUninitializedPixelBlocks(opts *bind.TransactOpts, _x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "buyUninitializedPixelBlocks", _x, _y, _price)
}

// BuyUninitializedPixelBlocks is a paid mutator transaction binding the contract method 0xa70753b3.
//
// Solidity: function buyUninitializedPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractSession) BuyUninitializedPixelBlocks(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyUninitializedPixelBlocks(&_Contract.TransactOpts, _x, _y, _price)
}

// BuyUninitializedPixelBlocks is a paid mutator transaction binding the contract method 0xa70753b3.
//
// Solidity: function buyUninitializedPixelBlocks(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactorSession) BuyUninitializedPixelBlocks(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.BuyUninitializedPixelBlocks(&_Contract.TransactOpts, _x, _y, _price)
}

// EndDutchAuction is a paid mutator transaction binding the contract method 0xf2556289.
//
// Solidity: function endDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractTransactor) EndDutchAuction(opts *bind.TransactOpts, _x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "endDutchAuction", _x, _y)
}

// EndDutchAuction is a paid mutator transaction binding the contract method 0xf2556289.
//
// Solidity: function endDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractSession) EndDutchAuction(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EndDutchAuction(&_Contract.TransactOpts, _x, _y)
}

// EndDutchAuction is a paid mutator transaction binding the contract method 0xf2556289.
//
// Solidity: function endDutchAuction(_x uint256, _y uint256) returns()
func (_Contract *ContractTransactorSession) EndDutchAuction(_x *big.Int, _y *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.EndDutchAuction(&_Contract.TransactOpts, _x, _y)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Contract *ContractTransactor) SafeTransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "safeTransferFrom", _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Contract *ContractSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, _from, _to, _tokenId, _data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(_from address, _to address, _tokenId uint256, _data bytes) returns()
func (_Contract *ContractTransactorSession) SafeTransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contract.Contract.SafeTransferFrom(&_Contract.TransactOpts, _from, _to, _tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_Contract *ContractTransactor) SetApprovalForAll(opts *bind.TransactOpts, _to common.Address, _approved bool) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setApprovalForAll", _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_Contract *ContractSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, _to, _approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(_to address, _approved bool) returns()
func (_Contract *ContractTransactorSession) SetApprovalForAll(_to common.Address, _approved bool) (*types.Transaction, error) {
	return _Contract.Contract.SetApprovalForAll(&_Contract.TransactOpts, _to, _approved)
}

// SetPixelBlockPrice is a paid mutator transaction binding the contract method 0x88cb8a73.
//
// Solidity: function setPixelBlockPrice(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactor) SetPixelBlockPrice(opts *bind.TransactOpts, _x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPixelBlockPrice", _x, _y, _price)
}

// SetPixelBlockPrice is a paid mutator transaction binding the contract method 0x88cb8a73.
//
// Solidity: function setPixelBlockPrice(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractSession) SetPixelBlockPrice(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPixelBlockPrice(&_Contract.TransactOpts, _x, _y, _price)
}

// SetPixelBlockPrice is a paid mutator transaction binding the contract method 0x88cb8a73.
//
// Solidity: function setPixelBlockPrice(_x uint256, _y uint256, _price uint256) returns()
func (_Contract *ContractTransactorSession) SetPixelBlockPrice(_x *big.Int, _y *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPixelBlockPrice(&_Contract.TransactOpts, _x, _y, _price)
}

// SetPixelBlockPrices is a paid mutator transaction binding the contract method 0x35de5e89.
//
// Solidity: function setPixelBlockPrices(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactor) SetPixelBlockPrices(opts *bind.TransactOpts, _x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "setPixelBlockPrices", _x, _y, _price)
}

// SetPixelBlockPrices is a paid mutator transaction binding the contract method 0x35de5e89.
//
// Solidity: function setPixelBlockPrices(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractSession) SetPixelBlockPrices(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPixelBlockPrices(&_Contract.TransactOpts, _x, _y, _price)
}

// SetPixelBlockPrices is a paid mutator transaction binding the contract method 0x35de5e89.
//
// Solidity: function setPixelBlockPrices(_x uint256[], _y uint256[], _price uint256[]) returns()
func (_Contract *ContractTransactorSession) SetPixelBlockPrices(_x []*big.Int, _y []*big.Int, _price []*big.Int) (*types.Transaction, error) {
	return _Contract.Contract.SetPixelBlockPrices(&_Contract.TransactOpts, _x, _y, _price)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _price uint256) returns()
func (_Contract *ContractTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferFrom", _from, _to, _tokenId, _price)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _price uint256) returns()
func (_Contract *ContractSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, _from, _to, _tokenId, _price)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(_from address, _to address, _tokenId uint256, _price uint256) returns()
func (_Contract *ContractTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int, _price *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.TransferFrom(&_Contract.TransactOpts, _from, _to, _tokenId, _price)
}

// TransferTaxes is a paid mutator transaction binding the contract method 0xda30ef9f.
//
// Solidity: function transferTaxes(user address) returns(bool)
func (_Contract *ContractTransactor) TransferTaxes(opts *bind.TransactOpts, user common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "transferTaxes", user)
}

// TransferTaxes is a paid mutator transaction binding the contract method 0xda30ef9f.
//
// Solidity: function transferTaxes(user address) returns(bool)
func (_Contract *ContractSession) TransferTaxes(user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferTaxes(&_Contract.TransactOpts, user)
}

// TransferTaxes is a paid mutator transaction binding the contract method 0xda30ef9f.
//
// Solidity: function transferTaxes(user address) returns(bool)
func (_Contract *ContractTransactorSession) TransferTaxes(user common.Address) (*types.Transaction, error) {
	return _Contract.Contract.TransferTaxes(&_Contract.TransactOpts, user)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(value uint256) returns()
func (_Contract *ContractTransactor) Withdraw(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "withdraw", value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(value uint256) returns()
func (_Contract *ContractSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(value uint256) returns()
func (_Contract *ContractTransactorSession) Withdraw(value *big.Int) (*types.Transaction, error) {
	return _Contract.Contract.Withdraw(&_Contract.TransactOpts, value)
}

// ContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contract contract.
type ContractApprovalIterator struct {
	Event *ContractApproval // Event containing the contract specifics and raw log

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
func (it *ContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApproval)
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
		it.Event = new(ContractApproval)
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
func (it *ContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApproval represents a Approval event raised by the Contract contract.
type ContractApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Contract *ContractFilterer) FilterApproval(opts *bind.FilterOpts, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (*ContractApprovalIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalIterator{contract: _Contract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(_owner indexed address, _approved indexed address, _tokenId indexed uint256)
func (_Contract *ContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractApproval, _owner []common.Address, _approved []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _approvedRule []interface{}
	for _, _approvedItem := range _approved {
		_approvedRule = append(_approvedRule, _approvedItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Approval", _ownerRule, _approvedRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApproval)
				if err := _Contract.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ContractApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contract contract.
type ContractApprovalForAllIterator struct {
	Event *ContractApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractApprovalForAll)
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
		it.Event = new(ContractApprovalForAll)
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
func (it *ContractApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractApprovalForAll represents a ApprovalForAll event raised by the Contract contract.
type ContractApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_Contract *ContractFilterer) FilterApprovalForAll(opts *bind.FilterOpts, _owner []common.Address, _operator []common.Address) (*ContractApprovalForAllIterator, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractApprovalForAllIterator{contract: _Contract.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: e ApprovalForAll(_owner indexed address, _operator indexed address, _approved bool)
func (_Contract *ContractFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractApprovalForAll, _owner []common.Address, _operator []common.Address) (event.Subscription, error) {

	var _ownerRule []interface{}
	for _, _ownerItem := range _owner {
		_ownerRule = append(_ownerRule, _ownerItem)
	}
	var _operatorRule []interface{}
	for _, _operatorItem := range _operator {
		_operatorRule = append(_operatorRule, _operatorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "ApprovalForAll", _ownerRule, _operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractApprovalForAll)
				if err := _Contract.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ContractBeginDutchAuctionIterator is returned from FilterBeginDutchAuction and is used to iterate over the raw logs and unpacked data for BeginDutchAuction events raised by the Contract contract.
type ContractBeginDutchAuctionIterator struct {
	Event *ContractBeginDutchAuction // Event containing the contract specifics and raw log

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
func (it *ContractBeginDutchAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBeginDutchAuction)
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
		it.Event = new(ContractBeginDutchAuction)
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
func (it *ContractBeginDutchAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBeginDutchAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBeginDutchAuction represents a BeginDutchAuction event raised by the Contract contract.
type ContractBeginDutchAuction struct {
	PixelId   [32]byte
	TokenId   *big.Int
	AuctionId [32]byte
	Initiator common.Address
	X         *big.Int
	Y         *big.Int
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBeginDutchAuction is a free log retrieval operation binding the contract event 0xbe5c52e6db6f0b0b69e8357b5fcfa0c6efe7caca0602b28e4219885042fc264c.
//
// Solidity: e BeginDutchAuction(pixelId indexed bytes32, tokenId indexed uint256, auctionId indexed bytes32, initiator address, x uint256, y uint256, startTime uint256, endTime uint256)
func (_Contract *ContractFilterer) FilterBeginDutchAuction(opts *bind.FilterOpts, pixelId [][32]byte, tokenId []*big.Int, auctionId [][32]byte) (*ContractBeginDutchAuctionIterator, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BeginDutchAuction", pixelIdRule, tokenIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractBeginDutchAuctionIterator{contract: _Contract.contract, event: "BeginDutchAuction", logs: logs, sub: sub}, nil
}

// WatchBeginDutchAuction is a free log subscription operation binding the contract event 0xbe5c52e6db6f0b0b69e8357b5fcfa0c6efe7caca0602b28e4219885042fc264c.
//
// Solidity: e BeginDutchAuction(pixelId indexed bytes32, tokenId indexed uint256, auctionId indexed bytes32, initiator address, x uint256, y uint256, startTime uint256, endTime uint256)
func (_Contract *ContractFilterer) WatchBeginDutchAuction(opts *bind.WatchOpts, sink chan<- *ContractBeginDutchAuction, pixelId [][32]byte, tokenId []*big.Int, auctionId [][32]byte) (event.Subscription, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BeginDutchAuction", pixelIdRule, tokenIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBeginDutchAuction)
				if err := _Contract.contract.UnpackLog(event, "BeginDutchAuction", log); err != nil {
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

// ContractBuyPixelIterator is returned from FilterBuyPixel and is used to iterate over the raw logs and unpacked data for BuyPixel events raised by the Contract contract.
type ContractBuyPixelIterator struct {
	Event *ContractBuyPixel // Event containing the contract specifics and raw log

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
func (it *ContractBuyPixelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractBuyPixel)
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
		it.Event = new(ContractBuyPixel)
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
func (it *ContractBuyPixelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractBuyPixelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractBuyPixel represents a BuyPixel event raised by the Contract contract.
type ContractBuyPixel struct {
	Id     [32]byte
	Seller common.Address
	Buyer  common.Address
	X      *big.Int
	Y      *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBuyPixel is a free log retrieval operation binding the contract event 0x6e19798a5611264094483d6addf1831a553450c5f3244d22683e503558e6a949.
//
// Solidity: e BuyPixel(id indexed bytes32, seller indexed address, buyer indexed address, x uint256, y uint256, price uint256)
func (_Contract *ContractFilterer) FilterBuyPixel(opts *bind.FilterOpts, id [][32]byte, seller []common.Address, buyer []common.Address) (*ContractBuyPixelIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "BuyPixel", idRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &ContractBuyPixelIterator{contract: _Contract.contract, event: "BuyPixel", logs: logs, sub: sub}, nil
}

// WatchBuyPixel is a free log subscription operation binding the contract event 0x6e19798a5611264094483d6addf1831a553450c5f3244d22683e503558e6a949.
//
// Solidity: e BuyPixel(id indexed bytes32, seller indexed address, buyer indexed address, x uint256, y uint256, price uint256)
func (_Contract *ContractFilterer) WatchBuyPixel(opts *bind.WatchOpts, sink chan<- *ContractBuyPixel, id [][32]byte, seller []common.Address, buyer []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "BuyPixel", idRule, sellerRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractBuyPixel)
				if err := _Contract.contract.UnpackLog(event, "BuyPixel", log); err != nil {
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

// ContractEndDutchAuctionIterator is returned from FilterEndDutchAuction and is used to iterate over the raw logs and unpacked data for EndDutchAuction events raised by the Contract contract.
type ContractEndDutchAuctionIterator struct {
	Event *ContractEndDutchAuction // Event containing the contract specifics and raw log

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
func (it *ContractEndDutchAuctionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractEndDutchAuction)
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
		it.Event = new(ContractEndDutchAuction)
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
func (it *ContractEndDutchAuctionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractEndDutchAuctionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractEndDutchAuction represents a EndDutchAuction event raised by the Contract contract.
type ContractEndDutchAuction struct {
	PixelId [32]byte
	TokenId *big.Int
	Claimer common.Address
	X       *big.Int
	Y       *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEndDutchAuction is a free log retrieval operation binding the contract event 0xe7d4e7416c380ec3ff1a2559464d5abf9a67dda168935fed9f4b2f0218b9ace4.
//
// Solidity: e EndDutchAuction(pixelId indexed bytes32, tokenId indexed uint256, claimer indexed address, x uint256, y uint256)
func (_Contract *ContractFilterer) FilterEndDutchAuction(opts *bind.FilterOpts, pixelId [][32]byte, tokenId []*big.Int, claimer []common.Address) (*ContractEndDutchAuctionIterator, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "EndDutchAuction", pixelIdRule, tokenIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return &ContractEndDutchAuctionIterator{contract: _Contract.contract, event: "EndDutchAuction", logs: logs, sub: sub}, nil
}

// WatchEndDutchAuction is a free log subscription operation binding the contract event 0xe7d4e7416c380ec3ff1a2559464d5abf9a67dda168935fed9f4b2f0218b9ace4.
//
// Solidity: e EndDutchAuction(pixelId indexed bytes32, tokenId indexed uint256, claimer indexed address, x uint256, y uint256)
func (_Contract *ContractFilterer) WatchEndDutchAuction(opts *bind.WatchOpts, sink chan<- *ContractEndDutchAuction, pixelId [][32]byte, tokenId []*big.Int, claimer []common.Address) (event.Subscription, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var claimerRule []interface{}
	for _, claimerItem := range claimer {
		claimerRule = append(claimerRule, claimerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "EndDutchAuction", pixelIdRule, tokenIdRule, claimerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractEndDutchAuction)
				if err := _Contract.contract.UnpackLog(event, "EndDutchAuction", log); err != nil {
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

// ContractSetPixelPriceIterator is returned from FilterSetPixelPrice and is used to iterate over the raw logs and unpacked data for SetPixelPrice events raised by the Contract contract.
type ContractSetPixelPriceIterator struct {
	Event *ContractSetPixelPrice // Event containing the contract specifics and raw log

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
func (it *ContractSetPixelPriceIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractSetPixelPrice)
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
		it.Event = new(ContractSetPixelPrice)
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
func (it *ContractSetPixelPriceIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractSetPixelPriceIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractSetPixelPrice represents a SetPixelPrice event raised by the Contract contract.
type ContractSetPixelPrice struct {
	Id     [32]byte
	Seller common.Address
	X      *big.Int
	Y      *big.Int
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSetPixelPrice is a free log retrieval operation binding the contract event 0x3937af36e37e156c0f45c37b6a60e512cb2eb852e6521822e474de9db240b8b3.
//
// Solidity: e SetPixelPrice(id indexed bytes32, seller indexed address, x uint256, y uint256, price uint256)
func (_Contract *ContractFilterer) FilterSetPixelPrice(opts *bind.FilterOpts, id [][32]byte, seller []common.Address) (*ContractSetPixelPriceIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "SetPixelPrice", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return &ContractSetPixelPriceIterator{contract: _Contract.contract, event: "SetPixelPrice", logs: logs, sub: sub}, nil
}

// WatchSetPixelPrice is a free log subscription operation binding the contract event 0x3937af36e37e156c0f45c37b6a60e512cb2eb852e6521822e474de9db240b8b3.
//
// Solidity: e SetPixelPrice(id indexed bytes32, seller indexed address, x uint256, y uint256, price uint256)
func (_Contract *ContractFilterer) WatchSetPixelPrice(opts *bind.WatchOpts, sink chan<- *ContractSetPixelPrice, id [][32]byte, seller []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var sellerRule []interface{}
	for _, sellerItem := range seller {
		sellerRule = append(sellerRule, sellerItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "SetPixelPrice", idRule, sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractSetPixelPrice)
				if err := _Contract.contract.UnpackLog(event, "SetPixelPrice", log); err != nil {
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

// ContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contract contract.
type ContractTransferIterator struct {
	Event *ContractTransfer // Event containing the contract specifics and raw log

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
func (it *ContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractTransfer)
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
		it.Event = new(ContractTransfer)
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
func (it *ContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractTransfer represents a Transfer event raised by the Contract contract.
type ContractTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Contract *ContractFilterer) FilterTransfer(opts *bind.FilterOpts, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (*ContractTransferIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractTransferIterator{contract: _Contract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(_from indexed address, _to indexed address, _tokenId indexed uint256)
func (_Contract *ContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractTransfer, _from []common.Address, _to []common.Address, _tokenId []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}
	var _tokenIdRule []interface{}
	for _, _tokenIdItem := range _tokenId {
		_tokenIdRule = append(_tokenIdRule, _tokenIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "Transfer", _fromRule, _toRule, _tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractTransfer)
				if err := _Contract.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ContractUpdateAuctionBidIterator is returned from FilterUpdateAuctionBid and is used to iterate over the raw logs and unpacked data for UpdateAuctionBid events raised by the Contract contract.
type ContractUpdateAuctionBidIterator struct {
	Event *ContractUpdateAuctionBid // Event containing the contract specifics and raw log

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
func (it *ContractUpdateAuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractUpdateAuctionBid)
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
		it.Event = new(ContractUpdateAuctionBid)
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
func (it *ContractUpdateAuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractUpdateAuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractUpdateAuctionBid represents a UpdateAuctionBid event raised by the Contract contract.
type ContractUpdateAuctionBid struct {
	PixelId   [32]byte
	TokenId   *big.Int
	AuctionId [32]byte
	Bidder    common.Address
	AmountBet *big.Int
	TimeBet   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUpdateAuctionBid is a free log retrieval operation binding the contract event 0x40cc8a83f6c546057aa75aabf21f3dd1e7011bcbc7a54d9c11d05ae76cbc8b5b.
//
// Solidity: e UpdateAuctionBid(pixelId indexed bytes32, tokenId indexed uint256, auctionId indexed bytes32, bidder address, amountBet uint256, timeBet uint256)
func (_Contract *ContractFilterer) FilterUpdateAuctionBid(opts *bind.FilterOpts, pixelId [][32]byte, tokenId []*big.Int, auctionId [][32]byte) (*ContractUpdateAuctionBidIterator, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "UpdateAuctionBid", pixelIdRule, tokenIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractUpdateAuctionBidIterator{contract: _Contract.contract, event: "UpdateAuctionBid", logs: logs, sub: sub}, nil
}

// WatchUpdateAuctionBid is a free log subscription operation binding the contract event 0x40cc8a83f6c546057aa75aabf21f3dd1e7011bcbc7a54d9c11d05ae76cbc8b5b.
//
// Solidity: e UpdateAuctionBid(pixelId indexed bytes32, tokenId indexed uint256, auctionId indexed bytes32, bidder address, amountBet uint256, timeBet uint256)
func (_Contract *ContractFilterer) WatchUpdateAuctionBid(opts *bind.WatchOpts, sink chan<- *ContractUpdateAuctionBid, pixelId [][32]byte, tokenId []*big.Int, auctionId [][32]byte) (event.Subscription, error) {

	var pixelIdRule []interface{}
	for _, pixelIdItem := range pixelId {
		pixelIdRule = append(pixelIdRule, pixelIdItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}
	var auctionIdRule []interface{}
	for _, auctionIdItem := range auctionId {
		auctionIdRule = append(auctionIdRule, auctionIdItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "UpdateAuctionBid", pixelIdRule, tokenIdRule, auctionIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractUpdateAuctionBid)
				if err := _Contract.contract.UnpackLog(event, "UpdateAuctionBid", log); err != nil {
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
