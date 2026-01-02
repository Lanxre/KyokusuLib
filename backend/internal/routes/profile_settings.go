package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *ProfileSettingRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/profile").Subrouter()

	s.HandleFunc("/update", middleware.AuthMiddleware(a.Handler.UpdateProfile, cfg.JWTSecret)).Methods("PUT", "POST")
	
	s.Handle("/avatar", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.UpdateAvatar), cfg.JWTSecret)).Methods("POST")
	s.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.UpdateBanner), cfg.JWTSecret)).Methods("POST")
    s.Handle("/avatar", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DeleteAvatar), cfg.JWTSecret)).Methods("DELETE")
    s.Handle("/banner", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DeleteBanner), cfg.JWTSecret)).Methods("DELETE")

	s.HandleFunc("/password", middleware.AuthMiddleware(a.Handler.ChangePassword, cfg.JWTSecret)).Methods("POST")
	s.HandleFunc("/email", middleware.AuthMiddleware(a.Handler.ChangeEmail, cfg.JWTSecret)).Methods("POST")

	s.HandleFunc("/settings", middleware.AuthMiddleware(a.Handler.GetProfileSettings, cfg.JWTSecret)).Methods("GET")
	s.HandleFunc("/settings", middleware.AuthMiddleware(a.Handler.UpdateProfileSettings, cfg.JWTSecret)).Methods("PUT")
	s.HandleFunc("/settings/notify", middleware.AuthMiddleware(a.Handler.GetNotifySettings, cfg.JWTSecret)).Methods("GET")
	s.HandleFunc("/settings/notify", middleware.AuthMiddleware(a.Handler.UpdateNotifySettings, cfg.JWTSecret)).Methods("PATCH")
	s.HandleFunc("/settings/interface", middleware.AuthMiddleware(a.Handler.UpdateUserInterface, cfg.JWTSecret)).Methods("PATCH")
	
}
