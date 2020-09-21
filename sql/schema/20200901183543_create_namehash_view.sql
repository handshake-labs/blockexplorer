-- +goose Up
-- +goose StatementBegin
create view namehash as
select distinct on(out1.covenant_name_hash, out2.covenant_name)
out1.covenant_name_hash,
out2.covenant_name
from tx_outputs out1, tx_outputs out2 where
out1.covenant_name_hash = out2.covenant_name_hash
and out1.covenant_name IS NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop view namehash;
-- +goose StatementEnd
