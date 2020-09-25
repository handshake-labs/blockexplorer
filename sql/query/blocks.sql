-- name: InsertBlock :exec
INSERT INTO blocks (hash, height, weight, size, version, hash_merkle_root, witness_root, tree_root, reserved_root, mask, time, bits, difficulty, chainwork, nonce, extra_nonce)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);

-- name: GetBlocks :many
SELECT *, (
  SELECT COUNT(*)
  FROM transactions
  WHERE block_hash = blocks.hash
)::integer AS txs_count
FROM blocks
ORDER BY height DESC
LIMIT $1 OFFSET $2;

-- name: GetBlockByHeight :one
SELECT *
FROM blocks
WHERE height = $1;

-- name: GetBlockByHash :one
SELECT *, (
  SELECT COUNT(*)
  FROM transactions
  WHERE block_hash = blocks.hash
)::integer AS txs_count
FROM blocks
WHERE hash = $1;

-- name: GetBlockHashByHeight :one
SELECT hash
FROM blocks
WHERE height = $1;

-- name: GetBlocksMaxHeight :one
SELECT COALESCE(MAX(height), -1)::integer
FROM blocks;

-- name: DeleteBlocksAfterHeight :exec
DELETE FROM blocks
WHERE height > $1;
