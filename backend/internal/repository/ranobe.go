package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type RanobeRepository struct {
	DB *sql.DB
}

func NewRanobeRepository(db *sql.DB) *RanobeRepository {
	return &RanobeRepository{DB: db}
}

func (r *RanobeRepository) Create(ctx context.Context, ranobe *db.Ranobe) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO ranobe (title, description, release_date, country, translation_status, views)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id`

	err = tx.QueryRowContext(ctx, query,
		ranobe.Title,
		ranobe.Description,
		ranobe.Year,
		ranobe.Country,
		ranobe.TranslationStatus,
		ranobe.Views,
	).Scan(&ranobe.ID)

	if err != nil {
		return err
	}

	if len(ranobe.Authors) > 0 {
		if err := r.linkAuthors(ctx, tx, ranobe.ID, ranobe.Authors); err != nil {
			return err
		}
	}

	if len(ranobe.Genres) > 0 {
		if err := r.linkGenres(ctx, tx, ranobe.ID, ranobe.Genres); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *RanobeRepository) GetByID(ctx context.Context, id int) (*db.Ranobe, error) {
	ranobe := &db.Ranobe{}

	query := `
		SELECT 
			r.id, r.title, r.description, r.translation_status, r.release_date, r.views,
			COALESCE(AVG(rt.rating), 0) as rating,
			COALESCE(array_remove(array_agg(DISTINCT a.name), NULL), '{}') as authors,
			COALESCE(array_remove(array_agg(DISTINCT g.name), NULL), '{}') as genres
		FROM ranobe r
		LEFT JOIN ranobe_ratings rt ON r.id = rt.ranobe_id
		LEFT JOIN ranobe_authors ra ON r.id = ra.ranobe_id
		LEFT JOIN authors a ON ra.author_id = a.id
		LEFT JOIN ranobe_genres rg ON r.id = rg.ranobe_id
		LEFT JOIN genres g ON rg.genre_id = g.id
		WHERE r.id = $1
		GROUP BY r.id`

	err := r.DB.QueryRowContext(ctx, query, id).Scan(
		&ranobe.ID,
		&ranobe.Title,
		&ranobe.Description,
		&ranobe.TranslationStatus,
		&ranobe.Year,
		&ranobe.Views,
		&ranobe.Rating,
		&ranobe.Authors,
		&ranobe.Genres,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return ranobe, nil
}

func (r *RanobeRepository) GetFullByID(ctx context.Context, id int) (*db.Ranobe, error) {
	ranobe, err := r.GetByID(ctx, id)
	if err != nil || ranobe == nil {
		return ranobe, err
	}

	volumesQuery := `SELECT id, title, volume_number FROM ranobe_volumes WHERE ranobe_id = $1 ORDER BY volume_number ASC`
	rows, err := r.DB.QueryContext(ctx, volumesQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var volumes []db.RanobeVolume
	for rows.Next() {
		var v db.RanobeVolume
		if err := rows.Scan(&v.ID, &v.Title, &v.Number); err != nil {
			return nil, err
		}
		
		chapters, err := r.getChaptersByVolumeID(ctx, v.ID)
		if err != nil {
			return nil, err
		}
		v.Chapters = chapters
		volumes = append(volumes, v)
	}

	ranobe.Volumes = volumes
	return ranobe, nil
}

func (r *RanobeRepository) linkAuthors(ctx context.Context, tx *sql.Tx, ranobeID int, authors []string) error {
	for _, name := range authors {
		var authorID int
		err := tx.QueryRowContext(ctx, `
			INSERT INTO authors (name) VALUES ($1) 
			ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name 
			RETURNING id`, name).Scan(&authorID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO ranobe_authors (ranobe_id, author_id) VALUES ($1, $2) 
			ON CONFLICT DO NOTHING`, ranobeID, authorID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RanobeRepository) linkGenres(ctx context.Context, tx *sql.Tx, ranobeID int, genres []string) error {
	for _, name := range genres {
		var genreID int
		err := tx.QueryRowContext(ctx, `
			INSERT INTO genres (name) VALUES ($1) 
			ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name 
			RETURNING id`, name).Scan(&genreID)
		if err != nil {
			return err
		}

		_, err = tx.ExecContext(ctx, `
			INSERT INTO ranobe_genres (ranobe_id, genre_id) VALUES ($1, $2) 
			ON CONFLICT DO NOTHING`, ranobeID, genreID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *RanobeRepository) CreateVolume(ctx context.Context, ranobeID int, vol *db.RanobeVolume) error {
	query := `INSERT INTO ranobe_volumes (ranobe_id, title, volume_number) VALUES ($1, $2, $3) RETURNING id`
	return r.DB.QueryRowContext(ctx, query, ranobeID, vol.Title, vol.Number).Scan(&vol.ID)
}

func (r *RanobeRepository) CreateChapter(ctx context.Context, volumeID int, ch *db.RanobeChapter) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO ranobe_chapters (ranobe_volume_id, title, chapter_number, content) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`
	
	err = tx.QueryRowContext(ctx, query, volumeID, ch.Title, ch.Number, ch.Content).Scan(&ch.ID)
	if err != nil {
		return err
	}

	if len(ch.Images) > 0 {
		stmt, err := tx.PrepareContext(ctx, `INSERT INTO ranobe_chapter_images (chapter_id, caption, image_url) VALUES ($1, $2, $3)`)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, img := range ch.Images {
			if _, err := stmt.ExecContext(ctx, ch.ID, img.Caption, img.ImageURL); err != nil {
				return err
			}
		}
	}

	return tx.Commit()
}

func (r *RanobeRepository) getChaptersByVolumeID(ctx context.Context, volumeID int) ([]db.RanobeChapter, error) {
	query := `SELECT id, title, chapter_number FROM ranobe_chapters WHERE ranobe_volume_id = $1 ORDER BY chapter_number ASC`
	rows, err := r.DB.QueryContext(ctx, query, volumeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var chapters []db.RanobeChapter
	for rows.Next() {
		var c db.RanobeChapter
		if err := rows.Scan(&c.ID, &c.Title, &c.Number); err != nil {
			return nil, err
		}
		chapters = append(chapters, c)
	}
	return chapters, nil
}

func (r *RanobeRepository) GetChapterByID(ctx context.Context, chapterID int) (*db.RanobeChapter, error) {
	ch := &db.RanobeChapter{}
	query := `SELECT id, title, chapter_number, content FROM ranobe_chapters WHERE id = $1`
	
	err := r.DB.QueryRowContext(ctx, query, chapterID).Scan(&ch.ID, &ch.Title, &ch.Number, &ch.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	imgQuery := `SELECT id, caption, image_url FROM ranobe_chapter_images WHERE chapter_id = $1 ORDER BY id`
	rows, err := r.DB.QueryContext(ctx, imgQuery, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var img db.RanobeChapterImages
		if err := rows.Scan(&img.ID, &img.Caption, &img.ImageURL); err != nil {
			return nil, err
		}
		ch.Images = append(ch.Images, img)
	}

	return ch, nil
}

func (r *RanobeRepository) RateRanobe(ctx context.Context, ranobeID int, userID int, rating int) error {
	if rating < 1 || rating > 10 {
		return errors.New("rating must be between 1 and 10")
	}
	
	query := `
		INSERT INTO ranobe_ratings (ranobe_id, user_id, rating)
		VALUES ($1, $2, $3)
		ON CONFLICT (ranobe_id, user_id)
		DO UPDATE SET rating = EXCLUDED.rating, created_at = NOW()`
	
	_, err := r.DB.ExecContext(ctx, query, ranobeID, userID, rating)
	return err
}