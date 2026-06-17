package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
)

func (a *NovelaStatisticsRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()

	s.Handle("/novelas/statistics/total", http.HandlerFunc(a.Handler.GetTotalNovelaStatistics)).Methods("GET")
	s.Handle("/novelas/statistics/general", http.HandlerFunc(a.Handler.GetGeneralStatistics)).Methods("GET")
	s.Handle("/novelas/statistics/monthly", http.HandlerFunc(a.Handler.GetMonthlyStatistics)).Methods("GET")
}