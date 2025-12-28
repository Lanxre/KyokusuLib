-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    name VARCHAR(255),
    picture TEXT,
    password_hash TEXT,
    role VARCHAR(20) NOT NULL DEFAULT 'user',
    status VARCHAR(50) DEFAULT 'offline',
    last_login TIMESTAMP DEFAULT CURRENT_TIMESTAMP

    CONSTRAINT users_role_check CHECK (role IN ('user', 'admin', 'moderator'))
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS users;
-- +goose StatementEnd
