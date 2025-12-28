package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (a *EmailRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/email").Subrouter()

	// send email
	s.HandleFunc("/", a.Handler.Send).Methods("POST")
}
