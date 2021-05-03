-- name: GetReservedName :one
SELECT 
CONVERT_FROM(origin_name, 'SQL_ASCII')::text as origin_name,
CONVERT_FROM(name, 'SQL_ASCII')::text as name,
name_hash,
claim_amount
FROM reserved_names
WHERE name = $1;

-- name: GetNameCountsByHash :one
SELECT
  (COUNT(*) FILTER (WHERE covenant_action = 'BID'))::integer AS bids_count,
  COUNT(covenant_record_data)::integer AS records_count,
  (COUNT(*) FILTER (WHERE covenant_action = 'CLAIM' OR covenant_action = 'RENEW' OR covenant_action = 'TRANSFER' OR covenant_action = 'FINALIZE' OR covenant_action = 'REVOKE'))::integer AS actions_count
FROM tx_outputs
WHERE covenant_name_hash = sqlc.arg('name_hash')::bytea;

-- name: GetNameBidsByHash :many
SELECT
  DISTINCT ON (blocks.height, bid_txid, lockup_outputs.index)
  bids.txid AS bid_txid, 
  COALESCE(blocks.height, -1)::integer AS block_height_not_null,
  COALESCE(reveals.txid, '\x00')::bytea AS reveal_txid,
  COALESCE(reveal_blocks.height, -1)::integer AS reveal_height_not_null,
  COALESCE(reveal_outputs.index, -1)::integer AS reveal_index_not_null,
  lockup_outputs.value as lockup_value,
  COALESCE(reveal_outputs.value, -1) as reveal_value_not_null
FROM                                                  
  transactions as bids
  JOIN tx_inputs as lockup_inputs ON lockup_inputs.txid=bids.txid
  LEFT JOIN blocks ON (bids.block_hash = blocks.hash)
  JOIN tx_outputs as lockup_outputs ON lockup_outputs.txid=bids.txid AND lockup_outputs.covenant_action = 'BID'
  LEFT JOIN tx_inputs reveal_inputs ON
     reveal_inputs.hash_prevout = lockup_outputs.txid AND
     reveal_inputs.index_prevout = lockup_outputs.index
  LEFT JOIN tx_outputs reveal_outputs ON reveal_outputs.covenant_action = 'REVEAL'  AND reveal_outputs.covenant_name_hash = lockup_outputs.covenant_name_hash AND
     reveal_inputs.txid = reveal_outputs.txid AND
     reveal_inputs.index = reveal_outputs.index 
  LEFT JOIN transactions AS reveals ON reveal_inputs.txid = reveals.txid AND reveal_outputs.txid = reveals.txid
  LEFT JOIN blocks as reveal_blocks ON (reveals.block_hash = reveal_blocks.hash)
WHERE lockup_outputs.covenant_name_hash = sqlc.arg('name_hash')::bytea
ORDER BY blocks.height DESC NULLS FIRST
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;

-- name: GetNameRecordsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null,
  tx_outputs.covenant_record_data::bytea AS data
FROM
  tx_outputs
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE tx_outputs.covenant_record_data IS NOT NULL AND tx_outputs.covenant_name_hash = sqlc.arg('name_hash')::bytea
ORDER BY (blocks.height, transactions.index, tx_outputs.index) DESC NULLS FIRST
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;

-- name: GetLastNameBlockHeightByActionAndHash :one
SELECT
  COALESCE(blocks.height, -1)::integer AS block_height_not_null
FROM
  tx_outputs
  INNER JOIN transactions ON (tx_outputs.txid = transactions.txid)
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE covenant_action = $1 AND covenant_name_hash = $2
ORDER BY blocks.height DESC NULLS FIRST
LIMIT 1;

-- name: GetNameOtherActionsByHash :many
SELECT
  transactions.txid AS txid,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null,
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
  tx_outputs.covenant_name_hash = sqlc.arg('name_hash')::bytea  AND
  tx_outputs.covenant_record_data IS NULL
ORDER BY (blocks.height, transactions.index, tx_outputs.index) DESC NULLS FIRST
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;
