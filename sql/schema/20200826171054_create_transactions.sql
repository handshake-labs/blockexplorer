-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transactions (
    txid bytea PRIMARY KEY CHECK (LENGTH(txid) = 32),
    witness_tx bytea CHECK (LENGTH(witness_tx) = 32) NOT NULL, --wtxid, witness data of transaction
    fee bigint NOT NULL,
    rate bigint NOT NULL,
    block_hash bytea REFERENCES blocks (hash) ON DELETE CASCADE,
    index integer,
    "version" integer NOT NULL,
    locktime integer NOT NULL,
    "size" bigint NOT NULL
);

CREATE UNIQUE INDEX transactions_block_hash_index ON transactions (block_hash, index) WHERE block_hash IS NOT NULL;

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE transactions;
