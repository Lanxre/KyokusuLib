package db

import (
	"encoding/json"
	"time"
)

type UserActivity struct {
	ID            int     			`json:"id"`
	UserID        int     			`json:"user_id"`
	ActivityType  string    		`json:"activity_type"`
	TargetID      int     			`json:"target_id"`
	Metadata      json.RawMessage 	`json:"metadata"`
	Timestamp     time.Time 	  	`json:"timestamp"`
}