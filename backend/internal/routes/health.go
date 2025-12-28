package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (h *HealthRoutes) Register(cfg *config.Config, r *mux.Router) {
	r.HandleFunc("/api/health", h.Handler.HealthCheckHandler).Methods("GET")
}
