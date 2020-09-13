package node

import (
	"context"
	// "github.com/handshake-labs/blockexplorer/pkg/types"
)

func (client *Client) GetInfo(ctx context.Context, height int32) (*Block, error) {
	block := new(Block)
	err := client.rpc(ctx, "getinfo", nil, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}
