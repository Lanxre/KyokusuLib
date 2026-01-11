package db

type BookmarkCategory string

const (
    CategoryPlanned   BookmarkCategory = "planned"
    CategoryReading   BookmarkCategory = "reading"
    CategoryCompleted BookmarkCategory = "completed"
    CategoryOnHold    BookmarkCategory = "on_hold"
    CategoryDropped   BookmarkCategory = "dropped"
)

type Bookmark struct {
    UserID   int              `json:"user_id"`
    NovelaID int              `json:"novela_id"`
    Category BookmarkCategory `json:"category"`
}

