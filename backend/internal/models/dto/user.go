package dto

import (
	"time"
)

type UserDTO struct {
	Id 		  string
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Avatar    string    `json:"avatar"`
	Token 	  string
}

type GetUserDTO struct {
	ID           			 		int       	`json:"id"`
	Name         			 		string    	`json:"name"`
	Picture      			 		string    	`json:"picture"`
	Banner      			 		string    	`json:"banner"`
	Role         			 		string    	`json:"role"`
	Status       			 		string    	`json:"status"`
	
	About							string		`json:"about"`
	Birthday						*time.Time	`json:"birthday"`
	Gender							string		`json:"gender"`
	
	IsPublic             		    bool        `json:"is_public"`
	LastLogin					    time.Time	`json:"last_login"`
	CreateAt 						time.Time 	`json:"create_at"`
	ActiveTag						string		`json:"active_tag"`
	AllTags							[]UserTagDTO	`json:"tags"`
	Settings						PublicUserSettingsDTO	`json:"settings"`
	UserLevel						UserLevelDTO	`json:"user_level"`
}

type UserLevelDTO struct {
	Level         int    `json:"level"`
    Experience    int64  `json:"experience"`
    LevelTitle    string `json:"level_title"`
    XPForNext     int64  `json:"xp_needed_for_next"`
}

type UpdateUserStatusDTO struct {
	Status 	 string `json:"status"`
	LastActive int64 `json:"last_active"`
}

type UserTagDTO struct {
	ID 	  int `json:"tag_id"`
	Tag   string `json:"tag"`
}

type UpdateUserTagDTO struct {
	ID int 	   `json:"id"`
	Tag string `json:"tag"`
}

type PublicUserSettingsDTO struct {
	IsShowTag						bool		`json:"is_show_tag"`
}
