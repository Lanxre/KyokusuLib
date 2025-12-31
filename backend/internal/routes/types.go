package routes

import "github.com/lanxre/kyokusulib/internal/handlers"

type AuthRoutes struct {
	Handler *handlers.AuthHandler
}

type HealthRoutes struct {
	Handler *handlers.HealthHandler
}

type EmailRoutes struct {
	Handler *handlers.EmailHandler
}

type ProfileSettingRoutes struct {
	Handler *handlers.ProfileSettingHandler
}

type SocialNetworkRoutes struct {
	Handler *handlers.SocialNetworkHandler
}

type UserRoutes struct {
	Handler *handlers.UserHandler
}

type RanobeRoutes struct {
	Handler *handlers.RanobeHandler
}

type AuthorRoutes struct {
	Handler *handlers.AuthorHandler
}

type UserActivityRoutes struct {
	Handler *handlers.UserActivityHandler
}