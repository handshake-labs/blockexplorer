//go:generate easytags $GOFILE json:camel

package node

import "github.com/handshake-labs/blockexplorer/pkg/types"

// import "log"

type Block struct {
	Hash           types.Bytes   `json:"hash"`
	PrevBlockHash  types.Bytes   `json:"previousblockhash"`
	Height         int32         `json:"height"`
	Weight         int32         `json:"weight"`
	Size           int64         `json:"size"`
	Version        int32         `json:"version"`
	HashMerkleRoot types.Bytes   `json:"merkleRoot"`
	WitnessRoot    types.Bytes   `json:"witnessRoot"`
	TreeRoot       types.Bytes   `json:"treeRoot"`
	ReservedRoot   types.Bytes   `json:"reservedRoot"`
	Mask           types.Bytes   `json:"mask"`
	Time           int32         `json:"time"`
	Bits           types.Bytes   `json:"bits"`
	Difficulty     float64       `json:"difficulty"`
	Chainwork      types.Bytes   `json:"chainwork"`
	Nonce          int64         `json:"nonce"`
	ExtraNonce     types.Bytes   `json:"extraNonce"`
	Transactions   []Transaction `json:"tx"`
}

type Transaction struct {
	Txid       types.Bytes `json:"txid"`
	WitnessTx  types.Bytes `json:"hash"`
	Fee        int64       `json:"fee"`
	Rate       int64       `json:"rate"`
	BlockHash  types.Bytes `json:"blockHash"`
	IndexBlock int32       `json:"index"`
	Version    int32       `json:"version"`
	Locktime   int32       `json:"locktime"`
	Size       int64       `json:"size"`
	TxInputs   []TxInput   `json:"vin"`
	TxOutputs  []TxOutput  `json:"vout"`
}

type TxInput struct {
	HashPrevout  types.Bytes `json:"txid"`
	IndexPrevout int64       `json:"vout"`
	Sequence     int64       `json:"sequence"`
}

type TxOutput struct {
	Index     int32       `json:"n"`
	Value     types.Money `json:"value"`
	MyAddress struct {    // to prevent conflict with method of the same name
		String string `json:"string"`
	} `json:"address"`
	Covenant struct {
		CovenantAction string        `json:"action"`
		CovenantItems  []types.Bytes `json:"items"`
	} `json:"covenant"`
}

func (txOutput *TxOutput) Address() string {
	// log.Printf("%+v", txOutput)
	return txOutput.MyAddress.String
}

type MempoolTx struct {
	Txid      types.Bytes `json:"hash"`
	WitnessTx types.Bytes `json:"witnessHash"`
	Mtime     int64       `json:"mtime"`
	Version   int32       `json:"version"`
	// Fee        int64       `json:"fee"`
	// Rate       int64       `json:"rate"`
	// BlockHash  types.Bytes `json:"blockHash"`
	// IndexBlock int32       `json:"index"`
	Locktime int32 `json:"locktime"`
	// Size       int64       `json:"size"`
	// TxInputs   []TxInput   `json:"vin"`
	// 	WitnessTx  types.Bytes `json:"hash"`
}

// {
//             "inputs": [
//                 {
//                     "prevout": {
//                         "hash": "df06bf8381130ae14644ceec4272297e9db112df4bfaab028099d5b7f7e1e3b0",
//                         "index": 1
//                     },
//                     "witness": [
//                         "1211e854e99719a1bce9fcb6dd27d6bf4f30d696d8d57cabe06315fa20c58a5c6d55c66055d3d182c6fac050a4f077fc166928c09eea2dd2532671b1820c8b3901",
//                         "03de837a91d582f70a89d26e4fb370336635f30fe29045fa4af13c8d274acbc2d9"
//                     ],
//                     "sequence": 4294967295,
//                     "address": "hs1qr237ngcel5tv52pzpvef8z68ry0xvr82x256pl"
//                 }
//             ],
//             "outputs": [
//                 {
//                     "value": 12000000,
//                     "address": "hs1qgkn49gwdfj4a6pkphj6lqsm0hl34my79ruwqd5",
//                     "covenant": {
//                         "type": 3,
//                         "action": "BID",
//                         "items": [
//                             "c59a58701522b312950107294004b1e113d9e70b532a19869b822015c645b775",
//                             "147a0000",
//                             "796773",
//                             "7698624a090205698b3295bbc2a237b10d4064a92336339e97d3dfdf61384dc2"
//                         ]
//                     }
//                 },
//                 {
//                     "value": 22679559199,
//                     "address": "hs1q3rs94lzjady5zaj38uqrj8vxhf0rsnh63nqmju",
//                     "covenant": {
//                         "type": 0,
//                         "action": "NONE",
//                         "items": []
//                     }
//                 }
//             ],
//             "locktime": 0,
//             "hex": "0000000001df06bf8381130ae14644ceec4272297e9db112df4bfaab028099d5b7f7e1e3b001000000ffffffff02001bb70000000000001445a752a1cd4cabdd06c1bcb5f0436fbfe35d93c5030420c59a58701522b312950107294004b1e113d9e70b532a19869b822015c645b77504147a000003796773207698624a090205698b3295bbc2a237b10d4064a92336339e97d3dfdf61384dc21f9cce4705000000001488e05afc52eb494176513f00391d86ba5e384efa00000000000002411211e854e99719a1bce9fcb6dd27d6bf4f30d696d8d57cabe06315fa20c58a5c6d55c66055d3d182c6fac050a4f077fc166928c09eea2dd2532671b1820c8b39012103de837a91d582f70a89d26e4fb370336635f30fe29045fa4af13c8d274acbc2d9"
//         }
