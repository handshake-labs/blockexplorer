package node

import (
	"context"
)

func (client *Client) GetMempool(ctx context.Context) (*[]MempoolTx, error) {
	var txs []MempoolTx
	err := client.rpc(ctx, "getexplicitmempool", nil, &txs)
	if err != nil {
		return nil, err
	}
	return &txs, nil
}
