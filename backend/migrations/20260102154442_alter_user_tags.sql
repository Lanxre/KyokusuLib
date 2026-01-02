-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_tags
    ALTER COLUMN tag DROP NOT NULL,
    ALTER COLUMN tag DROP DEFAULT,
    ADD CONSTRAINT unique_tag UNIQUE (tag);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_tags
    ALTER COLUMN tag SET NOT NULL,
    ALTER COLUMN tag SET DEFAULT '',
    DROP CONSTRAINT unique_tag;

-- +goose StatementEnd
