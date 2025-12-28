-- +goose Up
-- +goose StatementBegin
DO $$ BEGIN
    CREATE TYPE gender_type AS ENUM ('male', 'female', 'hidden');
EXCEPTION
    WHEN duplicate_object THEN NULL;
END $$;

CREATE TABLE IF NOT EXISTS user_profiles (
    user_id INTEGER PRIMARY KEY,
    name VARCHAR(250),
    picture TEXT,
    about TEXT,
    birthday DATE,
    gender gender_type DEFAULT 'hidden',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS user_profile_settings (
    user_id INTEGER PRIMARY KEY,
    theme VARCHAR(10) DEFAULT 'dark',
    FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

INSERT INTO user_profiles (user_id, name, picture, about, birthday, gender)
SELECT id, name, picture, about, birthday::DATE, gender
FROM users
ON CONFLICT (user_id) DO NOTHING;

ALTER TABLE users
    DROP COLUMN IF EXISTS name,
    DROP COLUMN IF EXISTS picture,
    DROP COLUMN IF EXISTS about,
    DROP COLUMN IF EXISTS birthday,
    DROP COLUMN IF EXISTS gender;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE users
    ADD COLUMN IF NOT EXISTS name VARCHAR(250),
    ADD COLUMN IF NOT EXISTS picture TEXT,
    ADD COLUMN IF NOT EXISTS about TEXT,
    ADD COLUMN IF NOT EXISTS birthday DATE,
    ADD COLUMN IF NOT EXISTS gender gender_type DEFAULT 'hidden';

UPDATE users u
SET 
    name = p.name,
    picture = p.picture,
    about = p.about,
    birthday = p.birthday,
    gender = p.gender
FROM user_profiles p
WHERE u.id = p.user_id;

DROP TABLE IF EXISTS user_profile_settings;
DROP TABLE IF EXISTS user_profiles;
-- +goose StatementEnd