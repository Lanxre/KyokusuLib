-- +goose Up
-- +goose StatementBegin

CREATE TABLE genres (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) UNIQUE NOT NULL
);

CREATE TABLE authors (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE ranobe (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    release_date DATE,
    country VARCHAR(100),
    views INTEGER DEFAULT 0,
    translation_status VARCHAR(50)
);

CREATE TABLE ranobe_genres (
    ranobe_id INTEGER REFERENCES ranobe(id) ON DELETE CASCADE,
    genre_id INTEGER REFERENCES genres(id) ON DELETE CASCADE,
    PRIMARY KEY (ranobe_id, genre_id)
);

CREATE TABLE ranobe_authors (
    ranobe_id INTEGER REFERENCES ranobe(id) ON DELETE CASCADE,
    author_id INTEGER REFERENCES authors(id) ON DELETE CASCADE,
    PRIMARY KEY (ranobe_id, author_id)
);

CREATE TABLE ranobe_volumes (
    id SERIAL PRIMARY KEY,
    ranobe_id INTEGER REFERENCES ranobe(id) ON DELETE CASCADE,
    volume_number INTEGER NOT NULL,
    title VARCHAR(255),
    UNIQUE (ranobe_id, volume_number)
);

CREATE TABLE ranobe_chapters (
    id SERIAL PRIMARY KEY,
    ranobe_volume_id INTEGER REFERENCES ranobe_volumes(id) ON DELETE CASCADE,
    chapter_number INTEGER NOT NULL,
    title VARCHAR(255),
    content TEXT NOT NULL,
    UNIQUE (ranobe_volume_id, chapter_number)
);

CREATE TABLE ranobe_chapter_images (
    id SERIAL PRIMARY KEY,
    chapter_id INTEGER REFERENCES ranobe_chapters(id) ON DELETE CASCADE,
    image_url VARCHAR(255) NOT NULL,
    caption TEXT
);

CREATE TABLE ranobe_ratings (
    ranobe_id INTEGER REFERENCES ranobe(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    rating INTEGER NOT NULL CHECK (rating >= 1 AND rating <= 10),
    created_at TIMESTAMPTZ DEFAULT NOW(),
    PRIMARY KEY (ranobe_id, user_id) 
);

CREATE INDEX idx_ranobe_ratings_ranobe_id ON ranobe_ratings(ranobe_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS ranobe_ratings;
DROP TABLE IF EXISTS ranobe_chapter_images;
DROP TABLE IF EXISTS ranobe_chapters;
DROP TABLE IF EXISTS ranobe_volumes;
DROP TABLE IF EXISTS ranobe_genres;
DROP TABLE IF EXISTS ranobe;
DROP TABLE IF EXISTS genres;
DROP TABLE IF EXISTS authors;

-- +goose StatementEnd