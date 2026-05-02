-- +goose Up
-- +goose StatementBegin

CREATE TABLE bookmark_categories (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE, -- NULL means default for everyone
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO bookmark_categories (user_id, name) VALUES 
    (NULL, 'planned'),
    (NULL, 'reading'),
    (NULL, 'completed'),
    (NULL, 'on_hold'),
    (NULL, 'dropped');

ALTER TABLE user_novela_bookmarks ADD COLUMN category_id INT REFERENCES bookmark_categories(id) ON DELETE CASCADE;

UPDATE user_novela_bookmarks b
SET category_id = c.id
FROM bookmark_categories c
WHERE b.category::text = c.name AND c.user_id IS NULL;

ALTER TABLE user_novela_bookmarks DROP COLUMN category;
DROP TYPE novela_bookmark_category;

CREATE UNIQUE INDEX idx_user_bookmark_categories_name ON bookmark_categories (user_id, name) WHERE user_id IS NOT NULL;
CREATE UNIQUE INDEX idx_default_bookmark_categories_name ON bookmark_categories (name) WHERE user_id IS NULL;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

CREATE TYPE novela_bookmark_category AS ENUM (
    'planned',
    'reading',
    'completed',
    'on_hold',
    'dropped'
);

ALTER TABLE user_novela_bookmarks ADD COLUMN category novela_bookmark_category DEFAULT 'planned';

UPDATE user_novela_bookmarks b
SET category = c.name::novela_bookmark_category
FROM bookmark_categories c
WHERE b.category_id = c.id;

ALTER TABLE user_novela_bookmarks DROP COLUMN category_id;
DROP TABLE bookmark_categories;

-- +goose StatementEnd
