-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN create_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
DROP COLUMN create_at;
-- +goose StatementEnd
