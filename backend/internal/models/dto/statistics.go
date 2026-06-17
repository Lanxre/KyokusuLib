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

type GeneralStatistics struct {
	BookmarkCount int `json:"bookmarkCount"`
	ReadCount     int `json:"readCount"`
	RatingCount   int `json:"ratingCount"`
	CommentCount  int `json:"commentCount"`
	VolumeCount   int `json:"volumeCount"`
	ChapterCount  int `json:"chapterCount"`
}

type ShortNovelaMonthly struct {
	NovelaID     int    `json:"novelaId"`
	Title        string `json:"title"`
	PosterURL    string `json:"posterUrl"`
}

type NovelaMonthlySeries struct {
	Novela       ShortNovelaMonthly `json:"novela"`
	MonthlyReads []int              `json:"monthlyReads"`
}
