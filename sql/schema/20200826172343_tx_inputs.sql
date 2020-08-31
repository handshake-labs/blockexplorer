-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tx_inputs (
    txid bytea NOT NULL,
    "index" bigint NOT NULL,
    hash_prevout bytea NOT NULL CHECK (LENGTH(hash_prevout) = 32), --tx id of previous output tx
    index_prevout bigint NOT NULL, -- txin_witness bytea NOT NULL CHECK (LENGTH(txin_witness) = 32),
    "sequence" bigint NOT NULL,
    block_hash bytea NOT NULL,
    FOREIGN KEY (txid, block_hash) REFERENCES transactions (txid, block_hash) ON DELETE CASCADE
);

CREATE INDEX tx_inputs_txid_index ON tx_inputs USING btree (txid);

--to select inputs belonging to some given transaction
CREATE INDEX tx_inputs_hash_prevout_index ON tx_inputs
USING btree (hash_prevout);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tx_inputs;

