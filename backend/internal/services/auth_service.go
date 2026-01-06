package service

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"log"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	Repo *repository.UserRepository
	UserProfileRepo *repository.UserProfileRepository
}

func NewAuthService(repo *repository.UserRepository, userProfileRepo *repository.UserProfileRepository) *AuthService {
	return &AuthService{Repo: repo, UserProfileRepo: userProfileRepo}
}

func (s *AuthService) RegisterUser(dto *dto.RegisterDTO) (*db.User, error) {
	if existing, _ := s.Repo.GetByEmail(dto.Email); existing != nil {
		return nil, errors.New("user already exists")
	}

	hashed, err := s.hashPassword(dto.Password)
	if err != nil {
		return nil, err
	}

	token, _ := s.generateToken()
	expiresAt := time.Now().Add(24 * time.Hour)

	user := &db.User{
		Email:                      dto.Email,
		Name:                       dto.Username,
		PasswordHash:               hashed,
		Role:                       "user",
		Status:                     "online",
		IsVerified:                 false,
		VerificationToken:          token,
		VerificationTokenExpiresAt: &expiresAt,
		Gender:                     db.HIDDEN_GENERE,
	}
	
	if err := s.Repo.Create(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (s *AuthService) LoginWithPassword(dto *dto.LoginDTO) (*dto.GetUserDTO, error) {
	user, err := s.Repo.GetByEmail(dto.Email)
	if err != nil || user == nil {
		return nil, errors.New("invalid email or password")
	}
	
	if !user.IsVerified {
		return nil, errors.New("email not verified")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(dto.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}
	
	userTags, err := s.GetUserTags(user.ID)
	if err != nil {
		return nil, err
	}
	
	userLevel, err := s.UserProfileRepo.GetUserLevel(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}
	
	userDTO := s.CreateUserDTO(user, userTags, userLevel)
		
	_ = s.Repo.UpdateStatus(user.ID, "online")
	user.Status = "online"
	
	return userDTO, nil
}

func (s AuthService) CreateUserDTO(user *db.User, userTags []dto.UserTagDTO,userLevel *db.UserLevel) *dto.GetUserDTO {
	return  &dto.GetUserDTO{
		ID:       user.ID,
		Email:    user.Email,
		Name:     user.Name,
		Picture:  user.Picture,
		Role:     user.Role,
		Status:   user.Status,
		About:    user.About,
		Birthday: user.Birthday,
		Gender:   string(user.Gender),
		IsPublic: user.IsPublic,
		LastLogin: user.LastLogin,
		CreateAt: user.CreateAt,
		Banner:   user.Banner,
		ActiveTag:      user.Tag,
		AllTags:        userTags,
		Settings:       dto.PublicUserSettingsDTO{
						IsShowTag: user.IsShowTag,
		},
		UserLevel:      dto.UserLevelDTO{
			Level:      userLevel.Level,
			Experience: userLevel.Experience,
			LevelTitle: userLevel.LevelTitle,
			XPForNext:  userLevel.XPForNext,
		},
	}
}

func (s *AuthService) LoginUser(gUser *dto.UserDTO) (*dto.GetUserDTO, error) {
	user, _ := s.Repo.GetByEmail(gUser.Email)
	if user == nil {
		user = &db.User{
			Email:      gUser.Email,
			Name:       gUser.Username,
			Picture:    gUser.Avatar,
			Role:       "user",
			Status:     "online",
			IsVerified: true,
		}
		if err := s.Repo.Create(user); err != nil {
			return nil, err
		}
		return s.CreateUserDTO(user, nil, nil), nil
	}

	if !user.IsVerified {
		_ = s.Repo.MarkUserVerified(user.ID)
		user.IsVerified = true
	}

	if user.Picture == "" && gUser.Avatar != "" {
		s.Repo.UpdateAvatar(user.ID, gUser.Avatar)
	}
	
	userTags, err := s.GetUserTags(user.ID)
	if err != nil {
		return nil, err
	}
	
	userLevel, err := s.UserProfileRepo.GetUserLevel(context.Background(), user.ID)
	if err != nil {
		return nil, err
	}
	_ = s.Repo.UpdateStatus(user.ID, "online")
	user.Status = "online"
	
	userDTO := s.CreateUserDTO(user, userTags, userLevel)

	return userDTO, nil
}

func (s *AuthService) RequestPasswordReset(email string) (string, error) {
	user, err := s.Repo.GetByEmail(email)
	if err != nil || user == nil {
		return "", errors.New("user not found")
	}

	token, _ := s.generateToken()
	expiresAt := time.Now().Add(1 * time.Hour)

	user.ResetToken = token
	user.ResetTokenExpiresAt = &expiresAt

	if err := s.Repo.Update(user); err != nil {
		return "", err
	}

	return token, nil
}

func (s *AuthService) ResetPassword(token, newPassword string) error {
	user, err := s.Repo.GetByResetToken(token)
	if err != nil || user == nil {
		return errors.New("invalid or expired token")
	}

	if user.ResetTokenExpiresAt != nil && user.ResetTokenExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	hashed, err := s.hashPassword(newPassword)
	if err != nil {
		return err
	}

	user.PasswordHash = hashed
	user.ResetToken = ""
	user.ResetTokenExpiresAt = nil

	return s.Repo.Update(user)
}

func (s *AuthService) VerifyUser(token string) error {
	user, err := s.Repo.GetByVerificationToken(token)
	if err != nil {
		return errors.New("invalid token")
	}

	if user.VerificationTokenExpiresAt != nil && user.VerificationTokenExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	return s.Repo.MarkUserVerified(user.ID)
}

func (s *AuthService) UpdateStatus(ctx context.Context, userID int, status bool) error {
	user, err := s.Repo.GetByID(userID)
	if err != nil || user == nil {
		return errors.New("user not found")
	}
	
	if status {
		err = s.Repo.UpdateStatus(userID, "online")
	} else {
		err = s.Repo.UpdateStatus(userID, "offline")
	}

	return err
}

func (s *AuthService) GetUserTags(userId int) ([]dto.UserTagDTO, error) {
	tags, err := s.Repo.GetUserTags(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	
	userTags := make([]dto.UserTagDTO, len(tags))
	for i, tag := range tags {
		userTags[i] = dto.UserTagDTO{
			ID:    tag.TagID,
			Tag:   tag.Tag,
		}
	}
	
	return userTags, nil
}


func (s *AuthService) hashPassword(pwd string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(b), err
}

func (s *AuthService) generateToken() (string, error) {
	b := make([]byte, 32)
	rand.Read(b)
	return hex.EncodeToString(b), nil
}

func (s *AuthService) StartCleanupWorker(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(24 * time.Hour)
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				log.Println("User cleanup worker stopped")
				return
			case <-ticker.C:
				log.Println("Running cleanup task...")
				if err := s.Repo.DeleteExpiredUnverified(); err != nil {
					log.Printf("Error cleaning up users: %v", err)
				}
			}
		}
	}()
}