package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/handlers"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

type TeamRoutes struct {
	Handler *handlers.TeamHandler
}

func (tr *TeamRoutes) Register(cfg *config.Config, r *mux.Router) {
	api := r.PathPrefix("/api/teams").Subrouter()

	api.HandleFunc("", middleware.AuthMiddleware(tr.Handler.Create, cfg.JWTSecret)).Methods("POST")
	api.HandleFunc("", tr.Handler.GetTeams).Methods("GET")
	api.HandleFunc("/{slug}", middleware.DefaultMiddleware(tr.Handler.Get, cfg.JWTSecret)).Methods("GET")
	api.HandleFunc("/{slug}/members", tr.Handler.GetMembers).Methods("GET")
	api.HandleFunc("/{slug}/subscribers", tr.Handler.GetSubscribers).Methods("GET")
	api.HandleFunc("/{slug}", middleware.AuthMiddleware(tr.Handler.Update, cfg.JWTSecret)).Methods("PATCH")
	api.HandleFunc("/{slug}/join", middleware.AuthMiddleware(tr.Handler.Join, cfg.JWTSecret)).Methods("POST")
	api.HandleFunc("/{slug}/leave", middleware.AuthMiddleware(tr.Handler.Leave, cfg.JWTSecret)).Methods("POST")
	
	api.HandleFunc("/subscribe", middleware.AuthMiddleware(tr.Handler.Subscribe, cfg.JWTSecret)).Methods("POST")
	api.HandleFunc("/unsubscribe", middleware.AuthMiddleware(tr.Handler.Unsubscribe, cfg.JWTSecret)).Methods("POST")

}