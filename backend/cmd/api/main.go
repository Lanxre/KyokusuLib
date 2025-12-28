package main

import (
	"context"
	"log"
	"net/http"

	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/handlers"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/routes"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/storage"
	"github.com/lanxre/kyokusulib/internal/utils/static"
	"github.com/rs/cors"
)

func main() {

	cfg := config.Load()

	db := storage.NewPostgresDB(cfg.DatabaseURL)
	defer db.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	userRepo := repository.NewUserRepository(db)
	userProfileSettingsRepo := repository.NewUserProfileSettingRepository(db)
	userSocialRepo := repository.NewUserSocialsRepository(db)

	authService := service.NewAuthService(userRepo)
	userService := service.NewUserService(userRepo)
	emailService := service.NewEmailService(cfg.KyokusuEmailName, cfg.KyokusuEmailPass)
	profileSettingService := service.NewProfileSettingService(userRepo, userProfileSettingsRepo)
	socialService := service.NewSocialsService(userSocialRepo)

	authService.StartCleanupWorker(ctx)

	healthHandler := &handlers.HealthHandler{}
	authHandler := handlers.NewAuthHandler(cfg, authService, emailService, socialService)
	userHandler := handlers.NewUserHandler(userService)
	emailHandler := handlers.NewEmailHandler(cfg, emailService)
	profileSettingHandler := handlers.NewProfileSettingHandler(profileSettingService)
	socialNetworkHandler := handlers.NewSocialNetworkHandler(cfg, socialService)

	rts := []routes.Route{
		&routes.HealthRoutes{Handler: healthHandler},
		&routes.AuthRoutes{Handler: authHandler},
		&routes.EmailRoutes{Handler: emailHandler},
		&routes.ProfileSettingRoutes{Handler: profileSettingHandler},
		&routes.SocialNetworkRoutes{Handler: socialNetworkHandler},
		&routes.UserRoutes{Handler: userHandler},
	}

	r := routes.NewRouter(cfg, rts...)

	static.CreateStaticDirs(r)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "X-Requested-With"},
		AllowCredentials: true,
		Debug:            true,
	})

	handler := c.Handler(r)

	srv := &http.Server{
		Handler:      handler,
		Addr:         cfg.Address,
		WriteTimeout: cfg.WriteTimeout,
		ReadTimeout:  cfg.ReadTimeout,
	}

	log.Printf("Server starting on %s ...", cfg.Address)
	if err := srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}