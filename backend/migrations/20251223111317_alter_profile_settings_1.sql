-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_profile_settings
    ADD COLUMN is_app_notify BOOLEAN DEFAULT TRUE,
    ADD COLUMN is_new_published_notify BOOLEAN DEFAULT TRUE;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user_profile_settings
    DROP COLUMN is_app_notify,
    DROP COLUMN is_app_notify;
-- +goose StatementEnd
