-- name: InsertTransaction :exec
INSERT INTO transactions (hash, block_hash, witness_tx, fee, rate, version, locktime, size)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetTransactionsByBlockHash :many
SELECT *, (COUNT(*) OVER())::smallint as count
FROM transactions
WHERE block_hash = sqlc.arg(block_hash)::bytea
ORDER BY hash
LIMIT sqlc.arg('limit')::smallint OFFSET sqlc.arg('offset')::smallint;
