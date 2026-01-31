package dto

type CommentResponse struct {
	ID        int                `json:"id"`
	ParentID  *int               `json:"parent_id"`
	Content   string             `json:"content"`
	CreatedAt string             `json:"created_at"`
	User      CommentUserAuthor  `json:"user"`
	Replies   []*CommentResponse `json:"replies"`
	LikeCount int                `json:"like_count"`
	HasLike   *bool               `json:"has_like"`
	HasReport *bool               `json:"has_report"`
}

type CommentUserAuthor struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Picture string `json:"picture"`
}

type CreateCommentRequest struct {
	NovelaID int    `json:"novela_id" validate:"required"`
	ParentID *int   `json:"parent_id"`
	Content  string `json:"content" validate:"required,min=1,max=5000"`
}

type UpdateCommentRequest struct {
	Content   string `json:"content" validate:"required,min=1,max=5000"`
	UpdatedAt string `json:"updated_at" validate:"required"`
}

type ProfileUserCommentResponse struct {
	ID        int 		`json:"id"`
	NovelaID  int		`json:"novela_id"`
	NovelaTitle string 	`json:"novela_title"`
	Content   string 	`json:"content"`
	CreatedAt string	`json:"created_at"`
	UpdatedAt string 	`json:"updated_at"`
}