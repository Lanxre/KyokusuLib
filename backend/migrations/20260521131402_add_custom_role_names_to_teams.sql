-- +goose Up
-- +goose StatementBegin
ALTER TABLE publisher_teams ADD COLUMN owner_role_name VARCHAR(100) DEFAULT 'Владелец';
ALTER TABLE publisher_teams ADD COLUMN moderator_role_name VARCHAR(100) DEFAULT 'Модератор';
ALTER TABLE publisher_teams ADD COLUMN member_role_name VARCHAR(100) DEFAULT 'Участник';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE publisher_teams DROP COLUMN owner_role_name;
ALTER TABLE publisher_teams DROP COLUMN moderator_role_name;
ALTER TABLE publisher_teams DROP COLUMN member_role_name;
-- +goose StatementEnd
