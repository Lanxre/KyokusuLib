package constants

type StatisticsPeriodSort string
type StatisticsFieldSort  string

const (
	Day   StatisticsPeriodSort = "day"
	Week  StatisticsPeriodSort = "week"
	Month StatisticsPeriodSort = "month"
)

const (
	Bookmarks StatisticsFieldSort = "bookmarks"
	Reads     StatisticsFieldSort = "reads"
	Rating    StatisticsFieldSort = "rating"
	Comments  StatisticsFieldSort = "comments"
	Volumes   StatisticsFieldSort = "volumes"
	Chapters  StatisticsFieldSort = "chapters"
	ID        StatisticsFieldSort = "id"
)

func (s StatisticsPeriodSort) IsValid() bool {
	switch s {
	case Day, Week, Month:
		return true
	}
	return false
}

func (s StatisticsFieldSort) IsValid() bool {
	switch s {
	case Bookmarks, Reads, Rating, Comments, Volumes, Chapters, ID:
		return true
	}
	return false
}

func (s StatisticsPeriodSort) String() string {
	return string(s)
}

func (s StatisticsFieldSort) String() string {
	return string(s)
}
