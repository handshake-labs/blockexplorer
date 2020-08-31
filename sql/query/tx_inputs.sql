-- name: InsertTxInput :exec
INSERT INTO tx_inputs (txid, index, hash_prevout, index_prevout, sequence, block_hash)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetTxInputsByTxHash :many
SELECT *
FROM tx_inputs
WHERE txid = $1
ORDER BY index;
