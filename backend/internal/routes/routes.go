package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func RegisterRoutes(r *mux.Router, cfg *config.Config, routes ...Route) {
    for _, route := range routes {
        route.Register(cfg, r)
    }
}