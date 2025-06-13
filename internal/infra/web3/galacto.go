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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"PublishAction\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"userWallet\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"userID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"username\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actionID\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"actionName\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"rewardAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"performedAt\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getBalanceBreakdown\",\"inputs\":[{\"name\":\"wallet\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"fromStaff\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"fromTransfer\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isMember\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"join\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"staff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFromStaff\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ActionPublished\",\"inputs\":[{\"name\":\"id\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"userWallet\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"userID\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"username\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"actionID\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"actionName\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"rewardAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"performedAt\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Joined\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransferFromStaff\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyMember\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"InvalidAmount\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NotMember\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]},{\"type\":\"error\",\"name\":\"NotStaff\",\"inputs\":[{\"name\":\"message\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x6080604052346103bb57604080519081016001600160401b038111828210176102ce576040908152600782526647616c6163746f60c81b602083015280519081016001600160401b038111828210176102ce57604052600381526211d0d560ea1b602082015281516001600160401b0381116102ce57600354600181811c911680156103b1575b60208210146102b057601f811161034e575b50602092601f82116001146102ed57928192935f926102e2575b50508160011b915f199060031b1c1916176003555b80516001600160401b0381116102ce57600454600181811c911680156102c4575b60208210146102b057601f811161024d575b50602091601f82116001146101ed579181925f926101e2575b50508160011b915f199060031b1c1916176004555b600580546001600160a01b0319163317905530156101cf5760025469d3c21bcecceda100000081018091116101bb57600255305f525f60205260405f2069d3c21bcecceda1000000815401905560405169d3c21bcecceda100000081525f7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60203093a3604051610e9c90816103c08239f35b634e487b7160e01b5f52601160045260245ffd5b63ec442f0560e01b5f525f60045260245ffd5b015190505f80610113565b601f1982169260045f52805f20915f5b8581106102355750836001951061021d575b505050811b01600455610128565b01515f1960f88460031b161c191690555f808061020f565b919260206001819286850151815501940192016101fd565b60045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b601f830160051c810191602084106102a6575b601f0160051c01905b81811061029b57506100fa565b5f815560010161028e565b9091508190610285565b634e487b7160e01b5f52602260045260245ffd5b90607f16906100e8565b634e487b7160e01b5f52604160045260245ffd5b015190505f806100b2565b601f1982169360035f52805f20915f5b868110610336575083600195961061031e575b505050811b016003556100c7565b01515f1960f88460031b161c191690555f8080610310565b919260206001819286850151815501940192016102fd565b60035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b601f830160051c810191602084106103a7575b601f0160051c01905b81811061039c5750610098565b5f815560010161038f565b9091508190610386565b90607f1690610086565b5f80fdfe6080806040526004361015610012575f80fd5b5f3560e01c90816306fdde0314610a4d57508063095ea7b3146109cb5780630bbae1b3146108f057806315b5709a146108a557806318160ddd1461088857806323b872dd146107a9578063313ce5671461078e5780636739b9f4146104c657806370a082311461048f57806395d89b4114610374578063a230c52414610337578063a9059cbb146101e5578063b688a3631461013d578063d11345e0146101155763dd62ed3e146100c1575f80fd5b34610111576040366003190112610111576100da610b46565b6100e2610b5c565b6001600160a01b039182165f908152600160209081526040808320949093168252928352819020549051908152f35b5f80fd5b34610111575f366003190112610111576005546040516001600160a01b039091168152602090f35b34610111575f36600319011261011157335f52600760205260ff60405f20541661019f57335f52600760205260405f20600160ff19825416179055337f7073afa60b48e833a1a66ebb442bc8e0d19e9f3cc05f34bdcd1da22c8d87f2755f80a2005b604051633a6ff68d60e01b815260206004820152601960248201527f596f752061726520616c72656164792061206d656d6265722e000000000000006044820152606490fd5b34610111576040366003190112610111576102a0610201610b46565b60243590335f52600760205261021d60ff60405f205416610bed565b6001600160a01b038116610232811515610c45565b61023e81331415610c85565b610249831515610cd2565b335f525f6020526102608360405f20541015610d29565b8290335f52600660205283600160405f200154105f1461033157335f526006602052600160405f2001545b806102e6575b50816102ab575b505033610dbc565b602060405160018152f35b335f52600660205260405f206102c2838254610daf565b90555f5260066020526102dd600160405f2001918254610d6e565b90558380610298565b61032a919250335f526006602052600160405f2001610306828254610daf565b9055825f526006602052600160405f2001610322828254610d6e565b905584610daf565b9085610291565b8361028b565b34610111576020366003190112610111576001600160a01b03610358610b46565b165f526007602052602060ff60405f2054166040519015158152f35b34610111575f366003190112610111576040515f6004548060011c90600181168015610485575b602083108114610471578285529081156104555750600114610400575b50819003601f01601f191681019067ffffffffffffffff8211818310176103ec576103e882918260405282610b1c565b0390f35b634e487b7160e01b5f52604160045260245ffd5b905060045f527f8a35acfbc15ff81a39ae7d344fd709f28e8600b4aa8c65c6b64bfe7fe36bd19b5f905b82821061043f575060209150820101826103b8565b600181602092548385880101520191019061042a565b90506020925060ff191682840152151560051b820101826103b8565b634e487b7160e01b5f52602260045260245ffd5b91607f169161039b565b34610111576020366003190112610111576001600160a01b036104b0610b46565b165f525f602052602060405f2054604051908152f35b34610111576101003660031901126101115760043567ffffffffffffffff8111610111576104f8903690600401610b72565b90610501610b5c565b9160443567ffffffffffffffff811161011157610522903690600401610b72565b939060643567ffffffffffffffff811161011157610544903690600401610b72565b9060843567ffffffffffffffff811161011157610565903690600401610b72565b909160a43567ffffffffffffffff811161011157610587903690600401610b72565b94909560c435976105a360018060a01b03600554163314610ba0565b6001600160a01b03165f81815260076020526040902054909b9060ff16156107375788156106e0578b5f52600660205260405f206105e28a8254610d6e565b90558b156106cd5788600254906105f891610d6e565b6002558b5f525f60205260405f208981540190558b6040518a81527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60205f92a36040519a8b9a60e08c5260e08c019061065192610d8f565b908a820360208c015261066392610d8f565b9088820360408a015261067592610d8f565b90868203606088015261068792610d8f565b90848203608086015261069992610d8f565b9060a083015260e43560c0830152037f747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c491a2005b63ec442f0560e01b5f525f60045260245ffd5b60405163589956a560e11b815260206004820152602860248201527f52657761726420616d6f756e74206d757374206265206772656174657220746860448201526730b7103d32b9379760c11b6064820152608490fd5b60405162b3883360e61b815260206004820152602960248201527f55736572206d7573742062652061206d656d62657220746f207075626c6973686044820152681030b1ba34b7b7399760b91b6064820152608490fd5b34610111575f36600319011261011157602060405160128152f35b34610111576060366003190112610111576107c2610b46565b6107ca610b5c565b6001600160a01b0382165f818152600160209081526040808320338452909152902054909260443592915f198110610808575b506102a09350610dbc565b83811061086d57841561085a573315610847576102a0945f52600160205260405f2060018060a01b0333165f526020528360405f2091039055846107fd565b634a1406b160e11b5f525f60045260245ffd5b63e602df0560e01b5f525f60045260245ffd5b8390637dc7a0d960e11b5f523360045260245260445260645ffd5b34610111575f366003190112610111576020600254604051908152f35b346101115760203660031901126101115760406001600160a01b036108c8610b46565b16805f526006602052815f2054905f5260066020526001825f20015482519182526020820152f35b3461011157604036600319011261011157610909610b46565b7f8091ad0f6e4ac8d4debf0ec435f6321d71ce46af83248ee4e3860328e24faf1b602060243561094460018060a01b03600554163314610ba0565b335f526007825261095b60ff60405f205416610bed565b6001600160a01b038416936109aa908290610977871515610c45565b61098387301415610c85565b61098e821515610cd2565b305f525f85526109a48260405f20541015610d29565b30610dbc565b835f526006825260405f206109c0828254610d6e565b9055604051908152a2005b34610111576040366003190112610111576109e4610b46565b60243590331561085a576001600160a01b031690811561084757335f52600160205260405f20825f526020528060405f20556040519081527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560203392a3602060405160018152f35b34610111575f366003190112610111575f6003548060011c90600181168015610b12575b602083108114610471578285529081156104555750600114610abd5750819003601f01601f191681019067ffffffffffffffff8211818310176103ec576103e882918260405282610b1c565b905060035f527fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b5f905b828210610afc575060209150820101826103b8565b6001816020925483858801015201910190610ae7565b91607f1691610a71565b602060409281835280519182918282860152018484015e5f828201840152601f01601f1916010190565b600435906001600160a01b038216820361011157565b602435906001600160a01b038216820361011157565b9181601f840112156101115782359167ffffffffffffffff8311610111576020838186019501011161011157565b15610ba757565b60405163cc92809560e01b815260206004820152601c60248201527f4f6e6c792073746166662063616e206578656375746520746869732e000000006044820152606490fd5b15610bf457565b60405162b3883360e61b8152602060048201526024808201527f596f75206d757374206a6f696e20666972737420746f206578656375746520746044820152633434b99760e11b6064820152608490fd5b15610c4c57565b604051630b0f5aa160e11b815260206004820152601060248201526f24b73b30b634b21030b2323932b9b99760811b6044820152606490fd5b15610c8c57565b604051630b0f5aa160e11b815260206004820152601a60248201527f43616e6e6f74207472616e7366657220746f20697473656c662e0000000000006044820152606490fd5b15610cd957565b60405163589956a560e11b815260206004820152602160248201527f416d6f756e74206d7573742062652067726561746572207468616e207a65726f6044820152601760f91b6064820152608490fd5b15610d3057565b604051637979dc8760e01b815260206004820152601560248201527424b739bab33334b1b4b2b73a103130b630b731b29760591b6044820152606490fd5b91908201809211610d7b57565b634e487b7160e01b5f52601160045260245ffd5b908060209392818452848401375f828201840152601f01601f1916010190565b91908203918211610d7b57565b6001600160a01b0316908115610e53576001600160a01b03169182156106cd57815f525f60205260405f2054818110610e3a57817fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef92602092855f525f84520360405f2055845f525f825260405f20818154019055604051908152a3565b8263391434e360e21b5f5260045260245260445260645ffd5b634b637e8f60e11b5f525f60045260245ffdfea2646970667358221220c81cf0810f355aa09f5f8ce840a86fb97f9442c447094d4692b272aedb184eeb64736f6c634300081e0033",
}

// GalactoABI is the input ABI used to generate the binding from.
// Deprecated: Use GalactoMetaData.ABI instead.
var GalactoABI = GalactoMetaData.ABI

// GalactoBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GalactoMetaData.Bin instead.
var GalactoBin = GalactoMetaData.Bin

// DeployGalacto deploys a new Ethereum contract, binding an instance of Galacto to it.
func DeployGalacto(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Galacto, error) {
	parsed, err := GalactoMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GalactoBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Galacto{GalactoCaller: GalactoCaller{contract: contract}, GalactoTransactor: GalactoTransactor{contract: contract}, GalactoFilterer: GalactoFilterer{contract: contract}}, nil
}

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

// PublishAction is a paid mutator transaction binding the contract method 0x6739b9f4.
//
// Solidity: function PublishAction(string id, address userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoTransactor) PublishAction(opts *bind.TransactOpts, id string, userWallet common.Address, userID string, username string, actionID string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.contract.Transact(opts, "PublishAction", id, userWallet, userID, username, actionID, actionName, rewardAmount, performedAt)
}

// PublishAction is a paid mutator transaction binding the contract method 0x6739b9f4.
//
// Solidity: function PublishAction(string id, address userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoSession) PublishAction(id string, userWallet common.Address, userID string, username string, actionID string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.PublishAction(&_Galacto.TransactOpts, id, userWallet, userID, username, actionID, actionName, rewardAmount, performedAt)
}

// PublishAction is a paid mutator transaction binding the contract method 0x6739b9f4.
//
// Solidity: function PublishAction(string id, address userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt) returns()
func (_Galacto *GalactoTransactorSession) PublishAction(id string, userWallet common.Address, userID string, username string, actionID string, actionName string, rewardAmount *big.Int, performedAt *big.Int) (*types.Transaction, error) {
	return _Galacto.Contract.PublishAction(&_Galacto.TransactOpts, id, userWallet, userID, username, actionID, actionName, rewardAmount, performedAt)
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
	ActionID     string
	ActionName   string
	RewardAmount *big.Int
	PerformedAt  *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterActionPublished is a free log retrieval operation binding the contract event 0x747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c4.
//
// Solidity: event ActionPublished(string id, address indexed userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt)
func (_Galacto *GalactoFilterer) FilterActionPublished(opts *bind.FilterOpts, userWallet []common.Address) (*GalactoActionPublishedIterator, error) {

	var userWalletRule []interface{}
	for _, userWalletItem := range userWallet {
		userWalletRule = append(userWalletRule, userWalletItem)
	}

	logs, sub, err := _Galacto.contract.FilterLogs(opts, "ActionPublished", userWalletRule)
	if err != nil {
		return nil, err
	}
	return &GalactoActionPublishedIterator{contract: _Galacto.contract, event: "ActionPublished", logs: logs, sub: sub}, nil
}

// WatchActionPublished is a free log subscription operation binding the contract event 0x747490d6d248bd45e018ab467642f5ea04202b4dc9ce1d91eaea4b5a2610c4c4.
//
// Solidity: event ActionPublished(string id, address indexed userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt)
func (_Galacto *GalactoFilterer) WatchActionPublished(opts *bind.WatchOpts, sink chan<- *GalactoActionPublished, userWallet []common.Address) (event.Subscription, error) {

	var userWalletRule []interface{}
	for _, userWalletItem := range userWallet {
		userWalletRule = append(userWalletRule, userWalletItem)
	}

	logs, sub, err := _Galacto.contract.WatchLogs(opts, "ActionPublished", userWalletRule)
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
// Solidity: event ActionPublished(string id, address indexed userWallet, string userID, string username, string actionID, string actionName, uint256 rewardAmount, uint256 performedAt)
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
