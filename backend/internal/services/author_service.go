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
	ext := strings.ToLower(filepath.Ext(header.Filename))
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return "", errors.New("unsupported file format")
	}

	srcImage, err := imaging.Decode(file)
	if err != nil {
		return "", err
	}

	dstImage := imaging.Fill(srcImage, 400, 400, imaging.Center, imaging.Lanczos)

	newFilename := fmt.Sprintf("%s.jpg", uuid.New().String())
	savePath := filepath.Join(static.UPLOAD_AUTHOR_DIR, newFilename)

	err = imaging.Save(dstImage, savePath, imaging.JPEGQuality(80))
	if err != nil {
		return "", err
	}

	avatarURL := fmt.Sprintf("/uploads/authors/%s", newFilename)
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
