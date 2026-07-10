package service

import (
	"context"
	"errors"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type UserTagService struct {
	UserTagRepo *repository.UserTagRepository
}

func NewUserTagService(userTagRepo *repository.UserTagRepository) *UserTagService {
	return &UserTagService{
		UserTagRepo: userTagRepo,
	}
}

func (s *UserTagService) GetDefinitions(ctx context.Context) ([]dto.UserTagDTO, error) {
	tags, err := s.UserTagRepo.GetUserTags(ctx)
	if err != nil {
		return nil, err
	}
	
	userTags := make([]dto.UserTagDTO, len(tags))
	for i, tag := range tags {
		userTags[i] = dto.UserTagDTO{
			ID:    tag.ID,
			Tag:   tag.Tag,
		}
	}
	
	return userTags, nil
}

func (s *UserTagService) UpdateUserTags(ctx context.Context, userID int, tagIds []int) error {
	exist, err := s.UserTagRepo.IsExist(ctx, tagIds)
	if err != nil {
		return err
	}

	if !exist {
		return errors.New("Not has such tag id")
	}

	s.UserTagRepo.UpdateUserTags(ctx, userID, tagIds)

	return nil
}