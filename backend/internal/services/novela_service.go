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
	Repo                     *repository.NovelaRepository
	RatingRepo         *repository.NovelaRatingRepository
	BookmarkRepo *repository.NovelaBookmarkRepository
	LikeRepo    *repository.NovelaLikeRepository
}

func NewNovelaService(repo *repository.NovelaRepository,
	novelaRatingRepo *repository.NovelaRatingRepository,
	novelaBookmarkRepo *repository.NovelaBookmarkRepository,
	novelaLikeRepo *repository.NovelaLikeRepository,
) *NovelaService {
	return &NovelaService{
		Repo: repo,
		RatingRepo: novelaRatingRepo,
		BookmarkRepo: novelaBookmarkRepo,
		LikeRepo: novelaLikeRepo,
	}
}

func (s *NovelaService) GetNovelaById(ctx context.Context, id, userID int) (*dto.NovelaResponse, error) {
	tx, err := s.Repo.DB.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	novela, err := s.Repo.GetFullByID(tx, ctx, id, userID)
	if err != nil || novela == nil {
		return nil, ErrNovelaNotFound
	}

	ratingData, _ := s.RatingRepo.GetRating(tx, ctx, novela.ID)
	bookmarkData, _ := s.BookmarkRepo.GetBookmarkStats(tx, ctx, novela.ID)

	res := s.novelaToDto(novela)
	res.RatingDetails = s.mapRatingToDto(ratingData)
	res.BookmarkDetails = s.mapBookmarkToDto(bookmarkData)

	return res, nil
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

	return s.BookmarkRepo.SetBookmark(context, dbBookmark)
}

func (s *NovelaService) RemoveBookmark(ctx context.Context, userID, novelaID int) error {
	return s.BookmarkRepo.RemoveBookmark(ctx, userID, novelaID)
}

func (s *NovelaService) SetLike(ctx context.Context, userID int, req dto.UpdateLikeRequest) error {
	return s.LikeRepo.SetLike(ctx, &db.NovelaLike{
		UserID:   userID,
		NovelaID: req.NovelaID,
		HasLiked: req.HasLiked,
	})
}

func (s *NovelaService) SetRating(ctx context.Context, userID int, req dto.UpdateRatingRequest) error {
	return s.RatingRepo.SetRating(ctx, &db.NovelaRating{
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

func (s *NovelaService) GetNovelas(ctx context.Context, userID int, f dto.NovelaFilters) ([]dto.NovelaResponse, int, error) {
	if f.Limit <= 0 || f.Limit > 100 { f.Limit = 20 }

	tx, err := s.Repo.DB.BeginTx(ctx, &sql.TxOptions{ReadOnly: true})
	if err != nil {
		return nil, 0, err
	}
	defer tx.Rollback()

	dbNovelas, total, err := s.Repo.GetNovelas(tx, ctx, userID, f)
	if err != nil {
		return nil, 0, err
	}

	res := make([]dto.NovelaResponse, 0, len(dbNovelas))
	for _, n := range dbNovelas {

		d := *s.novelaToDto(&n)
		ratingData, _ := s.RatingRepo.GetRating(tx, ctx, d.ID)
		bookmarkData, _ := s.BookmarkRepo.GetBookmarkStats(tx, ctx, d.ID)

		d.RatingDetails = s.mapRatingToDto(ratingData)
		d.BookmarkDetails = s.mapBookmarkToDto(bookmarkData)

		res = append(res, d)
	}

	return res, total, tx.Commit()
}

func (s *NovelaService) GetUserNovelaBookmarks(ctx context.Context, userID int, category string) ([]dto.UserNovelaBookmark, error) {
	novelas, err := s.Repo.GetUserNovelaBookmarks(ctx, userID, db.BookmarkCategory(category))
	if err != nil {
		return nil, errors.New("Invalid category type ")
	}

	var ns []dto.UserNovelaBookmark
	for _, n := range novelas {
		ns = append(ns, dto.UserNovelaBookmark{
			ID:        n.ID,
			Title:     n.Title,
			PosterURL: n.PosterURL,
			Type:      n.Type,
			Rating:    n.Rating,
		})
	}

	return ns, nil
}

func (s *NovelaService) novelaToDto(n *db.Novela) *dto.NovelaResponse {
	authors := make([]dto.NovelaAuthor, 0)
	for _, a := range n.Authors {
		authors = append(authors, dto.NovelaAuthor{ID: a.ID, Name: a.Name, Role: a.Role})
	}

	volumes := make([]dto.NovelaVolume, 0)
	for _, v := range n.Volumes {
		chapters := make([]dto.NovelaChapter, 0)
		for _, ch := range v.Chapters {
			imgs := make([]dto.NovelaChapterImage, 0)
			for _, img := range ch.Images {
				imgs = append(imgs, dto.NovelaChapterImage{ID: img.ID, ImageURL: img.ImageURL, Caption: img.Caption})
			}
			chapters = append(chapters, dto.NovelaChapter{ID: ch.ID, Title: ch.Title, Number: ch.Number, Images: imgs})
		}
		volumes = append(volumes, dto.NovelaVolume{ID: v.ID, Title: v.Title, Number: v.Number, Chapters: chapters})
	}

	var bookmark *string
	if n.Bookmark != nil {
		str := string(*n.Bookmark)
		bookmark = &str
	}

	return &dto.NovelaResponse{
		ID: n.ID, Title: n.Title, AlternativeTitles: n.AlternativeTitles, Description: n.Description,
		Type: n.Type, AgeRating: n.AgeRating, ReleaseDate: n.ReleaseDate.Format("2006-01-02"),
		Status: n.Status, TranslationStatus: n.TranslationStatus, PosterURL: n.PosterURL,
		Country: n.Country, Views: n.Views, Genres: n.Genres, Categories: n.Categories,
		Authors: authors, Volumes: volumes, Bookmark: bookmark, HasLiked: n.HasLiked,
		LikeCount: n.LikeCount, UserRating: n.UserRating,
	}
}

func (s *NovelaService) mapRatingToDto(data *db.NovelaRatingSummary) dto.NovelaRatingCategory {
	items := make([]dto.NCItem, 0)
	if data == nil {
		return dto.NovelaRatingCategory{NCItems: items}
	}
	
	for i := 10; i >= 1; i-- {
		count := 0
		if val, ok := data.Distribution[i]; ok { count = val }
		items = append(items, dto.NCItem{Value: i, Count: count})
	}
	return dto.NovelaRatingCategory{Total: data.TotalCount, TotalRating: data.AverageRating, NCItems: items}
}

func (s *NovelaService) mapBookmarkToDto(data *db.NovelaBookmarkSummary) dto.NovelaBookmarkCategory {
	items := make([]dto.NCItem, 0)
	if data == nil {
		return dto.NovelaBookmarkCategory{NCItems: items}
	}

	for _, cat := range db.BookmarkCategories {
		count := 0
		if val, ok := data.Distribution[string(cat)]; ok { count = val }
		items = append(items, dto.NCItem{Value: string(cat), Count: count})
	}
	return dto.NovelaBookmarkCategory{Total: data.TotalCount, NCItems: items}
}
