-- +goose Up
-- +goose StatementBegin
-- +goose StatementEnd
ALTER TABLE user_socials
    ADD COLUMN discord_refresh_token TEXT,
    ADD COLUMN google_refresh_token TEXT;
-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_socials
    DROP COLUMN IF EXISTS discord_refresh_token,
    DROP COLUMN IF EXISTS google_refresh_token;
-- +goose StatementEnd
