-- +goose Up
-- +goose StatementBegin
ALTER TABLE bookmark_categories ADD COLUMN visible BOOLEAN NOT NULL DEFAULT TRUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE bookmark_categories DROP COLUMN visible;
-- +goose StatementEnd
