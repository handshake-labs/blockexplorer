-- +goose Up
-- +goose StatementBegin
CREATE VIEW viewblock AS SELECT (hash, "height", weight, blocks.size, blocks.version, hash_merkle_root, witness_root, tree_root,
  reserved_root, mask, "time", bits, difficulty, chainwork, nonce, extra_nonce, count(transactions), sum())
from blocks, transactions, tx_inputs, tx_outputs
where blocks.hash=transactions.block_hash AND tx_inputs.txid=transactions.txid AND tx_outputs.txid=transactions.txid;

-- +goose Down
-- +goose StatementBegin
DROP VIEW viewblock;
-- +goose StatementEnd
