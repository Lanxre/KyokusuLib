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
	api.HandleFunc("/{slug}", tr.Handler.Get).Methods("GET")
	api.HandleFunc("/{slug}", middleware.AuthMiddleware(tr.Handler.Update, cfg.JWTSecret)).Methods("PATCH")
	api.HandleFunc("/{slug}/join", middleware.AuthMiddleware(tr.Handler.Join, cfg.JWTSecret)).Methods("POST")
}