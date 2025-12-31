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
}

type UpdateUserStatusDTO struct {
	Status 	 string `json:"status"`
	LastActive int64 `json:"last_active"`
}
