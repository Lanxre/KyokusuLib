package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *NovelaRoutes) Register(cfg *config.Config, r *mux.Router) {
	novelaRouter := r.PathPrefix("/api").Subrouter()
	
	novelaRouter.Handle("/novela/{id:[0-9]+}", http.HandlerFunc(a.Handler.GetNovela)).Methods("GET")
	novelaRouter.Handle("/novela", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.CreateNovela), cfg.JWTSecret)).Methods("POST")

	novelaRouter.Handle("/novela/bookmark", http.HandlerFunc(a.Handler.SetBookmark)).Methods("POST")
	novelaRouter.Handle("/novela/bookmark/{id:[0-9]+}", http.HandlerFunc(a.Handler.RemoveBookmark)).Methods("DELETE")
}