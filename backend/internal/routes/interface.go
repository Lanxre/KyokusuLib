package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

type Route interface {
	Register(cfg *config.Config, r *mux.Router)
}
