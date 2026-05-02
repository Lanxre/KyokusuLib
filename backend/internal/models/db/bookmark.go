package db

type BookmarkCategory struct {
	ID     int    `json:"id"`
	UserID *int   `json:"user_id"`
	Name   string `json:"name"`
}

type BookmarkCategoryCount struct {
	ID     int    `json:"id"`
	UserID *int   `json:"user_id"`
	Name   string `json:"name"`
	Count  int    `json:"count"`
}

type Bookmark struct {
	UserID     int `json:"user_id"`
	NovelaID   int `json:"novela_id"`
	CategoryID int `json:"category_id"`
}
