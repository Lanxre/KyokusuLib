package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (a *RanobeRoutes) Register(cfg *config.Config, r *mux.Router) {
	ranobeRouter := r.PathPrefix("/api/ranobe").Subrouter()
	
	ranobeRouter.Handle("/{id:[0-9]+}", http.HandlerFunc(a.Handler.GetRanobe)).Methods("GET")
}