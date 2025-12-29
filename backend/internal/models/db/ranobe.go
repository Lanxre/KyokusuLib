package db

import "time"

type Ranobe struct {
	ID          			int `json:"id"`
	Title       			string `json:"title"`
	Description 			string `json:"description"`
	Authors                 []string `json:"authors"`
	Genres      			[]string `json:"genres"`
	TranslationStatus       string `json:"translation_status"`
	Country 				string `json:"country"`
	Year       				*time.Time `json:"release_date"`
	Views      				int `json:"views"`
	Rating     				float64 `json:"rating"`
	
	Volumes []RanobeVolume  `json:"volumes"`
}

type RanobeVolume struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Number int `json:"number"`
	Chapters []RanobeChapter `json:"chapters"`
}

type RanobeChapter struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Number int `json:"number"`
	Content string `json:"content"`
	Images []RanobeChapterImages `json:"images"`
}

type RanobeChapterImages struct {
	ID int `json:"id"`
	Caption string `json:"caption"`
	ImageURL string `json:"image_url"`
}
