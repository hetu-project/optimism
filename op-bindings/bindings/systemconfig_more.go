// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SystemConfigStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initialized\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":1001,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_initializing\",\"offset\":1,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":1002,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":1003,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":1004,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":1005,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"overhead\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_uint256\"},{\"astId\":1006,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"scalar\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_uint256\"},{\"astId\":1007,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"batcherHash\",\"offset\":0,\"slot\":\"103\",\"type\":\"t_bytes32\"},{\"astId\":1008,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"gasLimit\",\"offset\":0,\"slot\":\"104\",\"type\":\"t_uint64\"},{\"astId\":1009,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"_resourceConfig\",\"offset\":0,\"slot\":\"105\",\"type\":\"t_struct(ResourceConfig)1011_storage\"},{\"astId\":1010,\"contract\":\"src/L1/SystemConfig.sol:SystemConfig\",\"label\":\"startBlock\",\"offset\":0,\"slot\":\"106\",\"type\":\"t_uint256\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\",\"base\":\"t_uint256\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_struct(ResourceConfig)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct ResourceMetering.ResourceConfig\",\"numberOfBytes\":\"32\"},\"t_uint128\":{\"encoding\":\"inplace\",\"label\":\"uint128\",\"numberOfBytes\":\"16\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_uint64\":{\"encoding\":\"inplace\",\"label\":\"uint64\",\"numberOfBytes\":\"8\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var SystemConfigStorageLayout = new(solc.StorageLayout)

var SystemConfigDeployedBin = "0x608060405234801561001057600080fd5b50600436106102415760003560e01c8063935f029e11610145578063cc731b02116100bd578063f45e65d81161008c578063f8c68de011610071578063f8c68de014610598578063fd32aa0f146105a0578063ffa1ad74146105a857600080fd5b8063f45e65d81461057b578063f68016b71461058457600080fd5b8063cc731b0214610423578063dac6e63a14610557578063e81b2c6d1461055f578063f2fde38b1461056857600080fd5b8063bc49ce5f11610114578063c71973f6116100f9578063c71973f6146103f5578063c9b26f6114610408578063ca3ea2011461041b57600080fd5b8063bc49ce5f146103e5578063c4e8ddfa146103ed57600080fd5b8063935f029e146103af5780639b7d7f0a146103c2578063a7119869146103ca578063b40a817c146103d257600080fd5b80634add321d116101d857806354fd4d50116101a757806361d157681161018c57806361d1576814610381578063715018a6146103895780638da5cb5b1461039157600080fd5b806354fd4d50146103305780635d73369c1461037957600080fd5b80634add321d146102cd5780634d9f1559146102ee5780634f16540b146102f65780635228a6ac1461031d57600080fd5b806318d139181161021457806318d139181461029f57806319f5cea8146102b45780631fd19ee1146102bc57806348cd4cb1146102c457600080fd5b806306c9265714610246578063078f29cf146102615780630a49cb031461028e5780630c18c16214610296575b600080fd5b61024e6105b0565b6040519081526020015b60405180910390f35b6102696105de565b60405173ffffffffffffffffffffffffffffffffffffffff9091168152602001610258565b610269610617565b61024e60655481565b6102b26102ad3660046116c5565b610647565b005b61024e61065b565b610269610686565b61024e606a5481565b6102d56106b0565b60405167ffffffffffffffff9091168152602001610258565b6102696106d6565b61024e7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c0881565b6102b261032b366004611855565b610706565b61036c6040518060400160405280600681526020017f312e31302e30000000000000000000000000000000000000000000000000000081525081565b60405161025891906119f8565b61024e610a9c565b61024e610ac7565b6102b2610af2565b60335473ffffffffffffffffffffffffffffffffffffffff16610269565b6102b26103bd366004611a0b565b610b06565b610269610b1c565b610269610b4c565b6102b26103e0366004611a2d565b610b7c565b61024e610b8d565b610269610bb8565b6102b2610403366004611a48565b610be8565b6102b2610416366004611a64565b610bf9565b61024e610c0a565b6104e76040805160c081018252600080825260208201819052918101829052606081018290526080810182905260a0810191909152506040805160c08101825260695463ffffffff8082168352640100000000820460ff9081166020850152650100000000008304169383019390935266010000000000008104831660608301526a0100000000000000000000810490921660808201526e0100000000000000000000000000009091046fffffffffffffffffffffffffffffffff1660a082015290565b6040516102589190600060c08201905063ffffffff80845116835260ff602085015116602084015260ff6040850151166040840152806060850151166060840152806080850151166080840152506fffffffffffffffffffffffffffffffff60a08401511660a083015292915050565b610269610c35565b61024e60675481565b6102b26105763660046116c5565b610c65565b61024e60665481565b6068546102d59067ffffffffffffffff1681565b61024e610d19565b61024e610d44565b61024e600081565b6105db60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611aac565b81565b600061061261060e60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611aac565b5490565b905090565b600061061261060e60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611aac565b61064f610d73565b61065881610df4565b50565b6105db60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611aac565b60006106127f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c085490565b6069546000906106129063ffffffff6a0100000000000000000000820481169116611ac3565b600061061261060e60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611aac565b600054600390610100900460ff16158015610728575060005460ff8083169116105b6107b9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a656400000000000000000000000000000000000060648201526084015b60405180910390fd5b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00001660ff8316176101001790556107f2610eb1565b6107fb8b610c65565b61080488610f50565b61080e8a8a610f78565b61081787611009565b61082086610df4565b61085361084e60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611aac565b849055565b61088761088160017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611aac565b83519055565b6108be6108b560017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611aac565b60208401519055565b6108f56108ec60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611aac565b60408401519055565b61092c61092360017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611aac565b60608401519055565b61096361095a60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611aac565b60808401519055565b61099a61099160017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611aac565b60a08401519055565b6109a3846110e7565b6109ac85611111565b6109b46106b0565b67ffffffffffffffff168767ffffffffffffffff161015610a31576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016107b0565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16905560405160ff821681527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15050505050505050505050565b6105db60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611aac565b6105db60017fe52a667f71ec761b9b381c7b76ca9b852adf7e8905da0e0ad49986a0a6871816611aac565b610afa610d73565b610b046000611585565b565b610b0e610d73565b610b188282610f78565b5050565b600061061261060e60017fa04c5bb938ca6fc46d95553abf0a76345ce3e722a30bf4f74928b8e7d852320d611aac565b600061061261060e60017f383f291819e6d54073bc9a648251d97421076bdd101933c0c022219ce9580637611aac565b610b84610d73565b61065881611009565b6105db60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611aac565b600061061261060e60017f46adcbebc6be8ce551740c29c47c8798210f23f7f4086c41752944352568d5a8611aac565b610bf0610d73565b61065881611111565b610c01610d73565b61065881610f50565b6105db60017f5574e74ed6237feee050139e5d3796c837cd70baa1eb558f13c1d72e3ca80205611aac565b600061061261060e60017f71ac12829d66ee73d8d95bff50b3589745ce57edae70a3fb111a2342464dc598611aac565b610c6d610d73565b73ffffffffffffffffffffffffffffffffffffffff8116610d10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f646472657373000000000000000000000000000000000000000000000000000060648201526084016107b0565b61065881611585565b6105db60017f9904ba90dde5696cda05c9e0dab5cbaa0fea005ace4d11218a02ac668dad6377611aac565b6105db60017f4b6c74f9e688cb39801f2112c14a8c57232a3fc5202e1444126d4bce86eb19ad611aac565b9055565b60335473ffffffffffffffffffffffffffffffffffffffff163314610b04576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016107b0565b610e1d7f65a7ed542fb37fe237fdfbdd70b31598523fe5b32879e307bae27a0bd9581c08829055565b6040805173ffffffffffffffffffffffffffffffffffffffff8316602082015260009101604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152919052905060035b60007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610ea591906119f8565b60405180910390a35050565b600054610100900460ff16610f48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107b0565b610b046115fc565b6067819055604080516020808201849052825180830390910181529082019091526000610e74565b606582905560668190556040805160208101849052908101829052600090606001604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190529050600160007f1d2b0bda21d56b8bd12d4f94ebacffdfb35f5e226f84b461103bb8beab6353be83604051610ffc91906119f8565b60405180910390a3505050565b6110116106b0565b67ffffffffffffffff168167ffffffffffffffff16101561108e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016107b0565b606880547fffffffffffffffffffffffffffffffffffffffffffffffff00000000000000001667ffffffffffffffff83169081179091556040805160208082019390935281518082039093018352810190526002610e74565b80158015906110f65750606a54155b1561110057606a55565b606a546000036106585743606a5550565b8060a001516fffffffffffffffffffffffffffffffff16816060015163ffffffff1611156111c1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f53797374656d436f6e6669673a206d696e206261736520666565206d7573742060448201527f6265206c657373207468616e206d61782062617365000000000000000000000060648201526084016107b0565b6001816040015160ff1611611258576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a2064656e6f6d696e61746f72206d757374206260448201527f65206c6172676572207468616e2031000000000000000000000000000000000060648201526084016107b0565b6068546080820151825167ffffffffffffffff909216916112799190611aef565b63ffffffff1611156112e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f53797374656d436f6e6669673a20676173206c696d697420746f6f206c6f770060448201526064016107b0565b6000816020015160ff161161137e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f53797374656d436f6e6669673a20656c6173746963697479206d756c7469706c60448201527f6965722063616e6e6f742062652030000000000000000000000000000000000060648201526084016107b0565b8051602082015163ffffffff82169160ff9091169061139e908290611b0e565b6113a89190611b58565b63ffffffff161461143b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f53797374656d436f6e6669673a20707265636973696f6e206c6f73732077697460448201527f6820746172676574207265736f75726365206c696d697400000000000000000060648201526084016107b0565b805160698054602084015160408501516060860151608087015160a09097015163ffffffff9687167fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009095169490941764010000000060ff94851602177fffffffffffffffffffffffffffffffffffffffffffff0000000000ffffffffff166501000000000093909216929092027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff1617660100000000000091851691909102177fffff0000000000000000000000000000000000000000ffffffffffffffffffff166a010000000000000000000093909416929092027fffff00000000000000000000000000000000ffffffffffffffffffffffffffff16929092176e0100000000000000000000000000006fffffffffffffffffffffffffffffffff90921691909102179055565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b600054610100900460ff16611693576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e6700000000000000000000000000000000000000000060648201526084016107b0565b610b0433611585565b803573ffffffffffffffffffffffffffffffffffffffff811681146116c057600080fd5b919050565b6000602082840312156116d757600080fd5b6116e08261169c565b9392505050565b803567ffffffffffffffff811681146116c057600080fd5b60405160c0810167ffffffffffffffff81118282101715611749577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405290565b803563ffffffff811681146116c057600080fd5b803560ff811681146116c057600080fd5b600060c0828403121561178657600080fd5b60405160c0810181811067ffffffffffffffff821117156117d0577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040529050806117df8361174f565b81526117ed60208401611763565b60208201526117fe60408401611763565b604082015261180f6060840161174f565b60608201526118206080840161174f565b608082015260a08301356fffffffffffffffffffffffffffffffff8116811461184857600080fd5b60a0919091015292915050565b6000806000806000806000806000808a8c0361028081121561187657600080fd5b61187f8c61169c565b9a5060208c0135995060408c0135985060608c013597506118a260808d016116e7565b96506118b060a08d0161169c565b95506118bf8d60c08e01611774565b94506101808c013593506118d66101a08d0161169c565b925060c07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe408201121561190857600080fd5b506119116116ff565b61191e6101c08d0161169c565b815261192d6101e08d0161169c565b602082015261193f6102008d0161169c565b60408201526119516102208d0161169c565b60608201526119636102408d0161169c565b60808201526119756102608d0161169c565b60a0820152809150509295989b9194979a5092959850565b6000815180845260005b818110156119b357602081850181015186830182015201611997565b818111156119c5576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006116e0602083018461198d565b60008060408385031215611a1e57600080fd5b50508035926020909101359150565b600060208284031215611a3f57600080fd5b6116e0826116e7565b600060c08284031215611a5a57600080fd5b6116e08383611774565b600060208284031215611a7657600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015611abe57611abe611a7d565b500390565b600067ffffffffffffffff808316818516808303821115611ae657611ae6611a7d565b01949350505050565b600063ffffffff808316818516808303821115611ae657611ae6611a7d565b600063ffffffff80841680611b4c577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910492915050565b600063ffffffff80831681851681830481118215151615611b7b57611b7b611a7d565b0294935050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SystemConfigStorageLayoutJSON), SystemConfigStorageLayout); err != nil {
		panic(err)
	}

	layouts["SystemConfig"] = SystemConfigStorageLayout
	deployedBytecodes["SystemConfig"] = SystemConfigDeployedBin
}
