-- +goose Up
-- +goose StatementBegin
DROP TABLE IF EXISTS user_catalog_filters;

CREATE TABLE user_catalog_filters (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    name TEXT NOT NULL DEFAULT '',
    filters JSONB NOT NULL DEFAULT '{}',
    created_at TIMESTAMPTZ DEFAULT NOW()
);

CREATE INDEX idx_user_catalog_filters_user_id ON user_catalog_filters(user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_catalog_filters;

CREATE TABLE user_catalog_filters (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    filters JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd
