package dto

import "time"

type ShortNovelaInfo struct {
	NovelaID    int			`json:"novelaId"`
	Title       string		`json:"title"`
	Description string		`json:"description"`
	PosterURL   string		`json:"posterUrl"`
	ReleaseDate time.Time	`json:"releaseDate"`
}

type TotatStatistics struct {
	Novela ShortNovelaInfo `json:"novela"`

	TotalBookmarkCount int     `json:"bookmarkCount"`
	TotalReadCount     int     `json:"readCount"`
	TotalRatingCount   int     `json:"ratingCount"`
	TotalCommentCount  int     `json:"commentCount"`
	TotalVolumeCount   int     `json:"volumeCount"`
	TotalChapterCount  int     `json:"chapterCount"`
	Rating             float64 `json:"rating"`
}

type TotalStatisticsResponse struct {
	Data  *[]TotatStatistics `json:"data"`
	Total int                `json:"total"`
}
