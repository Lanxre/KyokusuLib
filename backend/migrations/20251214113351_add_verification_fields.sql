-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
ADD COLUMN is_verified BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN verification_token VARCHAR(255),
ADD COLUMN verification_token_expires_at TIMESTAMP;

CREATE INDEX idx_users_verification_token ON users(verification_token);
UPDATE users SET is_verified = TRUE WHERE is_verified = FALSE;


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP INDEX IF EXISTS idx_users_verification_token;

ALTER TABLE users 
DROP COLUMN verification_token_expires_at,
DROP COLUMN verification_token,
DROP COLUMN is_verified;
-- +goose StatementEnd
