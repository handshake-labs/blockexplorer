// Code generated by sqlc. DO NOT EDIT.
// source: addresses.sql

package db

import (
	"context"

	"github.com/handshake-labs/blockexplorer/pkg/types"
)

const addressExists = `-- name: AddressExists :one
SELECT EXISTS(SELECT 1 FROM tx_outputs WHERE tx_outputs.address = $1::text)
`

func (q *Queries) AddressExists(ctx context.Context, address string) (bool, error) {
	row := q.db.QueryRowContext(ctx, addressExists, address)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const getAddressInfo = `-- name: GetAddressInfo :one

SELECT
  COALESCE(SUM(tx_outputs.value), 0)::bigint AS value_total,
  COALESCE(SUM(tx_outputs.value) filter (WHERE tx_inputs.txid IS NOT NULL), 0)::bigint AS value_used,
  COUNT(tx_outputs.txid) AS tx_outputs_total,
  COUNT(tx_inputs.hash_prevout) AS tx_outputs_used
FROM tx_outputs
LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
WHERE tx_outputs.address = $1::text
`

type GetAddressInfoRow struct {
	ValueTotal     int64
	ValueUsed      int64
	TxOutputsTotal int64
	TxOutputsUsed  int64
}

//This query takes a lot of time, perhaps can be optimized further
func (q *Queries) GetAddressInfo(ctx context.Context, address string) (GetAddressInfoRow, error) {
	row := q.db.QueryRowContext(ctx, getAddressInfo, address)
	var i GetAddressInfoRow
	err := row.Scan(
		&i.ValueTotal,
		&i.ValueUsed,
		&i.TxOutputsTotal,
		&i.TxOutputsUsed,
	)
	return i, err
}

const getTxOutputsByAddress = `-- name: GetTxOutputsByAddress :many

SELECT
  DISTINCT tx_outputs.txid, tx_outputs.index, tx_outputs.value, tx_outputs.address, tx_outputs.covenant_action, tx_outputs.covenant_name_hash, tx_outputs.covenant_height, tx_outputs.covenant_name, tx_outputs.covenant_bid_hash, tx_outputs.covenant_nonce, tx_outputs.covenant_record_data, tx_outputs.covenant_block_hash, tx_outputs.covenant_version, tx_outputs.covenant_address, tx_outputs.covenant_claim_height, tx_outputs.covenant_renewal_count,
  COALESCE(tx_inputs.txid, '\x')::bytea AS hash_prevout_not_null,
  COALESCE(tx_inputs.index, -1) AS index_prevout_not_null,
  COALESCE(bl2.height, -1)::integer AS spend_height_not_null, --height of -1 means mempool, so i need -2 to indicate the block does not exist 
  COALESCE(blocks.height, 2147483647)::integer AS height_not_null,
  COALESCE(CONVERT_FROM(t2.covenant_name, 'SQL_ASCII'), '')::text AS name
FROM tx_outputs
  LEFT JOIN tx_outputs t2 ON (tx_outputs.covenant_name_hash = t2.covenant_name_hash AND t2.covenant_name IS NOT NULL)
  LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
  JOIN transactions ON tx_outputs.txid = transactions.txid
  LEFT JOIN blocks ON transactions.block_hash = blocks.hash --for height of receive
  LEFT JOIN transactions tx2 ON tx_inputs.txid = tx2.txid
  LEFT JOIN blocks  bl2 ON tx2.block_hash = bl2.hash --for height of spend
WHERE tx_outputs.address = $1::text
ORDER BY height_not_null DESC NULLS LAST
LIMIT $3::integer OFFSET $2::integer
`

type GetTxOutputsByAddressParams struct {
	Address string
	Offset  int32
	Limit   int32
}

type GetTxOutputsByAddressRow struct {
	Txid                 types.Bytes
	Index                int32
	Value                int64
	Address              string
	CovenantAction       CovenantAction
	CovenantNameHash     *types.Bytes
	CovenantHeight       *types.Bytes
	CovenantName         *types.Bytes
	CovenantBidHash      *types.Bytes
	CovenantNonce        *types.Bytes
	CovenantRecordData   *types.Bytes
	CovenantBlockHash    *types.Bytes
	CovenantVersion      *types.Bytes
	CovenantAddress      *types.Bytes
	CovenantClaimHeight  *types.Bytes
	CovenantRenewalCount *types.Bytes
	HashPrevoutNotNull   types.Bytes
	IndexPrevoutNotNull  int64
	SpendHeightNotNull   int32
	HeightNotNull        int32
	Name                 string
}

//This query can be optimized to be very quick by removing join for the name,
//however as it's still quicker than the GetAddressInfo I've left the name for the sake of simplicity
func (q *Queries) GetTxOutputsByAddress(ctx context.Context, arg GetTxOutputsByAddressParams) ([]GetTxOutputsByAddressRow, error) {
	rows, err := q.db.QueryContext(ctx, getTxOutputsByAddress, arg.Address, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTxOutputsByAddressRow{}
	for rows.Next() {
		var i GetTxOutputsByAddressRow
		if err := rows.Scan(
			&i.Txid,
			&i.Index,
			&i.Value,
			&i.Address,
			&i.CovenantAction,
			&i.CovenantNameHash,
			&i.CovenantHeight,
			&i.CovenantName,
			&i.CovenantBidHash,
			&i.CovenantNonce,
			&i.CovenantRecordData,
			&i.CovenantBlockHash,
			&i.CovenantVersion,
			&i.CovenantAddress,
			&i.CovenantClaimHeight,
			&i.CovenantRenewalCount,
			&i.HashPrevoutNotNull,
			&i.IndexPrevoutNotNull,
			&i.SpendHeightNotNull,
			&i.HeightNotNull,
			&i.Name,
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
