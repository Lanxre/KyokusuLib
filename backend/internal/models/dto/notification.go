package dto

import "time"

type Notification struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"userId"`
	Title     string    `json:"title"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"isRead"`
	CreatedAt time.Time `json:"createdAt"`
}

type NotificationStats struct {
	UnreadCount int64 	`json:"unreadCount"`
}
