-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TYPE covenant_action
AS ENUM (
  'NONE',
  'CLAIM',
  'OPEN',
  'BID',
  'REVEAL',
  'REDEEM',
  'REGISTER',
  'UPDATE',
  'RENEW',
  'TRANSFER',
  'FINALIZE',
  'REVOKE'
);

CREATE TABLE tx_outputs (
  tx_hash bytea NOT NULL REFERENCES transactions (hash) ON DELETE CASCADE,
  index integer NOT NULL CHECK(index >= 0),
  value bigint NOT NULL,
  address text NOT NULL CHECK(LENGTH(address) <= 90),
  covenant_action covenant_action NOT NULL,
  covenant_name_hash bytea CHECK (LENGTH(covenant_name_hash) = 32),
  covenant_height bytea CHECK (LENGTH(covenant_height) = 4),
  covenant_name bytea CHECK (LENGTH(covenant_name) <= 63),
  covenant_bid_hash bytea CHECK (LENGTH(covenant_bid_hash) = 32),
  covenant_nonce bytea CHECK (LENGTH(covenant_nonce) = 32),
  covenant_record_data bytea CHECK (LENGTH(covenant_record_data) <= 512),
  covenant_block_hash bytea CHECK (LENGTH(covenant_block_hash) = 32),
  covenant_version bytea CHECK (LENGTH(covenant_version) = 1),
  covenant_address bytea CHECK (LENGTH(covenant_address) = 20),
  covenant_claim_height bytea CHECK (LENGTH(covenant_claim_height) = 4),
  covenant_renewal_count bytea CHECK (LENGTH(covenant_renewal_count) = 4),
  PRIMARY KEY (tx_hash, index)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tx_outputs;
DROP TYPE covenant_action;
