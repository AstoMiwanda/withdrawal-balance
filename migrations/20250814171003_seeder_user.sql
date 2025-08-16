-- +goose Up
-- +goose StatementBegin
INSERT INTO users (id, name, phone, email, created_at, updated_at)
VALUES (1, 'Asto', '081399998888', 'asto@test-mail.com', DEFAULT, DEFAULT);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DELETE
FROM users
WHERE id = 1;
-- +goose StatementEnd
