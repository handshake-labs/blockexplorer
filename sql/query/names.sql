-- name: GetReservedName :one
SELECT *
FROM reserved_names
WHERE name = $1;

-- name: GetNameCountsByHash :one
SELECT
  (COUNT(*) FILTER (WHERE covenant_action = 'BID'))::integer AS bids_count,
  COUNT(covenant_record_data)::integer AS records_count
FROM tx_outputs
WHERE covenant_name_hash = sqlc.arg('name_hash')::bytea;

-- name: GetNameBidsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height,
  lockups.value AS lockup_value,
  reveals.value AS reveal_value
FROM
  tx_outputs lockups
  INNER JOIN transactions ON (lockups.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
  LEFT JOIN tx_outputs reveals ON (lockups.covenant_name_hash = reveals.covenant_name_hash AND lockups.address = reveals.address)
  LEFT JOIN tx_inputs ON (reveals.txid = tx_inputs.txid AND lockups.index = tx_inputs.index)
WHERE lockups.covenant_action = 'BID' AND reveals.covenant_action = 'REVEAL' AND lockups.covenant_name_hash = sqlc.arg('name_hash')::bytea
ORDER BY (blocks.height, transactions.index, lockups.index) DESC NULLS FIRST
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;

-- name: GetNameRecordsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height,
  tx_outputs.covenant_record_data::bytea AS data
FROM
  tx_outputs
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE tx_outputs.covenant_record_data IS NOT NULL AND tx_outputs.covenant_name_hash = sqlc.arg('name_hash')::bytea
ORDER BY (blocks.height, transactions.index, tx_outputs.index) DESC NULLS FIRST
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;
