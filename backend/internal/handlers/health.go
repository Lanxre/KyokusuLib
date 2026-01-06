package handlers

import (
	"encoding/json"
	"net/http"
	"time"
)

type HealthHandler struct{}

type HealthResponse struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	Uptime    string `json:"uptime"`
}

func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

var startTime = time.Now()

func (h *HealthHandler) HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Status:    "UP",
		Timestamp: time.Now().Format(time.RFC3339),
		Uptime:    time.Since(startTime).String(),
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
