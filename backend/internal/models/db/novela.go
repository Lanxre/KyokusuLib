package db

import (
	"time"
)

type Novela struct {
	ID                int       `json:"id"`
	Title             string    `json:"title"`
	AlternativeTitles []string  `json:"alternative_titles"`
	Description       string    `json:"description"`
	Type              string    `json:"type"`
	AgeRating         string    `json:"age_rating"`
	ReleaseDate       time.Time `json:"release_date"`
	Status            string    `json:"status"`
	TranslationStatus string    `json:"translation_status"`
	PosterURL         string    `json:"poster_url"`
	Country           string    `json:"country"`
	Views             int       `json:"views"`
	Rating            float64   `json:"rating"`

	
	Genres     		[]string      		`json:"genres"`
	Categories 		[]string      		`json:"categories"`
	Authors    		[]NovelaAuthor 		`json:"authors"`
	
	Volumes    		[]NovelaVolume 		`json:"volumes"`
	
	Bookmark      	*BookmarkCategory 	`json:"bookmark"`
	BookmarkCount 	int 				`json:"bookmark_count"`
	HasLiked      	bool				`json:"has_liked"`
}

type NovelaAuthor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type NovelaVolume struct {
	ID       int             `json:"id"`
	Title    string          `json:"title"`
	Number   int             `json:"number"`
	Chapters []NovelaChapter `json:"chapters"`
}

type NovelaChapter struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Number float64 `json:"number"`
	Content string `json:"content,omitempty"`
	Images  []NovelaChapterImage `json:"images"` 
}

type NovelaChapterImage struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	Caption  string `json:"caption"`
}
type NovelaLike struct {
	NovelaID int  `json:"novela_id"`
	UserID   int  `json:"user_id"`
	HasLiked bool `json:"has_liked"`
}