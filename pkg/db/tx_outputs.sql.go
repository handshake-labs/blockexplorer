// Code generated by sqlc. DO NOT EDIT.
// source: tx_outputs.sql

package db

import (
	"context"

	"github.com/handshake-labs/blockexplorer/pkg/types"
)

const getTxOutputsByTxid = `-- name: GetTxOutputsByTxid :many
SELECT DISTINCT ON(t1.index) t1.txid, t1.index, t1.value, t1.address, t1.covenant_action, t1.covenant_name_hash, t1.covenant_height, t1.covenant_name, t1.covenant_bid_hash, t1.covenant_nonce, t1.covenant_record_data, t1.covenant_block_hash, t1.covenant_version, t1.covenant_address, t1.covenant_claim_height, t1.covenant_renewal_count, t2.covenant_name AS name
FROM tx_outputs t1 LEFT JOIN tx_outputs t2 ON (t1.covenant_name_hash = t2.covenant_name_hash AND t2.covenant_name IS NOT NULL)
WHERE t1.txid = $1
ORDER BY t1.index
`

type GetTxOutputsByTxidRow struct {
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
	Name                 *types.Bytes
}

func (q *Queries) GetTxOutputsByTxid(ctx context.Context, txid types.Bytes) ([]GetTxOutputsByTxidRow, error) {
	rows, err := q.db.QueryContext(ctx, getTxOutputsByTxid, txid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetTxOutputsByTxidRow{}
	for rows.Next() {
		var i GetTxOutputsByTxidRow
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

const insertBIDTxOutput = `-- name: InsertBIDTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash) VALUES ($1, $2, $3, $4, $5, 'BID', $6, $7, $8)
`

type InsertBIDTxOutputParams struct {
	Txid            types.Bytes
	Index           int32
	Value           int64
	Address         string
	CovenantAction  CovenantAction
	CovenantHeight  *types.Bytes
	CovenantName    *types.Bytes
	CovenantBidHash *types.Bytes
}

func (q *Queries) InsertBIDTxOutput(ctx context.Context, arg InsertBIDTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertBIDTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantName,
		arg.CovenantBidHash,
	)
	return err
}

const insertCLAIMTxOutput = `-- name: InsertCLAIMTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name) VALUES ($1, $2, $3, $4, $5, 'CLAIM', $6, $7)
`

type InsertCLAIMTxOutputParams struct {
	Txid           types.Bytes
	Index          int32
	Value          int64
	Address        string
	CovenantAction CovenantAction
	CovenantHeight *types.Bytes
	CovenantName   *types.Bytes
}

func (q *Queries) InsertCLAIMTxOutput(ctx context.Context, arg InsertCLAIMTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertCLAIMTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantName,
	)
	return err
}

const insertFINALIZETxOutput = `-- name: InsertFINALIZETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_claim_height, covenant_renewal_count, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'FINALIZE', $6, $7, $8, $9, $10)
`

type InsertFINALIZETxOutputParams struct {
	Txid                 types.Bytes
	Index                int32
	Value                int64
	Address              string
	CovenantAction       CovenantAction
	CovenantHeight       *types.Bytes
	CovenantName         *types.Bytes
	CovenantClaimHeight  *types.Bytes
	CovenantRenewalCount *types.Bytes
	CovenantBlockHash    *types.Bytes
}

func (q *Queries) InsertFINALIZETxOutput(ctx context.Context, arg InsertFINALIZETxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertFINALIZETxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantName,
		arg.CovenantClaimHeight,
		arg.CovenantRenewalCount,
		arg.CovenantBlockHash,
	)
	return err
}

const insertNONETxOutput = `-- name: InsertNONETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action) VALUES ($1, $2, $3, $4, 'NONE')
`

type InsertNONETxOutputParams struct {
	Txid    types.Bytes
	Index   int32
	Value   int64
	Address string
}

func (q *Queries) InsertNONETxOutput(ctx context.Context, arg InsertNONETxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertNONETxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
	)
	return err
}

const insertOPENTxOutput = `-- name: InsertOPENTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name) VALUES ($1, $2, $3, $4, $5, 'OPEN', $6, $7)
`

type InsertOPENTxOutputParams struct {
	Txid           types.Bytes
	Index          int32
	Value          int64
	Address        string
	CovenantAction CovenantAction
	CovenantHeight *types.Bytes
	CovenantName   *types.Bytes
}

func (q *Queries) InsertOPENTxOutput(ctx context.Context, arg InsertOPENTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertOPENTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantName,
	)
	return err
}

const insertREDEEMTxOutput = `-- name: InsertREDEEMTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height) VALUES ($1, $2, $3, $4, $5, 'REDEEM', $6)
`

type InsertREDEEMTxOutputParams struct {
	Txid           types.Bytes
	Index          int32
	Value          int64
	Address        string
	CovenantAction CovenantAction
	CovenantHeight *types.Bytes
}

func (q *Queries) InsertREDEEMTxOutput(ctx context.Context, arg InsertREDEEMTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertREDEEMTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
	)
	return err
}

const insertREGISTERTxOutput = `-- name: InsertREGISTERTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_record_data, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'REGISTER', $6, $7, $8)
`

type InsertREGISTERTxOutputParams struct {
	Txid               types.Bytes
	Index              int32
	Value              int64
	Address            string
	CovenantAction     CovenantAction
	CovenantHeight     *types.Bytes
	CovenantRecordData *types.Bytes
	CovenantBlockHash  *types.Bytes
}

func (q *Queries) InsertREGISTERTxOutput(ctx context.Context, arg InsertREGISTERTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertREGISTERTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantRecordData,
		arg.CovenantBlockHash,
	)
	return err
}

const insertRENEWTxOutput = `-- name: InsertRENEWTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_block_hash) VALUES ($1, $2, $3, $4, $5, 'RENEW', $6, $7)
`

type InsertRENEWTxOutputParams struct {
	Txid              types.Bytes
	Index             int32
	Value             int64
	Address           string
	CovenantAction    CovenantAction
	CovenantHeight    *types.Bytes
	CovenantBlockHash *types.Bytes
}

func (q *Queries) InsertRENEWTxOutput(ctx context.Context, arg InsertRENEWTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertRENEWTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantBlockHash,
	)
	return err
}

const insertREVEALTxOutput = `-- name: InsertREVEALTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_nonce) VALUES ($1, $2, $3, $4, $5, 'REVEAL', $6, $7)
`

type InsertREVEALTxOutputParams struct {
	Txid           types.Bytes
	Index          int32
	Value          int64
	Address        string
	CovenantAction CovenantAction
	CovenantHeight *types.Bytes
	CovenantNonce  *types.Bytes
}

func (q *Queries) InsertREVEALTxOutput(ctx context.Context, arg InsertREVEALTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertREVEALTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantNonce,
	)
	return err
}

const insertREVOKETxOutput = `-- name: InsertREVOKETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height) VALUES ($1, $2, $3, $4, $5, 'REVOKE', $6)
`

type InsertREVOKETxOutputParams struct {
	Txid           types.Bytes
	Index          int32
	Value          int64
	Address        string
	CovenantAction CovenantAction
	CovenantHeight *types.Bytes
}

func (q *Queries) InsertREVOKETxOutput(ctx context.Context, arg InsertREVOKETxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertREVOKETxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
	)
	return err
}

const insertTRANSFERTxOutput = `-- name: InsertTRANSFERTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_version, covenant_address) VALUES ($1, $2, $3, $4, $5, 'TRANSFER', $6, $7, $8)
`

type InsertTRANSFERTxOutputParams struct {
	Txid            types.Bytes
	Index           int32
	Value           int64
	Address         string
	CovenantAction  CovenantAction
	CovenantHeight  *types.Bytes
	CovenantVersion *types.Bytes
	CovenantAddress *types.Bytes
}

func (q *Queries) InsertTRANSFERTxOutput(ctx context.Context, arg InsertTRANSFERTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertTRANSFERTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantVersion,
		arg.CovenantAddress,
	)
	return err
}

const insertTxOutput = `-- name: InsertTxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_name, covenant_bid_hash, covenant_nonce, covenant_record_data, covenant_block_hash, covenant_version, covenant_address, covenant_claim_height, covenant_renewal_count) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16)
`

type InsertTxOutputParams struct {
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
}

func (q *Queries) InsertTxOutput(ctx context.Context, arg InsertTxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertTxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantNameHash,
		arg.CovenantHeight,
		arg.CovenantName,
		arg.CovenantBidHash,
		arg.CovenantNonce,
		arg.CovenantRecordData,
		arg.CovenantBlockHash,
		arg.CovenantVersion,
		arg.CovenantAddress,
		arg.CovenantClaimHeight,
		arg.CovenantRenewalCount,
	)
	return err
}

const insertUPDATETxOutput = `-- name: InsertUPDATETxOutput :exec
INSERT INTO tx_outputs (txid, index, value, address, covenant_action, covenant_name_hash, covenant_height, covenant_record_data) VALUES ($1, $2, $3, $4, $5, 'UPDATE', $6, $7)
`

type InsertUPDATETxOutputParams struct {
	Txid               types.Bytes
	Index              int32
	Value              int64
	Address            string
	CovenantAction     CovenantAction
	CovenantHeight     *types.Bytes
	CovenantRecordData *types.Bytes
}

func (q *Queries) InsertUPDATETxOutput(ctx context.Context, arg InsertUPDATETxOutputParams) error {
	_, err := q.db.ExecContext(ctx, insertUPDATETxOutput,
		arg.Txid,
		arg.Index,
		arg.Value,
		arg.Address,
		arg.CovenantAction,
		arg.CovenantHeight,
		arg.CovenantRecordData,
	)
	return err
}
