-- +goose Up
-- +goose StatementBegin
ALTER TABLE novela_volumes ADD COLUMN status VARCHAR(50) DEFAULT 'approved';
ALTER TABLE novela_volumes ADD COLUMN created_by INTEGER REFERENCES users(id);

ALTER TABLE novela_chapters ADD COLUMN status VARCHAR(50) DEFAULT 'approved';
ALTER TABLE novela_chapters ADD COLUMN created_by INTEGER REFERENCES users(id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE novela_chapters DROP COLUMN created_by;
ALTER TABLE novela_chapters DROP COLUMN status;

ALTER TABLE novela_volumes DROP COLUMN created_by;
ALTER TABLE novela_volumes DROP COLUMN status;
-- +goose StatementEnd
