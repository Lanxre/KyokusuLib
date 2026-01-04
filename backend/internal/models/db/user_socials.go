package db

type SocialProvider string

const (
	DISCORD_PROVIDER SocialProvider  = "discord"
	GOOGLE_PROVIDER SocialProvider   = "google"
	TELEGRAM_PROVIDER SocialProvider = "telegram"

)

type UserSocials struct {
	UserID		int `json:"-"`
	
	DiscordID  string `json:"discord_id,omitempty"`
	GoogleID   string `json:"google_id,omitempty"`

	IsDiscordConnected  bool `json:"is_discord_connected"`
	IsGoogleConnected   bool `json:"is_google_connected"`

	DiscordRefreshToken string
	GoogleRefreshToken  string
}