package service

import (
	"context"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
	UserProfileRepo *repository.UserProfileRepository
}

func NewUserService(repo *repository.UserRepository, userProfileRepo *repository.UserProfileRepository) *UserService {
	return &UserService{Repo: repo, UserProfileRepo: userProfileRepo}
}

func (s *UserService) GetUserById(userId int) (*dto.GetUserDTO, error) {
	userDb, err := s.Repo.GetByID(userId)

	if err != nil || userDb == nil {
		return nil, err
	}
	
	userTags, err := s.GetUserTags(userId)
	if err != nil {
		return nil, err
	}
	
	userPublicSettings := dto.PublicUserSettingsDTO{
		IsShowTag: userDb.IsShowTag,
	}
	
	userLevel, err := s.UserProfileRepo.GetUserLevel(context.Background(), userId)

	if err != nil {
		return nil, err
	}
	
	return  &dto.GetUserDTO{
		ID:       userDb.ID,
		Name:     userDb.Name,
		Picture:  userDb.Picture,
		Role:     userDb.Role,
		Status:   userDb.Status,
		About:    userDb.About,
		Birthday: userDb.Birthday,
		Gender:   string(userDb.Gender),
		IsPublic: userDb.IsPublic,
		LastLogin: userDb.LastLogin,
		CreateAt: userDb.CreateAt,
		Banner:   userDb.Banner,
		ActiveTag:      userDb.Tag,
		AllTags:        userTags,
		Settings:       userPublicSettings,
		UserLevel:      dto.UserLevelDTO{
			Level:      userLevel.Level,
			Experience: userLevel.Experience,
			LevelTitle: userLevel.LevelTitle,
			XPForNext:  userLevel.XPForNext,
		},
	}, err
}

func (s *UserService) UpdateUserStatus(ctx context.Context, userId int, dto dto.UpdateUserStatusDTO) error {
 	lastActiveTime := time.UnixMilli(dto.LastActive)

    if dto.LastActive == 0 {
        lastActiveTime = time.Now()
    }
    
	return s.Repo.UpdateDtoStatus(ctx, userId, dto.Status, lastActiveTime)
}

func (s *UserService) UpdateUserTag(ctx context.Context, userId int, dto dto.UpdateUserTagDTO) error {
	return s.Repo.UpdateUserTag(ctx, userId, dto.ID)
}

func (s *UserService) GetUserTags(userId int) ([]dto.UserTagDTO, error) {
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
