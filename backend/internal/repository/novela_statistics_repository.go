package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/db"
)

type NovelaStatisticsRepository struct {
	DB *sql.DB
}

func NewNovelaStatisticsRepository(db *sql.DB) *NovelaStatisticsRepository {
	return &NovelaStatisticsRepository{DB: db}
}

func (r *NovelaStatisticsRepository) GetTotalStatistics(
	ctx context.Context,
	limit int,
	offset int,
	period constants.StatisticsPeriodSort,
) ([]db.TotalNovelaStatistics, error) {
	if limit <= 0 {
		limit = 10
	}
	if limit > 100 {
		limit = 100
	}
	if offset < 0 {
		offset = 0
	}

	startTime := r.startOfPeriod(period)

	query := `
		WITH bookmark_stats AS (
			SELECT
				novela_id,
				COUNT(*) AS total_bookmark_count
			FROM user_novela_bookmarks
			WHERE ($2::TIMESTAMPTZ IS NULL OR created_at >= $2)
			GROUP BY novela_id
		),
		rating_stats AS (
			SELECT
				novela_id,
				COUNT(*) AS total_rating_count,
				AVG(rating)::float8 AS rating
			FROM novela_ratings
			WHERE ($2::TIMESTAMPTZ IS NULL OR created_at >= $2)
			GROUP BY novela_id
		),
		comment_stats AS (
			SELECT
				novela_id,
				COUNT(*) AS total_comment_count
			FROM novela_comments
			WHERE ($2::TIMESTAMPTZ IS NULL OR created_at >= $2)
			GROUP BY novela_id
		),
		volume_stats AS (
			SELECT
				novela_id,
				COUNT(*) AS total_volume_count
			FROM novela_volumes
			GROUP BY novela_id
		),
		chapter_stats AS (
			SELECT
				nv.novela_id,
				COUNT(*) AS total_chapter_count
			FROM novela_volumes nv
			JOIN novela_chapters nch
				ON nch.novela_volume_id = nv.id
			GROUP BY nv.novela_id
		),
		read_stats AS (
			SELECT
				nv.novela_id,
				COUNT(DISTINCT rc.user_id) AS total_read_count
			FROM read_chapters rc
			JOIN novela_chapters nch
				ON nch.id = rc.chapter_id
			JOIN novela_volumes nv
				ON nv.id = nch.novela_volume_id
			WHERE ($2::TIMESTAMPTZ IS NULL OR rc.created_at >= $2)
			GROUP BY nv.novela_id
		)
		SELECT
			n.id,
			n.title,
			n.description,
			n.poster_url,
			n.release_date,
		
			COALESCE(bs.total_bookmark_count, 0) AS total_bookmark_count,
			COALESCE(rs_read.total_read_count, 0) AS total_read_count,
			COALESCE(rs_rate.total_rating_count, 0) AS total_rating_count,
			COALESCE(cs.total_comment_count, 0) AS total_comment_count,
			COALESCE(vs.total_volume_count, 0) AS total_volume_count,
			COALESCE(chs.total_chapter_count, 0) AS total_chapter_count,
			COALESCE(rs_rate.rating, 0.0) AS rating
		
		FROM novela n
		
		LEFT JOIN bookmark_stats bs
			ON bs.novela_id = n.id
		
		LEFT JOIN read_stats rs_read
			ON rs_read.novela_id = n.id
		
		LEFT JOIN rating_stats rs_rate
			ON rs_rate.novela_id = n.id
		
		LEFT JOIN comment_stats cs
			ON cs.novela_id = n.id
		
		LEFT JOIN volume_stats vs
			ON vs.novela_id = n.id
		
		LEFT JOIN chapter_stats chs
			ON chs.novela_id = n.id
		
		ORDER BY
			rating DESC,
			total_read_count DESC,
			total_bookmark_count DESC,
			total_comment_count DESC,
			total_volume_count DESC,
			total_chapter_count DESC,
			n.id DESC
		LIMIT $1
		OFFSET $3
		`

	rows, err := r.DB.QueryContext(ctx, query, limit, startTime, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	stats := make([]db.TotalNovelaStatistics, 0, limit)

	for rows.Next() {
		var s db.TotalNovelaStatistics
		var releaseDate sql.NullTime

		if err := rows.Scan(
			&s.Novela.NovelaID,
			&s.Novela.Title,
			&s.Novela.Description,
			&s.Novela.PosterURL,
			&releaseDate,
			&s.TotalBookmarkCount,
			&s.TotalReadCount,
			&s.TotalRatingCount,
			&s.TotalCommentCount,
			&s.TotalVolumeCount,
			&s.TotalChapterCount,
			&s.Rating,
		); err != nil {
			return nil, err
		}

		if releaseDate.Valid {
			s.Novela.ReleaseDate = releaseDate.Time
		}

		stats = append(stats, s)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return stats, nil
}

func (r *NovelaStatisticsRepository) startOfPeriod(period constants.StatisticsPeriodSort) *time.Time {
	now := time.Now()
	var day int

	switch period {
		case constants.Day:
			day = now.Day()
		case constants.Week:
			weekday := int(now.Weekday())
			if weekday == 0 {
				weekday = 7
			}
			day = now.Day() - (weekday - 1)
		case constants.Month:
			day = 1
		default:
			return nil
	}

	t := time.Date(now.Year(), now.Month(), day, 0, 0, 0, 0, now.Location())
	return &t
}
