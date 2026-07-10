package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *UserTagRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/user/tags").Subrouter()
	
	s.HandleFunc("", middleware.AuthMiddleware(middleware.RoleGuard(a.Handler.GetTags, middleware.RoleModerator), cfg.JWTSecret)).Methods("GET")
	s.HandleFunc("", middleware.AuthMiddleware(middleware.RoleGuard(a.Handler.UpdateUserTags, middleware.RoleModerator), cfg.JWTSecret)).Methods("PUT")
}