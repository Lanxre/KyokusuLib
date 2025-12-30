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
	rows, err := r.DB.QueryContext(ctx, "SELECT id, name, country, metier, picture, bio FROM authors")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*db.Author
	for rows.Next() {
		var author db.Author
		if err := rows.Scan(&author.ID, &author.Name, &author.Country, &author.Metier, &author.Picture, &author.Bio); err != nil {
			return nil, err
		}
		authors = append(authors, &author)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return authors, nil
}

func (r *AuthorRepository) CreateAuthor(ctx context.Context, author *db.Author) error {
	query := `
		INSERT INTO authors (name, country, metier, picture, bio)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (name) DO UPDATE
		SET name = EXCLUDED.name`

	_, err := r.DB.Exec(query, author.Name, author.Country, author.Metier, author.Picture, author.Bio)

	if err != nil {
		return err
	}
	
	return nil
}

func (r *AuthorRepository) GetAuthorByName(ctx context.Context, name string) (*db.Author, error) {
	query := `
		SELECT id, name, country, metier, picture, bio
		FROM authors
		WHERE name = $1`

	var authorDb db.Author
	err := r.DB.QueryRowContext(ctx, query, name).Scan(
		&authorDb.ID, 
		&authorDb.Name, 
		&authorDb.Country, 
		&authorDb.Metier, 
		&authorDb.Picture, 
		&authorDb.Bio,
	)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	return &authorDb, nil
}
