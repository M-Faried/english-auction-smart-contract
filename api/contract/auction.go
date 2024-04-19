// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// AuctionMetaData contains all meta data concerning the Auction contract.
var AuctionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_nftId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"highestBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Ended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"}],\"name\":\"Started\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Withdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allBids\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"end\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"endTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ended\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBid\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"highestBidder\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nft\",\"outputs\":[{\"internalType\":\"contractIERC721\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nftId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_openingBid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"start\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"started\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60e060405234801561000f575f80fd5b5060405161136a38038061136a83398181016040528101906100319190610139565b3373ffffffffffffffffffffffffffffffffffffffff1660808173ffffffffffffffffffffffffffffffffffffffff16815250508173ffffffffffffffffffffffffffffffffffffffff1660a08173ffffffffffffffffffffffffffffffffffffffff16815250508060c081815250505050610177565b5f80fd5b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f6100d5826100ac565b9050919050565b6100e5816100cb565b81146100ef575f80fd5b50565b5f81519050610100816100dc565b92915050565b5f819050919050565b61011881610106565b8114610122575f80fd5b50565b5f815190506101338161010f565b92915050565b5f806040838503121561014f5761014e6100a8565b5b5f61015c858286016100f2565b925050602061016d85828601610125565b9150509250929050565b60805160a05160c0516111916101d95f395f81816107900152818161088e0152610a9f01525f81816105f6015281816107320152610a4001525f818161061a0152818161063e0152818161076e015281816108b80152610b0b01526111915ff3fe6080604052600436106100c1575f3560e01c80638da5cb5b1161007e578063bd20160711610058578063bd20160714610209578063c6bc518214610245578063d57bde791461026f578063efbe1c1c14610299576100c1565b80638da5cb5b1461018d5780638fb4b573146101b757806391f90157146101df576100c1565b806312fa6feb146100c55780631998aeef146100ef5780631f2698ab146100f95780633197cbb6146101235780633ccfd60b1461014d57806347ccca0214610163575b5f80fd5b3480156100d0575f80fd5b506100d96102af565b6040516100e69190610bfc565b60405180910390f35b6100f76102c1565b005b348015610104575f80fd5b5061010d6104bd565b60405161011a9190610bfc565b60405180910390f35b34801561012e575f80fd5b506101376104cd565b6040516101449190610c2d565b60405180910390f35b348015610158575f80fd5b506101616104d3565b005b34801561016e575f80fd5b506101776105f4565b6040516101849190610cc0565b60405180910390f35b348015610198575f80fd5b506101a1610618565b6040516101ae9190610cf9565b60405180910390f35b3480156101c2575f80fd5b506101dd60048036038101906101d89190610d40565b61063c565b005b3480156101ea575f80fd5b506101f3610852565b6040516102009190610d9e565b60405180910390f35b348015610214575f80fd5b5061022f600480360381019061022a9190610de1565b610877565b60405161023c9190610c2d565b60405180910390f35b348015610250575f80fd5b5061025961088c565b6040516102669190610c2d565b60405180910390f35b34801561027a575f80fd5b506102836108b0565b6040516102909190610c2d565b60405180910390f35b3480156102a4575f80fd5b506102ad6108b6565b005b5f60019054906101000a900460ff1681565b5f8054906101000a900460ff1661030d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161030490610e66565b60405180910390fd5b5f60019054906101000a900460ff161561035c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161035390610ece565b60405180910390fd5b60025434116103a0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161039790610f5c565b60405180910390fd5b60015442106103e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103db90610ece565b60405180910390fd5b3360035f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550346002819055503460045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055503373ffffffffffffffffffffffffffffffffffffffff167fe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2346040516104b39190610c2d565b60405180910390a2565b5f8054906101000a900460ff1681565b60015481565b5f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205490505f60045f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20819055505f8111156105a3573373ffffffffffffffffffffffffffffffffffffffff166108fc8290811502906040515f60405180830381858888f193505050501580156105a1573d5f803e3d5ffd5b505b3373ffffffffffffffffffffffffffffffffffffffff167f7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5826040516105e99190610c2d565b60405180910390a250565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000081565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106c190610fc4565b60405180910390fd5b5f8054906101000a900460ff1615610717576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161070e9061102c565b60405180910390fd5b81600281905550804261072a9190611077565b6001819055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd7f0000000000000000000000000000000000000000000000000000000000000000307f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b81526004016107cd939291906110ca565b5f604051808303815f87803b1580156107e4575f80fd5b505af11580156107f6573d5f803e3d5ffd5b5050505060015f806101000a81548160ff0219169083151502179055507f87ac41d581680567c1ef44614ddfa5522f853ea15b877693a35b1e4157cc309d426001546040516108469291906110ff565b60405180910390a15050565b60035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6004602052805f5260405f205f915090505481565b7f000000000000000000000000000000000000000000000000000000000000000081565b60025481565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610944576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161093b90610fc4565b60405180910390fd5b5f8054906101000a900460ff16610990576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161098790610e66565b60405180910390fd5b5f60019054906101000a900460ff16156109df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016109d690610ece565b60405180910390fd5b600154421015610a24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a1b90610ece565b60405180910390fd5b60015f60016101000a81548160ff0219169083151502179055507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166323b872dd3060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006040518463ffffffff1660e01b8152600401610adc93929190611126565b5f604051808303815f87803b158015610af3575f80fd5b505af1158015610b05573d5f803e3d5ffd5b505050507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166108fc60025490811502906040515f60405180830381858888f19350505050158015610b6e573d5f803e3d5ffd5b5060035f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f0ab6ead5ff681783e38604d729771f8ba89300f334c9f52cba4909f79c6b5588600254604051610bd89190610c2d565b60405180910390a2565b5f8115159050919050565b610bf681610be2565b82525050565b5f602082019050610c0f5f830184610bed565b92915050565b5f819050919050565b610c2781610c15565b82525050565b5f602082019050610c405f830184610c1e565b92915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f819050919050565b5f610c88610c83610c7e84610c46565b610c65565b610c46565b9050919050565b5f610c9982610c6e565b9050919050565b5f610caa82610c8f565b9050919050565b610cba81610ca0565b82525050565b5f602082019050610cd35f830184610cb1565b92915050565b5f610ce382610c46565b9050919050565b610cf381610cd9565b82525050565b5f602082019050610d0c5f830184610cea565b92915050565b5f80fd5b610d1f81610c15565b8114610d29575f80fd5b50565b5f81359050610d3a81610d16565b92915050565b5f8060408385031215610d5657610d55610d12565b5b5f610d6385828601610d2c565b9250506020610d7485828601610d2c565b9150509250929050565b5f610d8882610c46565b9050919050565b610d9881610d7e565b82525050565b5f602082019050610db15f830184610d8f565b92915050565b610dc081610d7e565b8114610dca575f80fd5b50565b5f81359050610ddb81610db7565b92915050565b5f60208284031215610df657610df5610d12565b5b5f610e0384828501610dcd565b91505092915050565b5f82825260208201905092915050565b7f41756374696f6e20686173206e6f7420737461727465640000000000000000005f82015250565b5f610e50601783610e0c565b9150610e5b82610e1c565b602082019050919050565b5f6020820190508181035f830152610e7d81610e44565b9050919050565b7f41756374696f6e2068617320656e6465640000000000000000000000000000005f82015250565b5f610eb8601183610e0c565b9150610ec382610e84565b602082019050919050565b5f6020820190508181035f830152610ee581610eac565b9050919050565b7f426964207072696365206973206c6f776572207468616e2074686520637572725f8201527f656e742068696768657374206269640000000000000000000000000000000000602082015250565b5f610f46602f83610e0c565b9150610f5182610eec565b604082019050919050565b5f6020820190508181035f830152610f7381610f3a565b9050919050565b7f4f6e6c79206f776e65722063616e2063616c6c207468652066756e6374696f6e5f82015250565b5f610fae602083610e0c565b9150610fb982610f7a565b602082019050919050565b5f6020820190508181035f830152610fdb81610fa2565b9050919050565b7f41756374696f6e2068617320616c7265616479207374617274656400000000005f82015250565b5f611016601b83610e0c565b915061102182610fe2565b602082019050919050565b5f6020820190508181035f8301526110438161100a565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61108182610c15565b915061108c83610c15565b92508282019050808211156110a4576110a361104a565b5b92915050565b5f6110b482610c8f565b9050919050565b6110c4816110aa565b82525050565b5f6060820190506110dd5f8301866110bb565b6110ea6020830185610d8f565b6110f76040830184610c1e565b949350505050565b5f6040820190506111125f830185610c1e565b61111f6020830184610c1e565b9392505050565b5f6060820190506111395f830186610d8f565b6111466020830185610d8f565b6111536040830184610c1e565b94935050505056fea2646970667358221220567950b901070c3677cd268e523208d3c6f9c50bc826baf5d9a1aecb9b6cd18d64736f6c63430008190033",
}

// AuctionABI is the input ABI used to generate the binding from.
// Deprecated: Use AuctionMetaData.ABI instead.
var AuctionABI = AuctionMetaData.ABI

// AuctionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AuctionMetaData.Bin instead.
var AuctionBin = AuctionMetaData.Bin

// DeployAuction deploys a new Ethereum contract, binding an instance of Auction to it.
func DeployAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nft common.Address, _nftId *big.Int) (common.Address, *types.Transaction, *Auction, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AuctionBin), backend, _nft, _nftId)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// Auction is an auto generated Go binding around an Ethereum contract.
type Auction struct {
	AuctionCaller     // Read-only binding to the contract
	AuctionTransactor // Write-only binding to the contract
	AuctionFilterer   // Log filterer for contract events
}

// AuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type AuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AuctionSession struct {
	Contract     *Auction          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AuctionCallerSession struct {
	Contract *AuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AuctionTransactorSession struct {
	Contract     *AuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type AuctionRaw struct {
	Contract *Auction // Generic contract binding to access the raw methods on
}

// AuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AuctionCallerRaw struct {
	Contract *AuctionCaller // Generic read-only contract binding to access the raw methods on
}

// AuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AuctionTransactorRaw struct {
	Contract *AuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAuction creates a new instance of Auction, bound to a specific deployed contract.
func NewAuction(address common.Address, backend bind.ContractBackend) (*Auction, error) {
	contract, err := bindAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Auction{AuctionCaller: AuctionCaller{contract: contract}, AuctionTransactor: AuctionTransactor{contract: contract}, AuctionFilterer: AuctionFilterer{contract: contract}}, nil
}

// NewAuctionCaller creates a new read-only instance of Auction, bound to a specific deployed contract.
func NewAuctionCaller(address common.Address, caller bind.ContractCaller) (*AuctionCaller, error) {
	contract, err := bindAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionCaller{contract: contract}, nil
}

// NewAuctionTransactor creates a new write-only instance of Auction, bound to a specific deployed contract.
func NewAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*AuctionTransactor, error) {
	contract, err := bindAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AuctionTransactor{contract: contract}, nil
}

// NewAuctionFilterer creates a new log filterer instance of Auction, bound to a specific deployed contract.
func NewAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*AuctionFilterer, error) {
	contract, err := bindAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AuctionFilterer{contract: contract}, nil
}

// bindAuction binds a generic wrapper to an already deployed contract.
func bindAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AuctionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.AuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.AuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Auction *AuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Auction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Auction *AuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Auction *AuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Auction.Contract.contract.Transact(opts, method, params...)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCaller) AllBids(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "allBids", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// AllBids is a free data retrieval call binding the contract method 0xbd201607.
//
// Solidity: function allBids(address ) view returns(uint256)
func (_Auction *AuctionCallerSession) AllBids(arg0 common.Address) (*big.Int, error) {
	return _Auction.Contract.AllBids(&_Auction.CallOpts, arg0)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCaller) EndTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "endTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// EndTime is a free data retrieval call binding the contract method 0x3197cbb6.
//
// Solidity: function endTime() view returns(uint256)
func (_Auction *AuctionCallerSession) EndTime() (*big.Int, error) {
	return _Auction.Contract.EndTime(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCaller) Ended(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "ended")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// Ended is a free data retrieval call binding the contract method 0x12fa6feb.
//
// Solidity: function ended() view returns(bool)
func (_Auction *AuctionCallerSession) Ended() (bool, error) {
	return _Auction.Contract.Ended(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCaller) HighestBid(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBid")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBid is a free data retrieval call binding the contract method 0xd57bde79.
//
// Solidity: function highestBid() view returns(uint256)
func (_Auction *AuctionCallerSession) HighestBid() (*big.Int, error) {
	return _Auction.Contract.HighestBid(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCaller) HighestBidder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "highestBidder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// HighestBidder is a free data retrieval call binding the contract method 0x91f90157.
//
// Solidity: function highestBidder() view returns(address)
func (_Auction *AuctionCallerSession) HighestBidder() (common.Address, error) {
	return _Auction.Contract.HighestBidder(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCaller) Nft(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nft")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// Nft is a free data retrieval call binding the contract method 0x47ccca02.
//
// Solidity: function nft() view returns(address)
func (_Auction *AuctionCallerSession) Nft() (common.Address, error) {
	return _Auction.Contract.Nft(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCaller) NftId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "nftId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// NftId is a free data retrieval call binding the contract method 0xc6bc5182.
//
// Solidity: function nftId() view returns(uint256)
func (_Auction *AuctionCallerSession) NftId() (*big.Int, error) {
	return _Auction.Contract.NftId(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Auction *AuctionCallerSession) Owner() (common.Address, error) {
	return _Auction.Contract.Owner(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCaller) Started(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Auction.contract.Call(opts, &out, "started")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Started is a free data retrieval call binding the contract method 0x1f2698ab.
//
// Solidity: function started() view returns(bool)
func (_Auction *AuctionCallerSession) Started() (bool, error) {
	return _Auction.Contract.Started(&_Auction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactor) Bid(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "bid")
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x1998aeef.
//
// Solidity: function bid() payable returns()
func (_Auction *AuctionTransactorSession) Bid() (*types.Transaction, error) {
	return _Auction.Contract.Bid(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactor) End(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "end")
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// End is a paid mutator transaction binding the contract method 0xefbe1c1c.
//
// Solidity: function end() returns()
func (_Auction *AuctionTransactorSession) End() (*types.Transaction, error) {
	return _Auction.Contract.End(&_Auction.TransactOpts)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactor) Start(opts *bind.TransactOpts, _openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "start", _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Start is a paid mutator transaction binding the contract method 0x8fb4b573.
//
// Solidity: function start(uint256 _openingBid, uint256 _duration) returns()
func (_Auction *AuctionTransactorSession) Start(_openingBid *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _Auction.Contract.Start(&_Auction.TransactOpts, _openingBid, _duration)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Auction.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Auction *AuctionTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Auction.Contract.Withdraw(&_Auction.TransactOpts)
}

// AuctionBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Auction contract.
type AuctionBidIterator struct {
	Event *AuctionBid // Event containing the contract specifics and raw log

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
func (it *AuctionBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionBid)
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
		it.Event = new(AuctionBid)
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
func (it *AuctionBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionBid represents a Bid event raised by the Auction contract.
type AuctionBid struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterBid(opts *bind.FilterOpts, bidder []common.Address) (*AuctionBidIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionBidIterator{contract: _Auction.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *AuctionBid, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Bid", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionBid)
				if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0xe684a55f31b79eca403df938249029212a5925ec6be8012e099b45bc1019e5d2.
//
// Solidity: event Bid(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseBid(log types.Log) (*AuctionBid, error) {
	event := new(AuctionBid)
	if err := _Auction.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionEndedIterator is returned from FilterEnded and is used to iterate over the raw logs and unpacked data for Ended events raised by the Auction contract.
type AuctionEndedIterator struct {
	Event *AuctionEnded // Event containing the contract specifics and raw log

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
func (it *AuctionEndedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionEnded)
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
		it.Event = new(AuctionEnded)
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
func (it *AuctionEndedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionEndedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionEnded represents a Ended event raised by the Auction contract.
type AuctionEnded struct {
	HighestBidder common.Address
	Value         *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterEnded is a free log retrieval operation binding the contract event 0x0ab6ead5ff681783e38604d729771f8ba89300f334c9f52cba4909f79c6b5588.
//
// Solidity: event Ended(address indexed highestBidder, uint256 value)
func (_Auction *AuctionFilterer) FilterEnded(opts *bind.FilterOpts, highestBidder []common.Address) (*AuctionEndedIterator, error) {

	var highestBidderRule []interface{}
	for _, highestBidderItem := range highestBidder {
		highestBidderRule = append(highestBidderRule, highestBidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Ended", highestBidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionEndedIterator{contract: _Auction.contract, event: "Ended", logs: logs, sub: sub}, nil
}

// WatchEnded is a free log subscription operation binding the contract event 0x0ab6ead5ff681783e38604d729771f8ba89300f334c9f52cba4909f79c6b5588.
//
// Solidity: event Ended(address indexed highestBidder, uint256 value)
func (_Auction *AuctionFilterer) WatchEnded(opts *bind.WatchOpts, sink chan<- *AuctionEnded, highestBidder []common.Address) (event.Subscription, error) {

	var highestBidderRule []interface{}
	for _, highestBidderItem := range highestBidder {
		highestBidderRule = append(highestBidderRule, highestBidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Ended", highestBidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionEnded)
				if err := _Auction.contract.UnpackLog(event, "Ended", log); err != nil {
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

// ParseEnded is a log parse operation binding the contract event 0x0ab6ead5ff681783e38604d729771f8ba89300f334c9f52cba4909f79c6b5588.
//
// Solidity: event Ended(address indexed highestBidder, uint256 value)
func (_Auction *AuctionFilterer) ParseEnded(log types.Log) (*AuctionEnded, error) {
	event := new(AuctionEnded)
	if err := _Auction.contract.UnpackLog(event, "Ended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionStartedIterator is returned from FilterStarted and is used to iterate over the raw logs and unpacked data for Started events raised by the Auction contract.
type AuctionStartedIterator struct {
	Event *AuctionStarted // Event containing the contract specifics and raw log

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
func (it *AuctionStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionStarted)
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
		it.Event = new(AuctionStarted)
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
func (it *AuctionStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionStarted represents a Started event raised by the Auction contract.
type AuctionStarted struct {
	StartTime *big.Int
	EndTime   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStarted is a free log retrieval operation binding the contract event 0x87ac41d581680567c1ef44614ddfa5522f853ea15b877693a35b1e4157cc309d.
//
// Solidity: event Started(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) FilterStarted(opts *bind.FilterOpts) (*AuctionStartedIterator, error) {

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Started")
	if err != nil {
		return nil, err
	}
	return &AuctionStartedIterator{contract: _Auction.contract, event: "Started", logs: logs, sub: sub}, nil
}

// WatchStarted is a free log subscription operation binding the contract event 0x87ac41d581680567c1ef44614ddfa5522f853ea15b877693a35b1e4157cc309d.
//
// Solidity: event Started(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) WatchStarted(opts *bind.WatchOpts, sink chan<- *AuctionStarted) (event.Subscription, error) {

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Started")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionStarted)
				if err := _Auction.contract.UnpackLog(event, "Started", log); err != nil {
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

// ParseStarted is a log parse operation binding the contract event 0x87ac41d581680567c1ef44614ddfa5522f853ea15b877693a35b1e4157cc309d.
//
// Solidity: event Started(uint256 startTime, uint256 endTime)
func (_Auction *AuctionFilterer) ParseStarted(log types.Log) (*AuctionStarted, error) {
	event := new(AuctionStarted)
	if err := _Auction.contract.UnpackLog(event, "Started", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AuctionWithdrawnIterator is returned from FilterWithdrawn and is used to iterate over the raw logs and unpacked data for Withdrawn events raised by the Auction contract.
type AuctionWithdrawnIterator struct {
	Event *AuctionWithdrawn // Event containing the contract specifics and raw log

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
func (it *AuctionWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AuctionWithdrawn)
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
		it.Event = new(AuctionWithdrawn)
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
func (it *AuctionWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AuctionWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AuctionWithdrawn represents a Withdrawn event raised by the Auction contract.
type AuctionWithdrawn struct {
	Bidder common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdrawn is a free log retrieval operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) FilterWithdrawn(opts *bind.FilterOpts, bidder []common.Address) (*AuctionWithdrawnIterator, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.FilterLogs(opts, "Withdrawn", bidderRule)
	if err != nil {
		return nil, err
	}
	return &AuctionWithdrawnIterator{contract: _Auction.contract, event: "Withdrawn", logs: logs, sub: sub}, nil
}

// WatchWithdrawn is a free log subscription operation binding the contract event 0x7084f5476618d8e60b11ef0d7d3f06914655adb8793e28ff7f018d4c76d505d5.
//
// Solidity: event Withdrawn(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) WatchWithdrawn(opts *bind.WatchOpts, sink chan<- *AuctionWithdrawn, bidder []common.Address) (event.Subscription, error) {

	var bidderRule []interface{}
	for _, bidderItem := range bidder {
		bidderRule = append(bidderRule, bidderItem)
	}

	logs, sub, err := _Auction.contract.WatchLogs(opts, "Withdrawn", bidderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AuctionWithdrawn)
				if err := _Auction.contract.UnpackLog(event, "Withdrawn", log); err != nil {
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
// Solidity: event Withdrawn(address indexed bidder, uint256 value)
func (_Auction *AuctionFilterer) ParseWithdrawn(log types.Log) (*AuctionWithdrawn, error) {
	event := new(AuctionWithdrawn)
	if err := _Auction.contract.UnpackLog(event, "Withdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
