package dto

import "time"

type ShortNovelaInfo struct {
	NovelaID    int
	Title       string
	Description string
	PosterURL   string
	ReleaseDate time.Time
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
