-- +goose Up
-- +goose StatementBegin
ALTER TABLE publisher_teams ADD COLUMN team_type VARCHAR(20) NOT NULL DEFAULT 'open';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE publisher_teams DROP COLUMN team_type;
-- +goose StatementEnd
