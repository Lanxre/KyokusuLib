package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/handlers"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

type ParseRoutes struct {
	Handler *handlers.ParseHandler
}

func (r *ParseRoutes) Register(cfg *config.Config, router *mux.Router) {
	s := router.PathPrefix("/api/parse").Subrouter()

	s.Handle("/ranobehub/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(r.Handler.ParseRanobeHub), cfg.JWTSecret)).Methods("POST")
}
