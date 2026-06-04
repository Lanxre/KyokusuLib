-- +goose Up
-- +goose StatementBegin
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- For novela_volumes
ALTER TABLE novela_volumes ADD COLUMN uuid_id UUID DEFAULT uuid_generate_v4();
ALTER TABLE novela_chapters ADD COLUMN new_volume_id UUID;
UPDATE novela_chapters SET new_volume_id = (SELECT uuid_id FROM novela_volumes WHERE id = novela_chapters.novela_volume_id);

ALTER TABLE novela_chapters DROP CONSTRAINT IF EXISTS ranobe_chapters_ranobe_volume_id_fkey;
ALTER TABLE novela_volumes DROP CONSTRAINT IF EXISTS ranobe_volumes_pkey CASCADE;

ALTER TABLE novela_volumes DROP COLUMN id;
ALTER TABLE novela_volumes RENAME COLUMN uuid_id TO id;
ALTER TABLE novela_volumes ADD PRIMARY KEY (id);

ALTER TABLE novela_chapters DROP COLUMN novela_volume_id;
ALTER TABLE novela_chapters RENAME COLUMN new_volume_id TO novela_volume_id;
ALTER TABLE novela_chapters ADD CONSTRAINT fk_volume FOREIGN KEY (novela_volume_id) REFERENCES novela_volumes(id) ON DELETE CASCADE;

-- For novela_chapters
ALTER TABLE novela_chapters ADD COLUMN uuid_id UUID DEFAULT uuid_generate_v4();
ALTER TABLE novela_chapter_images ADD COLUMN new_chapter_id UUID;
UPDATE novela_chapter_images SET new_chapter_id = (SELECT uuid_id FROM novela_chapters WHERE id = novela_chapter_images.chapter_id);

ALTER TABLE novela_chapter_images DROP CONSTRAINT IF EXISTS ranobe_chapter_images_chapter_id_fkey;
ALTER TABLE novela_chapters DROP CONSTRAINT IF EXISTS ranobe_chapters_pkey CASCADE;

ALTER TABLE novela_chapters DROP COLUMN id;
ALTER TABLE novela_chapters RENAME COLUMN uuid_id TO id;
ALTER TABLE novela_chapters ADD PRIMARY KEY (id);

ALTER TABLE novela_chapter_images DROP COLUMN chapter_id;
ALTER TABLE novela_chapter_images RENAME COLUMN new_chapter_id TO chapter_id;
ALTER TABLE novela_chapter_images ADD CONSTRAINT fk_chapter FOREIGN KEY (chapter_id) REFERENCES novela_chapters(id) ON DELETE CASCADE;

-- Re-add unique constraint for chapters
ALTER TABLE novela_chapters ADD CONSTRAINT novela_chapters_volume_id_chapter_number_key UNIQUE (novela_volume_id, chapter_number);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Down migration is complex due to UUID to INT conversion, skipping for now as per common practice in this project unless requested.
-- +goose StatementEnd
