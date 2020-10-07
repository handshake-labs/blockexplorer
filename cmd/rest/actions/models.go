//go:generate easytags $GOFILE json:camel

package actions

import (
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
)

type Block struct {
	Hash           types.Bytes `json:"hash"`
	Height         int32       `json:"height"`
	Weight         int32       `json:"weight"`
	Size           int64       `json:"size"`
	Version        int32       `json:"version"`
	HashMerkleRoot types.Bytes `json:"hashMerkleRoot"`
	WitnessRoot    types.Bytes `json:"witnessRoot"`
	TreeRoot       types.Bytes `json:"treeRoot"`
	ReservedRoot   types.Bytes `json:"reservedRoot"`
	Mask           types.Bytes `json:"mask"`
	Time           int32       `json:"time"`
	Bits           types.Bytes `json:"bits"`
	Difficulty     float64     `json:"difficulty"`
	Chainwork      types.Bytes `json:"chainwork"`
	Nonce          int64       `json:"nonce"`
	ExtraNonce     types.Bytes `json:"extraNonce"`
	TxsCount       int32       `json:"txsCount"`
}

type Transaction struct {
	Txid        types.Bytes `json:"txid"`
	BlockHeight *int32      `json:"block_height"`
	WitnessTx   types.Bytes `json:"witnessTx"`
	Fee         int64       `json:"fee"`
	Rate        int64       `json:"rate"`
	Version     int32       `json:"version"`
	Locktime    int32       `json:"locktime"`
	Size        int64       `json:"size"`
	Height      int64       `json:"height"`
	TxInputs    []TxInput   `json:"inputs"`
	TxOutputs   []TxOutput  `json:"outputs"`
}

type TxInput struct {
	HashPrevout  types.Bytes `json:"hashPrevout"`
	IndexPrevout int64       `json:"indexPrevout"`
	Sequence     int64       `json:"sequence"`
}

type TxOutput struct {
	Value                int64             `json:"value"`
	Address              string            `json:"address"`
	CovenantAction       db.CovenantAction `json:"covenantAction"`
	CovenantNameHash     types.Bytes       `json:"covenantNameHash,omitempty"`
	CovenantHeight       types.Bytes       `json:"covenantHeight,omitempty"`
	CovenantName         types.Bytes       `json:"covenantName,omitempty"`
	CovenantBidHash      types.Bytes       `json:"covenantBidHash,omitempty"`
	CovenantNonce        types.Bytes       `json:"covenantNonce,omitempty"`
	CovenantRecordData   types.Bytes       `json:"covenantRecordData,omitempty"`
	CovenantBlockHash    types.Bytes       `json:"covenantBlockHash,omitempty"`
	CovenantVersion      types.Bytes       `json:"covenantVersion,omitempty"`
	CovenantAddress      types.Bytes       `json:"covenantAddress,omitempty"`
	CovenantClaimHeight  types.Bytes       `json:"covenantClaimHeight,omitempty"`
	CovenantRenewalCount types.Bytes       `json:"covenantRenewalCount,omitempty"`
	Name                 string            `json:"name,omitempty"`
}

type ReservedName struct {
	OriginName  string      `json:"originName"`
	Name        string      `json:"name"`
	NameHash    types.Bytes `json:"nameHash"`
	ClaimAmount int64       `json:"claimAmount"`
}

type NameBid struct {
	Txid        types.Bytes `json:"txid"`
	BlockHeight *int32      `json:"height"`
	LockupValue int64       `json:"lockup"`
	RevealValue *int64      `json:"reveal"`
}

type NameRecord struct {
	Txid        types.Bytes `json:"txid"`
	BlockHeight *int32      `json:"height"`
	Data        types.Bytes `json:"data"`
}

type AuctionState string

const (
	AuctionStateClosed AuctionState = "CLOSED"
	AuctionStateOpen   AuctionState = "OPEN"
	AuctionStateBid    AuctionState = "BID"
	AuctionStateReveal AuctionState = "REVEAL"
)
