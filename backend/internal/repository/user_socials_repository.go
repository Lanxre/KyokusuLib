package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserSocialsRepository struct {
	DB *sql.DB
}

func NewUserSocialsRepository(db *sql.DB) *UserSocialsRepository {
	return &UserSocialsRepository{DB: db}
}

func (r *UserSocialsRepository) GetUserSocials(ctx context.Context, userID int) (*db.UserSocials, error) {
	socials := &db.UserSocials{UserID: userID}

	var (
		discordID  sql.NullString
		googleID   sql.NullString
		telegramID sql.NullString
		googleRefreshToken sql.NullString
		discordRefreshToken sql.NullString
	)

	query := `
		SELECT discord_id, google_id, telegram_id, is_discord_connected, is_google_connected, is_telegram_connected, discord_refresh_token, google_refresh_token
		FROM user_socials
		WHERE user_id = $1`

	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&discordID, 
		&googleID, 
		&telegramID, 
		&socials.IsDiscordConnected, 
		&socials.IsGoogleConnected, 
		&socials.IsTelegramConnected,
		&discordRefreshToken,
		&googleRefreshToken,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return socials, nil
		}
		return nil, err
	}

	if discordID.Valid {
		socials.DiscordID = discordID.String
	}
	if googleID.Valid {
		socials.GoogleID = googleID.String
	}
	if telegramID.Valid {
		socials.TelegramID = telegramID.String
	}

	if discordRefreshToken.Valid {
		socials.DiscordRefreshToken = discordRefreshToken.String
	}
	
	if googleRefreshToken.Valid {
		socials.GoogleRefreshToken = googleRefreshToken.String
	}
	
	return socials, nil
}

func (r *UserSocialsRepository) LinkSocial(ctx context.Context, userID int, provider db.SocialProvider, socialID, refresh_token string) error {
	var column string
	var is_connected string
	var token string
	switch provider {
	case db.DISCORD_PROVIDER:
		column = "discord_id"
		is_connected = "is_discord_connected"
		token = "discord_refresh_token"
	case db.GOOGLE_PROVIDER:
		column = "google_id"
		is_connected = "is_google_connected"
		token = "google_refresh_token"
	case db.TELEGRAM_PROVIDER:
		column = "telegram_id"
		is_connected = "is_telegram_connected"
	default:
		return errors.New("unknown provider")
	}

	query := fmt.Sprintf(`
		INSERT INTO user_socials (user_id, %s, %s, %s)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (user_id) 
		DO UPDATE SET %s = EXCLUDED.%s, %s = EXCLUDED.%s, %s = EXCLUDED.%s`, 
		column, 
		is_connected, 
		token, 
		column, 
		column, 
		is_connected, 
		is_connected, 
		token, 
		token,
	)

	_, err := r.DB.ExecContext(ctx, query, userID, socialID, true, refresh_token)
	return err
}

func (r *UserSocialsRepository) UnlinkSocial(ctx context.Context, userID int, provider db.SocialProvider) error {
	var is_connected string
	var token string
	switch provider {
	case db.DISCORD_PROVIDER:
		is_connected = "is_discord_connected"
		token = "discord_refresh_token"
	case db.GOOGLE_PROVIDER:
		is_connected = "is_google_connected"
		token = "google_refresh_token"
	case db.TELEGRAM_PROVIDER:
		is_connected = "is_telegram_connected"
	default:
		return errors.New("unknown provider")
	}

	query := fmt.Sprintf(`UPDATE user_socials SET %s = FALSE, %s = NULL WHERE user_id = $1`, is_connected, token)
	_, err := r.DB.ExecContext(ctx, query, userID)
	return err
}