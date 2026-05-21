-- +goose Up
-- +goose StatementBegin
ALTER TABLE publisher_teams ADD COLUMN banner_url VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE publisher_teams DROP COLUMN banner_url;
-- +goose StatementEnd
