package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	"github.com/lanxre/kyokusulib/internal/constants"
	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

var ErrFilterPresetNotFound = errors.New("filter preset not found")
var ErrMaxPresetsReached = errors.New("Max value on create filters: " + strconv.Itoa(constants.MaxFilterPresetsValue))

type CatalogService struct {
	catalogRepo *repository.CatalogRepository
}

func NewCatalogService(catalogRepo *repository.CatalogRepository) *CatalogService {
	return &CatalogService{
		catalogRepo: catalogRepo,
	}
}

func (s *CatalogService) SaveFilters(ctx context.Context, userID int, name string, filters json.RawMessage) (*db.CreateCatalogFilterPreset, error) {

	totalCount, err := s.catalogRepo.CountUserFilters(ctx, userID)

	if  err != nil {
		return nil, err
	}

	if totalCount >= constants.MaxFilterPresetsValue {
		return nil, ErrMaxPresetsReached
	}
	
	return s.catalogRepo.SaveFilters(ctx, userID, name, filters)
}

func (s *CatalogService) GetUserFilters(ctx context.Context, userID int) ([]*dto.CatalogFilterPreset, error) {
	presets, err := s.catalogRepo.GetUserFilters(ctx, userID)

	if err != nil {
		return nil, err
	}

	var dtos []*dto.CatalogFilterPreset

	for _, preset := range presets {
		dto := s.mapToCatalogFilterPreset(preset)
		dtos = append(dtos, dto)
	}

	return dtos, nil
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

func (s *CatalogService) mapToCatalogFilterPreset(preset *db.GetCatalogFilterPreset) *dto.CatalogFilterPreset {
	return &dto.CatalogFilterPreset{
		ID:        preset.ID,
		Name:      preset.Name,
		Filters:   preset.Filters,
		CreatedAt: preset.CreatedAt,
	}
}