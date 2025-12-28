package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (a *UserRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/user").Subrouter()

	s.HandleFunc("/{id:[0-9]+}", a.Handler.GetUserById).Methods("GET")
}
