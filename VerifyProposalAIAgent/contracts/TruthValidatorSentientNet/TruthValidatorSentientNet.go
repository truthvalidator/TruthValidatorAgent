// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package TruthValidatorSentientNet

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

// TruthValidatorSentientNetVote is an auto generated low-level Go binding around an user-defined struct.
type TruthValidatorSentientNetVote struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}

// TruthValidatorSentientNetMetaData contains all meta data concerning the TruthValidatorSentientNet contract.
var TruthValidatorSentientNetMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalResult\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structTruthValidatorSentientNet.Vote[]\",\"name\":\"voterResults\",\"type\":\"tuple[]\"}],\"name\":\"ProposalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"ProposalSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVote\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structTruthValidatorSentientNet.Vote\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalVoters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"setVoteThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260036004553480156013575f80fd5b505f80546001600160a01b03191633179055611096806100325f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80638a241f4c1161006e5780638a241f4c1461014a5780639a2ddcb21461018b578063ab89437c146101ad578063bc3f931f146101d8578063d3c0715b146101f8578063f851a4401461020b575f80fd5b8063013cf08b146100b55780630c0512e9146100e257806312909485146100f9578063373d6d5e1461010e5780634fe437d51461012157806386b646f21461012a575b5f80fd5b6100c86100c3366004610b40565b61021d565b6040516100d9959493929190610b9a565b60405180910390f35b6100eb60015481565b6040519081526020016100d9565b61010c610107366004610b40565b6102d6565b005b61010c61011c366004610c6e565b610392565b6100eb60045481565b61013d610138366004610b40565b610460565b6040516100d99190610ca8565b610176610158366004610b40565b5f908152600260208190526040909120908101546003909101549091565b604080519283526020830191909152016100d9565b61019e610199366004610cf4565b6104c9565b6040516100d993929190610d2d565b6101c06101bb366004610d61565b61058c565b6040516001600160a01b0390911681526020016100d9565b6101eb6101e6366004610cf4565b6105c0565b6040516100d99190610db3565b61010c610206366004610dcc565b6106be565b5f546101c0906001600160a01b031681565b60026020525f90815260409020805460018201805491929161023e90610e25565b80601f016020809104026020016040519081016040528092919081815260200182805461026a90610e25565b80156102b55780601f1061028c576101008083540402835291602001916102b5565b820191905f5260205f20905b81548152906001019060200180831161029857829003601f168201915b50505050600283015460038401546004909401549293909290915060ff1685565b5f546001600160a01b0316331461033e5760405162461bcd60e51b815260206004820152602160248201527f4f6e6c792061646d696e2063616e2063616c6c20746869732066756e6374696f6044820152603760f91b60648201526084015b60405180910390fd5b5f811161038d5760405162461bcd60e51b815260206004820181905260248201527f5468726573686f6c64206d7573742062652067726561746572207468616e20306044820152606401610335565b600455565b600180545f91826103a283610e71565b909155506040805160a08101825282815260208082018681525f838501819052606084018190526080840181905285815260029092529290208151815591519293509160018201906103f49082610ed5565b506040828101516002830155606083015160038301556080909201516004909101805460ff19169115159190911790555181907f076921e45db4958d9ff79453aa2f12e0fa9664505544a9f664f8762b330cf7ca90610454908590610f95565b60405180910390a25050565b5f818152600560209081526040918290208054835181840281018401909452808452606093928301828280156104bd57602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161049f575b50505050509050919050565b600360209081525f9283526040808420909152908252902080546001820180546001600160a01b03831693600160a01b90930460ff1692919061050b90610e25565b80601f016020809104026020016040519081016040528092919081815260200182805461053790610e25565b80156105825780601f1061055957610100808354040283529160200191610582565b820191905f5260205f20905b81548152906001019060200180831161056557829003601f168201915b5050505050905083565b6005602052815f5260405f2081815481106105a5575f80fd5b5f918252602090912001546001600160a01b03169150829050565b60408051606080820183525f8083526020830152918101919091525f8381526003602090815260408083206001600160a01b03868116855290835292819020815160608101835281549485168152600160a01b90940460ff161515928401929092526001820180549184019161063590610e25565b80601f016020809104026020016040519081016040528092919081815260200182805461066190610e25565b80156106ac5780601f10610683576101008083540402835291602001916106ac565b820191905f5260205f20905b81548152906001019060200180831161068f57829003601f168201915b50505050508152505090505b92915050565b5f838152600260205260409020600481015460ff16156107185760405162461bcd60e51b8152602060048201526015602482015274141c9bdc1bdcd85b081a5cc8199a5b985b1a5e9959605a1b6044820152606401610335565b5f8481526003602090815260408083203384529091529020546001600160a01b0316156107775760405162461bcd60e51b815260206004820152600d60248201526c105b1c9958591e481d9bdd1959609a1b6044820152606401610335565b604080516060810182523380825285151560208084019182528385018781525f8a81526003835286812094815293909152939091208251815492511515600160a01b026001600160a81b03199093166001600160a01b039190911617919091178155915190919060018201906107ed9082610ed5565b50905050821561081257600281018054905f61080883610e71565b9190505550610829565b600381018054905f61082383610e71565b91905055505b5f8481526005602090815260408083208054600181018255908452919092200180546001600160a01b03191633908117909155905185917fc4ce3f0cdcc6c3a55938650ba7706e6b5d9f686f108fdce6a6bb68424fe1ed7591610890919087908790610d2d565b60405180910390a2600454816003015482600201546108af9190610fa7565b106108bd576108bd846108c3565b50505050565b5f818152600260205260409020600481015460ff16156109255760405162461bcd60e51b815260206004820152601d60248201527f50726f706f73616c20697320616c72656164792066696e616c697a65640000006044820152606401610335565b60048101805460ff19166001179055600381015460028201545f84815260056020526040812054929091119167ffffffffffffffff81111561096957610969610bd1565b6040519080825280602002602001820160405280156109b557816020015b60408051606080820183525f8083526020830152918101919091528152602001906001900390816109875790505b5090505f5b5f85815260056020526040902054811015610aff575f8581526005602052604081208054839081106109ee576109ee610fba565b5f918252602080832091909101548883526003825260408084206001600160a01b0392831680865290845293819020815160608101835281549384168152600160a01b90930460ff1615159383019390935260018301805494955091939084019190610a5990610e25565b80601f0160208091040260200160405190810160405280929190818152602001828054610a8590610e25565b8015610ad05780601f10610aa757610100808354040283529160200191610ad0565b820191905f5260205f20905b815481529060010190602001808311610ab357829003601f168201915b505050505081525050838381518110610aeb57610aeb610fba565b6020908102919091010152506001016109ba565b50837fba55beb7b3e66a185c23c4a9dce26e8dc8cb6ff2408fbd527b1f9aa36ce2de5e8383604051610b32929190610fce565b60405180910390a250505050565b5f60208284031215610b50575f80fd5b5035919050565b5f81518084525f5b81811015610b7b57602081850181015186830182015201610b5f565b505f602082860101526020601f19601f83011685010191505092915050565b85815260a060208201525f610bb260a0830187610b57565b6040830195909552506060810192909252151560809091015292915050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610bf4575f80fd5b813567ffffffffffffffff80821115610c0f57610c0f610bd1565b604051601f8301601f19908116603f01168101908282118183101715610c3757610c37610bd1565b81604052838152866020858801011115610c4f575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f60208284031215610c7e575f80fd5b813567ffffffffffffffff811115610c94575f80fd5b610ca084828501610be5565b949350505050565b602080825282518282018190525f9190848201906040850190845b81811015610ce85783516001600160a01b031683529284019291840191600101610cc3565b50909695505050505050565b5f8060408385031215610d05575f80fd5b8235915060208301356001600160a01b0381168114610d22575f80fd5b809150509250929050565b6001600160a01b038416815282151560208201526060604082018190525f90610d5890830184610b57565b95945050505050565b5f8060408385031215610d72575f80fd5b50508035926020909101359150565b60018060a01b0381511682526020810151151560208301525f604082015160606040850152610ca06060850182610b57565b602081525f610dc56020830184610d81565b9392505050565b5f805f60608486031215610dde575f80fd5b8335925060208401358015158114610df4575f80fd5b9150604084013567ffffffffffffffff811115610e0f575f80fd5b610e1b86828701610be5565b9150509250925092565b600181811c90821680610e3957607f821691505b602082108103610e5757634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201610e8257610e82610e5d565b5060010190565b601f821115610ed057805f5260205f20601f840160051c81016020851015610eae5750805b601f840160051c820191505b81811015610ecd575f8155600101610eba565b50505b505050565b815167ffffffffffffffff811115610eef57610eef610bd1565b610f0381610efd8454610e25565b84610e89565b602080601f831160018114610f36575f8415610f1f5750858301515b5f19600386901b1c1916600185901b178555610f8d565b5f85815260208120601f198616915b82811015610f6457888601518255948401946001909101908401610f45565b5085821015610f8157878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f610dc56020830184610b57565b808201808211156106b8576106b8610e5d565b634e487b7160e01b5f52603260045260245ffd5b5f60408201841515835260206040602085015281855180845260608601915060608160051b8701019350602087015f5b8281101561102c57605f1988870301845261101a868351610d81565b95509284019290840190600101610ffe565b50939897505050505050505056fea264697066735822122075f0f31ebb1b818c08ac310615388ebd3c5074cf7f90382d7e09c0fd5e47909864736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
}

// TruthValidatorSentientNetABI is the input ABI used to generate the binding from.
// Deprecated: Use TruthValidatorSentientNetMetaData.ABI instead.
var TruthValidatorSentientNetABI = TruthValidatorSentientNetMetaData.ABI

// TruthValidatorSentientNetBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use TruthValidatorSentientNetMetaData.Bin instead.
var TruthValidatorSentientNetBin = TruthValidatorSentientNetMetaData.Bin

// DeployTruthValidatorSentientNet deploys a new Ethereum contract, binding an instance of TruthValidatorSentientNet to it.
func DeployTruthValidatorSentientNet(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TruthValidatorSentientNet, error) {
	parsed, err := TruthValidatorSentientNetMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(TruthValidatorSentientNetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TruthValidatorSentientNet{TruthValidatorSentientNetCaller: TruthValidatorSentientNetCaller{contract: contract}, TruthValidatorSentientNetTransactor: TruthValidatorSentientNetTransactor{contract: contract}, TruthValidatorSentientNetFilterer: TruthValidatorSentientNetFilterer{contract: contract}}, nil
}

// TruthValidatorSentientNet is an auto generated Go binding around an Ethereum contract.
type TruthValidatorSentientNet struct {
	TruthValidatorSentientNetCaller     // Read-only binding to the contract
	TruthValidatorSentientNetTransactor // Write-only binding to the contract
	TruthValidatorSentientNetFilterer   // Log filterer for contract events
}

// TruthValidatorSentientNetCaller is an auto generated read-only Go binding around an Ethereum contract.
type TruthValidatorSentientNetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruthValidatorSentientNetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TruthValidatorSentientNetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruthValidatorSentientNetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TruthValidatorSentientNetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TruthValidatorSentientNetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TruthValidatorSentientNetSession struct {
	Contract     *TruthValidatorSentientNet // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// TruthValidatorSentientNetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TruthValidatorSentientNetCallerSession struct {
	Contract *TruthValidatorSentientNetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// TruthValidatorSentientNetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TruthValidatorSentientNetTransactorSession struct {
	Contract     *TruthValidatorSentientNetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// TruthValidatorSentientNetRaw is an auto generated low-level Go binding around an Ethereum contract.
type TruthValidatorSentientNetRaw struct {
	Contract *TruthValidatorSentientNet // Generic contract binding to access the raw methods on
}

// TruthValidatorSentientNetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TruthValidatorSentientNetCallerRaw struct {
	Contract *TruthValidatorSentientNetCaller // Generic read-only contract binding to access the raw methods on
}

// TruthValidatorSentientNetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TruthValidatorSentientNetTransactorRaw struct {
	Contract *TruthValidatorSentientNetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTruthValidatorSentientNet creates a new instance of TruthValidatorSentientNet, bound to a specific deployed contract.
func NewTruthValidatorSentientNet(address common.Address, backend bind.ContractBackend) (*TruthValidatorSentientNet, error) {
	contract, err := bindTruthValidatorSentientNet(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNet{TruthValidatorSentientNetCaller: TruthValidatorSentientNetCaller{contract: contract}, TruthValidatorSentientNetTransactor: TruthValidatorSentientNetTransactor{contract: contract}, TruthValidatorSentientNetFilterer: TruthValidatorSentientNetFilterer{contract: contract}}, nil
}

// NewTruthValidatorSentientNetCaller creates a new read-only instance of TruthValidatorSentientNet, bound to a specific deployed contract.
func NewTruthValidatorSentientNetCaller(address common.Address, caller bind.ContractCaller) (*TruthValidatorSentientNetCaller, error) {
	contract, err := bindTruthValidatorSentientNet(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetCaller{contract: contract}, nil
}

// NewTruthValidatorSentientNetTransactor creates a new write-only instance of TruthValidatorSentientNet, bound to a specific deployed contract.
func NewTruthValidatorSentientNetTransactor(address common.Address, transactor bind.ContractTransactor) (*TruthValidatorSentientNetTransactor, error) {
	contract, err := bindTruthValidatorSentientNet(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetTransactor{contract: contract}, nil
}

// NewTruthValidatorSentientNetFilterer creates a new log filterer instance of TruthValidatorSentientNet, bound to a specific deployed contract.
func NewTruthValidatorSentientNetFilterer(address common.Address, filterer bind.ContractFilterer) (*TruthValidatorSentientNetFilterer, error) {
	contract, err := bindTruthValidatorSentientNet(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetFilterer{contract: contract}, nil
}

// bindTruthValidatorSentientNet binds a generic wrapper to an already deployed contract.
func bindTruthValidatorSentientNet(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TruthValidatorSentientNetMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TruthValidatorSentientNet.Contract.TruthValidatorSentientNetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.TruthValidatorSentientNetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.TruthValidatorSentientNetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TruthValidatorSentientNet.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.contract.Transact(opts, method, params...)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) Admin() (common.Address, error) {
	return _TruthValidatorSentientNet.Contract.Admin(&_TruthValidatorSentientNet.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) Admin() (common.Address, error) {
	return _TruthValidatorSentientNet.Contract.Admin(&_TruthValidatorSentientNet.CallOpts)
}

// GetVote is a free data retrieval call binding the contract method 0xbc3f931f.
//
// Solidity: function getVote(uint256 _proposalId, address _voter) view returns((address,bool,string))
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) GetVote(opts *bind.CallOpts, _proposalId *big.Int, _voter common.Address) (TruthValidatorSentientNetVote, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "getVote", _proposalId, _voter)

	if err != nil {
		return *new(TruthValidatorSentientNetVote), err
	}

	out0 := *abi.ConvertType(out[0], new(TruthValidatorSentientNetVote)).(*TruthValidatorSentientNetVote)

	return out0, err

}

// GetVote is a free data retrieval call binding the contract method 0xbc3f931f.
//
// Solidity: function getVote(uint256 _proposalId, address _voter) view returns((address,bool,string))
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) GetVote(_proposalId *big.Int, _voter common.Address) (TruthValidatorSentientNetVote, error) {
	return _TruthValidatorSentientNet.Contract.GetVote(&_TruthValidatorSentientNet.CallOpts, _proposalId, _voter)
}

// GetVote is a free data retrieval call binding the contract method 0xbc3f931f.
//
// Solidity: function getVote(uint256 _proposalId, address _voter) view returns((address,bool,string))
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) GetVote(_proposalId *big.Int, _voter common.Address) (TruthValidatorSentientNetVote, error) {
	return _TruthValidatorSentientNet.Contract.GetVote(&_TruthValidatorSentientNet.CallOpts, _proposalId, _voter)
}

// GetVoteCounts is a free data retrieval call binding the contract method 0x8a241f4c.
//
// Solidity: function getVoteCounts(uint256 _proposalId) view returns(uint256 yesVotes, uint256 noVotes)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) GetVoteCounts(opts *bind.CallOpts, _proposalId *big.Int) (struct {
	YesVotes *big.Int
	NoVotes  *big.Int
}, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "getVoteCounts", _proposalId)

	outstruct := new(struct {
		YesVotes *big.Int
		NoVotes  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.YesVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVoteCounts is a free data retrieval call binding the contract method 0x8a241f4c.
//
// Solidity: function getVoteCounts(uint256 _proposalId) view returns(uint256 yesVotes, uint256 noVotes)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) GetVoteCounts(_proposalId *big.Int) (struct {
	YesVotes *big.Int
	NoVotes  *big.Int
}, error) {
	return _TruthValidatorSentientNet.Contract.GetVoteCounts(&_TruthValidatorSentientNet.CallOpts, _proposalId)
}

// GetVoteCounts is a free data retrieval call binding the contract method 0x8a241f4c.
//
// Solidity: function getVoteCounts(uint256 _proposalId) view returns(uint256 yesVotes, uint256 noVotes)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) GetVoteCounts(_proposalId *big.Int) (struct {
	YesVotes *big.Int
	NoVotes  *big.Int
}, error) {
	return _TruthValidatorSentientNet.Contract.GetVoteCounts(&_TruthValidatorSentientNet.CallOpts, _proposalId)
}

// GetVoters is a free data retrieval call binding the contract method 0x86b646f2.
//
// Solidity: function getVoters(uint256 _proposalId) view returns(address[])
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) GetVoters(opts *bind.CallOpts, _proposalId *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "getVoters", _proposalId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetVoters is a free data retrieval call binding the contract method 0x86b646f2.
//
// Solidity: function getVoters(uint256 _proposalId) view returns(address[])
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) GetVoters(_proposalId *big.Int) ([]common.Address, error) {
	return _TruthValidatorSentientNet.Contract.GetVoters(&_TruthValidatorSentientNet.CallOpts, _proposalId)
}

// GetVoters is a free data retrieval call binding the contract method 0x86b646f2.
//
// Solidity: function getVoters(uint256 _proposalId) view returns(address[])
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) GetVoters(_proposalId *big.Int) ([]common.Address, error) {
	return _TruthValidatorSentientNet.Contract.GetVoters(&_TruthValidatorSentientNet.CallOpts, _proposalId)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) ProposalCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "proposalCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) ProposalCounter() (*big.Int, error) {
	return _TruthValidatorSentientNet.Contract.ProposalCounter(&_TruthValidatorSentientNet.CallOpts)
}

// ProposalCounter is a free data retrieval call binding the contract method 0x0c0512e9.
//
// Solidity: function proposalCounter() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) ProposalCounter() (*big.Int, error) {
	return _TruthValidatorSentientNet.Contract.ProposalCounter(&_TruthValidatorSentientNet.CallOpts)
}

// ProposalVoters is a free data retrieval call binding the contract method 0xab89437c.
//
// Solidity: function proposalVoters(uint256 , uint256 ) view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) ProposalVoters(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "proposalVoters", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalVoters is a free data retrieval call binding the contract method 0xab89437c.
//
// Solidity: function proposalVoters(uint256 , uint256 ) view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) ProposalVoters(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _TruthValidatorSentientNet.Contract.ProposalVoters(&_TruthValidatorSentientNet.CallOpts, arg0, arg1)
}

// ProposalVoters is a free data retrieval call binding the contract method 0xab89437c.
//
// Solidity: function proposalVoters(uint256 , uint256 ) view returns(address)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) ProposalVoters(arg0 *big.Int, arg1 *big.Int) (common.Address, error) {
	return _TruthValidatorSentientNet.Contract.ProposalVoters(&_TruthValidatorSentientNet.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string content, uint256 yesVotes, uint256 noVotes, bool isFinalized)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id          *big.Int
	Content     string
	YesVotes    *big.Int
	NoVotes     *big.Int
	IsFinalized bool
}, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "proposals", arg0)

	outstruct := new(struct {
		Id          *big.Int
		Content     string
		YesVotes    *big.Int
		NoVotes     *big.Int
		IsFinalized bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Content = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.YesVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.NoVotes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.IsFinalized = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string content, uint256 yesVotes, uint256 noVotes, bool isFinalized)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) Proposals(arg0 *big.Int) (struct {
	Id          *big.Int
	Content     string
	YesVotes    *big.Int
	NoVotes     *big.Int
	IsFinalized bool
}, error) {
	return _TruthValidatorSentientNet.Contract.Proposals(&_TruthValidatorSentientNet.CallOpts, arg0)
}

// Proposals is a free data retrieval call binding the contract method 0x013cf08b.
//
// Solidity: function proposals(uint256 ) view returns(uint256 id, string content, uint256 yesVotes, uint256 noVotes, bool isFinalized)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) Proposals(arg0 *big.Int) (struct {
	Id          *big.Int
	Content     string
	YesVotes    *big.Int
	NoVotes     *big.Int
	IsFinalized bool
}, error) {
	return _TruthValidatorSentientNet.Contract.Proposals(&_TruthValidatorSentientNet.CallOpts, arg0)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) VoteThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "voteThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) VoteThreshold() (*big.Int, error) {
	return _TruthValidatorSentientNet.Contract.VoteThreshold(&_TruthValidatorSentientNet.CallOpts)
}

// VoteThreshold is a free data retrieval call binding the contract method 0x4fe437d5.
//
// Solidity: function voteThreshold() view returns(uint256)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) VoteThreshold() (*big.Int, error) {
	return _TruthValidatorSentientNet.Contract.VoteThreshold(&_TruthValidatorSentientNet.CallOpts)
}

// VoterVotes is a free data retrieval call binding the contract method 0x9a2ddcb2.
//
// Solidity: function voterVotes(uint256 , address ) view returns(address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCaller) VoterVotes(opts *bind.CallOpts, arg0 *big.Int, arg1 common.Address) (struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}, error) {
	var out []interface{}
	err := _TruthValidatorSentientNet.contract.Call(opts, &out, "voterVotes", arg0, arg1)

	outstruct := new(struct {
		Voter      common.Address
		IsApproved bool
		Reason     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Voter = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.IsApproved = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Reason = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// VoterVotes is a free data retrieval call binding the contract method 0x9a2ddcb2.
//
// Solidity: function voterVotes(uint256 , address ) view returns(address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) VoterVotes(arg0 *big.Int, arg1 common.Address) (struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}, error) {
	return _TruthValidatorSentientNet.Contract.VoterVotes(&_TruthValidatorSentientNet.CallOpts, arg0, arg1)
}

// VoterVotes is a free data retrieval call binding the contract method 0x9a2ddcb2.
//
// Solidity: function voterVotes(uint256 , address ) view returns(address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetCallerSession) VoterVotes(arg0 *big.Int, arg1 common.Address) (struct {
	Voter      common.Address
	IsApproved bool
	Reason     string
}, error) {
	return _TruthValidatorSentientNet.Contract.VoterVotes(&_TruthValidatorSentientNet.CallOpts, arg0, arg1)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(uint256 _newThreshold) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactor) SetVoteThreshold(opts *bind.TransactOpts, _newThreshold *big.Int) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.contract.Transact(opts, "setVoteThreshold", _newThreshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(uint256 _newThreshold) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) SetVoteThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.SetVoteThreshold(&_TruthValidatorSentientNet.TransactOpts, _newThreshold)
}

// SetVoteThreshold is a paid mutator transaction binding the contract method 0x12909485.
//
// Solidity: function setVoteThreshold(uint256 _newThreshold) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactorSession) SetVoteThreshold(_newThreshold *big.Int) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.SetVoteThreshold(&_TruthValidatorSentientNet.TransactOpts, _newThreshold)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x373d6d5e.
//
// Solidity: function submitProposal(string _content) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactor) SubmitProposal(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.contract.Transact(opts, "submitProposal", _content)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x373d6d5e.
//
// Solidity: function submitProposal(string _content) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) SubmitProposal(_content string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.SubmitProposal(&_TruthValidatorSentientNet.TransactOpts, _content)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x373d6d5e.
//
// Solidity: function submitProposal(string _content) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactorSession) SubmitProposal(_content string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.SubmitProposal(&_TruthValidatorSentientNet.TransactOpts, _content)
}

// Vote is a paid mutator transaction binding the contract method 0xd3c0715b.
//
// Solidity: function vote(uint256 _proposalId, bool _isApproved, string _reason) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactor) Vote(opts *bind.TransactOpts, _proposalId *big.Int, _isApproved bool, _reason string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.contract.Transact(opts, "vote", _proposalId, _isApproved, _reason)
}

// Vote is a paid mutator transaction binding the contract method 0xd3c0715b.
//
// Solidity: function vote(uint256 _proposalId, bool _isApproved, string _reason) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetSession) Vote(_proposalId *big.Int, _isApproved bool, _reason string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.Vote(&_TruthValidatorSentientNet.TransactOpts, _proposalId, _isApproved, _reason)
}

// Vote is a paid mutator transaction binding the contract method 0xd3c0715b.
//
// Solidity: function vote(uint256 _proposalId, bool _isApproved, string _reason) returns()
func (_TruthValidatorSentientNet *TruthValidatorSentientNetTransactorSession) Vote(_proposalId *big.Int, _isApproved bool, _reason string) (*types.Transaction, error) {
	return _TruthValidatorSentientNet.Contract.Vote(&_TruthValidatorSentientNet.TransactOpts, _proposalId, _isApproved, _reason)
}

// TruthValidatorSentientNetProposalFinalizedIterator is returned from FilterProposalFinalized and is used to iterate over the raw logs and unpacked data for ProposalFinalized events raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetProposalFinalizedIterator struct {
	Event *TruthValidatorSentientNetProposalFinalized // Event containing the contract specifics and raw log

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
func (it *TruthValidatorSentientNetProposalFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruthValidatorSentientNetProposalFinalized)
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
		it.Event = new(TruthValidatorSentientNetProposalFinalized)
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
func (it *TruthValidatorSentientNetProposalFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruthValidatorSentientNetProposalFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruthValidatorSentientNetProposalFinalized represents a ProposalFinalized event raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetProposalFinalized struct {
	ProposalId   *big.Int
	FinalResult  bool
	VoterResults []TruthValidatorSentientNetVote
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterProposalFinalized is a free log retrieval operation binding the contract event 0xba55beb7b3e66a185c23c4a9dce26e8dc8cb6ff2408fbd527b1f9aa36ce2de5e.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool finalResult, (address,bool,string)[] voterResults)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) FilterProposalFinalized(opts *bind.FilterOpts, proposalId []*big.Int) (*TruthValidatorSentientNetProposalFinalizedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.FilterLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetProposalFinalizedIterator{contract: _TruthValidatorSentientNet.contract, event: "ProposalFinalized", logs: logs, sub: sub}, nil
}

// WatchProposalFinalized is a free log subscription operation binding the contract event 0xba55beb7b3e66a185c23c4a9dce26e8dc8cb6ff2408fbd527b1f9aa36ce2de5e.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool finalResult, (address,bool,string)[] voterResults)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) WatchProposalFinalized(opts *bind.WatchOpts, sink chan<- *TruthValidatorSentientNetProposalFinalized, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.WatchLogs(opts, "ProposalFinalized", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruthValidatorSentientNetProposalFinalized)
				if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
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

// ParseProposalFinalized is a log parse operation binding the contract event 0xba55beb7b3e66a185c23c4a9dce26e8dc8cb6ff2408fbd527b1f9aa36ce2de5e.
//
// Solidity: event ProposalFinalized(uint256 indexed proposalId, bool finalResult, (address,bool,string)[] voterResults)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) ParseProposalFinalized(log types.Log) (*TruthValidatorSentientNetProposalFinalized, error) {
	event := new(TruthValidatorSentientNetProposalFinalized)
	if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "ProposalFinalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TruthValidatorSentientNetProposalSubmittedIterator is returned from FilterProposalSubmitted and is used to iterate over the raw logs and unpacked data for ProposalSubmitted events raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetProposalSubmittedIterator struct {
	Event *TruthValidatorSentientNetProposalSubmitted // Event containing the contract specifics and raw log

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
func (it *TruthValidatorSentientNetProposalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruthValidatorSentientNetProposalSubmitted)
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
		it.Event = new(TruthValidatorSentientNetProposalSubmitted)
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
func (it *TruthValidatorSentientNetProposalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruthValidatorSentientNetProposalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruthValidatorSentientNetProposalSubmitted represents a ProposalSubmitted event raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetProposalSubmitted struct {
	ProposalId *big.Int
	Content    string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalSubmitted is a free log retrieval operation binding the contract event 0x076921e45db4958d9ff79453aa2f12e0fa9664505544a9f664f8762b330cf7ca.
//
// Solidity: event ProposalSubmitted(uint256 indexed proposalId, string content)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) FilterProposalSubmitted(opts *bind.FilterOpts, proposalId []*big.Int) (*TruthValidatorSentientNetProposalSubmittedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.FilterLogs(opts, "ProposalSubmitted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetProposalSubmittedIterator{contract: _TruthValidatorSentientNet.contract, event: "ProposalSubmitted", logs: logs, sub: sub}, nil
}

// WatchProposalSubmitted is a free log subscription operation binding the contract event 0x076921e45db4958d9ff79453aa2f12e0fa9664505544a9f664f8762b330cf7ca.
//
// Solidity: event ProposalSubmitted(uint256 indexed proposalId, string content)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) WatchProposalSubmitted(opts *bind.WatchOpts, sink chan<- *TruthValidatorSentientNetProposalSubmitted, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.WatchLogs(opts, "ProposalSubmitted", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruthValidatorSentientNetProposalSubmitted)
				if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
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

// ParseProposalSubmitted is a log parse operation binding the contract event 0x076921e45db4958d9ff79453aa2f12e0fa9664505544a9f664f8762b330cf7ca.
//
// Solidity: event ProposalSubmitted(uint256 indexed proposalId, string content)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) ParseProposalSubmitted(log types.Log) (*TruthValidatorSentientNetProposalSubmitted, error) {
	event := new(TruthValidatorSentientNetProposalSubmitted)
	if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "ProposalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TruthValidatorSentientNetVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetVoteCastIterator struct {
	Event *TruthValidatorSentientNetVoteCast // Event containing the contract specifics and raw log

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
func (it *TruthValidatorSentientNetVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruthValidatorSentientNetVoteCast)
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
		it.Event = new(TruthValidatorSentientNetVoteCast)
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
func (it *TruthValidatorSentientNetVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruthValidatorSentientNetVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruthValidatorSentientNetVoteCast represents a VoteCast event raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetVoteCast struct {
	ProposalId *big.Int
	Voter      common.Address
	IsApproved bool
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xc4ce3f0cdcc6c3a55938650ba7706e6b5d9f686f108fdce6a6bb68424fe1ed75.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) FilterVoteCast(opts *bind.FilterOpts, proposalId []*big.Int) (*TruthValidatorSentientNetVoteCastIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.FilterLogs(opts, "VoteCast", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetVoteCastIterator{contract: _TruthValidatorSentientNet.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xc4ce3f0cdcc6c3a55938650ba7706e6b5d9f686f108fdce6a6bb68424fe1ed75.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *TruthValidatorSentientNetVoteCast, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _TruthValidatorSentientNet.contract.WatchLogs(opts, "VoteCast", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruthValidatorSentientNetVoteCast)
				if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "VoteCast", log); err != nil {
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

// ParseVoteCast is a log parse operation binding the contract event 0xc4ce3f0cdcc6c3a55938650ba7706e6b5d9f686f108fdce6a6bb68424fe1ed75.
//
// Solidity: event VoteCast(uint256 indexed proposalId, address voter, bool isApproved, string reason)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) ParseVoteCast(log types.Log) (*TruthValidatorSentientNetVoteCast, error) {
	event := new(TruthValidatorSentientNetVoteCast)
	if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
