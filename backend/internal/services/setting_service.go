package service

import (
	"context"
	"database/sql"
	"errors"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/utils/files"
	"golang.org/x/crypto/bcrypt"
)

type ProfileSettingService struct {
	Repo                *repository.UserRepository
	ProfileSettingsRepo *repository.UserProfileSettingRepository
}

func NewProfileSettingService(repo *repository.UserRepository, ps *repository.UserProfileSettingRepository) *ProfileSettingService {
	return &ProfileSettingService{
		Repo:                repo,
		ProfileSettingsRepo: ps,
	}
}

func (s *ProfileSettingService) UpdateProfile(ctx context.Context, userID int, input dto.UpdateProfileDTO) error {
	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	user.Name = input.Nickname
	user.About = input.About
	user.Gender = db.UserGenere(input.Gender)
	user.IsPublic = input.IsPublic
	if input.Birthday != "" {
		parsed, _ := time.Parse("2006-01-02", input.Birthday)
		user.Birthday = &parsed
	} else {
		user.Birthday = nil
	}
	return s.Repo.UpdateProfile(user)
}

func (s *ProfileSettingService) DeleteOldAvatar(ctx context.Context, userID int) error {
	avatarURL, err := s.Repo.GetAvatarUrl(userID)
	if err == nil && strings.HasPrefix(avatarURL, "/uploads") {
		os.Remove(filepath.Join("./", avatarURL))
		s.Repo.UpdateAvatar(userID, "")
	}

	return err
}

func (s *ProfileSettingService) DeleteOldBanner(ctx context.Context, userID int) error {
	bannerURL, err := s.Repo.GetBannerUrl(userID)
	if err == nil && strings.HasPrefix(bannerURL, "/uploads") {
		os.Remove(filepath.Join("./", bannerURL))
		s.Repo.UpdateBanner(userID, "")
	}

	return err
}

func (s *ProfileSettingService) UploadAndSetAvatar(ctx context.Context, userID int, file multipart.File, header *multipart.FileHeader) (string, error) {
	avatarURL, err := files.UploadImage(ctx, file, header, "avatars", 400, 400)
    if err != nil {
        return "", err
    }
    

	err = s.Repo.UpdateAvatar(userID, avatarURL)
	if err != nil {
		os.Remove("." + avatarURL)
		return "", err
	}

	return avatarURL, nil
}

func (s *ProfileSettingService) UploadAndSetBanner(ctx context.Context, userID int, file multipart.File, header *multipart.FileHeader) (string, error) {
	bannerURL, err := files.UploadImage(ctx, file, header, "banners", 1200, 320)
	if err != nil {
		return "", err
	}

	err = s.Repo.UpdateBanner(userID, bannerURL)
	if err != nil {
		os.Remove("." + bannerURL)
		return "", err
	}

	return bannerURL, nil
}

func (s *ProfileSettingService) ChangePassword(ctx context.Context, userID int, oldPassword, newPassword string) error {
	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.PasswordHash = string(hashedPassword)

	return s.Repo.Update(user)
}

func (s *ProfileSettingService) ChangeEmail(ctx context.Context, userID int, oldPassword, email string) error {
	user, err := s.Repo.GetByID(userID)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(oldPassword)); err != nil {
		return errors.New("invalid old password")
	}

	user.Email = email

	return s.Repo.Update(user)
}

func (s *ProfileSettingService) GetSettings(ctx context.Context, userID int) (*db.UserProfileSetting, error) {
	settings, err := s.ProfileSettingsRepo.GetUserProfileSettings(ctx, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			defaultSettings := &db.UserProfileSetting{
				Theme: string(db.DarkTheme),
			}

			if err := s.ProfileSettingsRepo.UpsertUserInterfaceSettings(ctx, userID, defaultSettings); err != nil {
				return nil, err
			}

			return defaultSettings, nil
		}
		return nil, err
	}

	return settings, nil
}

func (s *ProfileSettingService) UpdateTheme(ctx context.Context, userID int, theme string) error {
	if theme != string(db.DarkTheme) && theme != string(db.LightTheme) {
		return errors.New("invalid theme value")
	}

	settings := &db.UserProfileSetting{
		Theme: theme,
	}

	return s.ProfileSettingsRepo.UpsertUserInterfaceSettings(ctx, userID, settings)
}

func (s *ProfileSettingService) GetNotifySettings(ctx context.Context, userID int) (*dto.NotifySettingsPatchDTO, error) {
	settings, err := s.GetSettings(ctx, userID)
	if err != nil {
		return nil, err
	}

	ps := &dto.NotifySettingsPatchDTO{
		IsAppNotify:          &settings.IsAppNotify,
		IsNewPublishedNotify: &settings.IsNewPublishedNotify,
	}

	return ps, nil

}

func (s *ProfileSettingService) UpdateNotifySettings(ctx context.Context, userID int, settings dto.NotifySettingsPatchDTO) error {
	return s.ProfileSettingsRepo.UpdateNotifySettings(ctx, userID, settings)
}

func (s *ProfileSettingService) ResetSettings(ctx context.Context, userID int) error {
	return s.ProfileSettingsRepo.DeleteUserProfileSettings(ctx, userID)
}

func (s *ProfileSettingService) UpdateInterfaceSettings(ctx context.Context, userID int, settings dto.UserInterfacePatchDTO) error {
	return s.ProfileSettingsRepo.UpdateUserInterfaceSettings(ctx, userID, settings)
}

func (s *ProfileSettingService) UpdateReaderSettings(ctx context.Context, userID int, settings dto.ReaderSettingsPatchDTO) error {
	return s.ProfileSettingsRepo.UpdateReaderSettings(ctx, userID, settings)
}
