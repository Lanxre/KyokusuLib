-- +goose Up
-- +goose StatementBegin
ALTER TABLE novela_chapters
    ADD COLUMN created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE novela_chapters
    DROP COLUMN created_at;
-- +goose StatementEnd
