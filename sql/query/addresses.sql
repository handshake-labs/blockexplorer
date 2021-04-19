--This query can be optimized to be very quick by removing join for the name,
--however as it's still quicker than the GetAddressInfo I've left the name for the sake of simplicity

-- name: GetTxOutputsByAddress :many
SELECT
  DISTINCT tx_outputs.*,
  COALESCE(tx_inputs.txid, '\x') AS hash_prevout_not_null,
  COALESCE(tx_inputs.index, -1) AS index_prevout_not_null,
  COALESCE(bl2.height, -1) AS spend_height_not_null, --height of -1 means mempool, so i need -2 to indicate the block does not exist 
  blocks.height AS height,
  COALESCE(CONVERT_FROM(t2.covenant_name, 'SQL_ASCII'), '')::text AS name
FROM tx_outputs
  LEFT JOIN tx_outputs t2 ON (tx_outputs.covenant_name_hash = t2.covenant_name_hash AND t2.covenant_name IS NOT NULL)
  LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
  JOIN transactions ON tx_outputs.txid = transactions.txid
  JOIN blocks ON transactions.block_hash = blocks.hash --for height of receive
  LEFT JOIN transactions tx2 ON tx_inputs.txid = tx2.txid
  LEFT JOIN blocks  bl2 ON tx2.block_hash = bl2.hash --for height of spend
WHERE tx_outputs.address = sqlc.arg('address')::text
ORDER BY blocks.height DESC 
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;


--This query takes a lot of time, perhaps can be optimized further

-- name: GetAddressInfo :one
SELECT
  COALESCE(SUM(tx_outputs.value), 0)::bigint AS value_total,
  COALESCE(SUM(tx_outputs.value) filter (WHERE tx_inputs.txid IS NOT NULL), 0)::bigint AS value_used,
  COUNT(tx_outputs.txid) AS tx_outputs_total,
  COUNT(tx_inputs.hash_prevout) AS tx_outputs_used
FROM tx_outputs
LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
WHERE tx_outputs.address = sqlc.arg('address')::text;

-- name: AddressExists :one
SELECT EXISTS(SELECT 1 FROM tx_outputs WHERE tx_outputs.address = sqlc.arg('address')::text); 
