package dto

type NovelaSort string

const (
	SortNew      NovelaSort = "new"
	SortPopular  NovelaSort = "popular"
	SortRating   NovelaSort = "rating"
	SortUpdated  NovelaSort = "updated"
	SortTrending NovelaSort = "trending" 
)

type CreateNovelaDTO struct {
	Title             string `validate:"required,min=2,max=500"`
	Description       string `validate:"required,max=5000"`
	Type              string `validate:"required,oneof=Ранобе Веб-новелла Манга Манхва"`
	AgeRating         string `validate:"required,oneof=0+ 6+ 12+ 16+ 18+"`
	ReleaseYear       string `validate:"required,len=4,numeric"`
	Status            string `validate:"required"`
	TranslationStatus string `validate:"required"`
	Country           string `validate:"required"`
	
	Genres            []string `validate:"required,min=1"`
	Categories        []string `validate:"required,min=1"`
}

type NovelaResponse struct {
	ID                int                    `json:"id"`
	Title             string                 `json:"title"`
	AlternativeTitles []string               `json:"alternative_titles"`
	Description       string                 `json:"description"`
	Type              string                 `json:"type"`
	AgeRating         string                 `json:"age_rating"`
	ReleaseDate       string                 `json:"release_date"`
	Status            string                 `json:"status"`
	TranslationStatus string                 `json:"translation_status"`
	PosterURL         string                 `json:"poster_url"`
	Country           string                 `json:"country"`
	Views             int                    `json:"views"`
	Genres            []string               `json:"genres"`
	Categories        []string               `json:"categories"`
	Authors           []NovelaAuthor         `json:"authors"`
	Volumes           []NovelaVolume         `json:"volumes"`
	Bookmark          *string                `json:"bookmark"`
	HasLiked          bool                   `json:"has_liked"`
	LikeCount         int                    `json:"like_count"`
	UserRating        int                    `json:"user_rating"`
	RatingDetails     NovelaRatingCategory   `json:"rating_details"`
	BookmarkDetails   NovelaBookmarkCategory `json:"bookmark_details"`
}
type NovelaRatingCategory struct {
	Total        int       `json:"total"`
	TotalRating float64   `json:"total_rating"`
	NCItems      []NCItem  `json:"nc_items"`
}

type NovelaBookmarkCategory struct {
	Total        int       `json:"total"`
	NCItems      []NCItem  `json:"nc_items"`
}

type NCItem struct {
	Value any `json:"value"`
	Count int `json:"count"`
}

type NovelaAuthor struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

type NovelaVolume struct {
	ID       string          `json:"id"`
	Title    string          `json:"title"`
	Number   int             `json:"number"`
	Chapters []NovelaChapter `json:"chapters"`
}

type NovelaChapter struct {
	ID     string    `json:"id"`
	Title  string `json:"title"`
	Number float64 `json:"number"`
	Content string `json:"content,omitempty"`
	Images  []NovelaChapterImage `json:"images"`
}

type NovelaChapterImage struct {
	ID       int    `json:"id"`
	ImageURL string `json:"image_url"`
	Caption  string `json:"caption"`
	Position int    `json:"position"`
}

type UpdateNovelaRequest struct {
	Title             string   `json:"title" validate:"required"`
	AlternativeTitles []string `json:"alternative_titles"`
	Description       string   `json:"description"`
	Type              string   `json:"type"`
	AgeRating         string   `json:"age_rating"`
	ReleaseDate       string   `json:"release_date"`
	Status            string   `json:"status"`
	Country           string   `json:"country"`
	TranslationStatus string   `json:"translation_status"`
	Genres            []string `json:"genres"`
	Categories        []string `json:"categories"`
	Authors           []int    `json:"authors"`
}

type NovelaFilters struct {
	Search     string     `json:"search"`
	Genres     []string   `json:"genres"`
	Categories []string   `json:"categories"`
	Type       string     `json:"type"`
	Status     string     `json:"status"`
	Sort       NovelaSort `json:"sort"`
	Limit      int        `json:"limit"`
	Offset     int        `json:"offset"`
	AuthorID   int        `json:"author_id"`
}

type UserNovelaBookmark struct {
	ID     		int    		`json:"id"`
	Title  		string 		`json:"title"`
	PosterURL 	string 		`json:"poster_url"`
	Type 		string 		`json:"type"`
	Rating     	float64     `json:"rating"`
}

type AddNovelaTeamRequest struct {
	TeamID int `json:"team_id" validate:"required"`
}

type AddVolumeRequest struct {
	VolumeNumber int    `json:"volume_number" validate:"required"`
	Title        string `json:"title"`
}

type AddChapterRequest struct {
	ChapterNumber float64 `json:"chapter_number" validate:"required"`
	Title         string  `json:"title"`
	Content       string  `json:"content" validate:"required"`
}

type AddChapterImageRequest struct {
	ImageURL string `json:"image_url" validate:"required"`
	Caption  string `json:"caption"`
	Position int    `json:"position"`
}
