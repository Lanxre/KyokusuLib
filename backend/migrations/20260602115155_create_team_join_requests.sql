-- +goose Up
-- +goose StatementBegin
CREATE TABLE team_join_requests (
    team_id INTEGER REFERENCES publisher_teams(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (team_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS team_join_requests;
-- +goose StatementEnd