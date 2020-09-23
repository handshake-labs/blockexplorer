package actions

import (
	"log"

	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
)

type GetTransactionsByBlockHashParams struct {
	BlockHash types.Bytes `json:"hash"`
	Limit     int8        `json:"limit"`
	Offset    int32       `json:"offset"`
}

type GetTransactionsByBlockHashResult struct {
	Transactions []Transaction `json:"txs"`
}

func GetTransactionsByBlockHash(ctx *Context, params *GetTransactionsByBlockHashParams) (*GetTransactionsByBlockHashResult, error) {
	transactions, err := ctx.db.GetTransactionsByBlockHash(ctx, db.GetTransactionsByBlockHashParams{
		BlockHash: &params.BlockHash,
		Limit:     int32(params.Limit),
		Offset:    params.Offset,
	})
	if err != nil {
		return nil, err
	}
	result := GetTransactionsByBlockHashResult{}
	for _, transaction := range transactions {
		txInputs, err := ctx.db.GetTxInputsByTxid(ctx, transaction.Txid)
		if err != nil {
			return nil, err
		}
		txOutputs, err := ctx.db.GetTxOutputsByTxid(ctx, transaction.Txid)
		if err != nil {
			return nil, err
		}
		var resultTransaction Transaction
		copier.Copy(&resultTransaction, &transaction)
		copier.Copy(&resultTransaction.TxInputs, &txInputs)
		copier.Copy(&resultTransaction.TxOutputs, &txOutputs)
		result.Transactions = append(result.Transactions, resultTransaction)
	}
	return &result, nil
}

type GetTransactionByTxidParams struct {
	Txid types.Bytes `json:"txid"`
}

func GetTransactionByTxid(ctx *Context, params *GetTransactionByTxidParams) (*Transaction, error) {
	log.Println(params.Txid)
	transaction, err := ctx.db.GetTransactionByTxid(ctx, params.Txid)
	if err != nil {
		return nil, err
	}
	txInputs, err := ctx.db.GetTxInputsByTxid(ctx, transaction.Txid)
	if err != nil {
		return nil, err
	}
	txOutputs, err := ctx.db.GetTxOutputsByTxid(ctx, transaction.Txid)
	if err != nil {
		return nil, err
	}
	var resultTransaction Transaction
	copier.Copy(&resultTransaction, &transaction)
	copier.Copy(&resultTransaction.TxInputs, &txInputs)
	copier.Copy(&resultTransaction.TxOutputs, &txOutputs)
	return &resultTransaction, nil
}
