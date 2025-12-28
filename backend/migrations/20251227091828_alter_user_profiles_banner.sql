-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_profiles
    ADD COLUMN banner TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_profiles
    DROP COLUMN banner;
-- +goose StatementEnd
