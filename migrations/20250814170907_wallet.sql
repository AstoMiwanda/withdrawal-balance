-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS wallets (
    id BIGINT PRIMARY KEY AUTO_INCREMENT COMMENT 'ID unik untuk dompet',
    user_id BIGINT NOT NULL,
    balance DECIMAL(15, 2) NOT NULL DEFAULT 0.00,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
) COMMENT='Menyimpan saldo dompet pengguna';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS wallets;
-- +goose StatementEnd
