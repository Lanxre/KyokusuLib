package repository

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
)

type UserProfileSettingRepository struct {
	DB *sql.DB
}

func NewUserProfileSettingRepository(db *sql.DB) *UserProfileSettingRepository {
	return &UserProfileSettingRepository{DB: db}
}

func (r *UserProfileSettingRepository) GetUserProfileSettings(ctx context.Context, userID int) (*db.UserProfileSetting, error) {
	settings := &db.UserProfileSetting{}

	query := `
		SELECT theme, is_app_notify, is_new_published_notify, is_show_tag, is_show_bookmark
		FROM user_profile_settings
		WHERE user_id = $1`

	err := r.DB.QueryRowContext(ctx, query, userID).Scan(
		&settings.Theme,
		&settings.IsAppNotify,
		&settings.IsNewPublishedNotify,
		&settings.IsShowTag,
		&settings.IsShowBookmark,
	)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

func (r *UserProfileSettingRepository) UpsertUserInterfaceSettings(ctx context.Context, userID int, settings *db.UserProfileSetting) error {
	query := `
		INSERT INTO user_profile_settings (user_id, theme)
		VALUES ($1, $2)
		ON CONFLICT (user_id) 
		DO UPDATE SET theme = EXCLUDED.theme`

	_, err := r.DB.ExecContext(ctx, query, userID, settings.Theme)
	return err
}

func (r *UserProfileSettingRepository) DeleteUserProfileSettings(ctx context.Context, userID int) error {
	query := `DELETE FROM user_profile_settings WHERE user_id = $1`
	_, err := r.DB.ExecContext(ctx, query, userID)
	return err
}

func (r *UserProfileSettingRepository) UpdateUserInterfaceSettings(ctx context.Context, userID int, settings dto.UserInterfacePatchDTO) error {
	var (
		setClauses []string
		args       []any
		argID      = 1
	)
	
	if settings.IsShowTag != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_show_tag = $%d", argID))
		args = append(args, *settings.IsShowTag)
		argID++
	}

	if settings.IsShowBookmark != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_show_bookmark = $%d", argID))
		args = append(args, *settings.IsShowBookmark)
		argID++
	}

	if len(setClauses) == 0 {
		return nil
	}

	args = append(args, userID)

	query := fmt.Sprintf(
		"UPDATE user_profile_settings SET %s WHERE user_id = $%d",
		strings.Join(setClauses, ", "),
		argID,
	)

	result, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	
	rowsAffected, err := result.RowsAffected()
	
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}

func (r *UserProfileSettingRepository) UpdateNotifySettings(ctx context.Context, userID int, settings dto.NotifySettingsPatchDTO) error {
	var (
		setClauses []string
		args       []any
		argID      = 1
	)
	
	if settings.IsAppNotify != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_app_notify = $%d", argID))
		args = append(args, *settings.IsAppNotify)
		argID++
	}

	if settings.IsNewPublishedNotify != nil {
		setClauses = append(setClauses, fmt.Sprintf("is_new_published_notify = $%d", argID))
		args = append(args, *settings.IsNewPublishedNotify)
		argID++
	}

	if len(setClauses) == 0 {
		return nil
	}

	args = append(args, userID)

	query := fmt.Sprintf(
		"UPDATE user_profile_settings SET %s WHERE user_id = $%d",
		strings.Join(setClauses, ", "),
		argID,
	)

	result, err := r.DB.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}

	
	rowsAffected, err := result.RowsAffected()
	
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return nil
	}

	return nil
}