-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE transactions (
  hash bytea NOT NULL PRIMARY KEY CHECK (LENGTH(hash) = 32),
  block_hash bytea NOT NULL REFERENCES blocks (hash) ON DELETE CASCADE,
  witness_tx bytea UNIQUE CHECK (LENGTH(witness_tx) = 32) NOT NULL,
  fee bigint NOT NULL,
  rate bigint NOT NULL,
  version integer NOT NULL,
  locktime integer NOT NULL,
  size bigint NOT NULL
);

CREATE INDEX ON transactions (block_hash);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE transactions;
