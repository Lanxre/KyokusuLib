package response

import (
	"encoding/json"
	"net/http"
)

func JSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}

func Error(w http.ResponseWriter, status int, message string) {
	JSON(w, status, map[string]string{"error": message})
}

func ErrorWithDetails(w http.ResponseWriter, status int, entity any) {
	JSON(w, status, entity)
}

func Success(w http.ResponseWriter, status int, entity any) {
	JSON(w, status, map[string]any{"message": entity})
}
