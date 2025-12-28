package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func NewRouter(cfg *config.Config, routeGroups ...Route) *mux.Router {
	r := mux.NewRouter()

	for _, group := range routeGroups {
		group.Register(cfg, r)
	}

	return r
}
