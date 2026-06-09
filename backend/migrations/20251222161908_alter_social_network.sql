-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_socials ADD COLUMN IF NOT EXISTS discord_id NUMERIC;
ALTER TABLE user_socials ADD COLUMN IF NOT EXISTS google_id NUMERIC;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_socials
    DROP COLUMN discord_id,
    DROP COLUMN google_id;
-- +goose StatementEnd
