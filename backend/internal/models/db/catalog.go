package db

import (
	"encoding/json"
	"time"
)

type CreateCatalogFilterPreset struct {
	ID        int             
	UserID    int             
	Name      string          
	Filters   json.RawMessage 
	CreatedAt time.Time       
}

type GetCatalogFilterPreset struct {
	ID        int
	Name      string          
	Filters   json.RawMessage 
	CreatedAt time.Time  
}