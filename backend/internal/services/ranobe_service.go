package service

import (
	"context"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type RanobeService struct {
	Repo *repository.RanobeRepository
}

func NewRanobeService(repo *repository.RanobeRepository) *RanobeService {
	return &RanobeService{
		Repo: repo,
	}
}

func (s *RanobeService) GetRanobeById(ctx context.Context, id int) (*db.Ranobe, error){
	ranobe, err := s.Repo.GetFullByID(ctx, id)
	return ranobe, err
}