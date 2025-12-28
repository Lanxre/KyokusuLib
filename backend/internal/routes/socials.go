package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *SocialNetworkRoutes) Register(cfg *config.Config, r *mux.Router) {
	s := r.PathPrefix("/api").Subrouter()

	s.Handle("/socials", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetSocials), cfg.JWTSecret)).Methods("GET")
	
	s.Handle("/socials/google/login", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GoogleLogin), cfg.JWTSecret)).Methods("GET")
	s.Handle("/socials/google/callback", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GoogleCallback), cfg.JWTSecret)).Methods("GET")
	s.Handle("/socials/google/logout", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.UnlinkGoogle), cfg.JWTSecret)).Methods("GET")

	s.Handle("/socials/discord/login", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DiscordLogin), cfg.JWTSecret)).Methods("GET")
	s.Handle("/socials/discord/callback", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DiscordCallback), cfg.JWTSecret)).Methods("GET")
	s.Handle("/socials/discord/logout", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.UnlinkDiscord), cfg.JWTSecret)).Methods("GET")

}