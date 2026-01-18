package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type NovelaLikeRepository struct {
	DB *sql.DB
}

func NewNovelaLikeRepository(db *sql.DB) *NovelaLikeRepository {
	return &NovelaLikeRepository{DB: db}
}

func (r *NovelaLikeRepository) SetLike(ctx context.Context, novelaLike *db.NovelaLike) error {
	query := `
		INSERT INTO user_novela_likes (user_id, novela_id, has_liked, updated_at) 
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP)
		ON CONFLICT (user_id, novela_id) 
		DO UPDATE SET
			has_liked = EXCLUDED.has_liked,
			updated_at = CURRENT_TIMESTAMP`
	_, err := r.DB.ExecContext(ctx, query, novelaLike.UserID, novelaLike.NovelaID, novelaLike.HasLiked)
	return err
}