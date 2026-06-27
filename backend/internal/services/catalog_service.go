package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"

	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

var ErrFilterPresetNotFound = errors.New("filter preset not found")

type CatalogService struct {
	catalogRepo *repository.CatalogRepository
}

func NewCatalogService(catalogRepo *repository.CatalogRepository) *CatalogService {
	return &CatalogService{
		catalogRepo: catalogRepo,
	}
}

func (s *CatalogService) SaveFilters(ctx context.Context, userID int, name string, filters json.RawMessage) (*dto.CatalogFilterPreset, error) {
	return s.catalogRepo.SaveFilters(ctx, userID, name, filters)
}

func (s *CatalogService) GetUserFilters(ctx context.Context, userID int) ([]*dto.CatalogFilterPreset, error) {
	return s.catalogRepo.GetUserFilters(ctx, userID)
}

func (s *CatalogService) GetFilterByID(ctx context.Context, filterID, userID int) (*dto.CatalogFilterPreset, error) {
	preset, err := s.catalogRepo.GetFilterByID(ctx, filterID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrFilterPresetNotFound
		}
		return nil, err
	}
	return preset, nil
}

func (s *CatalogService) DeleteFilter(ctx context.Context, filterID, userID int) error {
	err := s.catalogRepo.DeleteFilter(ctx, filterID, userID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrFilterPresetNotFound
		}
		return err
	}
	return nil
}
