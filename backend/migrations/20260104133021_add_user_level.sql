-- +goose Up
-- +goose StatementBegin
ALTER TABLE user_profiles
ADD COLUMN level INTEGER NOT NULL DEFAULT 1,
ADD COLUMN experience BIGINT NOT NULL DEFAULT 0;

CREATE TABLE level_definitions (
    level INTEGER PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    total_xp_required BIGINT NOT NULL,
    UNIQUE (level)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE level_definitions;
ALTER TABLE user_profiles
DROP COLUMN level,
DROP COLUMN experience;
-- +goose StatementEnd
