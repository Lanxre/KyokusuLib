package service

import (
	"context"
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/utils/files"
	"github.com/lanxre/kyokusulib/internal/utils/static"
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

func (s *AuthorService) GetAuthors(ctx context.Context) ([]*db.Author, error) {
	return s.repo.GetAuthors(ctx)
}
