package db

import "time"

type NovelaComment struct {
	ID        int        `db:"id"`
	NovelaID  int        `db:"novela_id"`
	UserID    int        `db:"user_id"`
	ParentID  *int       `db:"parent_id"`
	Content   string     `db:"content"`
	CreatedAt time.Time  `db:"created_at"`
	UpdatedAt time.Time  `db:"updated_at"`
	UserName  string     `db:"user_name"`
	UserImage string     `db:"user_image"`
	LikeCount int        `db:"like_count"`
	HasLike   *bool       `db:"has_like"`
}