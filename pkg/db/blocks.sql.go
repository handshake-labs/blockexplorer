// Code generated by sqlc. DO NOT EDIT.
// source: blocks.sql

package db

import (
	"context"

	"github.com/handshake-labs/blockexplorer/pkg/types"
)

const deleteBlocksAfterHeight = `-- name: DeleteBlocksAfterHeight :exec
DELETE FROM blocks
WHERE height > $1
`

func (q *Queries) DeleteBlocksAfterHeight(ctx context.Context, height int32) error {
	_, err := q.db.ExecContext(ctx, deleteBlocksAfterHeight, height)
	return err
}

const getBlockByHash = `-- name: GetBlockByHash :one
SELECT hash, height, weight, size, version, hash_merkle_root, witness_root, tree_root, reserved_root, mask, time, bits, difficulty, chainwork, nonce, extra_nonce, orphan
FROM blocks
WHERE hash = $1
`

func (q *Queries) GetBlockByHash(ctx context.Context, hash types.Bytes) (Block, error) {
	row := q.db.QueryRowContext(ctx, getBlockByHash, hash)
	var i Block
	err := row.Scan(
		&i.Hash,
		&i.Height,
		&i.Weight,
		&i.Size,
		&i.Version,
		&i.HashMerkleRoot,
		&i.WitnessRoot,
		&i.TreeRoot,
		&i.ReservedRoot,
		&i.Mask,
		&i.Time,
		&i.Bits,
		&i.Difficulty,
		&i.Chainwork,
		&i.Nonce,
		&i.ExtraNonce,
		&i.Orphan,
	)
	return i, err
}

const getBlockByHeight = `-- name: GetBlockByHeight :one
SELECT hash, height, weight, size, version, hash_merkle_root, witness_root, tree_root, reserved_root, mask, time, bits, difficulty, chainwork, nonce, extra_nonce, orphan
FROM blocks
WHERE height = $1
`

func (q *Queries) GetBlockByHeight(ctx context.Context, height int32) (Block, error) {
	row := q.db.QueryRowContext(ctx, getBlockByHeight, height)
	var i Block
	err := row.Scan(
		&i.Hash,
		&i.Height,
		&i.Weight,
		&i.Size,
		&i.Version,
		&i.HashMerkleRoot,
		&i.WitnessRoot,
		&i.TreeRoot,
		&i.ReservedRoot,
		&i.Mask,
		&i.Time,
		&i.Bits,
		&i.Difficulty,
		&i.Chainwork,
		&i.Nonce,
		&i.ExtraNonce,
		&i.Orphan,
	)
	return i, err
}

const getBlockHashByHeight = `-- name: GetBlockHashByHeight :one
SELECT hash
FROM blocks
WHERE height = $1
`

func (q *Queries) GetBlockHashByHeight(ctx context.Context, height int32) (types.Bytes, error) {
	row := q.db.QueryRowContext(ctx, getBlockHashByHeight, height)
	var hash types.Bytes
	err := row.Scan(&hash)
	return hash, err
}

const getBlocksMaxHeight = `-- name: GetBlocksMaxHeight :one
SELECT COALESCE(MAX(height), -1)::integer
FROM blocks
`

func (q *Queries) GetBlocksMaxHeight(ctx context.Context) (int32, error) {
	row := q.db.QueryRowContext(ctx, getBlocksMaxHeight)
	var column_1 int32
	err := row.Scan(&column_1)
	return column_1, err
}

const insertBlock = `-- name: InsertBlock :exec
INSERT INTO blocks (hash, height, weight, size, version, hash_merkle_root, witness_root, tree_root, reserved_root, mask, time, bits, difficulty, chainwork, nonce, extra_nonce)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
`

type InsertBlockParams struct {
	Hash           types.Bytes
	Height         int32
	Weight         int32
	Size           int64
	Version        int32
	HashMerkleRoot types.Bytes
	WitnessRoot    types.Bytes
	TreeRoot       types.Bytes
	ReservedRoot   types.Bytes
	Mask           types.Bytes
	Time           int32
	Bits           types.Bytes
	Difficulty     float64
	Chainwork      types.Bytes
	Nonce          int64
	ExtraNonce     types.Bytes
}

func (q *Queries) InsertBlock(ctx context.Context, arg InsertBlockParams) error {
	_, err := q.db.ExecContext(ctx, insertBlock,
		arg.Hash,
		arg.Height,
		arg.Weight,
		arg.Size,
		arg.Version,
		arg.HashMerkleRoot,
		arg.WitnessRoot,
		arg.TreeRoot,
		arg.ReservedRoot,
		arg.Mask,
		arg.Time,
		arg.Bits,
		arg.Difficulty,
		arg.Chainwork,
		arg.Nonce,
		arg.ExtraNonce,
	)
	return err
}
