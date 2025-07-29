-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `account` (
    `account_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `document_number` VARCHAR(11) NOT NULL UNIQUE KEY
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
