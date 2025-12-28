-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_socials
    ADD COLUMN discord_id   NUMERIC,
    ADD COLUMN google_id    NUMERIC;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_socials
    DROP COLUMN discord_id,
    DROP COLUMN google_id;
-- +goose StatementEnd
