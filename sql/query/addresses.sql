-- name: GetTxOutputsByAddress :many
SELECT
  DISTINCT tx_outputs.*,
  COALESCE(tx_inputs.txid, '') AS spend_txid,
  COALESCE(tx_inputs.index, -1) AS spend_index,
  blocks.height AS height,
  COALESCE(CONVERT_FROM(t2.covenant_name, 'SQL_ASCII'), '')::text AS name
FROM tx_outputs
LEFT JOIN tx_outputs t2 ON (tx_outputs.covenant_name_hash = t2.covenant_name_hash AND t2.covenant_name IS NOT NULL)
LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
JOIN transactions ON tx_outputs.txid = transactions.txid
JOIN blocks ON transactions.block_hash = blocks.hash
WHERE tx_outputs.address = sqlc.arg('address')::text
ORDER BY blocks.height DESC 
LIMIT sqlc.arg('limit')::integer OFFSET sqlc.arg('offset')::integer;

-- name: GetAddressInfo :one
SELECT
  COALESCE(SUM(tx_outputs.value), 0)::bigint AS value_total,
  COALESCE(SUM(tx_outputs.value) filter (WHERE tx_inputs.txid IS NOT NULL), 0)::bigint AS value_used,
  COUNT(*) AS tx_outputs_total,
  COUNT(tx_inputs.*) AS tx_outputs_used
FROM tx_outputs
LEFT JOIN tx_inputs ON tx_outputs.txid = tx_inputs.hash_prevout AND tx_outputs.index = tx_inputs.index_prevout
WHERE tx_outputs.address = sqlc.arg('address')::text;
