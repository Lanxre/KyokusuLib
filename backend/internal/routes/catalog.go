package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *CatalogRoutes) Register(cfg *config.Config, r *mux.Router) {
	catalogRouter := r.PathPrefix("/api").Subrouter()

	catalogRouter.Handle("/catalog/filters", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.SaveFilters), cfg.JWTSecret)).Methods("POST")
	catalogRouter.Handle("/catalog/filters", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetUserFilters), cfg.JWTSecret)).Methods("GET")
	catalogRouter.Handle("/catalog/filters/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetFilterByID), cfg.JWTSecret)).Methods("GET")
	catalogRouter.Handle("/catalog/filters/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DeleteFilter), cfg.JWTSecret)).Methods("DELETE")
}
