-- name: InsertTransaction :exec
INSERT INTO transactions (txid, witness_tx, fee, rate, block_hash, index_block, "version", locktime, "size")
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);


-- name: GetTransactionByTxid :one
SELECT
    *
FROM
    transactions
WHERE
    txid = $1;


-- name: GetTransactionsByBlockHash :many
SELECT *, COUNT(*) OVER() as count
FROM transactions
WHERE block_hash = $1
ORDER BY index_block
LIMIT $2 OFFSET $3;
