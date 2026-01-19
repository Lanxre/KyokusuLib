package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type NovelaRatingRepository struct {
	DB *sql.DB
}

func NewNovelaRatingRepository(db *sql.DB) *NovelaRatingRepository {
	return &NovelaRatingRepository{DB: db}
}

func (r *NovelaRatingRepository) SetRating(ctx context.Context, rating *db.NovelaRating) error {
	query := `
		INSERT INTO novela_ratings (user_id, novela_id, rating)
		VALUES ($1, $2, $3)
		ON CONFLICT (user_id, novela_id)
		DO UPDATE SET
			rating = EXCLUDED.rating`
	_, err := r.DB.ExecContext(ctx, query, rating.UserID, rating.NovelaID, rating.Rating)
	return err
}

func (r *NovelaRatingRepository) GetRating(tx *sql.Tx, ctx context.Context, novelaID int) (*db.NovelaRatingSummary, error) {
	query := `
			SELECT
		    SUM(count) AS total_count,
		    COALESCE(SUM(q)::numeric / NULLIF(SUM(count), 0), 0) AS avg_rating,
		    jsonb_object_agg(rating, count) AS distribution
		FROM (
		    SELECT
		        rating,
		        COUNT(*) AS count,
		        (COUNT(*) * rating) AS q
		    FROM novela_ratings
		    WHERE novela_id = $1
		    GROUP BY rating
		) AS s;`

	var summary db.NovelaRatingSummary
	var distJSON []byte

	err := tx.QueryRowContext(ctx, query, novelaID).Scan(
		&summary.TotalCount,
		&summary.AverageRating,
		&distJSON,
	)

	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	if len(distJSON) > 0 {
		json.Unmarshal(distJSON, &summary.Distribution)
	}

	return &summary, nil
}
