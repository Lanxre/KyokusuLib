-- +goose Up
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS isPublic BOOLEAN DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN IF EXISTS isPublic;
-- +goose StatementEnd
