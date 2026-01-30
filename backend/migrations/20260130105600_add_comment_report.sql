-- +goose Up
-- +goose StatementBegin
CREATE TABLE novela_comments_reports (
    comment_id INT NOT NULL REFERENCES novela_comments(id) ON DELETE CASCADE,
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    reason VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (comment_id, user_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS novela_comments_reports;
-- +goose StatementEnd
