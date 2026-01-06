package service

import (
	"context"
	"mime/multipart"

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

func (s *NovelaService) UploadPoster(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	return "", nil
    // return s.FileStorage.SaveImage(file, header, "posters", 600, 900)
}

func (s *NovelaService) Create(ctx context.Context, novela *db.Novela) error {
	return s.Repo.Create(ctx, novela)
}
