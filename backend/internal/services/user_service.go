package service

import (
	"context"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type UserService struct {
	Repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{Repo: repo}
}

func (s *UserService) GetUserById(userId int) (*dto.GetUserDTO, error) {
	userDb, err := s.Repo.GetByID(userId)

	if err != nil {
		return nil, err
	}
	
	userTags, err := s.GetUserTags(userId)
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
	}, err
}

func (s *UserService) UpdateUserStatus(ctx context.Context, userId int, dto dto.UpdateUserStatusDTO) error {
 	lastActiveTime := time.UnixMilli(dto.LastActive)

    if dto.LastActive == 0 {
        lastActiveTime = time.Now()
    }
    
	return s.Repo.UpdateDtoStatus(ctx, userId, dto.Status, lastActiveTime)
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
