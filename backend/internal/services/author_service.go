package service

import (
	"context"
	"mime/multipart"
	"regexp"
	"strings"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/utils/files"
)

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: *repo}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, author *db.Author) error {
	return s.repo.CreateAuthor(ctx, author)
}

func (s *AuthorService) UploadAuthorImage(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	avatarURL, err := files.UploadImage(ctx, file, header, "authors", 400, 400)
    if err != nil {
        return "", err
    }
    
	return avatarURL, nil
}

func (s *AuthorService) GetAuthorByName(ctx context.Context, name string) (*db.Author, error) {
	name = strings.TrimSpace(name)
	name = regexp.MustCompile(`[^\w\s-]`).ReplaceAllString(name, "")
	name = regexp.MustCompile(`\s+`).ReplaceAllString(name, " ")
	return s.repo.GetAuthorByName(ctx, name)
}

func (s *AuthorService) GetAuthors(ctx context.Context, search string, limit int) ([]*db.Author, error) {
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	
	authorsDb, err := s.repo.GetAuthors(ctx, search, limit)
	
	if err != nil {
		return nil, err
	}

	authorsDto := make([]*dto.AuthorDTO, len(authorsDb))
	for i, author := range authorsDb {
		authorsDto[i] = s.toDTO(author)
	}

	return authorsDb, nil
}

func (s *AuthorService) toDTO(author *db.Author) *dto.AuthorDTO {
	return &dto.AuthorDTO{
		ID:       author.ID,
		Name:     author.Name,
		Country:  author.Country,
		Metier:   author.Metier,
		Picture:  author.Picture,
		Bio:      author.Bio,
	}
}
