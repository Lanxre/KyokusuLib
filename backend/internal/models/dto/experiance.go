package dto

type LevelDefinitions struct {
	Level int					`json:"level"`
	Title string				`json:"title"`
	Total_XP_Required int64		`json:"total_xp_required"`
}