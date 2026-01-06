package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (a *NovelaRoutes) Register(cfg *config.Config, r *mux.Router) {
	ranobeRouter := r.PathPrefix("/api/novela").Subrouter()
	
	ranobeRouter.Handle("/{id:[0-9]+}", http.HandlerFunc(a.Handler.GetNovela)).Methods("GET")
}