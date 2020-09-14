-- name: InsertTransaction :exec
INSERT INTO transactions (txid, witness_tx, fee, rate, block_hash, index, "version", locktime, "size")
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);


-- name: GetTransactionByTxid :one
SELECT transactions.*, blocks.height FROM transactions, blocks WHERE transactions.block_hash=blocks.hash AND transactions.txid = $1;

-- name: GetTransactionsByBlockHash :many
SELECT *, (COUNT(*) OVER())::smallint as count
FROM transactions
WHERE block_hash = sqlc.arg(block_hash)::bytea
ORDER BY index
LIMIT sqlc.arg('limit')::smallint OFFSET sqlc.arg('offset')::smallint;
