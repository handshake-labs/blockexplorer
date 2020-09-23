package actions

import (
	"github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/jinzhu/copier"
)

type GetMempoolTxsParams struct {
	Limit  int8  `json:"limit"`
	Offset int32 `json:"offset"`
}

type GetMempoolTxsResult struct {
	Transactions []Transaction `json:"txs""`
}

func GetMempoolTxs(ctx *Context, params *GetMempoolTxsParams) (*GetMempoolTxsResult, error) {
	transactions, err := ctx.db.GetMempoolTransactions(ctx, db.GetMempoolTransactionsParams{
		Limit:  int32(params.Limit),
		Offset: params.Offset,
	})
	if err != nil {
		return nil, err
	}
	result := GetMempoolTxsResult{}
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
