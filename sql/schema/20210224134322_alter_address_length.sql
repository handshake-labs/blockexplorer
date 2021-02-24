-- +goose Up
-- +goose StatementBegin
ALTER TABLE tx_outputs DROP CONSTRAINT tx_outputs_covenant_address_check;
ALTER TABLE tx_outputs ADD CONSTRAINT tx_outputs_covenant_address CHECK (LENGTH(covenant_address) <= 40); 
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
