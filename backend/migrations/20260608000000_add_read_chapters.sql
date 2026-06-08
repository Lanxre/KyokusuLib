-- +goose Up
-- +goose StatementBegin
CREATE TABLE read_chapters (
    user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    chapter_id UUID NOT NULL REFERENCES novela_chapters(id) ON DELETE CASCADE,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id, chapter_id)
);

CREATE INDEX idx_read_chapters_user_id ON read_chapters(user_id);
CREATE INDEX idx_read_chapters_chapter_id ON read_chapters(chapter_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS read_chapters;
-- +goose StatementEnd
