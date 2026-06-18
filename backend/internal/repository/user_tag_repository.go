package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserTagRepository struct {
	DB *sql.DB
}

func NewUserTagRepository(db *sql.DB) *UserTagRepository {
	return &UserTagRepository{DB: db}
}


func (r *UserTagRepository) GetUserTags(ctx context.Context) ([]db.UserDbTag, error) {
	query := "select id, tag, created_at from user_tags"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []db.UserDbTag
	for rows.Next() {
		var tag db.UserDbTag
		if err := rows.Scan(&tag.ID, &tag.Tag, &tag.CreatedAt); err != nil {
			return nil, err
		}
		tags = append(tags, tag)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tags, nil
}

func (r *UserTagRepository) GetUserTagById(ctx context.Context, id int) (*db.UserDbTag, error) {
	query := "select id, tag, created_at from user_tags where id = $1"
	row := r.DB.QueryRow(query, id)

	var tag db.UserDbTag
	if err := row.Scan(&tag.ID, &tag.Tag, &tag.CreatedAt); err != nil {
		return nil, err
	}

	return &tag, nil
}