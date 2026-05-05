-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_profile_settings
ADD COLUMN IF NOT EXISTS is_show_bookmark BOOLEAN DEFAULT true;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_profile_settings
DROP COLUMN IF EXISTS is_show_bookmark;
-- +goose StatementEnd
