-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `transaction` (
    `transaction_id` INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `account_id` INT NOT NULL,
    `operation_type_id` INT NOT NULL,
    `amount` DECIMAL(17,2) NOT NULL,
    `event_date` DATETIME NOT NULL DEFAULT NOW(),
    FOREIGN KEY (`account_id`) REFERENCES `account` (`account_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
    FOREIGN KEY (`operation_type_id`) REFERENCES `operation_type` (`operation_type_id`) ON DELETE RESTRICT ON UPDATE CASCADE,
    INDEX (`event_date`)
) ENGINE=InnoDB DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- +goose StatementEnd
