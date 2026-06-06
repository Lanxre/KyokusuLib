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

func (r *AuthorRepository) GetAuthors(ctx context.Context, search string, limit int, offset int) ([]*db.Author, error) {
	var query string
	var rows *sql.Rows
	var err error

	var country sql.NullString
	var metier sql.NullString
	var picture sql.NullString
	var bio sql.NullString

	if search != "" {
		query = "SELECT id, name, country, metier, picture, bio FROM authors WHERE name ILIKE $1 ORDER BY id DESC LIMIT $2 OFFSET $3"
		rows, err = r.DB.QueryContext(ctx, query, "%"+search+"%", limit, offset)
	} else {
		query = "SELECT id, name, country, metier, picture, bio FROM authors ORDER BY id DESC LIMIT $1 OFFSET $2"
		rows, err = r.DB.QueryContext(ctx, query, limit, offset)
	}
	
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var authors []*db.Author
	for rows.Next() {
		var author db.Author
		if err := rows.Scan(&author.ID, &author.Name, &country, &metier, &picture, &bio); err != nil {
			return nil, err
		}

		if country.Valid {
			author.Country = country.String
		}
		if metier.Valid {
			author.Metier = metier.String
		}
		if picture.Valid {
			author.Picture = picture.String
		}
		if bio.Valid {
			author.Bio = bio.String
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
	var country sql.NullString
	var metier sql.NullString
	var picture sql.NullString
	var bio sql.NullString
	
	query := `
		SELECT id, name, country, metier, picture, bio
		FROM authors
		WHERE name = $1`

	var authorDb db.Author
	err := r.DB.QueryRowContext(ctx, query, name).Scan(
		&authorDb.ID, 
		&authorDb.Name, 
		&country, 
		&metier, 
		&picture, 
		&bio,
	)
	
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if country.Valid {
		authorDb.Country = country.String
	}
	if metier.Valid {
		authorDb.Metier = metier.String
	}
	if picture.Valid {
		authorDb.Picture = picture.String
	}
	if bio.Valid {
		authorDb.Bio = bio.String
	}

	return &authorDb, nil
}

func (r *AuthorRepository) GetAuthorById(ctx context.Context, id string) (*db.Author, error) {
	var country sql.NullString
	var metier sql.NullString
	var picture sql.NullString
	var bio sql.NullString
	
	query := `
		SELECT id, name, country, metier, picture, bio
		FROM authors
		WHERE id = $1`

	var authorDb db.Author
	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&authorDb.ID, 
		&authorDb.Name, 
		&country, 
		&metier, 
		&picture, 
		&bio,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}

	if country.Valid {
		authorDb.Country = country.String
	}
	if metier.Valid {
		authorDb.Metier = metier.String
	}
	if picture.Valid {
		authorDb.Picture = picture.String
	}
	if bio.Valid {
		authorDb.Bio = bio.String
	}

	return &authorDb, nil
}

func (r *AuthorRepository) UpdateAuthor(ctx context.Context, id string, author *db.Author) error {
	query := `
		UPDATE authors
		SET name = $1, country = $2, metier = $3, picture = COALESCE(NULLIF($4, ''), picture), bio = $5
		WHERE id = $6`

	res, err := r.DB.ExecContext(ctx, query, author.Name, author.Country, author.Metier, author.Picture, author.Bio, id)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
