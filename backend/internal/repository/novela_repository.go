package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lib/pq"
)

type NovelaRepository struct {
	DB *sql.DB
}

func NewNovelaRepository(db *sql.DB) *NovelaRepository {
	return &NovelaRepository{DB: db}
}

func (r *NovelaRepository) Create(ctx context.Context, n *db.Novela) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO novela (
			title, alternative_titles, description, type, age_rating, 
			release_date, status, country, translation_status, poster_url, views
		)
		SELECT $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
		WHERE NOT EXISTS (
			SELECT 1 FROM novela 
			WHERE title ILIKE $1 
			   OR $1 ILIKE ANY(alternative_titles)
			   OR title = ANY($2)
			   OR alternative_titles && $2
		)
		RETURNING id`

	err = tx.QueryRowContext(ctx, query,
		n.Title,
		n.AlternativeTitles,
		n.Description,
		n.Type,
		n.AgeRating,
		n.ReleaseDate,
		n.Status,
		n.Country,
		n.TranslationStatus,
		n.PosterURL,
		0,
	).Scan(&n.ID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("novela with this title already exists")
		}
		return fmt.Errorf("failed to insert novela: %w", err)
	}

	if len(n.Genres) > 0 {
		if err := r.linkTags(ctx, tx, n.ID, n.Genres, "genres", "novela_genres", "genre_id"); err != nil {
			return err
		}
	}

	if len(n.Categories) > 0 {
		if err := r.linkTags(ctx, tx, n.ID, n.Categories, "categories", "novela_categories", "category_id"); err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (r *NovelaRepository) GetFullByID(ctx context.Context, id, userID int) (*db.Novela, error) {
	n := &db.Novela{}

	var (
		authorsJSON []byte
		volumesJSON []byte
	)

	query := `
		SELECT 
			n.id, 
			n.title, 
			COALESCE(n.alternative_titles, '{}')::text[],
			n.description, 
			n.type, 
			n.age_rating, 
			n.release_date, 
			n.status,
			n.country, 
			n.translation_status, 
			n.poster_url, 
			n.views,
			COALESCE((SELECT AVG(rating) FROM novela_ratings WHERE novela_id = n.id), 0),
			
			COALESCE((
				SELECT array_agg(g.name) 
				FROM novela_genres ng 
				JOIN genres g ON ng.genre_id = g.id 
				WHERE ng.novela_id = n.id
			), '{}')::text[],

			COALESCE((
				SELECT array_agg(c.name) 
				FROM novela_categories nc 
				JOIN categories c ON nc.category_id = c.id 
				WHERE nc.novela_id = n.id
			), '{}')::text[],

			COALESCE((
				SELECT json_agg(json_build_object('id', a.id, 'name', a.name, 'role', na.role))
				FROM novela_authors na 
				JOIN authors a ON na.author_id = a.id 
				WHERE na.novela_id = n.id
			), '[]'),

			COALESCE((
				SELECT json_agg(
					json_build_object(
						'id', v.id,
						'title', v.title,
						'number', v.volume_number,
						'chapters', COALESCE((
							SELECT json_agg(
								json_build_object(
									'id', ch.id,
									'title', ch.title,
									'number', ch.chapter_number,
									'images', COALESCE((
										SELECT json_agg(
											json_build_object('id', img.id, 'image_url', img.image_url, 'caption', img.caption)
											ORDER BY img.id
										) FROM novela_chapter_images img WHERE img.chapter_id = ch.id
									), '[]'::json)
								) ORDER BY ch.chapter_number
							) FROM novela_chapters ch WHERE ch.novela_volume_id = v.id
						), '[]'::json)
					) ORDER BY v.volume_number
				) FROM novela_volumes v WHERE v.novela_id = n.id
			), '[]'),

			CASE 
				WHEN $2 > 0 THEN (SELECT category FROM user_novela_bookmarks WHERE novela_id = n.id AND user_id = $2)
				ELSE NULL 
			END as user_category,

			COALESCE((SELECT COUNT(*) FROM user_novela_bookmarks WHERE novela_id = n.id), 0),

			COALESCE((SELECT has_liked FROM user_novela_likes WHERE novela_id = n.id AND user_id = $2), FALSE),
			COALESCE((SELECT COUNT(*) FROM user_novela_likes WHERE novela_id = n.id AND has_liked = TRUE), 0)
				

		FROM novela n
		WHERE n.id = $1`

	err := r.DB.QueryRowContext(ctx, query, id, userID).Scan(
		&n.ID,
		&n.Title,
		pq.Array(&n.AlternativeTitles),
		&n.Description,
		&n.Type,
		&n.AgeRating,
		&n.ReleaseDate,
		&n.Status,
		&n.Country,
		&n.TranslationStatus,
		&n.PosterURL,
		&n.Views,
		&n.Rating,
		pq.Array(&n.Genres),
		pq.Array(&n.Categories),
		&authorsJSON,
		&volumesJSON,
		&n.Bookmark,
		&n.BookmarkCount,
		&n.HasLiked,
		&n.LikeCount,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if len(authorsJSON) > 0 {
		if err := json.Unmarshal(authorsJSON, &n.Authors); err != nil {
			return nil, err
		}
	}
	if len(volumesJSON) > 0 {
		if err := json.Unmarshal(volumesJSON, &n.Volumes); err != nil {
			return nil, err
		}
	}

	if n.Bookmark != nil {
		*n.Bookmark = db.BookmarkCategory(*n.Bookmark)
	}

	return n, nil
}

func (r *NovelaRepository) Update(ctx context.Context, n *db.Novela) error {
	query := `
		UPDATE novela 
		SET title=$1, alternative_titles=$2, description=$3, type=$4, age_rating=$5, 
		    release_date=$6, status=$7, translation_status=$8, poster_url=$9, updated_at=NOW()
		WHERE id=$10`

	_, err := r.DB.ExecContext(ctx, query,
		n.Title,
		n.AlternativeTitles,
		n.Description,
		n.Type,
		n.AgeRating,
		n.ReleaseDate,
		n.Status,
		n.TranslationStatus,
		n.PosterURL,
		n.ID,
	)
	return err
}

func (r *NovelaRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM novela WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *NovelaRepository) LinkAuthor(ctx context.Context, novelaID int, authorName string, role string) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var authorID int
	err = tx.QueryRowContext(ctx, `
		INSERT INTO authors (name) VALUES ($1) 
		ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name 
		RETURNING id`, authorName).Scan(&authorID)

	if err != nil {
		return fmt.Errorf("failed to upsert author: %w", err)
	}
	query := `
		INSERT INTO novela_authors (novela_id, author_id, role) 
		VALUES ($1, $2, $3)
		ON CONFLICT (novela_id, author_id) 
		DO UPDATE SET role = EXCLUDED.role`

	_, err = tx.ExecContext(ctx, query, novelaID, authorID, role)
	if err != nil {
		return fmt.Errorf("failed to link author: %w", err)
	}

	return tx.Commit()
}
func (r *NovelaRepository) linkTags(ctx context.Context, tx *sql.Tx, novelaID int, tags []string, tableName, linkTable, fkCol string) error {
	for _, name := range tags {
		var tagID int
		upsertQuery := fmt.Sprintf(`
			INSERT INTO %s (name) VALUES ($1) 
			ON CONFLICT (name) DO UPDATE SET name=EXCLUDED.name 
			RETURNING id`, tableName)

		err := tx.QueryRowContext(ctx, upsertQuery, name).Scan(&tagID)
		if err != nil {
			return err
		}

		linkQuery := fmt.Sprintf(`
			INSERT INTO %s (novela_id, %s) VALUES ($1, $2) 
			ON CONFLICT DO NOTHING`, linkTable, fkCol)

		_, err = tx.ExecContext(ctx, linkQuery, novelaID, tagID)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *NovelaRepository) CreateVolume(ctx context.Context, novelaID int, title string, number int) (int, error) {
	query := `
		INSERT INTO novela_volumes (novela_id, title, volume_number) 
		VALUES ($1, $2, $3) 
		RETURNING id`

	var id int
	err := r.DB.QueryRowContext(ctx, query, novelaID, title, number).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("failed to create volume: %w", err)
	}
	return id, nil
}

func (r *NovelaRepository) CreateChapter(ctx context.Context, volumeID int, ch *db.NovelaChapter) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		INSERT INTO novela_chapters (novela_volume_id, title, chapter_number, content) 
		VALUES ($1, $2, $3, $4) 
		RETURNING id`

	err = tx.QueryRowContext(ctx, query, volumeID, ch.Title, ch.Number, ch.Content).Scan(&ch.ID)
	if err != nil {
		return fmt.Errorf("failed to create chapter: %w", err)
	}

	if len(ch.Images) > 0 {
		imgQuery := `INSERT INTO novela_chapter_images (chapter_id, image_url, caption) VALUES ($1, $2, $3)`

		stmt, err := tx.PrepareContext(ctx, imgQuery)
		if err != nil {
			return err
		}
		defer stmt.Close()

		for _, img := range ch.Images {
			if _, err := stmt.ExecContext(ctx, ch.ID, img.ImageURL, img.Caption); err != nil {
				return fmt.Errorf("failed to save image: %w", err)
			}
		}
	}

	return tx.Commit()
}

func (r *NovelaRepository) GetChapterByID(ctx context.Context, chapterID int) (*db.NovelaChapter, error) {
	ch := &db.NovelaChapter{}

	query := `SELECT id, title, chapter_number, content FROM novela_chapters WHERE id = $1`
	err := r.DB.QueryRowContext(ctx, query, chapterID).Scan(&ch.ID, &ch.Title, &ch.Number, &ch.Content)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	imgQuery := `SELECT id, image_url, caption FROM novela_chapter_images WHERE chapter_id = $1 ORDER BY id`
	rows, err := r.DB.QueryContext(ctx, imgQuery, chapterID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var img db.NovelaChapterImage
		var caption sql.NullString

		if err := rows.Scan(&img.ID, &img.ImageURL, &caption); err != nil {
			return nil, err
		}
		img.Caption = caption.String
		ch.Images = append(ch.Images, img)
	}

	return ch, nil
}

func (r *NovelaRepository) FindByTitle(ctx context.Context, title string) (*db.Novela, error) {
	novela := &db.Novela{}

	query := `
		SELECT id, title, alternative_titles
		FROM novela
		WHERE title ILIKE $1 OR $1 ILIKE ANY(alternative_titles)
		LIMIT 1`

	err := r.DB.QueryRowContext(ctx, query, title).Scan(&novela.ID, &novela.Title, pq.Array(&novela.AlternativeTitles))

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return novela, nil
}

func (r *NovelaRepository) SetBookmark(ctx context.Context, bookmark *db.Bookmark) error {
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

func (r *NovelaRepository) RemoveBookmark(ctx context.Context, userID, novelaID int) error {
	query := `DELETE FROM user_novela_bookmarks WHERE user_id = $1 AND novela_id = $2`
	_, err := r.DB.ExecContext(ctx, query, userID, novelaID)
	return err
}

func (r *NovelaRepository) SetLike(ctx context.Context, novelaLike *db.NovelaLike) error {
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