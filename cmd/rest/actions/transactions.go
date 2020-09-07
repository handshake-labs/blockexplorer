package actions

import (
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
	// "log"
)

type GetTransactionsByBlockHashParams struct {
	BlockHash types.Bytes `json:"hash"`
	Page      int16       `json:"page"`
}

type GetTransactionsByBlockHashResult struct {
	Transactions []Transaction `json:"txs"`
	Count        int16         `json:"count"`
	Limit        int16         `json:"limit"`
}

func GetTransactionsByBlockHash(ctx *Context, params *GetTransactionsByBlockHashParams) (*GetTransactionsByBlockHashResult, error) {
	result := GetTransactionsByBlockHashResult{}
	result.Limit = 50
	var page int16 = params.Page
	if page < 0 {
		page = 0
	}
	transactions, err := ctx.db.GetTransactionsByBlockHash(ctx, db.GetTransactionsByBlockHashParams{
		BlockHash: params.BlockHash,
		Limit:     result.Limit,
		Offset:    page * result.Limit,
	})
	if err != nil {
		return nil, err
	}
	// log.Printf("%+v", transactions[0])
	if len(transactions) == 0 {
		return &result, nil
	}
	result.Count = transactions[0].Count
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
