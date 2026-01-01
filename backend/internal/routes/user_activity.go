package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *UserActivityRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/user").Subrouter()
	
	s.Handle("/activities", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetUserActivities), cfg.JWTSecret)).Methods("GET")
	s.HandleFunc("/activities/{id:[0-9]+}", a.Handler.GetUserActivityById).Methods("GET")
	
	s.Handle("/activities", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.CreateUserActivity), cfg.JWTSecret)).Methods("POST")
}