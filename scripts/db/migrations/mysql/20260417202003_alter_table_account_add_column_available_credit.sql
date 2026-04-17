-- +goose Up
-- +goose StatementBegin
ALTER TABLE `account` ADD COLUMN `available_credit` DECIMAL(17,2) DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
