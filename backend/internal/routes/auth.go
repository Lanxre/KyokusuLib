package routes

import (
	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *AuthRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api/auth").Subrouter()
	a.RegisterBase(s, cfg)
	a.RegisterResetPassword(s)
	a.RegisterGoogleRoutes(s)
	a.RegisterDiscordRoutes(s)
}

func (a *AuthRoutes) RegisterBase(s *mux.Router, cfg *config.Config) {
	s.HandleFunc("/me", middleware.AuthMiddleware(a.Handler.GetMe, cfg.JWTSecret)).Methods("GET")
	s.HandleFunc("/logout", a.Handler.Logout).Methods("POST")
	s.HandleFunc("/register", a.Handler.Register).Methods("POST")
	s.HandleFunc("/login", a.Handler.Login).Methods("POST")
	s.HandleFunc("/verify-email", a.Handler.VerifyEmail).Methods("GET")
}

func (a *AuthRoutes) RegisterGoogleRoutes(s *mux.Router) {
	s.HandleFunc("/google/login", a.Handler.GoogleLogin).Methods("GET")
	s.HandleFunc("/google/callback", a.Handler.GoogleCallback).Methods("GET")
}

func (a *AuthRoutes) RegisterDiscordRoutes(s *mux.Router) {
	s.HandleFunc("/discord/login", a.Handler.DiscordLogin).Methods("GET")
	s.HandleFunc("/discord/callback", a.Handler.DiscordCallback).Methods("GET")
}

func (a *AuthRoutes) RegisterResetPassword(s *mux.Router) {
	s.HandleFunc("/forgot-password", a.Handler.ForgotPassword).Methods("POST")
	s.HandleFunc("/reset-password", a.Handler.ResetPassword).Methods("POST")
}