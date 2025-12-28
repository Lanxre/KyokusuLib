-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN reset_token VARCHAR(255),
ADD COLUMN reset_token_expires_at TIMESTAMP;

CREATE INDEX idx_users_reset_token ON users(reset_token);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_reset_token;

ALTER TABLE users 
DROP COLUMN reset_token_expires_at,
DROP COLUMN reset_token;
-- +goose StatementEnd
