package service

import (
	"context"
	"errors"
	"time"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/db"
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

func (s *UserService) GetUsers(ctx context.Context, search string, limit int, offset int) ([]*dto.GetUserDTO, error) {
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	
	usersDb, err := s.Repo.GetUsers(ctx, search, limit, offset)
	if err != nil {
		return nil, err
	}

	usersDto := make([]*dto.GetUserDTO, len(usersDb))
	for i, user := range usersDb {
		userLevel, err := s.UserProfileRepo.GetUserLevel(context.Background(), user.ID)
		if err != nil {
			return nil, err
		}

		userTags, err := s.GetUserTags(user.ID)

		if err != nil {
			return nil, err
		}

		totalComments, readChapters, err := s.Repo.GetUserStats(context.Background(), user.ID)
		if err != nil {
			return nil, err
		}
		
		usersDto[i] = toUserDTO(user, userLevel, userTags, dto.PublicUserSettingsDTO{
			IsShowTag:     user.IsShowTag,
			IsShowBookmark: user.IsShowBookmark,
		}, dto.UserStatsDTO{
			TotalComments: totalComments,
			ReadChapters:  readChapters,
		})
	}
	return usersDto, nil
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
	
	userLevel, err := s.UserProfileRepo.GetUserLevel(context.Background(), userId)
	if err != nil {
		return nil, err
	}

	totalComments, readChapters, err := s.Repo.GetUserStats(context.Background(), userId)
	if err != nil {
		return nil, err
	}
	
	return toUserDTO(userDb, userLevel, userTags, dto.PublicUserSettingsDTO{
		IsShowTag:     userDb.IsShowTag,
		IsShowBookmark: userDb.IsShowBookmark,
	}, dto.UserStatsDTO{
		TotalComments: totalComments,
		ReadChapters:  readChapters,
	}), nil
}

func (s *UserService) UpdateUserStatus(ctx context.Context, userID int, dto dto.UpdateUserStatusDTO) error {
    lastActive := time.Now()
    if dto.LastActive != 0 {
        lastActive = time.UnixMilli(dto.LastActive)
    }

    status := constants.UserStatus(dto.Status)
    if !status.IsValid() {
        return errors.New("Invalid user status")
    }

    return s.Repo.UpdateDtoStatus(ctx, userID, status, lastActive)
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

func toUserDTO(user *db.User, level *db.UserLevel, tags []dto.UserTagDTO, settings dto.PublicUserSettingsDTO, stats dto.UserStatsDTO) *dto.GetUserDTO {
	return &dto.GetUserDTO{
		ID:        user.ID,
		Email:     user.Email,
		Name:      user.Name,
		Picture:   user.Picture,
		Banner:    user.Banner,
		Role:      user.Role,
		Status:    user.Status,
		About:     user.About,
		Birthday:  user.Birthday,
		Gender:    string(user.Gender),
		IsPublic:  user.IsPublic,
		LastLogin: user.LastLogin,
		CreateAt:  user.CreateAt,
		ActiveTag: user.Tag,
		AllTags:   tags,
		Settings:  settings,
		UserStats: stats,
		UserLevel: dto.UserLevelDTO{
			Level:       level.Level,
			Experience:  level.Experience,
			LevelTitle:  level.LevelTitle,
			XPForNext:   level.XPForNext,
		},
	}
}