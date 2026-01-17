-- +goose Up
-- +goose StatementBegin
CREATE TABLE publisher_teams (
    id SERIAL PRIMARY KEY,
    owner_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(255) NOT NULL UNIQUE,
    slug VARCHAR(255) NOT NULL UNIQUE,
    description TEXT,
    avatar_url TEXT,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    updated_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE TABLE team_members (
    team_id INTEGER REFERENCES publisher_teams(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    joined_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE team_subscribers (
    team_id INTEGER REFERENCES publisher_teams(id) ON DELETE CASCADE,
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (team_id, user_id)
);

CREATE TABLE novela_teams (
    novela_id INTEGER REFERENCES novela(id) ON DELETE CASCADE,
    team_id INTEGER REFERENCES publisher_teams(id) ON DELETE CASCADE,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (novela_id, team_id)
);

CREATE INDEX idx_publisher_teams_slug ON publisher_teams(slug);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS novela_teams;
DROP TABLE IF EXISTS team_subscribers;
DROP TABLE IF EXISTS team_members;
DROP TABLE IF EXISTS publisher_teams;
-- +goose StatementEnd