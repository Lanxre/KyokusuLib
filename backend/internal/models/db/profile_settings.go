package db

type UserTheme string

const (
	DarkTheme  UserTheme = "dark"
	LightTheme UserTheme = "light"
)

type UserProfileSetting struct {
	Theme string `json:"theme"`
	IsAppNotify bool `json:"is_app_notify"`
	IsNewPublishedNotify bool `json:"is_new_published_notify"`
	IsShowTag bool `json:"is_show_tag"`
	IsShowBookmark bool `json:"is_show_bookmark"`
	FontSize int `json:"font_size"`
	LineWeight int `json:"line_weight"`
	FontFamily string `json:"font_family"`
	AutoScroll bool `json:"auto_scroll"`
}
