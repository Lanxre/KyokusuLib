package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *AuthorRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()
	
	s.Handle("/author", middleware.AuthMiddleware(
		middleware.RoleGuard(
			http.HandlerFunc(a.Handler.GetAuthors), middleware.RoleModerator,
		), cfg.JWTSecret)).Methods("GET")
	
	s.Handle("/author", middleware.AuthMiddleware(
		middleware.RoleGuard(
			http.HandlerFunc(a.Handler.CreateAuthor), middleware.RoleModerator,
		), cfg.JWTSecret)).Methods("POST")
}