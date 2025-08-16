-- +goose Up
-- +goose StatementBegin
INSERT INTO wallets (id, user_id, balance, created_at, updated_at)
VALUES (1, 1, 1000000.00, DEFAULT, DEFAULT);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM wallets
WHERE id = 1;
-- +goose StatementEnd
