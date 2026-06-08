-- +goose Up
-- +goose StatementBegin
ALTER TABLE read_chapters 
    ADD COLUMN scroll_position INT NOT NULL DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE read_chapters 
    DROP COLUMN IF EXISTS scroll_position;
-- +goose StatementEnd
