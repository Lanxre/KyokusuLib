package dto

type CommentResponse struct {
	ID        int                `json:"id"`
	ParentID  *int               `json:"parent_id"`
	Content   string             `json:"content"`
	CreatedAt string             `json:"created_at"`
	User      CommentUserAuthor  `json:"user"`
	Replies   []*CommentResponse `json:"replies"`
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