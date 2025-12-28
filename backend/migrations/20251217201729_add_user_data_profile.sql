-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
    CREATE TYPE gender_type AS ENUM ('male', 'female', 'hidden');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

ALTER TABLE users
    ADD COLUMN IF NOT EXISTS about TEXT,
    ADD COLUMN IF NOT EXISTS birthday TIMESTAMP,
    ADD COLUMN IF NOT EXISTS gender gender_type DEFAULT 'hidden',
    ADD COLUMN IF NOT EXISTS updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    DROP COLUMN IF EXISTS about,
    DROP COLUMN IF EXISTS birthday,
    DROP COLUMN IF EXISTS gender,
    DROP COLUMN IF EXISTS updated_at;

DROP TYPE IF EXISTS gender_type;
-- +goose StatementEnd