package dto

type CreateNovelaDTO struct {
	Title             string `validate:"required,min=2,max=255"`
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
	ID                int       		`json:"id"`
	Title             string    		`json:"title"`
	AlternativeTitles []string  		`json:"alternative_titles"`
	Description       string    		`json:"description"`
	Type              string    		`json:"type"`
	AgeRating         string    		`json:"age_rating"`
	ReleaseDate       string 			`json:"release_date"`
	Status            string    		`json:"status"`
	TranslationStatus string    		`json:"translation_status"`
	PosterURL         string    		`json:"poster_url"`
	Country           string    		`json:"country"`
	Views             int       		`json:"views"`
	Rating            float64   		`json:"rating"`
	Genres     		  []string      	`json:"genres"`
	Categories 		  []string      	`json:"categories"`
	Authors    		  []NovelaAuthor 	`json:"authors"`
	Volumes    		  []NovelaVolume 	`json:"volumes"`
	Bookmark  		  *string			`json:"bookmark"`
	BookmarkCount	  int				`json:"bookmark_count"`
	HasLiked  		  bool				`json:"has_liked"`
	LikeCount  		  int				`json:"like_count"`
	UserRating  	  int				`json:"user_rating"`
	RatingCount  	  int				`json:"rating_count"`
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