package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type NovelaBookmarkRepository struct {
	DB *sql.DB
}

func NewNovelaBookmarkRepository(db *sql.DB) *NovelaBookmarkRepository {
	return &NovelaBookmarkRepository{DB: db}
}

func (r *NovelaBookmarkRepository) SetBookmark(ctx context.Context, bookmark *db.Bookmark) error {
	query := `
		INSERT INTO user_novela_bookmarks (user_id, novela_id, category, updated_at) 
		VALUES ($1, $2, $3, CURRENT_TIMESTAMP) 
		ON CONFLICT (user_id, novela_id) 
		DO UPDATE SET 
			category = EXCLUDED.category, 
			updated_at = CURRENT_TIMESTAMP`

	_, err := r.DB.ExecContext(ctx, query, bookmark.UserID, bookmark.NovelaID, bookmark.Category)
	return err
}

func (r *NovelaBookmarkRepository) RemoveBookmark(ctx context.Context, userID, novelaID int) error {
	query := `DELETE FROM user_novela_bookmarks WHERE user_id = $1 AND novela_id = $2`
	_, err := r.DB.ExecContext(ctx, query, userID, novelaID)
	return err
}