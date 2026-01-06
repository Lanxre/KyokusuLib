package service

import (
	"context"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type NovelaService struct {
	Repo *repository.NovelaRepository
}

func NewNovelaService(repo *repository.NovelaRepository) *NovelaService {
	return &NovelaService{
		Repo: repo,
	}
}

func (s *NovelaService) GetNovelaById(ctx context.Context, id int) (*db.Novela, error){
	ranobe, err := s.Repo.GetFullByID(ctx, id)
	return ranobe, err
}