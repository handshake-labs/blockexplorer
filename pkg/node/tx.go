package node

import (
	"context"

)

func (client *Client) GetTxByTxid(ctx context.Context, txid string) (*SingleTransaction, error) {
	tx := new(SingleTransaction)

	err := client.rest(ctx, "GET", []string{"tx",txid}, nil, &tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}


