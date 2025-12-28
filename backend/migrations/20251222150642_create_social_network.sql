-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_socials (
    user_id INTEGER PRIMARY KEY,

    discord_id TYPE NUMERIC USING discord_id::TEXT::NUMERIC,
    is_discord_connected BOOLEAN DEFAULT FALSE,
    
    google_id TYPE NUMERIC USING discord_id::TEXT::NUMERIC,
    is_google_connected BOOLEAN DEFAULT FALSE,
    
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_socials;
-- +goose StatementEnd
