-- name: InsertBlock :exec
INSERT INTO blocks (hash, height, weight, size, version, hash_merkle_root, witness_root, tree_root, reserved_root, mask, time, bits, difficulty, chainwork, nonce, extra_nonce)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16);

-- name: GetBlockByHeight :one
SELECT *
FROM blocks
WHERE height = $1;

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
