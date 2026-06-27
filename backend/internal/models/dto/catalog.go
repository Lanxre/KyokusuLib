package dto

import (
	"encoding/json"
	"time"
)

type CatalogFilters struct {
	Search 	 		    string 		`json:"search"`
	Genres 	   		  []string	 	`json:"genres"`
	Categories 		  []string		`json:"categories"`
	Type 	 		    string		`json:"type"`
	Status 	 		    string		`json:"status"`
	TranslationStatus	string		`json:"translation_status"`
	ChaptersFrom 		int			`json:"chapters_from"`
	ChaptersTo 			int			`json:"chapters_to"`
	Sort 				NovelaSort  `json:"sort"`
	Limit 				int			`json:"limit"`
	Offset 				int			`json:"offset"`
}

type SaveFilterPresetRequest struct {
	Name    string          `json:"name" validate:"required,min=1,max=100"`
	Filters json.RawMessage `json:"filters" validate:"required"`
}

type CatalogFilterPreset struct {
	ID        int              `json:"id"`
	UserID    int              `json:"user_id"`
	Name      string           `json:"name"`
	Filters   json.RawMessage  `json:"filters"`
	CreatedAt time.Time        `json:"created_at"`
}