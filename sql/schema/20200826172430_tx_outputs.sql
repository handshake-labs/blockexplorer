-- +goose Up
-- SQL in this section is executed when the migration is applied.
CREATE TYPE covenant_action
AS ENUM ('NONE',
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
    'REVOKE');

CREATE TABLE tx_outputs (
    txid bytea NOT NULL, --tx hash in which output occured
    index integer NOT NULL CHECK(index >= 0),
    value bigint NOT NULL,
    block_hash bytea NOT NULL,
    address varchar(90) NOT NULL, --can be null for coinbase transactions
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
    FOREIGN KEY (txid, block_hash) REFERENCES transactions (txid, block_hash) ON DELETE CASCADE,
    PRIMARY KEY (txid, index)
);

-- to select outputs of some transaction
CREATE INDEX tx_outputs_covenant_name_hash_index ON tx_outputs
USING btree (covenant_name_hash);

--to select outputs of some name by its hash
CREATE INDEX tx_outputs_covenant_name_index ON tx_outputs
USING btree (covenant_name);

CREATE INDEX tx_outputs_address_index ON tx_outputs
USING btree (address);

--to select outputs leading to some address
CREATE INDEX tx_outputs_txid_index ON tx_outputs
USING btree (txid);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE tx_outputs;

DROP TYPE covenant_action;
