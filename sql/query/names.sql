-- name: GetNameOpeningBlock :one
SELECT blocks.height
FROM tx_outputs, transactions, blocks WHERE
covenant_name = $1 AND
covenant_action = 'OPEN' AND
tx_outputs.txid = transactions.txid AND
blocks.hash = transactions.block_hash;


-- get auction history of a name with reveals, input parameter - hash of the name 
-- name: GetAuctionHistoryByNameHash :many
SELECT height, transactions.txid, A.value AS lockup, B.value AS reveal, A.covenant_action
FROM transactions, blocks, tx_outputs A LEFT OUTER JOIN tx_outputs B ON (A.covenant_name_hash = B.covenant_name_hash AND B.covenant_action =
  'REVEAL' AND A.address = B.address AND A.covenant_action = 'BID') WHERE A.txid=transactions.txid AND
transactions.block_hash = blocks.hash AND A.covenant_name_hash=$1 ORDER BY height DESC;


-- domain top list by lockup with count and revealed values 
-- name: GetTopList :many

create view domains as
-- SELECT max(blocks.height) AS open_height, bids.covenant_name, bids.covenant_name_hash AS name, max(lockups.value) AS max_lockup, max(reveals.value) AS max_revealed,
SELECT max(blocks.height) AS open_height, bids.covenant_name_hash, bids.covenant_name AS name, max(lockups.value) AS max_lockup, max(reveals.value) AS max_revealed,
count(distinct(reveals.txid, reveals.index)) AS bidcount
FROM transactions, blocks, tx_outputs lockups, tx_outputs opens, tx_outputs reveals LEFT OUTER JOIN
tx_outputs bids ON (bids.covenant_name_hash = reveals.covenant_name_hash) WHERE
lockups.covenant_name_hash = bids.covenant_name_hash AND
reveals.covenant_action = 'REVEAL' AND
bids.covenant_action='BID' AND
opens.covenant_name_hash = bids.covenant_name_hash AND --covenant_name = $1 AND
opens.covenant_action = 'OPEN' AND
opens.txid = transactions.txid AND
blocks.hash = transactions.block_hash
GROUP BY
bids.covenant_name, bids.covenant_name_hash;

SELECT bids.covenant_name AS name, max(lockups.value) AS max_lockup, max(reveals.value) AS max_revealed,
-- count(distinct(reveals.txid, reveals.index_out)) AS bidcount FROM tx_outputs lockups, tx_outputs reveals LEFT OUTER JOIN
count(distinct(reveals.txid, reveals.index)) AS bidcount
FROM tx_outputs lockups, tx_outputs reveals LEFT OUTER JOIN
tx_outputs bids ON (reveals.covenant_name_hash = bids.covenant_name_hash) WHERE lockups.covenant_name_hash =
bids.covenant_name_hash AND reveals.covenant_action = 'REVEAL' AND (bids.covenant_action='BID') GROUP BY
bids.covenant_name ORDER BY max_lockup DESC;


SELECT height, transactions.txid, A.covenant_name, A.value AS lockup, B.value AS reveal, A.covenant_action, A.covenant_record_data
FROM transactions, blocks, tx_outputs A LEFT OUTER JOIN tx_outputs B ON (A.covenant_name_hash = B.covenant_name_hash AND B.covenant_action =
  'REVEAL' AND A.address = B.address AND A.covenant_action = 'BID') WHERE A.txid=transactions.txid AND
transactions.block_hash = blocks.hash AND A.covenant_name_hash='\x9db397b206f0abc65dfb5dc32ce5388ff34640d3c3405664cb0f30454413f5f2' ORDER BY height DESC;



create view records as select blocks.height, covenant_record_data, covenant_name_hash from tx_outputs, blocks where blocks.hash
= tx_outputs.block_hash and covenant_record_data is not null group by covenant_action, blocks.height, covenant_record_data, covenant_name_hash ORDER BY height desc;

