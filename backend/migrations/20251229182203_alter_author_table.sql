-- +goose Up
-- +goose StatementBegin
ALTER TABLE authors 
    ADD COLUMN country VARCHAR(30),
    ADD COLUMN metier VARCHAR(30),
    ADD COLUMN picture TEXT,
    ADD COLUMN bio TEXT;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE authors 
    DROP COLUMN country,
    DROP COLUMN metier,
    DROP COLUMN picture,
    DROP COLUMN bio;
-- +goose StatementEnd
