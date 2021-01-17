package node

import (
	"context"
)

func (client *Client) GetMempool(ctx context.Context) ([]Transaction, error) {
	var txs []Transaction
	err := client.rpc(ctx, "getexplicitmempool", nil, &txs)
	if err != nil {
		return nil, err
	}
	return txs, nil
}
