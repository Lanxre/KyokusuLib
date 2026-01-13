-- +goose Up
-- +goose StatementBegin
CREATE TABLE user_novela_likes (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    novela_id INT NOT NULL REFERENCES novela(id) ON DELETE CASCADE,
    has_liked BOOLEAN NOT NULL DEFAULT FALSE,

    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, novela_id)
);

CREATE INDEX idx_user_novela_likes_novela_id ON user_novela_likes(novela_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_novela_likes;
-- +goose StatementEnd
