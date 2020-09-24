-- +goose Up
CREATE MATERIALIZED VIEW auctions AS
SELECT distinct on (txid, lockups.index)
height,
transactions.txid,
lockups.covenant_name,
lockups.value AS lockup,
reveals.value AS reveal,
lockups.covenant_action,
lockups.covenant_record_data,
lockups.covenant_name_hash,
lockups.index
FROM transactions, blocks, tx_outputs lockups LEFT OUTER JOIN tx_outputs reveals ON
(lockups.covenant_name_hash = reveals.covenant_name_hash AND
reveals.covenant_action = 'REVEAL' AND
lockups.address = reveals.address AND
lockups.covenant_action = 'BID')
left join tx_inputs ins on (ins.txid = reveals.txid) AND (ins.index = lockups.index)
WHERE
transactions.block_hash IS NOT NULL AND
lockups.txid=transactions.txid AND
transactions.block_hash = blocks.hash;

CREATE UNIQUE INDEX auctions_index ON auctions (txid, index);


-- +goose Down
DROP MATERIALIZED VIEW auctions;


