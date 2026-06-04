-- +goose Up
ALTER TABLE novela_chapter_images ADD COLUMN position INTEGER DEFAULT 0;

-- +goose Down
ALTER TABLE novela_chapter_images DROP COLUMN position;
