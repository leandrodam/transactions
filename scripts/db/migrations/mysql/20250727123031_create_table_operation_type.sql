-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `operation_type` (
    `operation_type_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `description` VARCHAR(50) NOT NULL UNIQUE KEY
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
