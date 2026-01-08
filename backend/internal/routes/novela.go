package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *NovelaRoutes) Register(cfg *config.Config, r *mux.Router) {
	ranobeRouter := r.PathPrefix("/api").Subrouter()
	
	ranobeRouter.Handle("/novela/{id:[0-9]+}", http.HandlerFunc(a.Handler.GetNovela)).Methods("GET")
	ranobeRouter.Handle("/novela", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.CreateNovela), cfg.JWTSecret)).Methods("POST")
}