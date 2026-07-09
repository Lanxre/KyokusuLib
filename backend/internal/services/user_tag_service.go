package service

import (
	"context"

	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/models/dto"
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