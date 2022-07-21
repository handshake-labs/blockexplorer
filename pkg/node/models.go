//go:generate easytags $GOFILE json:camel

package node

import "github.com/handshake-labs/blockexplorer/pkg/types"

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

// there is a need for additonal type, because the default node field names are different in different methods
type SingleTransaction struct {
	Txid       types.Bytes `json:"hash"`
	WitnessTx  types.Bytes `json:"witnessHash"`
	Fee        int64       `json:"fee"`
	Rate       int64       `json:"rate"`
	BlockHash  types.Bytes `json:"blockHash"`
	IndexBlock int32       `json:"index"`
	Version    int32       `json:"version"`
	Locktime   int32       `json:"locktime"`
	Size       int64       `json:"size"`
	TxInputs   []SingleTxInput   `json:"inputs"`
	TxOutputs  []SingleTxOutput  `json:"outputs"`
}

type SingleTxInput struct {
	MyHashPrevout  struct {
		Hash types.Bytes `json:"hash"`
		IndexPrevout int64       `json:"index"`
	} `json:"prevout"`
	Sequence     int64       `json:"sequence"`
}

type SingleTxOutput struct {
	Index     int32       `json:"n"`
	Value     types.Money `json:"value"`
	MyAddress string `json:"address"`
	Covenant struct {
		CovenantAction string        `json:"action"`
		CovenantItems  []types.Bytes `json:"items"`
	} `json:"covenant"`
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
	return txOutput.MyAddress.String
}

func (txInput *SingleTxInput) HashPrevout() types.Bytes{
	return txInput.MyHashPrevout.Hash
}

type MyAddress struct {
	String string
}
