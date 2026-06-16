package app

import (
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/gorilla/mux"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/middleware"
	"github.com/lanxre/kyokusulib/internal/routes"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/fx"
)

func AsRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(routes.Route)),
		fx.ResultTags(`group:"routes"`),
	)
}


func NewValidator() *validator.Validate {
	return validator.New()
}

func NewEmailService(cfg *config.Config) *service.EmailService {
	return service.NewEmailService(cfg.KyokusuEmailName, cfg.KyokusuEmailPass)
}

func NewMuxRouter(cfg *config.Config, rts []routes.Route, log *slog.Logger, redis *redis.Client) *mux.Router {
	r := mux.NewRouter()

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()
			next.ServeHTTP(w, r)
			log.Info("request completed",
				slog.String("method", r.Method),
				slog.String("path", r.URL.Path),
				slog.String("remote_addr", r.RemoteAddr),
				slog.Duration("duration", time.Since(start)),
			)
		})
	})

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					log.Error("panic recovered",
						slog.Any("error", err),
						slog.String("stack", string(debug.Stack())),
					)
					http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				}
			}()
			next.ServeHTTP(w, r)
		})
	})

	r.Use(func(next http.Handler) http.Handler {
		protectedPaths := []string{
			"/api/auth/login",
			"/api/auth/register",
			"/api/auth/forgot-password",
			"/api/novela/comments",
		}
		rm := middleware.RateLimitMiddleware(redis, 60, time.Minute)
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method == http.MethodPost {
				for _, p := range protectedPaths {
					if len(r.URL.Path) >= len(p) && r.URL.Path[:len(p)] == p {
						rm(next).ServeHTTP(w, r)
						return
					}
				}
			}
			next.ServeHTTP(w, r)
		})
	})

	routes.RegisterRoutes(r, cfg, rts...)

	return r
}