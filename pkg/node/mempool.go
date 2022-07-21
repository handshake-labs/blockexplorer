package node

import (
	"context"
)

func (client *Client) GetMempool(ctx context.Context) ([]SingleTransaction, error) {
	var txids []string
	var txs []SingleTransaction
	err := client.rpc(ctx, "getrawmempool", nil, &txids)
	if err != nil {
		return nil, err
	}
	for _, txid := range txids {
		tx, err := client.GetTxByTxid(context.Background(), txid)
		if err != nil {
			return nil, err
		}
		txs = append(txs, *tx)
	}
	return txs, nil
}
