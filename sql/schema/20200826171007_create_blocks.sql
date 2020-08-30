-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TABLE blocks (
  hash bytea NOT NULL PRIMARY KEY CHECK (LENGTH(hash) = 32),
  height integer NOT NULL UNIQUE CHECK(height >= 0),
  weight integer NOT NULL,
  size bigint NOT NULL,
  version integer NOT NULL,
  hash_merkle_root bytea CHECK (LENGTH(hash_merkle_root) = 32) NOT NULL,
  witness_root bytea CHECK (LENGTH(witness_root) = 32) NOT NULL,
  tree_root bytea CHECK (LENGTH(tree_root) = 32) NOT NULL,
  reserved_root bytea CHECK (LENGTH(reserved_root) = 32) NOT NULL,
  mask bytea CHECK (LENGTH(mask) = 32) NOT NULL,
  time integer NOT NULL,
  bits bytea CHECK (LENGTH(bits) = 4) NOT NULL,
  difficulty double precision NOT NULL,
  chainwork bytea CHECK (LENGTH(chainwork) = 32) NOT NULL,
  nonce bigint NOT NULL,
  extra_nonce bytea CHECK (LENGTH(extra_nonce) = 24) NOT NULL
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE blocks;
