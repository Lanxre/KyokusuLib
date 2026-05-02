package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
)

func (a *NovelaRoutes) Register(cfg *config.Config, r *mux.Router) {
	novelaRouter := r.PathPrefix("/api").Subrouter()

	novelaRouter.Handle("/novela", middleware.DefaultMiddleware(http.HandlerFunc(a.Handler.GetNovelas), cfg.JWTSecret)).Methods("GET")
	novelaRouter.Handle("/novela/{id:[0-9]+}", middleware.DefaultMiddleware(http.HandlerFunc(a.Handler.GetNovela), cfg.JWTSecret)).Methods("GET")
	novelaRouter.Handle("/novela/{id:[0-9]+}", middleware.DefaultMiddleware(http.HandlerFunc(a.Handler.UpdateNovela), cfg.JWTSecret)).Methods("PUT")

	novelaRouter.Handle("/novela", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.CreateNovela), cfg.JWTSecret)).Methods("POST")

	novelaRouter.Handle("/novela/bookmark", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.SetBookmark), cfg.JWTSecret)).Methods("POST")
	novelaRouter.Handle("/novela/bookmark/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.RemoveBookmark), cfg.JWTSecret)).Methods("DELETE")

	novelaRouter.Handle("/novela/bookmarks", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetUserNovelaBookmarks), cfg.JWTSecret)).Methods("GET")

	novelaRouter.Handle("/novela/like", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.SetLike), cfg.JWTSecret)).Methods("POST")
	novelaRouter.Handle("/novela/rating", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.SetRating), cfg.JWTSecret)).Methods("POST")

	novelaRouter.Handle("/bookmarks/categories", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.GetBookmarkCategories), cfg.JWTSecret)).Methods("GET")
	novelaRouter.Handle("/bookmarks/categories", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.CreateBookmarkCategory), cfg.JWTSecret)).Methods("POST")
	novelaRouter.Handle("/bookmarks/categories/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.UpdateBookmarkCategory), cfg.JWTSecret)).Methods("PUT")
	novelaRouter.Handle("/bookmarks/categories/{id:[0-9]+}", middleware.AuthMiddleware(http.HandlerFunc(a.Handler.DeleteBookmarkCategory), cfg.JWTSecret)).Methods("DELETE")
}
