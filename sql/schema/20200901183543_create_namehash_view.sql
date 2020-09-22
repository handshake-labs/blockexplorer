-- +goose Up
-- +goose StatementBegin
create materialized view namehash as
select distinct on(out1.covenant_name_hash, out2.covenant_name)
out1.covenant_name_hash,
out2.covenant_name
from tx_outputs out1, tx_outputs out2 where
out1.covenant_name_hash = out2.covenant_name_hash
AND out2.covenant_name IS NOT NULL;
-- +goose StatementEnd
CREATE UNIQUE INDEX namehash_name_index ON namehash(covenant_name);

-- +goose Down
-- +goose StatementBegin
DROP materialized VIEW namehash;
-- +goose StatementEnd
