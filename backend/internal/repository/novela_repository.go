package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
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
		SELECT 
            $1::text, $2::text[], $3::text, $4::text, $5::text, 
            $6::timestamp, $7::text, $8::text, $9::text, $10::text, $11::int
		WHERE NOT EXISTS (
			SELECT 1 FROM novela 
			WHERE title ILIKE $1::text 
			   OR $1::text ILIKE ANY(alternative_titles)
			   OR title = ANY($2::text[])
			   OR alternative_titles && $2::text[]
		)
		RETURNING id`

	err = tx.QueryRowContext(ctx, query,
		n.Title,
		pq.Array(n.AlternativeTitles),
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
func (r *NovelaRepository) GetFullByID(tx *sql.Tx, ctx context.Context, id, userID int) (*db.Novela, error) {
	n := &db.Novela{}
	var (
		authorsJSON []byte
		volumesJSON []byte
		userRating  sql.NullInt32
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

			COALESCE((SELECT has_liked FROM user_novela_likes WHERE novela_id = n.id AND user_id = $1), FALSE) as has_liked,

			COALESCE((SELECT COUNT(*) FROM user_novela_likes WHERE novela_id = n.id AND has_liked = TRUE), 0),

			CASE
				WHEN $2 > 0 THEN (SELECT rating FROM novela_ratings WHERE novela_id = n.id AND user_id = $2)
				ELSE 0
			END as user_rating

		FROM novela n
		WHERE n.id = $1`

	err := tx.QueryRowContext(ctx, query, id, userID).Scan(
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
		pq.Array(&n.Genres),
		pq.Array(&n.Categories),
		&authorsJSON,
		&volumesJSON,
		&n.Bookmark,
		&n.HasLiked,
		&n.LikeCount,
		&userRating,
	)

	if err != nil {
		fmt.Println(err)
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
		val := db.BookmarkCategory(*n.Bookmark)
		n.Bookmark = &val
	}

	if userRating.Valid {
		n.UserRating = int(userRating.Int32)
	}

	return n, nil
}

func (r *NovelaRepository) UpdateFull(ctx context.Context, n *db.Novela) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	query := `
		UPDATE novela 
		SET title = $1, alternative_titles = $2, description = $3, type = $4, 
		    age_rating = $5, release_date = $6, status = $7, country = $8, 
		    translation_status = $9, poster_url = COALESCE(NULLIF($10, ''), poster_url),
			updated_at = NOW()
		WHERE id = $11`

	res, err := tx.ExecContext(ctx, query,
		n.Title, pq.Array(n.AlternativeTitles), n.Description, n.Type,
		n.AgeRating, n.ReleaseDate, n.Status, n.Country,
		n.TranslationStatus, n.PosterURL, n.ID,
	)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	for _, table := range []string{"novela_genres", "novela_categories", "novela_authors"} {
		if _, err := tx.ExecContext(ctx, fmt.Sprintf("DELETE FROM %s WHERE novela_id = $1", table), n.ID); err != nil {
			return err
		}
	}

	if err := r.linkTags(ctx, tx, n.ID, n.Genres, "genres", "novela_genres", "genre_id"); err != nil {
		return err
	}

	if err := r.linkTags(ctx, tx, n.ID, n.Categories, "categories", "novela_categories", "category_id"); err != nil {
		return err
	}

	for _, author := range n.Authors {
		if _, err := tx.ExecContext(ctx,
			`INSERT INTO novela_authors (novela_id, author_id, role) VALUES ($1, $2, $3)`,
			n.ID, author.ID, "Author",
		); err != nil {
			return err
		}
	}

	return tx.Commit()
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

func (r *NovelaRepository) UpdatePoster(ctx context.Context, id int, posterURL string) error {
	query := `UPDATE novela SET poster_url = $2 WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id, posterURL)
	return err
}

func (r *NovelaRepository) GetNovelas(tx *sql.Tx, ctx context.Context, userID int, f dto.NovelaFilters) ([]db.Novela, int, error) {
	var args []interface{}
	where := []string{"1=1"}
	argID := 1

	args = append(args, userID)
	argID++

	if f.Search != "" {
		where = append(where, fmt.Sprintf("(n.title ILIKE $%d OR $%d ILIKE ANY(n.alternative_titles))", argID, argID))
		args = append(args, "%"+f.Search+"%")
		argID++
	}

	if len(f.Genres) > 0 {
		where = append(where, fmt.Sprintf(`EXISTS (
			SELECT 1 FROM novela_genres ng JOIN genres g ON ng.genre_id = g.id 
			WHERE ng.novela_id = n.id AND g.name = ANY($%d)
		)`, argID))
		args = append(args, pq.Array(f.Genres))
		argID++
	}

	if len(f.Categories) > 0 {
		where = append(where, fmt.Sprintf(`EXISTS (
			SELECT 1 FROM novela_categories nc JOIN categories c ON nc.category_id = c.id 
			WHERE nc.novela_id = n.id AND c.name = ANY($%d)
		)`, argID))
		args = append(args, pq.Array(f.Categories))
		argID++
	}

	if f.Type != "" {
		where = append(where, fmt.Sprintf("n.type = $%d", argID))
		args = append(args, f.Type)
		argID++
	}

	orderBy := "n.created_at DESC"
	switch f.Sort {
	case dto.SortPopular:
		orderBy = "avg_rating DESC NULLS LAST"
	case dto.SortTrending:
		orderBy = "n.views DESC"
	case dto.SortUpdated:
		orderBy = "last_chapter_at DESC NULLS LAST"
	case dto.SortNew:
		orderBy = "n.created_at DESC"
	}

	query := fmt.Sprintf(`
    SELECT 
        n.id, n.title, COALESCE(n.alternative_titles, '{}')::text[], n.type, n.status, n.poster_url, n.views,
        COALESCE((SELECT array_agg(g.name) FROM novela_genres ng JOIN genres g ON ng.genre_id = g.id WHERE ng.novela_id = n.id), '{}')::text[],
        COALESCE((SELECT array_agg(c.name) FROM novela_categories nc JOIN categories c ON nc.category_id = c.id WHERE nc.novela_id = n.id), '{}')::text[],
        COALESCE((SELECT AVG(rating) FROM novela_ratings WHERE novela_id = n.id), 0) as avg_rating,
        (SELECT MAX(created_at) FROM novela_chapters ch JOIN novela_volumes v ON ch.novela_volume_id = v.id WHERE v.novela_id = n.id) as last_chapter_at,
        
        -- Поля пользователя
        (SELECT category FROM user_novela_bookmarks WHERE novela_id = n.id AND user_id = $1) as user_bookmark,
        
        COALESCE((SELECT has_liked FROM user_novela_likes WHERE novela_id = n.id AND user_id = $1), FALSE) as has_liked

    FROM novela n
    WHERE %s
    ORDER BY %s
    LIMIT $%d OFFSET $%d`, 
    strings.Join(where, " AND "), orderBy, argID, argID+1)

	finalArgs := append(args, f.Limit, f.Offset)
	rows, err := tx.QueryContext(ctx, query, finalArgs...)
	if err != nil {
		return nil, 0, err
	}
	defer rows.Close()

	var novelas []db.Novela
	for rows.Next() {
		var n db.Novela
		var lastChapterAt sql.NullTime
		err := rows.Scan(
			&n.ID,
			&n.Title,
			pq.Array(&n.AlternativeTitles),
			&n.Type,
			&n.Status,
			&n.PosterURL,
			&n.Views,
			pq.Array(&n.Genres),
			pq.Array(&n.Categories),
			&n.Rating,
			&lastChapterAt,
			&n.Bookmark,
			&n.HasLiked,
		)
		if err != nil {
			return nil, 0, err
		}
		novelas = append(novelas, n)
	}

	var total int
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM novela n WHERE %s", strings.Join(where, " AND "))
	if err := tx.QueryRowContext(ctx, countQuery, args[1:]...).Scan(&total); err != nil {
		return nil, 0, err
	}

	return novelas, total, nil
}
func (r *NovelaRepository) GetUserNovelaBookmarks(ctx context.Context, userID int, category db.BookmarkCategory) ([]db.UserNovelaBookmark, error) {
	query := `
		SELECT n.id, n.title, n.poster_url, n.type,
		       COALESCE((SELECT AVG(rating) FROM novela_ratings WHERE novela_id = n.id), 0) as rating
		FROM novela n
		JOIN user_novela_bookmarks b ON n.id = b.novela_id
		WHERE b.user_id = $1 AND b.category = $2
		ORDER BY b.updated_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, userID, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var novelas []db.UserNovelaBookmark
	for rows.Next() {
		var n db.UserNovelaBookmark
		if err := rows.Scan(&n.ID, &n.Title, &n.PosterURL, &n.Type, &n.Rating); err != nil {
			return nil, err
		}
		novelas = append(novelas, n)
	}
	
	return novelas, nil
}