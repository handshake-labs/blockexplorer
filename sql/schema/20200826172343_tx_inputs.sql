-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE tx_inputs (
  tx_hash bytea NOT NULL REFERENCES transactions (hash) ON DELETE CASCADE,
  index integer NOT NULL CHECK(index >= 0),
  hash_prevout bytea NOT NULL CHECK (LENGTH(hash_prevout) = 32),
  index_prevout bigint NOT NULL,
  sequence bigint NOT NULL,
  PRIMARY KEY (tx_hash, index)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tx_inputs;
