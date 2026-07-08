package service

import (
	"context"

	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserExperianceService struct {
	UserExpRepo *repository.UserExperianceRepository
}

func NewUserExperianceService(userExpRepo *repository.UserExperianceRepository) *UserExperianceService {
	return &UserExperianceService{
		UserExpRepo: userExpRepo,
	}
}

func (s *UserExperianceService) GetLevelDefinitions (ctx context.Context) ([]dto.LevelDefinitions, error) {
	defenitions, err := s.UserExpRepo.GetLevelDefinitions(ctx)
	if err != nil {
		return nil, err
	}

	levelDefinitions := make([]dto.LevelDefinitions, len(defenitions))
	for i, level := range defenitions {
		levelDefinitions[i] = s.mapToDtoLevelDefinitions(level)
	}

	return levelDefinitions, nil
}

func (s *UserExperianceService) mapToDtoLevelDefinitions(level db.UserExperianceDefinitions) dto.LevelDefinitions {
	return dto.LevelDefinitions{
		Level: level.Level,
		Title: level.Title,
		Total_XP_Required: level.Total_XP_Required,
	}
}

