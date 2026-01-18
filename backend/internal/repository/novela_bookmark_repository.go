package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

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

func (r *NovelaBookmarkRepository) GetBookmarkStats(tx *sql.Tx, ctx context.Context, novelaID int) (*db.NovelaBookmarkSummary, error) {
	// COALESCE(jsonb_object_agg(...), '{}') гарантирует, что мы не получим null в Go
	query := `
		SELECT 
			COALESCE(jsonb_object_agg(category, count), '{}'::jsonb)
		FROM (
			SELECT category, COUNT(*) as count 
			FROM user_novela_bookmarks 
			WHERE novela_id = $1 
			GROUP BY category
		) s`

	var summary db.NovelaBookmarkSummary
	var distJSON []byte

	err := tx.QueryRowContext(ctx, query, novelaID).Scan(
		&distJSON,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &db.NovelaBookmarkSummary{Distribution: make(map[string]int)}, nil
		}
		return nil, err
	}
	
	summary.Distribution = make(map[string]int)
	if len(distJSON) > 0 {
		json.Unmarshal(distJSON, &summary.Distribution)
	}

	for cat := range summary.Distribution {
		summary.TotalCount += summary.Distribution[cat]
	}

	return &summary, nil
}