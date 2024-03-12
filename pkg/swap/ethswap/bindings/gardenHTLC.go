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
	_ = abi.ConvertType
)

// GardenHTLCMetaData contains all meta data concerning the GardenHTLC contract.
var GardenHTLCMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token_\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"secret\",\"type\":\"bytes\"}],\"name\":\"Redeemed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"orderID\",\"type\":\"bytes32\"}],\"name\":\"Refunded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"}],\"name\":\"initiate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"secretHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"initiateWithSignature\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"orders\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isFulfilled\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initiatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"secret\",\"type\":\"bytes\"}],\"name\":\"redeem\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"orderID\",\"type\":\"bytes32\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x61018060405234801562000011575f80fd5b5060405162001bef38038062001bef833981016040819052620000349162000247565b818162000042825f620000fe565b6101205262000053816001620000fe565b61014052815160208084019190912060e052815190820120610100524660a052620000e060e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c05250506001600160a01b03166101605262000477565b5f6020835110156200011d57620001158362000136565b905062000130565b816200012a848262000357565b5060ff90505b92915050565b5f80829050601f815111156200016c578260405163305a27a960e01b81526004016200016391906200041f565b60405180910390fd5b8051620001798262000453565b179392505050565b634e487b7160e01b5f52604160045260245ffd5b5f5b83811015620001b157818101518382015260200162000197565b50505f910152565b5f82601f830112620001c9575f80fd5b81516001600160401b0380821115620001e657620001e662000181565b604051601f8301601f19908116603f0116810190828211818310171562000211576200021162000181565b816040528381528660208588010111156200022a575f80fd5b6200023d84602083016020890162000195565b9695505050505050565b5f805f606084860312156200025a575f80fd5b83516001600160a01b038116811462000271575f80fd5b60208501519093506001600160401b03808211156200028e575f80fd5b6200029c87838801620001b9565b93506040860151915080821115620002b2575f80fd5b50620002c186828701620001b9565b9150509250925092565b600181811c90821680620002e057607f821691505b602082108103620002ff57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000352575f81815260208120601f850160051c810160208610156200032d5750805b601f850160051c820191505b818110156200034e5782815560010162000339565b5050505b505050565b81516001600160401b0381111562000373576200037362000181565b6200038b81620003848454620002cb565b8462000305565b602080601f831160018114620003c1575f8415620003a95750858301515b5f19600386901b1c1916600185901b1785556200034e565b5f85815260208120601f198616915b82811015620003f157888601518255948401946001909101908401620003d0565b50858210156200040f57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b602081525f82518060208401526200043f81604085016020870162000195565b601f01601f19169190910160400192915050565b80516020808301519190811015620002ff575f1960209190910360031b1b16919050565b60805160a05160c05160e05161010051610120516101405161016051611707620004e85f395f8181610181015281816103250152818161087e0152610c1701525f61052301525f6104f901525f610ea001525f610e7801525f610dd301525f610dfd01525f610e2701526117075ff3fe608060405234801561000f575f80fd5b506004361061007a575f3560e01c806397ffc7ae1161005857806397ffc7ae146100ca5780639c3f1e90146100dd578063f7ff720714610169578063fc0c546a1461017c575f80fd5b80637249fbb61461007e5780637929d59d1461009357806384b0196e146100a6575b5f80fd5b61009161008c36600461136c565b6101bb565b005b6100916100a13660046113dc565b610358565b6100ae6104ec565b6040516100c19796959493929190611495565b60405180910390f35b6100916100d8366004611529565b610572565b61012e6100eb36600461136c565b600260208190525f91825260409091208054600182015492820154600383015460049093015460ff8316946001600160a01b036101009094048416949316929086565b6040805196151587526001600160a01b03958616602088015293909416928501929092526060840152608083015260a082015260c0016100c1565b61009161017736600461155f565b610643565b6101a37f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020016100c1565b5f81815260026020526040902060018101546001600160a01b03166102275760405162461bcd60e51b815260206004820152601f60248201527f47617264656e48544c433a206f72646572206e6f7420696e697469617465640060448201526064015b60405180910390fd5b805460ff16156102795760405162461bcd60e51b815260206004820152601b60248201527f47617264656e48544c433a206f726465722066756c66696c6c65640000000000604482015260640161021e565b438160030154826002015461028e91906115a7565b106102db5760405162461bcd60e51b815260206004820152601d60248201527f47617264656e48544c433a206f72646572206e6f742065787069726564000000604482015260640161021e565b805460ff1916600117815560405182907ffe509803c09416b28ff3d8f690c8b0c61462a892c46d5430c8fb20abe472daf0905f90a280546004820154610354916001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081169261010090920416906108af565b5050565b8585856001600160a01b0383166103815760405162461bcd60e51b815260040161021e906115c6565b5f82116103ca5760405162461bcd60e51b815260206004820152601760248201527647617264656e48544c433a207a65726f2065787069727960481b604482015260640161021e565b5f81116104135760405162461bcd60e51b815260206004820152601760248201527611d85c99195b92151310ce881e995c9bc8185b5bdd5b9d604a1b604482015260640161021e565b5f6104d186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250506040516104cb92506104b091507f1e45abfc3242e55884a137cdd9336f2a0dd2f4cba3cf9934a5d0d5d0eb43f36c908f908f908f908f906020019485526001600160a01b0393909316602085015260408401919091526060830152608082015260a00190565b60405160208183030381529060405280519060200120610917565b90610949565b90506104e0818b8b8b8b61096b565b50505050505050505050565b5f6060808280808361051e7f000000000000000000000000000000000000000000000000000000000000000083610c4b565b6105497f00000000000000000000000000000000000000000000000000000000000000006001610c4b565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b8383836001600160a01b03831661059b5760405162461bcd60e51b815260040161021e906115c6565b5f82116105e45760405162461bcd60e51b815260206004820152601760248201527647617264656e48544c433a207a65726f2065787069727960481b604482015260640161021e565b5f811161062d5760405162461bcd60e51b815260206004820152601760248201527611d85c99195b92151310ce881e995c9bc8185b5bdd5b9d604a1b604482015260640161021e565b61063a338888888861096b565b50505050505050565b5f83815260026020526040902060018101546001600160a01b03166106aa5760405162461bcd60e51b815260206004820152601f60248201527f47617264656e48544c433a206f72646572206e6f7420696e6974696174656400604482015260640161021e565b805460ff16156106fc5760405162461bcd60e51b815260206004820152601b60248201527f47617264656e48544c433a206f726465722066756c66696c6c65640000000000604482015260640161021e565b5f6002848460405161070f929190611607565b602060405180830381855afa15801561072a573d5f803e3d5ffd5b5050506040513d601f19601f8201168201806040525081019061074d9190611616565b825460408051602081018490526101009092046001600160a01b031690820152909150859060029060600160408051601f19818403018152908290526107929161162d565b602060405180830381855afa1580156107ad573d5f803e3d5ffd5b5050506040513d601f19601f820116820180604052508101906107d09190611616565b1461081d5760405162461bcd60e51b815260206004820152601c60248201527f47617264656e48544c433a20696e636f72726563742073656372657400000000604482015260640161021e565b815460ff19166001178255604051819086907f4c9a044220477b4e94dbb0d07ff6ff4ac30d443bef59098c4541b006954778e29061085e9088908890611648565b60405180910390a3600182015460048301546108a8916001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116929116906108af565b5050505050565b6040516001600160a01b03831660248201526044810182905261091290849063a9059cbb60e01b906064015b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152610cf4565b505050565b5f610943610923610dc7565b8360405161190160f01b8152600281019290925260228201526042902090565b92915050565b5f805f6109568585610ef5565b9150915061096381610f37565b509392505050565b836001600160a01b0316856001600160a01b0316036109dc5760405162461bcd60e51b815260206004820152602760248201527f47617264656e48544c433a2073616d6520696e69746961746f7220616e64207260448201526632b232b2b6b2b960c91b606482015260840161021e565b5f60028287604051602001610a049291909182526001600160a01b0316602082015260400190565b60408051601f1981840301815290829052610a1e9161162d565b602060405180830381855afa158015610a39573d5f803e3d5ffd5b5050506040513d601f19601f82011682018060405250810190610a5c9190611616565b5f81815260026020818152604092839020835160c081018552815460ff8116151582526001600160a01b0361010090910481169382019390935260018201549092169382018490529182015460608201526003820154608082015260049091015460a082015291925015610b125760405162461bcd60e51b815260206004820152601b60248201527f47617264656e48544c433a206475706c6963617465206f726465720000000000604482015260640161021e565b6040805160c0810182525f8082526001600160a01b038a811660208085019182528b83168587019081524360608701908152608087018d815260a088018d81528b88526002808652978a90208951815497516001600160a81b0319909816901515610100600160a81b031916176101009789169790970296909617865592516001860180546001600160a01b03191691909716179095555194830194909455915160038201559151600490920182905592519081529091859185917f01b41cbd4bbcc3c5b968a04d3fbdd8c1648a39ff6d9a3929b4840cea1142bc65910160405180910390a35f83815260026020526040902060040154610c41906001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016908a903090611083565b5050505050505050565b606060ff8314610c6557610c5e836110c1565b9050610943565b818054610c7190611676565b80601f0160208091040260200160405190810160405280929190818152602001828054610c9d90611676565b8015610ce85780601f10610cbf57610100808354040283529160200191610ce8565b820191905f5260205f20905b815481529060010190602001808311610ccb57829003601f168201915b50505050509050610943565b5f610d48826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166110fe9092919063ffffffff16565b905080515f1480610d68575080806020019051810190610d6891906116ae565b6109125760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b606482015260840161021e565b5f306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148015610e1f57507f000000000000000000000000000000000000000000000000000000000000000046145b15610e4957507f000000000000000000000000000000000000000000000000000000000000000090565b610ef0604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b905090565b5f808251604103610f29576020830151604084015160608501515f1a610f1d87828585611114565b94509450505050610f30565b505f905060025b9250929050565b5f816004811115610f4a57610f4a6116d4565b03610f525750565b6001816004811115610f6657610f666116d4565b03610fb35760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e61747572650000000000000000604482015260640161021e565b6002816004811115610fc757610fc76116d4565b036110145760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e67746800604482015260640161021e565b6003816004811115611028576110286116d4565b036110805760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b606482015260840161021e565b50565b6040516001600160a01b03808516602483015283166044820152606481018290526110bb9085906323b872dd60e01b906084016108db565b50505050565b60605f6110cd836111d1565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b606061110c84845f856111f8565b949350505050565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a083111561114957505f905060036111c8565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa15801561119a573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b0381166111c2575f600192509250506111c8565b91505f90505b94509492505050565b5f60ff8216601f81111561094357604051632cd44ac360e21b815260040160405180910390fd5b6060824710156112595760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b606482015260840161021e565b5f80866001600160a01b03168587604051611274919061162d565b5f6040518083038185875af1925050503d805f81146112ae576040519150601f19603f3d011682016040523d82523d5f602084013e6112b3565b606091505b50915091506112c4878383876112cf565b979650505050505050565b6060831561133d5782515f03611336576001600160a01b0385163b6113365760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604482015260640161021e565b508161110c565b61110c83838151156113525781518083602001fd5b8060405162461bcd60e51b815260040161021e91906116e8565b5f6020828403121561137c575f80fd5b5035919050565b80356001600160a01b0381168114611399575f80fd5b919050565b5f8083601f8401126113ae575f80fd5b50813567ffffffffffffffff8111156113c5575f80fd5b602083019150836020828501011115610f30575f80fd5b5f805f805f8060a087890312156113f1575f80fd5b6113fa87611383565b9550602087013594506040870135935060608701359250608087013567ffffffffffffffff81111561142a575f80fd5b61143689828a0161139e565b979a9699509497509295939492505050565b5f5b8381101561146257818101518382015260200161144a565b50505f910152565b5f8151808452611481816020860160208601611448565b601f01601f19169290920160200192915050565b60ff60f81b881681525f602060e0818401526114b460e084018a61146a565b83810360408501526114c6818a61146a565b606085018990526001600160a01b038816608086015260a0850187905284810360c086015285518082528387019250908301905f5b81811015611517578351835292840192918401916001016114fb565b50909c9b505050505050505050505050565b5f805f806080858703121561153c575f80fd5b61154585611383565b966020860135965060408601359560600135945092505050565b5f805f60408486031215611571575f80fd5b83359250602084013567ffffffffffffffff81111561158e575f80fd5b61159a8682870161139e565b9497909650939450505050565b8082018082111561094357634e487b7160e01b5f52601160045260245ffd5b60208082526021908201527f47617264656e48544c433a207a65726f20616464726573732072656465656d656040820152603960f91b606082015260800190565b818382375f9101908152919050565b5f60208284031215611626575f80fd5b5051919050565b5f825161163e818460208701611448565b9190910192915050565b60208152816020820152818360408301375f818301604090810191909152601f909201601f19160101919050565b600181811c9082168061168a57607f821691505b6020821081036116a857634e487b7160e01b5f52602260045260245ffd5b50919050565b5f602082840312156116be575f80fd5b815180151581146116cd575f80fd5b9392505050565b634e487b7160e01b5f52602160045260245ffd5b602081525f6116cd602083018461146a56fea164736f6c6343000815000a",
}

// GardenHTLCABI is the input ABI used to generate the binding from.
// Deprecated: Use GardenHTLCMetaData.ABI instead.
var GardenHTLCABI = GardenHTLCMetaData.ABI

// GardenHTLCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GardenHTLCMetaData.Bin instead.
var GardenHTLCBin = GardenHTLCMetaData.Bin

// DeployGardenHTLC deploys a new Ethereum contract, binding an instance of GardenHTLC to it.
func DeployGardenHTLC(auth *bind.TransactOpts, backend bind.ContractBackend, token_ common.Address, name string, version string) (common.Address, *types.Transaction, *GardenHTLC, error) {
	parsed, err := GardenHTLCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GardenHTLCBin), backend, token_, name, version)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GardenHTLC{GardenHTLCCaller: GardenHTLCCaller{contract: contract}, GardenHTLCTransactor: GardenHTLCTransactor{contract: contract}, GardenHTLCFilterer: GardenHTLCFilterer{contract: contract}}, nil
}

// GardenHTLC is an auto generated Go binding around an Ethereum contract.
type GardenHTLC struct {
	GardenHTLCCaller     // Read-only binding to the contract
	GardenHTLCTransactor // Write-only binding to the contract
	GardenHTLCFilterer   // Log filterer for contract events
}

// GardenHTLCCaller is an auto generated read-only Go binding around an Ethereum contract.
type GardenHTLCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenHTLCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GardenHTLCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenHTLCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GardenHTLCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GardenHTLCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GardenHTLCSession struct {
	Contract     *GardenHTLC       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GardenHTLCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GardenHTLCCallerSession struct {
	Contract *GardenHTLCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// GardenHTLCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GardenHTLCTransactorSession struct {
	Contract     *GardenHTLCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GardenHTLCRaw is an auto generated low-level Go binding around an Ethereum contract.
type GardenHTLCRaw struct {
	Contract *GardenHTLC // Generic contract binding to access the raw methods on
}

// GardenHTLCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GardenHTLCCallerRaw struct {
	Contract *GardenHTLCCaller // Generic read-only contract binding to access the raw methods on
}

// GardenHTLCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GardenHTLCTransactorRaw struct {
	Contract *GardenHTLCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGardenHTLC creates a new instance of GardenHTLC, bound to a specific deployed contract.
func NewGardenHTLC(address common.Address, backend bind.ContractBackend) (*GardenHTLC, error) {
	contract, err := bindGardenHTLC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GardenHTLC{GardenHTLCCaller: GardenHTLCCaller{contract: contract}, GardenHTLCTransactor: GardenHTLCTransactor{contract: contract}, GardenHTLCFilterer: GardenHTLCFilterer{contract: contract}}, nil
}

// NewGardenHTLCCaller creates a new read-only instance of GardenHTLC, bound to a specific deployed contract.
func NewGardenHTLCCaller(address common.Address, caller bind.ContractCaller) (*GardenHTLCCaller, error) {
	contract, err := bindGardenHTLC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCCaller{contract: contract}, nil
}

// NewGardenHTLCTransactor creates a new write-only instance of GardenHTLC, bound to a specific deployed contract.
func NewGardenHTLCTransactor(address common.Address, transactor bind.ContractTransactor) (*GardenHTLCTransactor, error) {
	contract, err := bindGardenHTLC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCTransactor{contract: contract}, nil
}

// NewGardenHTLCFilterer creates a new log filterer instance of GardenHTLC, bound to a specific deployed contract.
func NewGardenHTLCFilterer(address common.Address, filterer bind.ContractFilterer) (*GardenHTLCFilterer, error) {
	contract, err := bindGardenHTLC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCFilterer{contract: contract}, nil
}

// bindGardenHTLC binds a generic wrapper to an already deployed contract.
func bindGardenHTLC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GardenHTLCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GardenHTLC *GardenHTLCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GardenHTLC.Contract.GardenHTLCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GardenHTLC *GardenHTLCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GardenHTLC.Contract.GardenHTLCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GardenHTLC *GardenHTLCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GardenHTLC.Contract.GardenHTLCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GardenHTLC *GardenHTLCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GardenHTLC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GardenHTLC *GardenHTLCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GardenHTLC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GardenHTLC *GardenHTLCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GardenHTLC.Contract.contract.Transact(opts, method, params...)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GardenHTLC *GardenHTLCCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GardenHTLC.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GardenHTLC *GardenHTLCSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GardenHTLC.Contract.Eip712Domain(&_GardenHTLC.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GardenHTLC *GardenHTLCCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GardenHTLC.Contract.Eip712Domain(&_GardenHTLC.CallOpts)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bool isFulfilled, address initiator, address redeemer, uint256 initiatedAt, uint256 expiry, uint256 amount)
func (_GardenHTLC *GardenHTLCCaller) Orders(opts *bind.CallOpts, arg0 [32]byte) (struct {
	IsFulfilled bool
	Initiator   common.Address
	Redeemer    common.Address
	InitiatedAt *big.Int
	Expiry      *big.Int
	Amount      *big.Int
}, error) {
	var out []interface{}
	err := _GardenHTLC.contract.Call(opts, &out, "orders", arg0)

	outstruct := new(struct {
		IsFulfilled bool
		Initiator   common.Address
		Redeemer    common.Address
		InitiatedAt *big.Int
		Expiry      *big.Int
		Amount      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsFulfilled = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Initiator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Redeemer = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.InitiatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Expiry = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bool isFulfilled, address initiator, address redeemer, uint256 initiatedAt, uint256 expiry, uint256 amount)
func (_GardenHTLC *GardenHTLCSession) Orders(arg0 [32]byte) (struct {
	IsFulfilled bool
	Initiator   common.Address
	Redeemer    common.Address
	InitiatedAt *big.Int
	Expiry      *big.Int
	Amount      *big.Int
}, error) {
	return _GardenHTLC.Contract.Orders(&_GardenHTLC.CallOpts, arg0)
}

// Orders is a free data retrieval call binding the contract method 0x9c3f1e90.
//
// Solidity: function orders(bytes32 ) view returns(bool isFulfilled, address initiator, address redeemer, uint256 initiatedAt, uint256 expiry, uint256 amount)
func (_GardenHTLC *GardenHTLCCallerSession) Orders(arg0 [32]byte) (struct {
	IsFulfilled bool
	Initiator   common.Address
	Redeemer    common.Address
	InitiatedAt *big.Int
	Expiry      *big.Int
	Amount      *big.Int
}, error) {
	return _GardenHTLC.Contract.Orders(&_GardenHTLC.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GardenHTLC *GardenHTLCCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GardenHTLC.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GardenHTLC *GardenHTLCSession) Token() (common.Address, error) {
	return _GardenHTLC.Contract.Token(&_GardenHTLC.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GardenHTLC *GardenHTLCCallerSession) Token() (common.Address, error) {
	return _GardenHTLC.Contract.Token(&_GardenHTLC.CallOpts)
}

// Initiate is a paid mutator transaction binding the contract method 0x97ffc7ae.
//
// Solidity: function initiate(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash) returns()
func (_GardenHTLC *GardenHTLCTransactor) Initiate(opts *bind.TransactOpts, redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.contract.Transact(opts, "initiate", redeemer, expiry, amount, secretHash)
}

// Initiate is a paid mutator transaction binding the contract method 0x97ffc7ae.
//
// Solidity: function initiate(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash) returns()
func (_GardenHTLC *GardenHTLCSession) Initiate(redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Initiate(&_GardenHTLC.TransactOpts, redeemer, expiry, amount, secretHash)
}

// Initiate is a paid mutator transaction binding the contract method 0x97ffc7ae.
//
// Solidity: function initiate(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash) returns()
func (_GardenHTLC *GardenHTLCTransactorSession) Initiate(redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Initiate(&_GardenHTLC.TransactOpts, redeemer, expiry, amount, secretHash)
}

// InitiateWithSignature is a paid mutator transaction binding the contract method 0x7929d59d.
//
// Solidity: function initiateWithSignature(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash, bytes signature) returns()
func (_GardenHTLC *GardenHTLCTransactor) InitiateWithSignature(opts *bind.TransactOpts, redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _GardenHTLC.contract.Transact(opts, "initiateWithSignature", redeemer, expiry, amount, secretHash, signature)
}

// InitiateWithSignature is a paid mutator transaction binding the contract method 0x7929d59d.
//
// Solidity: function initiateWithSignature(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash, bytes signature) returns()
func (_GardenHTLC *GardenHTLCSession) InitiateWithSignature(redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.InitiateWithSignature(&_GardenHTLC.TransactOpts, redeemer, expiry, amount, secretHash, signature)
}

// InitiateWithSignature is a paid mutator transaction binding the contract method 0x7929d59d.
//
// Solidity: function initiateWithSignature(address redeemer, uint256 expiry, uint256 amount, bytes32 secretHash, bytes signature) returns()
func (_GardenHTLC *GardenHTLCTransactorSession) InitiateWithSignature(redeemer common.Address, expiry *big.Int, amount *big.Int, secretHash [32]byte, signature []byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.InitiateWithSignature(&_GardenHTLC.TransactOpts, redeemer, expiry, amount, secretHash, signature)
}

// Redeem is a paid mutator transaction binding the contract method 0xf7ff7207.
//
// Solidity: function redeem(bytes32 orderID, bytes secret) returns()
func (_GardenHTLC *GardenHTLCTransactor) Redeem(opts *bind.TransactOpts, orderID [32]byte, secret []byte) (*types.Transaction, error) {
	return _GardenHTLC.contract.Transact(opts, "redeem", orderID, secret)
}

// Redeem is a paid mutator transaction binding the contract method 0xf7ff7207.
//
// Solidity: function redeem(bytes32 orderID, bytes secret) returns()
func (_GardenHTLC *GardenHTLCSession) Redeem(orderID [32]byte, secret []byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Redeem(&_GardenHTLC.TransactOpts, orderID, secret)
}

// Redeem is a paid mutator transaction binding the contract method 0xf7ff7207.
//
// Solidity: function redeem(bytes32 orderID, bytes secret) returns()
func (_GardenHTLC *GardenHTLCTransactorSession) Redeem(orderID [32]byte, secret []byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Redeem(&_GardenHTLC.TransactOpts, orderID, secret)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 orderID) returns()
func (_GardenHTLC *GardenHTLCTransactor) Refund(opts *bind.TransactOpts, orderID [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.contract.Transact(opts, "refund", orderID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 orderID) returns()
func (_GardenHTLC *GardenHTLCSession) Refund(orderID [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Refund(&_GardenHTLC.TransactOpts, orderID)
}

// Refund is a paid mutator transaction binding the contract method 0x7249fbb6.
//
// Solidity: function refund(bytes32 orderID) returns()
func (_GardenHTLC *GardenHTLCTransactorSession) Refund(orderID [32]byte) (*types.Transaction, error) {
	return _GardenHTLC.Contract.Refund(&_GardenHTLC.TransactOpts, orderID)
}

// GardenHTLCEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GardenHTLC contract.
type GardenHTLCEIP712DomainChangedIterator struct {
	Event *GardenHTLCEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *GardenHTLCEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenHTLCEIP712DomainChanged)
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
		it.Event = new(GardenHTLCEIP712DomainChanged)
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
func (it *GardenHTLCEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenHTLCEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenHTLCEIP712DomainChanged represents a EIP712DomainChanged event raised by the GardenHTLC contract.
type GardenHTLCEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GardenHTLC *GardenHTLCFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GardenHTLCEIP712DomainChangedIterator, error) {

	logs, sub, err := _GardenHTLC.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GardenHTLCEIP712DomainChangedIterator{contract: _GardenHTLC.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GardenHTLC *GardenHTLCFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GardenHTLCEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GardenHTLC.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenHTLCEIP712DomainChanged)
				if err := _GardenHTLC.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GardenHTLC *GardenHTLCFilterer) ParseEIP712DomainChanged(log types.Log) (*GardenHTLCEIP712DomainChanged, error) {
	event := new(GardenHTLCEIP712DomainChanged)
	if err := _GardenHTLC.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenHTLCInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the GardenHTLC contract.
type GardenHTLCInitiatedIterator struct {
	Event *GardenHTLCInitiated // Event containing the contract specifics and raw log

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
func (it *GardenHTLCInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenHTLCInitiated)
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
		it.Event = new(GardenHTLCInitiated)
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
func (it *GardenHTLCInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenHTLCInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenHTLCInitiated represents a Initiated event raised by the GardenHTLC contract.
type GardenHTLCInitiated struct {
	OrderID    [32]byte
	SecretHash [32]byte
	Amount     *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0x01b41cbd4bbcc3c5b968a04d3fbdd8c1648a39ff6d9a3929b4840cea1142bc65.
//
// Solidity: event Initiated(bytes32 indexed orderID, bytes32 indexed secretHash, uint256 amount)
func (_GardenHTLC *GardenHTLCFilterer) FilterInitiated(opts *bind.FilterOpts, orderID [][32]byte, secretHash [][32]byte) (*GardenHTLCInitiatedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _GardenHTLC.contract.FilterLogs(opts, "Initiated", orderIDRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCInitiatedIterator{contract: _GardenHTLC.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0x01b41cbd4bbcc3c5b968a04d3fbdd8c1648a39ff6d9a3929b4840cea1142bc65.
//
// Solidity: event Initiated(bytes32 indexed orderID, bytes32 indexed secretHash, uint256 amount)
func (_GardenHTLC *GardenHTLCFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *GardenHTLCInitiated, orderID [][32]byte, secretHash [][32]byte) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _GardenHTLC.contract.WatchLogs(opts, "Initiated", orderIDRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenHTLCInitiated)
				if err := _GardenHTLC.contract.UnpackLog(event, "Initiated", log); err != nil {
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

// ParseInitiated is a log parse operation binding the contract event 0x01b41cbd4bbcc3c5b968a04d3fbdd8c1648a39ff6d9a3929b4840cea1142bc65.
//
// Solidity: event Initiated(bytes32 indexed orderID, bytes32 indexed secretHash, uint256 amount)
func (_GardenHTLC *GardenHTLCFilterer) ParseInitiated(log types.Log) (*GardenHTLCInitiated, error) {
	event := new(GardenHTLCInitiated)
	if err := _GardenHTLC.contract.UnpackLog(event, "Initiated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenHTLCRedeemedIterator is returned from FilterRedeemed and is used to iterate over the raw logs and unpacked data for Redeemed events raised by the GardenHTLC contract.
type GardenHTLCRedeemedIterator struct {
	Event *GardenHTLCRedeemed // Event containing the contract specifics and raw log

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
func (it *GardenHTLCRedeemedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenHTLCRedeemed)
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
		it.Event = new(GardenHTLCRedeemed)
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
func (it *GardenHTLCRedeemedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenHTLCRedeemedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenHTLCRedeemed represents a Redeemed event raised by the GardenHTLC contract.
type GardenHTLCRedeemed struct {
	OrderID    [32]byte
	SecretHash [32]byte
	Secret     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterRedeemed is a free log retrieval operation binding the contract event 0x4c9a044220477b4e94dbb0d07ff6ff4ac30d443bef59098c4541b006954778e2.
//
// Solidity: event Redeemed(bytes32 indexed orderID, bytes32 indexed secretHash, bytes secret)
func (_GardenHTLC *GardenHTLCFilterer) FilterRedeemed(opts *bind.FilterOpts, orderID [][32]byte, secretHash [][32]byte) (*GardenHTLCRedeemedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _GardenHTLC.contract.FilterLogs(opts, "Redeemed", orderIDRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCRedeemedIterator{contract: _GardenHTLC.contract, event: "Redeemed", logs: logs, sub: sub}, nil
}

// WatchRedeemed is a free log subscription operation binding the contract event 0x4c9a044220477b4e94dbb0d07ff6ff4ac30d443bef59098c4541b006954778e2.
//
// Solidity: event Redeemed(bytes32 indexed orderID, bytes32 indexed secretHash, bytes secret)
func (_GardenHTLC *GardenHTLCFilterer) WatchRedeemed(opts *bind.WatchOpts, sink chan<- *GardenHTLCRedeemed, orderID [][32]byte, secretHash [][32]byte) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}
	var secretHashRule []interface{}
	for _, secretHashItem := range secretHash {
		secretHashRule = append(secretHashRule, secretHashItem)
	}

	logs, sub, err := _GardenHTLC.contract.WatchLogs(opts, "Redeemed", orderIDRule, secretHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenHTLCRedeemed)
				if err := _GardenHTLC.contract.UnpackLog(event, "Redeemed", log); err != nil {
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

// ParseRedeemed is a log parse operation binding the contract event 0x4c9a044220477b4e94dbb0d07ff6ff4ac30d443bef59098c4541b006954778e2.
//
// Solidity: event Redeemed(bytes32 indexed orderID, bytes32 indexed secretHash, bytes secret)
func (_GardenHTLC *GardenHTLCFilterer) ParseRedeemed(log types.Log) (*GardenHTLCRedeemed, error) {
	event := new(GardenHTLCRedeemed)
	if err := _GardenHTLC.contract.UnpackLog(event, "Redeemed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GardenHTLCRefundedIterator is returned from FilterRefunded and is used to iterate over the raw logs and unpacked data for Refunded events raised by the GardenHTLC contract.
type GardenHTLCRefundedIterator struct {
	Event *GardenHTLCRefunded // Event containing the contract specifics and raw log

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
func (it *GardenHTLCRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GardenHTLCRefunded)
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
		it.Event = new(GardenHTLCRefunded)
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
func (it *GardenHTLCRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GardenHTLCRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GardenHTLCRefunded represents a Refunded event raised by the GardenHTLC contract.
type GardenHTLCRefunded struct {
	OrderID [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRefunded is a free log retrieval operation binding the contract event 0xfe509803c09416b28ff3d8f690c8b0c61462a892c46d5430c8fb20abe472daf0.
//
// Solidity: event Refunded(bytes32 indexed orderID)
func (_GardenHTLC *GardenHTLCFilterer) FilterRefunded(opts *bind.FilterOpts, orderID [][32]byte) (*GardenHTLCRefundedIterator, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _GardenHTLC.contract.FilterLogs(opts, "Refunded", orderIDRule)
	if err != nil {
		return nil, err
	}
	return &GardenHTLCRefundedIterator{contract: _GardenHTLC.contract, event: "Refunded", logs: logs, sub: sub}, nil
}

// WatchRefunded is a free log subscription operation binding the contract event 0xfe509803c09416b28ff3d8f690c8b0c61462a892c46d5430c8fb20abe472daf0.
//
// Solidity: event Refunded(bytes32 indexed orderID)
func (_GardenHTLC *GardenHTLCFilterer) WatchRefunded(opts *bind.WatchOpts, sink chan<- *GardenHTLCRefunded, orderID [][32]byte) (event.Subscription, error) {

	var orderIDRule []interface{}
	for _, orderIDItem := range orderID {
		orderIDRule = append(orderIDRule, orderIDItem)
	}

	logs, sub, err := _GardenHTLC.contract.WatchLogs(opts, "Refunded", orderIDRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GardenHTLCRefunded)
				if err := _GardenHTLC.contract.UnpackLog(event, "Refunded", log); err != nil {
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

// ParseRefunded is a log parse operation binding the contract event 0xfe509803c09416b28ff3d8f690c8b0c61462a892c46d5430c8fb20abe472daf0.
//
// Solidity: event Refunded(bytes32 indexed orderID)
func (_GardenHTLC *GardenHTLCFilterer) ParseRefunded(log types.Log) (*GardenHTLCRefunded, error) {
	event := new(GardenHTLCRefunded)
	if err := _GardenHTLC.contract.UnpackLog(event, "Refunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
