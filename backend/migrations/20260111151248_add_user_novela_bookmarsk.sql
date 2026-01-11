-- +goose Up
-- +goose StatementBegin
CREATE TYPE novela_bookmark_category AS ENUM (
    'planned',
    'reading',
    'completed',
    'on_hold',
    'dropped'
);

CREATE TABLE user_novela_bookmarks (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    novela_id INT NOT NULL REFERENCES novela(id) ON DELETE CASCADE,
    category novela_bookmark_category NOT NULL DEFAULT 'planned', 

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, novela_id)
);

CREATE INDEX idx_user_novela_bookmarks_novela_id ON user_novela_bookmarks(novela_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_novela_bookmarks;
DROP TYPE IF EXISTS novela_bookmark_category;
-- +goose StatementEnd
