package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type AuthorRepository struct {
	DB *sql.DB
}

func NewAuthorRepository(db *sql.DB) *AuthorRepository {
	return &AuthorRepository{DB: db}
}

func (r *AuthorRepository) GetAuthors(ctx context.Context) ([]*db.Author, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, name FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*db.Author
	for rows.Next() {
		var author db.Author
		if err := rows.Scan(&author.ID, &author.Name); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, name string) error {
	query := `
		INSERT INTO authors (name)
		VALUES ($1)
		ON CONFLICT (name) DO UPDATE
		SET name = EXCLUDED.name`

	_, err := r.DB.Exec(query, name)

	if err != nil {
		return err
	}
	
	return nil
}

func (r *AuthorRepository) GetAuthorByName(ctx context.Context, name string) (*db.Author, error) {
	query := `
		SELECT id, name
		FROM authors
		WHERE name = $1`

	var author db.Author
	err := r.DB.QueryRowContext(ctx, query, name).Scan(&author.ID, &author.Name)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	return &author, nil
}
