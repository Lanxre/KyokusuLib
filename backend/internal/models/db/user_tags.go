package db

import "time"

type UserDbTag struct {
	ID     	  int       `db:"id"`
	Tag    	  string    `db:"tag"`
	CreatedAt time.Time `db:"created_at"`
}