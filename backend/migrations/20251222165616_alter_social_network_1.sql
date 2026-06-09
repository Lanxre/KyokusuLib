-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_socials ADD COLUMN IF NOT EXISTS discord_refresh_token TEXT;
ALTER TABLE user_socials ADD COLUMN IF NOT EXISTS google_refresh_token TEXT;
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_socials
    DROP COLUMN IF EXISTS discord_refresh_token,
    DROP COLUMN IF EXISTS google_refresh_token;
-- +goose StatementEnd
