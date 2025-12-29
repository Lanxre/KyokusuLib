package service

import (
	"context"
	"regexp"
	"strings"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type AuthorService struct {
	repo repository.AuthorRepository
}

func NewAuthorService(repo *repository.AuthorRepository) *AuthorService {
	return &AuthorService{repo: *repo}
}

func (s *AuthorService) CreateAuthor(ctx context.Context, name string) error {
	return s.repo.CreateAuthor(ctx, name)
}

func (s *AuthorService) GetAuthorByName(ctx context.Context, name string) (*db.Author, error) {
	name = strings.TrimSpace(name)
	name = regexp.MustCompile(`[^\w\s-]`).ReplaceAllString(name, "")
	name = regexp.MustCompile(`\s+`).ReplaceAllString(name, " ")
	return s.repo.GetAuthorByName(ctx, name)
}

func (s *AuthorService) GetAuthors(ctx context.Context) ([]*db.Author, error) {
	return s.repo.GetAuthors(ctx)
}
