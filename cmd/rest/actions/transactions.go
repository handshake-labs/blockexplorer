package actions

import (
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/types"
	"github.com/jinzhu/copier"
)

type GetTransactionsByBlockHeightParams struct {
	BlockHeight int32 `json:"height"`
	Limit       int8  `json:"limit"`
	Offset      int32 `json:"offset"`
}

type GetTransactionsByBlockHeightResult struct {
	Transactions []Transaction `json:"txs"`
}

func GetTransactionsByBlockHeight(ctx *Context, params *GetTransactionsByBlockHeightParams) (*GetTransactionsByBlockHeightResult, error) {
	transactions, err := ctx.db.GetTransactionsByBlockHeight(ctx, db.GetTransactionsByBlockHeightParams{
		Height: params.BlockHeight,
		Limit:  int32(params.Limit),
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}
	if len(transactions) == 0 {
		return nil, err
	}
	result := GetTransactionsByBlockHeightResult{[]Transaction{}}
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

type GetTransactionByTxidResult Transaction

func GetTransactionByTxid(ctx *Context, params *GetTransactionByTxidParams) (*GetTransactionByTxidResult, error) {
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
	result := GetTransactionByTxidResult{}
	copier.Copy(&result, &transaction)
	copier.Copy(&result.TxInputs, &txInputs)
	copier.Copy(&result.TxOutputs, &txOutputs)
	return &result, nil
}
