-- +goose Up
ALTER TABLE user_profile_settings 
    ADD COLUMN font_size INT DEFAULT 18,
    ADD COLUMN line_weight FLOAT DEFAULT 1.6,
    ADD COLUMN font_family VARCHAR(255) DEFAULT 'font-sans',
    ADD COLUMN auto_scroll BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE user_profile_settings
    DROP COLUMN font_size,
    DROP COLUMN line_weight,
    DROP COLUMN font_family,
    DROP COLUMN auto_scroll;
