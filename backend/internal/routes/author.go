package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *AuthorRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()
	
	s.Handle("/author", http.HandlerFunc(a.Handler.GetAuthors)).Methods("GET")
	s.Handle("/author/{id:[0-9]+}", http.HandlerFunc(a.Handler.GetAuthorById)).Methods("GET")
	
	s.Handle("/author", middleware.AuthMiddleware(
		middleware.RoleGuard(
			http.HandlerFunc(a.Handler.CreateAuthor), middleware.RoleModerator,
		), cfg.JWTSecret)).Methods("POST")

	s.Handle("/author/{id:[0-9]+}", middleware.AuthMiddleware(
		middleware.RoleGuard(
			http.HandlerFunc(a.Handler.UpdateAuthor), middleware.RoleModerator,
		), cfg.JWTSecret)).Methods("PUT")
}