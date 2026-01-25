package main

import (
	"github.com/lanxre/kyokusulib/internal/app"
	"github.com/lanxre/kyokusulib/internal/config"
	"github.com/lanxre/kyokusulib/internal/handlers"
	"github.com/lanxre/kyokusulib/internal/lib/logger"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/routes"
	service "github.com/lanxre/kyokusulib/internal/services"
	"github.com/lanxre/kyokusulib/internal/storage"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
)

func main() {
	fx.New(
		fx.Provide(
			config.Load,
			logger.New,			
			storage.NewPostgresDB,
			app.NewValidator,
			

			repository.NewUserRepository,
			repository.NewUserProfileSettingRepository,
			repository.NewUserSocialsRepository,
			repository.NewUserActivityRepository,
			repository.NewUserProfileRepository,
			repository.NewAuthorRepository,
			repository.NewNovelaRatingRepository,
			repository.NewNovelaBookmarkRepository,
			repository.NewNovelaLikeRepository,
			repository.NewCommentsRepository,
			repository.NewNovelaRepository,
			repository.NewTeamRepository,

			service.NewAuthService,
			service.NewUserService,
			service.NewUserActivityService,
			service.NewProfileSettingService,
			service.NewSocialsService,
			service.NewAuthorService,
			service.NewNovelaService,
			service.NewCommentService,
			service.NewTeamService,
			app.NewEmailService,

			handlers.NewAuthHandler,
			handlers.NewHealthHandler,
			handlers.NewUserHandler,
			handlers.NewEmailHandler,
			handlers.NewProfileSettingHandler,
			handlers.NewSocialNetworkHandler,
			handlers.NewAuthorHandler,
			handlers.NewUserActivityHandler,
			handlers.NewNovelaHandler,
			handlers.NewCommentHandler,
			handlers.NewTeamHandler,

			app.AsRoute(func(h *handlers.HealthHandler) *routes.HealthRoutes { return &routes.HealthRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.AuthHandler) *routes.AuthRoutes { return &routes.AuthRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.EmailHandler) *routes.EmailRoutes { return &routes.EmailRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.ProfileSettingHandler) *routes.ProfileSettingRoutes {
				return &routes.ProfileSettingRoutes{Handler: h}
			}),
			app.AsRoute(func(h *handlers.SocialNetworkHandler) *routes.SocialNetworkRoutes {
				return &routes.SocialNetworkRoutes{Handler: h}
			}),
			app.AsRoute(func(h *handlers.UserHandler) *routes.UserRoutes { return &routes.UserRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.AuthorHandler) *routes.AuthorRoutes { return &routes.AuthorRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.UserActivityHandler) *routes.UserActivityRoutes {
				return &routes.UserActivityRoutes{Handler: h}
			}),
			app.AsRoute(func(h *handlers.NovelaHandler) *routes.NovelaRoutes { return &routes.NovelaRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.TeamHandler) *routes.TeamRoutes { return &routes.TeamRoutes{Handler: h} }),
			app.AsRoute(func(h *handlers.CommentHandler) *routes.CommentRoutes { return &routes.CommentRoutes{Handler: h} }),

			fx.Annotate(
				app.NewMuxRouter,
				fx.ParamTags("", `group:"routes"`),
			),
		),

		fx.Invoke(
			app.RegisterStaticFiles,
			app.StartBackgroundWorkers,
			app.StartHTTPServer,
		),
		fx.WithLogger(func() fxevent.Logger { return fxevent.NopLogger }),
	).Run()
}