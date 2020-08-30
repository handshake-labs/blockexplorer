-- name: InsertTransaction :exec
INSERT INTO transactions (hash, block_hash, witness_tx, fee, rate, version, locktime, size)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8);
