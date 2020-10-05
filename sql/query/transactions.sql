-- name: InsertTransaction :exec
INSERT INTO transactions (txid, witness_tx, fee, rate, block_hash, index, "version", locktime, "size")
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9);

-- name: GetTransactionByTxid :one
SELECT
  transactions.*,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null
FROM transactions
  LEFT JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE transactions.txid = $1;

-- name: GetTransactionsByBlockHeight :many
SELECT
  transactions.*,
  COALESCE(blocks.height, -1)::integer AS block_height_not_null
FROM
  transactions
  INNER JOIN blocks ON (transactions.block_hash = blocks.hash)
WHERE blocks.height = $1
ORDER BY transactions.index
LIMIT $2 OFFSET $3;

-- name: GetMempoolTransactions :many
SELECT *
FROM transactions
WHERE block_hash IS NULL
ORDER BY index
LIMIT $1 OFFSET $2;

-- name: DeleteMempool :exec
DELETE FROM transactions
WHERE block_hash IS NULL;
