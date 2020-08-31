-- +goose Up
-- +goose StatementBegin
CREATE VIEW viewblock AS SELECT (hash, "height", weight, "size", "version", hash_merkle_root, witness_root, tree_root,
  reserved_root, mask, "time", bits, difficulty, chainwork, nonce, extra_nonce, count(transactions), sum()
from blocks, transactions, tx_inputs, tx_outputs
where blocks.hash=transactions.block_hash AND tx_inputs.txid=transactions.txid AND tx_outputs.txid=transactions.txid;

-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE blocks (
    hash bytea NOT NULL PRIMARY KEY CHECK (LENGTH(hash) = 32),
    "height" integer UNIQUE NOT NULL CHECK (HEIGHT >= 0),
    weight integer NOT NULL,
    "size" bigint NOT NULL,
    "version" integer NOT NULL,
    hash_merkle_root bytea CHECK (LENGTH(hash_merkle_root) = 32) NOT NULL,
    witness_root bytea CHECK (LENGTH(witness_root) = 32) NOT NULL,
    tree_root bytea CHECK (LENGTH(tree_root) = 32) NOT NULL,
    reserved_root bytea CHECK (LENGTH(reserved_root) = 32) NOT NULL,
    mask bytea CHECK (LENGTH(mask) = 32) NOT NULL,
    "time" integer NOT NULL,
    bits bytea CHECK (LENGTH(bits) = 4) NOT NULL,
    difficulty double precision NOT NULL,
    chainwork bytea CHECK (LENGTH(chainwork) = 32) NOT NULL,
    nonce bigint NOT NULL,
    extra_nonce bytea CHECK (LENGTH(extra_nonce) = 24) NOT NULL,
    orphan boolean NOT NULL DEFAULT FALSE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP VIEW viewblock;
-- +goose StatementEnd
