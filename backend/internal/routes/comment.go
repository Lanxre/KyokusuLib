package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (h *CommentRoutes) Register(cfg *config.Config, r *mux.Router) {
	api := r.PathPrefix("/api/novela").Subrouter()

	api.HandleFunc("/comments", middleware.AuthMiddleware(h.Handler.CreateComment, cfg.JWTSecret)).Methods("POST")
	api.HandleFunc("/comments/{id:[0-9]+}", middleware.DefaultMiddleware(h.Handler.GetNovelaComments, cfg.JWTSecret)).Methods("GET")
	api.HandleFunc("/comments/{id:[0-9]+}", middleware.AuthMiddleware(h.Handler.DeleteComment, cfg.JWTSecret)).Methods("DELETE")
	api.HandleFunc("/comments/{id:[0-9]+}", middleware.AuthMiddleware(h.Handler.UpdateComment, cfg.JWTSecret)).Methods("PUT")
	
	api.HandleFunc("/comments/like/{id:[0-9]+}", middleware.AuthMiddleware(h.Handler.LikeComment, cfg.JWTSecret)).Methods("POST")
	api.HandleFunc("/comments/unlike/{id:[0-9]+}", middleware.AuthMiddleware(h.Handler.UnlikeComment, cfg.JWTSecret)).Methods("POST")
}