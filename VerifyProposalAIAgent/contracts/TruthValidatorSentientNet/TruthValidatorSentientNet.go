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
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"finalResult\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"indexed\":false,\"internalType\":\"structTruthValidatorSentientNet.Vote[]\",\"name\":\"voterResults\",\"type\":\"tuple[]\"}],\"name\":\"ProposalFinalized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"ProposalSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"VoteThresholdChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_voter\",\"type\":\"address\"}],\"name\":\"getVote\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"internalType\":\"structTruthValidatorSentientNet.Vote\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoteCounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"}],\"name\":\"getVoters\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposalVoters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"yesVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"noVotes\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isFinalized\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_newThreshold\",\"type\":\"uint256\"}],\"name\":\"setVoteThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_proposalId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_reason\",\"type\":\"string\"}],\"name\":\"vote\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voteThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"voterVotes\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isApproved\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405260036005553480156013575f80fd5b5060015f81905580546001600160a01b031916331790556111dd806100375f395ff3fe608060405234801561000f575f80fd5b50600436106100b1575f3560e01c80638a241f4c1161006e5780638a241f4c1461014a5780639a2ddcb214610189578063ab89437c146101ab578063bc3f931f146101d6578063d3c0715b146101f6578063f851a44014610209575f80fd5b8063013cf08b146100b55780630c0512e9146100e257806312909485146100f9578063373d6d5e1461010e5780634fe437d51461012157806386b646f21461012a575b5f80fd5b6100c86100c3366004610c87565b61021c565b6040516100d9959493929190610ce1565b60405180910390f35b6100eb60025481565b6040519081526020016100d9565b61010c610107366004610c87565b6102d5565b005b61010c61011c366004610db5565b6103d6565b6100eb60055481565b61013d610138366004610c87565b6104fc565b6040516100d99190610def565b610174610158366004610c87565b5f90815260036020819052604090912060028101549101549091565b604080519283526020830191909152016100d9565b61019c610197366004610e3b565b610565565b6040516100d993929190610e74565b6101be6101b9366004610ea8565b610628565b6040516001600160a01b0390911681526020016100d9565b6101e96101e4366004610e3b565b61065c565b6040516100d99190610efa565b61010c610204366004610f13565b61075a565b6001546101be906001600160a01b031681565b60036020525f90815260409020805460018201805491929161023d90610f6c565b80601f016020809104026020016040519081016040528092919081815260200182805461026990610f6c565b80156102b45780601f1061028b576101008083540402835291602001916102b4565b820191905f5260205f20905b81548152906001019060200180831161029757829003601f168201915b50505050600283015460038401546004909401549293909290915060ff1685565b6001546001600160a01b031633146103405760405162461bcd60e51b815260206004820152602360248201527f547275746856616c696461746f723a2063616c6c6572206973206e6f7420616460448201526236b4b760e91b60648201526084015b60405180910390fd5b5f811161039b5760405162461bcd60e51b8152602060048201526024808201527f547275746856616c696461746f723a207468726573686f6c64206d7573742062604482015263065203e360e41b6064820152608401610337565b60058190556040518181527fb85f4a3c2a2a4d5ee4d22634fd1ed5b17b5fb6e35ef5ca2d03049a0ef4e59d259060200160405180910390a150565b5f81511180156103e95750610118815111155b61042e5760405162461bcd60e51b8152602060048201526016602482015275092dcecc2d8d2c840c6dedce8cadce840d8cadccee8d60531b6044820152606401610337565b600280545f918261043e83610fb8565b909155506040805160a08101825282815260208082018681525f83850181905260608401819052608084018190528581526003909252929020815181559151929350916001820190610490908261101c565b506040828101516002830155606083015160038301556080909201516004909101805460ff19169115159190911790555181907f076921e45db4958d9ff79453aa2f12e0fa9664505544a9f664f8762b330cf7ca906104f09085906110dc565b60405180910390a25050565b5f8181526006602090815260409182902080548351818402810184019094528084526060939283018282801561055957602002820191905f5260205f20905b81546001600160a01b0316815260019091019060200180831161053b575b50505050509050919050565b600460209081525f9283526040808420909152908252902080546001820180546001600160a01b03831693600160a01b90930460ff169291906105a790610f6c565b80601f01602080910402602001604051908101604052809291908181526020018280546105d390610f6c565b801561061e5780601f106105f55761010080835404028352916020019161061e565b820191905f5260205f20905b81548152906001019060200180831161060157829003601f168201915b5050505050905083565b6006602052815f5260405f208181548110610641575f80fd5b5f918252602090912001546001600160a01b03169150829050565b60408051606080820183525f8083526020830152918101919091525f8381526004602090815260408083206001600160a01b03868116855290835292819020815160608101835281549485168152600160a01b90940460ff16151592840192909252600182018054918401916106d190610f6c565b80601f01602080910402602001604051908101604052809291908181526020018280546106fd90610f6c565b80156107485780601f1061071f57610100808354040283529160200191610748565b820191905f5260205f20905b81548152906001019060200180831161072b57829003601f168201915b50505050508152505090505b92915050565b60025f54036107ab5760405162461bcd60e51b815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610337565b60025f55620186a05a116108015760405162461bcd60e51b815260206004820152601960248201527f496e73756666696369656e742067617320666f7220766f7465000000000000006044820152606401610337565b5f838152600360205260409020600481015460ff161561085b5760405162461bcd60e51b8152602060048201526015602482015274141c9bdc1bdcd85b081a5cc8199a5b985b1a5e9959605a1b6044820152606401610337565b5f8481526004602090815260408083203384529091529020546001600160a01b0316156108ba5760405162461bcd60e51b815260206004820152600d60248201526c105b1c9958591e481d9bdd1959609a1b6044820152606401610337565b604080516060810182523380825285151560208084019182528385018781525f8a81526004835286812094815293909152939091208251815492511515600160a01b026001600160a81b03199093166001600160a01b03919091161791909117815591519091906001820190610930908261101c565b50905050821561095557600281018054905f61094b83610fb8565b919050555061096c565b600381018054905f61096683610fb8565b91905055505b5f8481526006602090815260408083208054600181018255908452919092200180546001600160a01b03191633908117909155905185917fc4ce3f0cdcc6c3a55938650ba7706e6b5d9f686f108fdce6a6bb68424fe1ed75916109d3919087908790610e74565b60405180910390a2600554816003015482600201546109f291906110ee565b10610a0057610a0084610a0a565b505060015f555050565b5f818152600360205260409020600481015460ff1615610a6c5760405162461bcd60e51b815260206004820152601d60248201527f50726f706f73616c20697320616c72656164792066696e616c697a65640000006044820152606401610337565b60048101805460ff19166001179055600381015460028201545f84815260066020526040812054929091119167ffffffffffffffff811115610ab057610ab0610d18565b604051908082528060200260200182016040528015610afc57816020015b60408051606080820183525f808352602083015291810191909152815260200190600190039081610ace5790505b5090505f5b5f85815260066020526040902054811015610c46575f858152600660205260408120805483908110610b3557610b35611101565b5f918252602080832091909101548883526004825260408084206001600160a01b0392831680865290845293819020815160608101835281549384168152600160a01b90930460ff1615159383019390935260018301805494955091939084019190610ba090610f6c565b80601f0160208091040260200160405190810160405280929190818152602001828054610bcc90610f6c565b8015610c175780601f10610bee57610100808354040283529160200191610c17565b820191905f5260205f20905b815481529060010190602001808311610bfa57829003601f168201915b505050505081525050838381518110610c3257610c32611101565b602090810291909101015250600101610b01565b50837fba55beb7b3e66a185c23c4a9dce26e8dc8cb6ff2408fbd527b1f9aa36ce2de5e8383604051610c79929190611115565b60405180910390a250505050565b5f60208284031215610c97575f80fd5b5035919050565b5f81518084525f5b81811015610cc257602081850181015186830182015201610ca6565b505f602082860101526020601f19601f83011685010191505092915050565b85815260a060208201525f610cf960a0830187610c9e565b6040830195909552506060810192909252151560809091015292915050565b634e487b7160e01b5f52604160045260245ffd5b5f82601f830112610d3b575f80fd5b813567ffffffffffffffff80821115610d5657610d56610d18565b604051601f8301601f19908116603f01168101908282118183101715610d7e57610d7e610d18565b81604052838152866020858801011115610d96575f80fd5b836020870160208301375f602085830101528094505050505092915050565b5f60208284031215610dc5575f80fd5b813567ffffffffffffffff811115610ddb575f80fd5b610de784828501610d2c565b949350505050565b602080825282518282018190525f9190848201906040850190845b81811015610e2f5783516001600160a01b031683529284019291840191600101610e0a565b50909695505050505050565b5f8060408385031215610e4c575f80fd5b8235915060208301356001600160a01b0381168114610e69575f80fd5b809150509250929050565b6001600160a01b038416815282151560208201526060604082018190525f90610e9f90830184610c9e565b95945050505050565b5f8060408385031215610eb9575f80fd5b50508035926020909101359150565b60018060a01b0381511682526020810151151560208301525f604082015160606040850152610de76060850182610c9e565b602081525f610f0c6020830184610ec8565b9392505050565b5f805f60608486031215610f25575f80fd5b8335925060208401358015158114610f3b575f80fd5b9150604084013567ffffffffffffffff811115610f56575f80fd5b610f6286828701610d2c565b9150509250925092565b600181811c90821680610f8057607f821691505b602082108103610f9e57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b5f60018201610fc957610fc9610fa4565b5060010190565b601f82111561101757805f5260205f20601f840160051c81016020851015610ff55750805b601f840160051c820191505b81811015611014575f8155600101611001565b50505b505050565b815167ffffffffffffffff81111561103657611036610d18565b61104a816110448454610f6c565b84610fd0565b602080601f83116001811461107d575f84156110665750858301515b5f19600386901b1c1916600185901b1785556110d4565b5f85815260208120601f198616915b828110156110ab5788860151825594840194600190910190840161108c565b50858210156110c857878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b602081525f610f0c6020830184610c9e565b8082018082111561075457610754610fa4565b634e487b7160e01b5f52603260045260245ffd5b5f60408201841515835260206040602085015281855180845260608601915060608160051b8701019350602087015f5b8281101561117357605f19888703018452611161868351610ec8565b95509284019290840190600101611145565b50939897505050505050505056fea26469706673582212208eb0ec5eb60985558511dc05d901cef604e37a5443432d5f857a2e1c62a6df1964736f6c637828302e382e32352d646576656c6f702e323032342e322e32342b636f6d6d69742e64626137353465630059",
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

// TruthValidatorSentientNetVoteThresholdChangedIterator is returned from FilterVoteThresholdChanged and is used to iterate over the raw logs and unpacked data for VoteThresholdChanged events raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetVoteThresholdChangedIterator struct {
	Event *TruthValidatorSentientNetVoteThresholdChanged // Event containing the contract specifics and raw log

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
func (it *TruthValidatorSentientNetVoteThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TruthValidatorSentientNetVoteThresholdChanged)
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
		it.Event = new(TruthValidatorSentientNetVoteThresholdChanged)
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
func (it *TruthValidatorSentientNetVoteThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TruthValidatorSentientNetVoteThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TruthValidatorSentientNetVoteThresholdChanged represents a VoteThresholdChanged event raised by the TruthValidatorSentientNet contract.
type TruthValidatorSentientNetVoteThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterVoteThresholdChanged is a free log retrieval operation binding the contract event 0xb85f4a3c2a2a4d5ee4d22634fd1ed5b17b5fb6e35ef5ca2d03049a0ef4e59d25.
//
// Solidity: event VoteThresholdChanged(uint256 newThreshold)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) FilterVoteThresholdChanged(opts *bind.FilterOpts) (*TruthValidatorSentientNetVoteThresholdChangedIterator, error) {

	logs, sub, err := _TruthValidatorSentientNet.contract.FilterLogs(opts, "VoteThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &TruthValidatorSentientNetVoteThresholdChangedIterator{contract: _TruthValidatorSentientNet.contract, event: "VoteThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchVoteThresholdChanged is a free log subscription operation binding the contract event 0xb85f4a3c2a2a4d5ee4d22634fd1ed5b17b5fb6e35ef5ca2d03049a0ef4e59d25.
//
// Solidity: event VoteThresholdChanged(uint256 newThreshold)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) WatchVoteThresholdChanged(opts *bind.WatchOpts, sink chan<- *TruthValidatorSentientNetVoteThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _TruthValidatorSentientNet.contract.WatchLogs(opts, "VoteThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TruthValidatorSentientNetVoteThresholdChanged)
				if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "VoteThresholdChanged", log); err != nil {
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

// ParseVoteThresholdChanged is a log parse operation binding the contract event 0xb85f4a3c2a2a4d5ee4d22634fd1ed5b17b5fb6e35ef5ca2d03049a0ef4e59d25.
//
// Solidity: event VoteThresholdChanged(uint256 newThreshold)
func (_TruthValidatorSentientNet *TruthValidatorSentientNetFilterer) ParseVoteThresholdChanged(log types.Log) (*TruthValidatorSentientNetVoteThresholdChanged, error) {
	event := new(TruthValidatorSentientNetVoteThresholdChanged)
	if err := _TruthValidatorSentientNet.contract.UnpackLog(event, "VoteThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
