package db

import (
	"context"
	"github.com/handshake-labs/blockexplorer/pkg/types"
)

const getReservedByNameHash = `-- name: GetReservedByNameHash :one
SELECT name, origin_name, name_hash, claim_amount FROM reserved_names WHERE name_hash = $1
`

func (q *Queries) GetReservedByNameHash(ctx context.Context, nameHash types.Bytes) (ReservedName, error) {
	row := q.db.QueryRowContext(ctx, getReservedByNameHash, nameHash)
	var i ReservedName
	err := row.Scan(
		&i.Name,
		&i.OriginName,
		&i.NameHash,
		&i.ClaimAmount,
	)
	return i, err
}

const getAuctionHistoryByName = `-- name: GetAuctionHistoryByName :many
SELECT
height,
txid,
covenant_name,
lockup,
reveal,
covenant_action,
covenant_record_data,
<<<<<<< HEAD
covenant_name_hash,
=======
covenant_name_hash
>>>>>>> mess
(COUNT(*) OVER())::smallint as count
FROM auctions
WHERE covenant_name = $1 
ORDER BY height DESC
LIMIT $3::smallint OFFSET $2::smallint;
`

type GetAuctionHistoryByNameParams struct {
	Name   string
	Offset int16
	Limit  int16
}

type AuctionHistoryRow struct {
	Height             int32
	Txid               types.Bytes
	CovenantName       *types.Bytes
	LockupValue        *int64
	RevealValue        *int64
	CovenantAction     string
	CovenantRecordData *types.Bytes
	CovenantNameHash   *types.Bytes
	Count              int16
}

type NameRow struct {
	OpenHeight       int32
	CovenantNameHash types.Bytes
	CovenantName     types.Bytes
	MaxLockup        int64
	MaxRevealed      int64
	BidCount         int16
	Count            int32
}

type RecordRow struct {
	Height             int32
	CovenantRecordData types.Bytes
	Count              int16
}

func (q *Queries) GetAuctionHistoryByName(ctx context.Context, arg GetAuctionHistoryByNameParams) ([]AuctionHistoryRow, error) {
	rows, err := q.db.QueryContext(ctx, getAuctionHistoryByName, arg.Name, arg.Offset, arg.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []AuctionHistoryRow{}
	for rows.Next() {
		var i AuctionHistoryRow
		if err := rows.Scan(
			&i.Height,
			&i.Txid,
			&i.CovenantName,
			&i.LockupValue,
			&i.RevealValue,
			&i.CovenantAction,
			&i.CovenantRecordData,
			&i.CovenantNameHash,
			&i.Count,
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

type GetNameRecordHistoryByNameHashParams struct {
	NameHash types.Bytes
	Offset   int16
	Limit    int16
}

const getNameRecordHistoryByNameHash = `-- name: GetNameRecordHistoryByNameHash :many 
SELECT height, covenant_record_data, (COUNT(*) OVER())::smallint as count
FROM records
WHERE covenant_name_hash = $1
ORDER BY height DESC
LIMIT $3::smallint OFFSET $2::smallint;
`

func (q *Queries) GetNameRecordHistoryByNameHash(ctx context.Context, params GetNameRecordHistoryByNameHashParams) ([]RecordRow, error) {
	rows, err := q.db.QueryContext(ctx, getNameRecordHistoryByNameHash, params.NameHash, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []RecordRow{}
	for rows.Next() {
		var i RecordRow
		if err := rows.Scan(
			&i.Height,
			&i.CovenantRecordData,
			&i.Count,
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

type GetMostExpensiveNamesParams struct {
	Offset int16
	Limit  int16
}

const getMostExpensiveNames = `-- name: GetMostExpensiveNames :many
SELECT
open_height,
covenant_name_hash,
name,
max_lockup,
max_revealed,
bidcount,
(COUNT(*) OVER())::bigint as count
FROM names
ORDER BY max_revealed DESC 
LIMIT $2::smallint OFFSET $1::smallint;
`

func (q *Queries) GetMostExpensiveNames(ctx context.Context, params GetMostExpensiveNamesParams) ([]NameRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostExpensiveNames, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NameRow{}
	for rows.Next() {
		var i NameRow
		if err := rows.Scan(
			&i.OpenHeight,
			&i.CovenantNameHash,
			&i.CovenantName,
			&i.MaxLockup,
			&i.MaxRevealed,
			&i.BidCount,
			&i.Count,
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

type NameVolumeRow struct {
	OpenHeight       int32
	CovenantNameHash types.Bytes
	CovenantName     types.Bytes
	MaxLockup        int64
	MaxRevealed      int64
	VolumeLockup     int64
	VolumeRevealed   int64
	BidCount         int16
	Count            int32
}

type GetMostLockupVolumeNamesParams struct {
	Offset int16
	Limit  int16
}

const getMostLockupVolumeNames = `-- name: GetMostLockupVolumeNames :many
SELECT
open_height,
covenant_name_hash,
name,
max_lockup,
max_revealed,
sum_lockup,
sum_revealed,
bidcount,
(COUNT(*) OVER())::bigint as count
FROM names
ORDER BY sum_lockup DESC 
LIMIT $2::smallint OFFSET $1::smallint;
`

func (q *Queries) GetMostLockupVolumeNames(ctx context.Context, params GetMostLockupVolumeNamesParams) ([]NameVolumeRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostLockupVolumeNames, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NameVolumeRow{}
	for rows.Next() {
		var i NameVolumeRow
		if err := rows.Scan(
			&i.OpenHeight,
			&i.CovenantNameHash,
			&i.CovenantName,
			&i.MaxLockup,
			&i.MaxRevealed,
			&i.VolumeLockup,
			&i.VolumeRevealed,
			&i.BidCount,
			&i.Count,
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

type GetMostRevealVolumeNamesParams struct {
	Offset int16
	Limit  int16
}

const getMostRevealVolumeNames = `-- name: GetMostRevealVolumeNames :many
SELECT
open_height,
covenant_name_hash,
name,
max_lockup,
max_revealed,
sum_lockup,
sum_revealed,
bidcount,
(COUNT(*) OVER())::bigint as count
FROM names
ORDER BY sum_revealed DESC 
LIMIT $2::smallint OFFSET $1::smallint;
`

func (q *Queries) GetMostRevealVolumeNames(ctx context.Context, params GetMostRevealVolumeNamesParams) ([]NameVolumeRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostRevealVolumeNames, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NameVolumeRow{}
	for rows.Next() {
		var i NameVolumeRow
		if err := rows.Scan(
			&i.OpenHeight,
			&i.CovenantNameHash,
			&i.CovenantName,
			&i.MaxLockup,
			&i.MaxRevealed,
			&i.VolumeLockup,
			&i.VolumeRevealed,
			&i.BidCount,
			&i.Count,
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

type GetMostBidsNamesParams struct {
	Offset int16
	Limit  int16
}

const getMostBidsNames = `-- name: GetMostBidsNames :many
SELECT
open_height,
covenant_name_hash,
name,
max_lockup,
max_revealed,
sum_lockup,
sum_revealed,
bidcount,
(COUNT(*) OVER())::bigint as count
FROM names
ORDER BY bidcount DESC 
LIMIT $2::smallint OFFSET $1::smallint;
`

func (q *Queries) GetMostBidsNames(ctx context.Context, params GetMostBidsNamesParams) ([]NameVolumeRow, error) {
	rows, err := q.db.QueryContext(ctx, getMostBidsNames, params.Offset, params.Limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []NameVolumeRow{}
	for rows.Next() {
		var i NameVolumeRow
		if err := rows.Scan(
			&i.OpenHeight,
			&i.CovenantNameHash,
			&i.CovenantName,
			&i.MaxLockup,
			&i.MaxRevealed,
			&i.VolumeLockup,
			&i.VolumeRevealed,
			&i.BidCount,
			&i.Count,
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
