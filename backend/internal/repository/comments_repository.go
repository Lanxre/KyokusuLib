package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
)

type CommentsRepository struct {
	DB *sql.DB
}

func NewCommentsRepository(db *sql.DB) *CommentsRepository {
	return &CommentsRepository{DB: db}
}

func (r *CommentsRepository) GetCommentsByNovelaID(ctx context.Context, novelaID, userID int) ([]db.NovelaComment, error) {
	query := `
		SELECT 
			c.id,
			c.parent_id,
			c.content,
			c.created_at,
			u.id          AS user_id,
			up.name       AS user_name,
			up.picture    AS user_image,
			COUNT(cl.user_id) AS like_count,
			BOOL_OR(cl.user_id = $2) AS has_like,
			BOOL_OR(cr.user_id = $2) AS has_report
		FROM novela_comments c
		JOIN users u ON c.user_id = u.id
		LEFT JOIN user_profiles up ON u.id = up.user_id
		LEFT JOIN like_comments cl ON cl.comment_id = c.id
		LEFT JOIN novela_comments_reports cr ON cr.comment_id = c.id
		WHERE c.novela_id = $1
		GROUP BY c.id, u.id, up.name, up.picture`

	rows, err := r.DB.QueryContext(ctx, query, novelaID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []db.NovelaComment
	for rows.Next() {
		var c db.NovelaComment
		if err := rows.Scan(
			&c.ID, 
			&c.ParentID, 
			&c.Content, 
			&c.CreatedAt, 
			&c.UserID, 
			&c.UserName, 
			&c.UserImage,
			&c.LikeCount,
			&c.HasLike,
			&c.HasReport,
		); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (r *CommentsRepository) CreateComment(ctx context.Context, userID int, req *dto.CreateCommentRequest) (int, error) {
	query := `INSERT INTO novela_comments (novela_id, user_id, parent_id, content) VALUES ($1, $2, $3, $4) RETURNING id`
	var id int
	err := r.DB.QueryRowContext(ctx, query, req.NovelaID, userID, req.ParentID, req.Content).Scan(&id)
	return id, err
}

func (r *CommentsRepository) DeleteComment(ctx context.Context, commentID, userID int) error {
	query := `DELETE FROM novela_comments WHERE id = $1 AND user_id = $2`
	_, err := r.DB.ExecContext(ctx, query, commentID, userID)
	return err
}

func (r *CommentsRepository) UpdateComment(ctx context.Context, commentID, userID int, req *dto.UpdateCommentRequest) error {
	query := `UPDATE novela_comments SET content = $1, updated_at = $2 WHERE id = $3 AND user_id = $4`
	_, err := r.DB.ExecContext(ctx, query, req.Content, req.UpdatedAt, commentID, userID)
	return err
}

func (r *CommentsRepository) SetCommentLike(ctx context.Context, commentID, userID int) error {
	query := `INSERT INTO like_comments (comment_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.DB.ExecContext(ctx, query, commentID, userID)
	return err
}

func (r *CommentsRepository) DeleteCommentLike(ctx context.Context, commentID, userID int) error {
	query := `DELETE FROM like_comments WHERE comment_id = $1 AND user_id = $2`
	_, err := r.DB.ExecContext(ctx, query, commentID, userID)
	return err
}

func (r *CommentsRepository) CreateCommentReport(ctx context.Context, commentID, userID int, reason string) error {
	query := `INSERT INTO novela_comments_reports (comment_id, user_id, reason) VALUES ($1, $2, $3)`
	_, err := r.DB.ExecContext(ctx, query, commentID, userID, reason)
	return err
}

func (r *CommentsRepository) GetCommentsByUserID(ctx context.Context, userID int) ([]db.SelectNovelaComment, error) {
	query := `SELECT c.id, c.novela_id, n.title as novela_title, c.content, c.created_at, c.updated_at 
			  FROM novela_comments c
			  JOIN novela n ON c.novela_id = n.id
			  WHERE c.user_id = $1 AND c.parent_id IS NULL
			  ORDER BY c.created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []db.SelectNovelaComment
	for rows.Next() {
		var c db.SelectNovelaComment
		if err := rows.Scan(&c.ID, &c.NovelaID, &c.NovelaTitle, &c.Content, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}