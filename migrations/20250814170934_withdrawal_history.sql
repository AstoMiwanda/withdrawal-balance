-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS withdrawal_histories (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID unik untuk riwayat penarikan',
    user_id BIGINT NOT NULL,
    wallet_id BIGINT NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    bank_account_number VARCHAR(255) NOT NULL,
    bank_name VARCHAR(255) NOT NULL,
    status VARCHAR(50) NOT NULL DEFAULT 'pending' COMMENT 'Status penarikan: pending, completed, failed',
    transaction_reference VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE RESTRICT,
    FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE RESTRICT,
    INDEX (user_id)
) COMMENT='Menyimpan riwayat penarikan saldo';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS withdrawal_histories;
-- +goose StatementEnd
