-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_catalog_filters (
    user_id INTEGER REFERENCES users(id) ON DELETE CASCADE,
    filters JSONB,
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_catalog_filters;
-- +goose StatementEnd