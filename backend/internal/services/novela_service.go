package service

import (
	"context"
	"database/sql"
	"errors"
	"mime/multipart"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/utils/files"
)

type NovelaService struct {
	Repo *repository.NovelaRepository
}

func NewNovelaService(repo *repository.NovelaRepository) *NovelaService {
	return &NovelaService{
		Repo: repo,
	}
}

func (s *NovelaService) GetNovelaById(ctx context.Context, id, userID int) (*dto.NovelaResponse, error) {
	ranobe, err := s.Repo.GetFullByID(ctx, id, userID)
	
	if err != nil {
		return nil, err
	}

	return s.novelaToDto(ranobe), nil
}

func (s *NovelaService) UploadPoster(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	return files.UploadImage(ctx, file, header, "/novelas/posters", 600, 900)
}

func (s *NovelaService) SavePoster(ctx context.Context, id int, posterURL string) error {
	return s.Repo.UpdatePoster(ctx, id, posterURL)
}

func (s *NovelaService) Create(ctx context.Context, novela *db.Novela) error {
	return s.Repo.Create(ctx, novela)
}

func (s *NovelaService) SetBookmark(context context.Context, userID int, req dto.UpdateBookmarkRequest) error {

	dbBookmark := &db.Bookmark{
		UserID:   userID,
		NovelaID: req.NovelaID,
		Category: db.BookmarkCategory(req.Category),
	}

	return s.Repo.SetBookmark(context, dbBookmark)
}

func (s *NovelaService) RemoveBookmark(ctx context.Context, userID, novelaID int) error {
	return s.Repo.RemoveBookmark(ctx, userID, novelaID)
}

func (s *NovelaService) SetLike(ctx context.Context, userID int, req dto.UpdateLikeRequest) error {
	return s.Repo.SetLike(ctx, &db.NovelaLike{
		UserID:   userID,
		NovelaID: req.NovelaID,
		HasLiked: req.HasLiked,
	})
}

func (s *NovelaService) SetRating(ctx context.Context, userID int, req dto.UpdateRatingRequest) error {
	return s.Repo.SetRating(ctx, &db.NovelaRating{
		UserID:   userID,
		NovelaID: req.NovelaID,
		Rating:   req.Rating,
	})
}

func (s *NovelaService) UpdateNovela(ctx context.Context, id int, req dto.UpdateNovelaRequest, posterURL string) error {
	releaseDate, _ := time.Parse("2006", req.ReleaseDate)
	novela := &db.Novela{
		ID:                id,
		Title:             req.Title,
		AlternativeTitles: req.AlternativeTitles,
		Description:       req.Description,
		Type:              req.Type,
		AgeRating:         req.AgeRating,
		ReleaseDate:       releaseDate,
		Status:            req.Status,
		Country:           req.Country,
		TranslationStatus: req.TranslationStatus,
		PosterURL:         posterURL,
		Genres:            req.Genres,
		Categories:        req.Categories,
	}

	for _, authorID := range req.Authors {
		novela.Authors = append(novela.Authors, db.NovelaAuthor{ID: authorID})
	}

	err := s.Repo.UpdateFull(ctx, novela)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.ErrUnsupported
		}
		return ErrInternal
	}

	return nil
}
func (s *NovelaService) novelaToDto(novela *db.Novela) *dto.NovelaResponse {
	if novela == nil {
		return nil
	}

	authors := make([]dto.NovelaAuthor, 0)
	for _, a := range novela.Authors {
		authors = append(authors, dto.NovelaAuthor{
			ID:   a.ID,
			Name: a.Name,
			Role: a.Role,
		})
	}

	volumes := make([]dto.NovelaVolume, 0)
	for _, v := range novela.Volumes {
		chapters := make([]dto.NovelaChapter, 0)
		for _, ch := range v.Chapters {
			images := make([]dto.NovelaChapterImage, 0)
			for _, img := range ch.Images {
				images = append(images, dto.NovelaChapterImage{
					ID:       img.ID,
					ImageURL: img.ImageURL,
					Caption:  img.Caption,
				})
			}

			chapters = append(chapters, dto.NovelaChapter{
				ID:     ch.ID,
				Title:  ch.Title,
				Number: ch.Number,
				Images: images,
			})
		}

		volumes = append(volumes, dto.NovelaVolume{
			ID:       v.ID,
			Title:    v.Title,
			Number:   v.Number,
			Chapters: chapters,
		})
	}

	var bookmark string
	if novela.Bookmark != nil {
		bookmark = string(*novela.Bookmark)
	}

	return &dto.NovelaResponse{
		ID:                novela.ID,
		Title:             novela.Title,
		AlternativeTitles: novela.AlternativeTitles,
		Description:       novela.Description,
		Type:              novela.Type,
		AgeRating:         novela.AgeRating,
		ReleaseDate:       novela.ReleaseDate.Format("2006-01-02"),
		Status:            novela.Status,
		TranslationStatus: novela.TranslationStatus,
		PosterURL:         novela.PosterURL,
		Country:           novela.Country,
		Views:             novela.Views,
		Rating:            novela.Rating,
		Genres:            novela.Genres,
		Categories:        novela.Categories,
		Authors:           authors,
		Volumes:           volumes,
		Bookmark: 		   &bookmark,
		BookmarkCount:     novela.BookmarkCount,
		HasLiked:          novela.HasLiked,
		LikeCount:         novela.LikeCount,
		UserRating:        novela.UserRating,
		RatingCount:       novela.RatingCount,
	}
}