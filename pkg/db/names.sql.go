// Code generated by sqlc. DO NOT EDIT.
// source: names.sql

package db

import (
	"context"

	"github.com/handshake-labs/blockexplorer/pkg/types"
)

const getLastNameBlockHeightByActionAndHash = `-- name: GetLastNameBlockHeightByActionAndHash :one
SELECT
  COALESCE(blocks.height, -1)::integer AS block_height_not_null
FROM
  tx_outputs
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE covenant_action = $1 AND covenant_name_hash = $2
ORDER BY blocks.height DESC NULLS FIRST
LIMIT 1
`

type GetLastNameBlockHeightByActionAndHashParams struct {
	CovenantAction   CovenantAction
	CovenantNameHash *types.Bytes
}

func (q *Queries) GetLastNameBlockHeightByActionAndHash(ctx context.Context, arg GetLastNameBlockHeightByActionAndHashParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, getLastNameBlockHeightByActionAndHash, arg.CovenantAction, arg.CovenantNameHash)
	var block_height_not_null int32
	err := row.Scan(&block_height_not_null)
	return block_height_not_null, err
}

const getNameBidsByHash = `-- name: GetNameBidsByHash :many
SELECT
  DISTINCT ON (block_height_not_null, bid_txid, lockup_outputs.index)
  bids.txid AS bid_txid, 
  COALESCE(blocks.height, -1)::integer AS block_height_not_null,
  COALESCE(reveals.txid, '\x00')::bytea AS reveal_txid,
  COALESCE(reveal_outputs.index, -1)::integer AS reveal_index_not_null,
  lockup_outputs.value as lockup_value,
  COALESCE(reveal_outputs.value, -1) as reveal_value_not_null
FROM                                                  
  transactions as bids
  JOIN tx_inputs as lockup_inputs ON lockup_inputs.txid=bids.txid
  JOIN blocks ON (bids.block_hash = blocks.hash)
  JOIN tx_outputs as lockup_outputs ON lockup_outputs.txid=bids.txid AND lockup_outputs.covenant_action = 'BID'
  LEFT JOIN tx_inputs reveal_inputs ON
     reveal_inputs.hash_prevout = lockup_outputs.txid AND
     reveal_inputs.index_prevout = lockup_outputs.index
  LEFT JOIN tx_outputs reveal_outputs ON reveal_outputs.covenant_action = 'REVEAL'  AND reveal_outputs.covenant_name_hash = lockup_outputs.covenant_name_hash AND
     reveal_inputs.txid = reveal_outputs.txid AND
     reveal_inputs.index = reveal_outputs.index 
  LEFT JOIN transactions AS reveals ON reveal_inputs.txid = reveals.txid AND reveal_outputs.txid = reveals.txid
WHERE lockup_outputs.covenant_name_hash = $1::bytea
ORDER BY block_height_not_null DESC NULLS FIRST
LIMIT $3::integer OFFSET $2::integer
`

type GetNameBidsByHashParams struct {
	NameHash types.Bytes
	Offset   int32
	Limit    int32
}

type GetNameBidsByHashRow struct {
	BidTxid            types.Bytes
	BlockHeightNotNull int32
	RevealTxid         types.Bytes
	RevealIndexNotNull int32
	LockupValue        int64
	RevealValueNotNull int64
}

func (q *Queries) GetNameBidsByHash(ctx context.Context, arg GetNameBidsByHashParams) ([]GetNameBidsByHashRow, error) {
	rows, err := q.db.QueryContext(ctx, getNameBidsByHash, arg.NameHash, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetNameBidsByHashRow{}
	for rows.Next() {
		var i GetNameBidsByHashRow
		if err := rows.Scan(
			&i.BidTxid,
			&i.BlockHeightNotNull,
			&i.RevealTxid,
			&i.RevealIndexNotNull,
			&i.LockupValue,
			&i.RevealValueNotNull,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNameCountsByHash = `-- name: GetNameCountsByHash :one
SELECT
  (COUNT(*) FILTER (WHERE covenant_action = 'BID'))::integer AS bids_count,
  COUNT(covenant_record_data)::integer AS records_count,
  (COUNT(*) FILTER (WHERE covenant_action = 'CLAIM' OR covenant_action = 'RENEW' OR covenant_action = 'TRANSFER' OR covenant_action = 'FINALIZE' OR covenant_action = 'REVOKE'))::integer AS actions_count
FROM tx_outputs
WHERE covenant_name_hash = $1::bytea
`

type GetNameCountsByHashRow struct {
	BidsCount    int32
	RecordsCount int32
	ActionsCount int32
}

func (q *Queries) GetNameCountsByHash(ctx context.Context, nameHash types.Bytes) (GetNameCountsByHashRow, error) {
	row := q.db.QueryRowContext(ctx, getNameCountsByHash, nameHash)
	var i GetNameCountsByHashRow
	err := row.Scan(&i.BidsCount, &i.RecordsCount, &i.ActionsCount)
	return i, err
}

const getNameOtherActionsByHash = `-- name: GetNameOtherActionsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height,
  tx_outputs.covenant_action AS covenant_action
FROM
  tx_outputs 
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE
  tx_outputs.covenant_action != 'OPEN' AND
  tx_outputs.covenant_action != 'BID' AND
  tx_outputs.covenant_action != 'REVEAL' AND
  tx_outputs.covenant_action != 'REDEEM' AND
  tx_outputs.covenant_name_hash = $1::bytea  AND
  tx_outputs.covenant_record_data IS NULL
ORDER BY (blocks.height, transactions.index, tx_outputs.index) DESC NULLS FIRST
LIMIT $3::integer OFFSET $2::integer
`

type GetNameOtherActionsByHashParams struct {
	NameHash types.Bytes
	Offset   int32
	Limit    int32
}

type GetNameOtherActionsByHashRow struct {
	Txid           types.Bytes
	BlockHeight    int32
	CovenantAction CovenantAction
}

func (q *Queries) GetNameOtherActionsByHash(ctx context.Context, arg GetNameOtherActionsByHashParams) ([]GetNameOtherActionsByHashRow, error) {
	rows, err := q.db.QueryContext(ctx, getNameOtherActionsByHash, arg.NameHash, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetNameOtherActionsByHashRow{}
	for rows.Next() {
		var i GetNameOtherActionsByHashRow
		if err := rows.Scan(&i.Txid, &i.BlockHeight, &i.CovenantAction); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getNameRecordsByHash = `-- name: GetNameRecordsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null,
  tx_outputs.covenant_record_data::bytea AS data
FROM
  tx_outputs
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE tx_outputs.covenant_record_data IS NOT NULL AND tx_outputs.covenant_name_hash = $1::bytea
ORDER BY (blocks.height, transactions.index, tx_outputs.index) DESC NULLS FIRST
LIMIT $3::integer OFFSET $2::integer
`

type GetNameRecordsByHashParams struct {
	NameHash types.Bytes
	Offset   int32
	Limit    int32
}

type GetNameRecordsByHashRow struct {
	Txid               types.Bytes
	BlockHeightNotNull int32
	Data               types.Bytes
}

func (q *Queries) GetNameRecordsByHash(ctx context.Context, arg GetNameRecordsByHashParams) ([]GetNameRecordsByHashRow, error) {
	rows, err := q.db.QueryContext(ctx, getNameRecordsByHash, arg.NameHash, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetNameRecordsByHashRow{}
	for rows.Next() {
		var i GetNameRecordsByHashRow
		if err := rows.Scan(&i.Txid, &i.BlockHeightNotNull, &i.Data); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReservedName = `-- name: GetReservedName :one
SELECT 
CONVERT_FROM(origin_name, 'SQL_ASCII')::text as origin_name,
CONVERT_FROM(name, 'SQL_ASCII')::text as name,
name_hash,
claim_amount
FROM reserved_names
WHERE name = $1
`

type GetReservedNameRow struct {
	OriginName  string
	Name        string
	NameHash    types.Bytes
	ClaimAmount int64
}

func (q *Queries) GetReservedName(ctx context.Context, name string) (GetReservedNameRow, error) {
	row := q.db.QueryRowContext(ctx, getReservedName, name)
	var i GetReservedNameRow
	err := row.Scan(
		&i.OriginName,
		&i.Name,
		&i.NameHash,
		&i.ClaimAmount,
	)
	return i, err
}
