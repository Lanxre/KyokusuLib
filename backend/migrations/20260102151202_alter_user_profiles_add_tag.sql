-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_tags (
    id SERIAL PRIMARY KEY,
    tag VARCHAR(255) NOT NULL DEFAULT 'Читатель',
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE user_profiles
ADD COLUMN tag_id INTEGER REFERENCES user_tags(id);

ALTER TABLE user_profile_settings
ADD COLUMN is_show_tag BOOLEAN NOT NULL DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_tags;
ALTER TABLE user_profiles DROP COLUMN tag_id;

ALTER TABLE user_profile_settings DROP COLUMN is_show_tag;
-- +goose StatementEnd
