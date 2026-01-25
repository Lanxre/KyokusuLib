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

func (r *CommentsRepository) GetCommentsByNovelaID(ctx context.Context, novelaID int) ([]db.NovelaComment, error) {
	query := `
		SELECT c.id, c.parent_id, c.content, c.created_at, u.id as user_id, up.name as user_name, up.picture as user_image
		FROM novela_comments c
		JOIN users u ON c.user_id = u.id
		LEFT JOIN user_profiles up ON u.id = up.user_id
		WHERE c.novela_id = $1
		ORDER BY c.created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, novelaID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []db.NovelaComment
	for rows.Next() {
		var c db.NovelaComment
		if err := rows.Scan(&c.ID, &c.ParentID, &c.Content, &c.CreatedAt, &c.UserID, &c.UserName, &c.UserImage); err != nil {
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