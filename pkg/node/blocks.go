package node

import (
	"context"

	"github.com/handshake-labs/blockexplorer/pkg/types"
)

func (client *Client) GetBlockByHeight(ctx context.Context, height int32) (*Block, error) {
	block := new(Block)
	err := client.rpc(ctx, "getblockbyheight", []interface{}{height, true, true}, block)
	if err != nil {
		return nil, err
	}
	return block, nil
}

func (client *Client) GetBlockHashByHeight(ctx context.Context, height int32) (types.Bytes, error) {
	var hash types.Bytes
	err := client.rpc(ctx, "getblockhash", []interface{}{height}, &hash)
	if err != nil {
		return nil, err
	}
	return hash, nil
}

func (client *Client) GetBlocksHeight(ctx context.Context) (int32, error) {
	var height int32
	err := client.rpc(ctx, "getblockcount", nil, &height)
	if err != nil {
		return -1, err
	}
	return height, nil
}


