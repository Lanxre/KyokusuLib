package db

import "time"

type UserGenere string

const (
	MALE_GENERE = "male"
	FEMALE_GENERE = "female"
	HIDDEN_GENERE = "hidden"
)

type User struct {
	ID           			 		int       `json:"id"`
	Email        			 		string    `json:"email"`
	PasswordHash 			 		string    `json:"-"`
	Name         			 		string    `json:"name"`
	Picture      			 		string    `json:"picture"`
	Role         			 		string    `json:"role"`
	Banner         			 		string    `json:"banner"`
	Status       			 		string    `json:"status"`
	LastLogin    			 		time.Time `json:"last_login"`
	UpdateAt 						time.Time `json:"update_at"`
	CreateAt 						time.Time `json:"create_at"`
	
	IsVerified               		bool       `json:"is_verified" db:"is_verified"`
	IsPublic             		    bool       `json:"is_public"`
    VerificationToken        		string     `json:"-" db:"verification_token"`
    VerificationTokenExpiresAt 		*time.Time `json:"-" db:"verification_token_expires_at"`

	ResetToken						string 	   `json:"-" db:"reset_token"`
	ResetTokenExpiresAt				*time.Time `json:"-" db:"reset_token_expires_at"`

	About							string		`json:"about"`
	Birthday						*time.Time	`json:"birthday"`
	Gender							UserGenere	`json:"gender"`
	Tag								string		`json:"tag"`
}




