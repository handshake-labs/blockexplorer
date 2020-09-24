-- name: InsertTransaction :exec
INSERT INTO transactions (txid, witness_tx, fee, rate, block_hash, index, "version", locktime, "size")
    VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetTransactionByTxid :one
SELECT transactions.*, COALESCE(blocks.height, -1) FROM transactions LEFT OUTER JOIN blocks ON transactions.block_hash=blocks.hash WHERE transactions.txid = $1;

-- name: GetTransactionsByBlockHash :many
SELECT *
FROM transactions
WHERE block_hash = $1
ORDER BY index
LIMIT $2 OFFSET $3;


-- name: GetMempoolTransactions :many
SELECT *
FROM transactions
WHERE block_hash IS NULL 
ORDER BY index
LIMIT $1 OFFSET $2;
