package constants

const (
	MostSearchedCategoriesLimit = 3

	UpLevelExperienceThreshold = 5
)

type TagId int
type TagName string

const (
	ReadTenChaptersTagId  TagId  = 1
	LevelTwoNewbieTagId   TagId  = 2
)

const (
	ReadTenChaptersTagName  TagName  = "read_ten_chapters"
	LevelTwoNewbieTagName   TagName  = "level_two_newbie"
)