package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (n *NotificationRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/notifications").Subrouter()

	s.Handle("/stream", middleware.AuthMiddleware(n.Handler.StreamNotifications, cfg.JWTSecret)).Methods("GET")

	s.Handle("", middleware.AuthMiddleware(n.Handler.List, cfg.JWTSecret)).Methods("GET")
	
	s.Handle("/stats", middleware.AuthMiddleware(n.Handler.NotificationStats, cfg.JWTSecret)).Methods("GET")

	s.Handle("/{id:[0-9]+}/read", middleware.AuthMiddleware(n.Handler.MarkRead, cfg.JWTSecret)).Methods("PATCH")
	s.Handle("/read-all", middleware.AuthMiddleware(n.Handler.MarkAllRead, cfg.JWTSecret)).Methods("PATCH")
	s.Handle("/{id:[0-9]+}", middleware.AuthMiddleware(n.Handler.Delete, cfg.JWTSecret)).Methods("DELETE")
}
