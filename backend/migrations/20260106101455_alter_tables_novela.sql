-- +goose Up
-- +goose StatementBegin

-- 1. Переименование таблиц
ALTER TABLE ranobe RENAME TO novela;
ALTER TABLE ranobe_genres RENAME TO novela_genres;
ALTER TABLE ranobe_authors RENAME TO novela_authors;
ALTER TABLE ranobe_volumes RENAME TO novela_volumes;
ALTER TABLE ranobe_chapters RENAME TO novela_chapters;
ALTER TABLE ranobe_chapter_images RENAME TO novela_chapter_images;
ALTER TABLE ranobe_ratings RENAME TO novela_ratings;

-- 2. Переименование колонок внешних ключей (чтобы код был чистым)
ALTER TABLE novela_genres RENAME COLUMN ranobe_id TO novela_id;
ALTER TABLE novela_authors RENAME COLUMN ranobe_id TO novela_id;
ALTER TABLE novela_volumes RENAME COLUMN ranobe_id TO novela_id;
ALTER TABLE novela_chapters RENAME COLUMN ranobe_volume_id TO novela_volume_id;
ALTER TABLE novela_ratings RENAME COLUMN ranobe_id TO novela_id;

-- 3. Создание таблицы Категорий (Categories)
CREATE TABLE categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

-- 4. Создание связующей таблицы для Категорий
CREATE TABLE novela_categories (
    novela_id INTEGER REFERENCES novela(id) ON DELETE CASCADE,
    category_id INTEGER REFERENCES categories(id) ON DELETE CASCADE,
    PRIMARY KEY (novela_id, category_id)
);

-- 5. Обновление таблицы связей авторов (Добавляем роль)
-- Это позволяет указать: 'Author', 'Artist', 'Original Creator' и т.д.
ALTER TABLE novela_authors
    ADD COLUMN role VARCHAR(50) DEFAULT 'author';

-- 6. Добавление новых колонок в таблицу novela
ALTER TABLE novela
    -- Массив альтернативных названий (PostgreSQL Array)
    ADD COLUMN alternative_titles TEXT[],
    -- Тип произведения (например: Ранобэ, Веб-новелла)
    ADD COLUMN type VARCHAR(50),
    -- Возрастной рейтинг (например: 16+, 18+)
    ADD COLUMN age_rating VARCHAR(20),
    -- Статус выхода оригинала (ongoing, completed, frozen)
    ADD COLUMN status VARCHAR(50) DEFAULT 'ongoing',
    -- Ссылка на обложку
    ADD COLUMN poster_url TEXT,
    -- Дата добавления и обновления
    ADD COLUMN created_at TIMESTAMPTZ DEFAULT NOW(),
    ADD COLUMN updated_at TIMESTAMPTZ DEFAULT NOW();

-- 7. Переименование индекса рейтинга (для порядка)
ALTER INDEX idx_ranobe_ratings_ranobe_id RENAME TO idx_novela_ratings_novela_id;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- 1. Удаление новых колонок из novela
ALTER TABLE novela 
    DROP COLUMN alternative_titles,
    DROP COLUMN type,
    DROP COLUMN age_rating,
    DROP COLUMN status,
    DROP COLUMN poster_url,
    DROP COLUMN created_at,
    DROP COLUMN updated_at;

-- 2. Удаление колонки роли у авторов
ALTER TABLE novela_authors 
    DROP COLUMN role;

-- 3. Удаление категорий
DROP TABLE IF EXISTS novela_categories;
DROP TABLE IF EXISTS categories;

-- 4. Возврат имен колонок
ALTER TABLE novela_ratings RENAME COLUMN novela_id TO ranobe_id;
ALTER TABLE novela_chapters RENAME COLUMN novela_volume_id TO ranobe_volume_id;
ALTER TABLE novela_volumes RENAME COLUMN novela_id TO ranobe_id;
ALTER TABLE novela_authors RENAME COLUMN novela_id TO ranobe_id;
ALTER TABLE novela_genres RENAME COLUMN novela_id TO ranobe_id;

-- 5. Возврат имен таблиц
ALTER TABLE novela_ratings RENAME TO ranobe_ratings;
ALTER TABLE novela_chapter_images RENAME TO ranobe_chapter_images;
ALTER TABLE novela_chapters RENAME TO ranobe_chapters;
ALTER TABLE novela_volumes RENAME TO ranobe_volumes;
ALTER TABLE novela_authors RENAME TO ranobe_authors;
ALTER TABLE novela_genres RENAME TO ranobe_genres;
ALTER TABLE novela RENAME TO ranobe;

-- 6. Возврат имени индекса
ALTER INDEX idx_novela_ratings_novela_id RENAME TO idx_ranobe_ratings_ranobe_id;
-- +goose StatementEnd