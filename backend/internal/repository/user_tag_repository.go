package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lib/pq"
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

func (r *UserTagRepository) UpdateUserTags(ctx context.Context, userID int, tagIds []int) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	if _, err := tx.ExecContext(ctx, `DELETE FROM users_user_tags WHERE user_id = $1`, userID); err != nil {
		return err
	}

	for _, tagID := range tagIds {
		if _, err := tx.ExecContext(ctx, `INSERT INTO users_user_tags (user_id, tag_id) VALUES ($1, $2)`, userID, tagID); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *UserTagRepository) IsExist(ctx context.Context, tagIDs []int) (bool, error) {
	query := `SELECT EXISTS (SELECT 1 FROM user_tags WHERE id = ANY($1))`

    var exists bool
    err := r.DB.QueryRowContext(ctx, query, pq.Array(tagIDs)).Scan(&exists)
    
    if err != nil {
        return false, err
    }

    return exists, nil
}