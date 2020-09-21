package actions

import (
	// "github.com/handshake-labs/blockexplorer/pkg/types"
	// "github.com/handshake-labs/blockexplorer/pkg/db"
	"github.com/handshake-labs/blockexplorer/pkg/node"
	// "log"
)

type GetMempoolTxsParams struct {
	Page int16 `json:"page"`
}

type GetMempoolTxsResult struct {
	MempoolTxs []node.MempoolTx `json:"txs"`
}

func GetMempoolTxs(ctx *Context, params *GetMempoolTxsParams) (*GetMempoolTxsResult, error) {
	// mempoolTxs, err := ctx.client.GetMempool(ctx)
	return nil, nil
}
