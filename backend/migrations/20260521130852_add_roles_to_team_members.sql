-- +goose Up
-- +goose StatementBegin
ALTER TABLE team_members ADD COLUMN role VARCHAR(50) DEFAULT 'member';
ALTER TABLE team_members ADD COLUMN custom_role_name VARCHAR(100);

UPDATE team_members
SET role = 'owner'
WHERE user_id = (SELECT owner_id FROM publisher_teams WHERE id = team_members.team_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE team_members DROP COLUMN custom_role_name;
ALTER TABLE team_members DROP COLUMN role;
-- +goose StatementEnd
