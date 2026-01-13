package dto

type UpdateBookmarkRequest struct {
	NovelaID int `json:"novela_id" validate:"required,gt=0"`
	Category string `json:"category" validate:"required,oneof=planned reading completed on_hold dropped"`
}

type DeleteBookmarkRequest struct {
	NovelaID int `json:"novela_id" validate:"required,gt=0"`
}

type BookmarkResponse struct {
	UserID    int    `json:"user_id"`
	NovelaID  int    `json:"novela_id"`
	Category  string `json:"category"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UpdateLikeRequest struct {
	NovelaID int  `json:"novela_id" validate:"required,gt=0"`
	HasLiked bool `json:"has_liked"`
}

type UpdateRatingRequest struct {
	NovelaID int `json:"novela_id" validate:"required,gt=0"`
	Rating   int `json:"rating" validate:"required,gt=0,lte=10"`
}