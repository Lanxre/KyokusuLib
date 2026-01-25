-- +goose Up
-- +goose StatementBegin
CREATE TABLE novela_comments (
    id SERIAL PRIMARY KEY,
    novela_id INT NOT NULL REFERENCES novela(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    parent_id INT REFERENCES novela_comments(id) ON DELETE CASCADE,
    content TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_comments_novela_id ON novela_comments(novela_id);
CREATE INDEX idx_comments_parent_id ON novela_comments(parent_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS novela_comments;
-- +goose StatementEnd
