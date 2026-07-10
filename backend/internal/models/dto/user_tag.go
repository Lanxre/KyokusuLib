package dto

type UpdateUserTags struct {
	UserID int        `json:"user_id"`
	TagIds []int  	  `json:"tag_ids"`
}