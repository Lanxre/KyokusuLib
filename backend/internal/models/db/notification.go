package db

import "time"

type Notification struct {
	ID        int64     `db:"id"`
	UserID    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Message   string    `db:"message"`
	IsRead    bool      `db:"is_read"`
	CreatedAt time.Time `db:"created_at"`
}
