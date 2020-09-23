-- name: InsertTxInput :exec
INSERT INTO tx_inputs (txid, index, hash_prevout, index_prevout, sequence)
VALUES ($1, $2, $3, $4, $5);

-- name: GetTxInputsByTxid :many
SELECT *
FROM tx_inputs
WHERE txid = $1
ORDER BY index;
