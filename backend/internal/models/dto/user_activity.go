package dto

import (
	"encoding/json"
	"time"
)

type GetUserActivity struct {
	ID            int	     		`json:"id"`
	ActivityType  string    		`json:"activity_type"`
	Metadata      json.RawMessage 	`json:"metadata"`
	Timestamp     time.Time 	  	`json:"timestamp"`
}

type CreateUserActivity struct {
	UserID        int            `json:"user_id"`
	ActivityType  string         `json:"activity_type"`
	TargetID      int            `json:"target_id"`
	Metadata      json.RawMessage `json:"metadata"`
}
