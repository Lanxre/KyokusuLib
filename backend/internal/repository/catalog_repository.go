package repository

import (
	"context"
	"database/sql"
	"encoding/json"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
)

type CatalogRepository struct {
	DB *sql.DB
}

func NewCatalogRepository(db *sql.DB) *CatalogRepository {
	return &CatalogRepository{DB: db}
}

func (r *CatalogRepository) SaveFilters(ctx context.Context, userID int, name string, filters json.RawMessage) (*db.CreateCatalogFilterPreset, error) {
	var preset db.CreateCatalogFilterPreset
	err := r.DB.QueryRowContext(ctx, `
		INSERT INTO user_catalog_filters (user_id, name, filters)
		VALUES ($1, $2, $3)
		RETURNING id, user_id, name, filters, created_at
	`, userID, name, filters).Scan(
		&preset.ID, &preset.UserID, &preset.Name, &preset.Filters, &preset.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &preset, nil
}

func (r *CatalogRepository) GetUserFilters(ctx context.Context, userID int) ([]*db.GetCatalogFilterPreset, error) {
	rows, err := r.DB.QueryContext(ctx, `
		SELECT id, name, filters, created_at
		FROM user_catalog_filters
		WHERE user_id = $1
		ORDER BY created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var presets []*db.GetCatalogFilterPreset
	for rows.Next() {
		var preset db.GetCatalogFilterPreset
		if err := rows.Scan(&preset.ID, &preset.Name, &preset.Filters, &preset.CreatedAt); err != nil {
			return nil, err
		}
		presets = append(presets, &preset)
	}
	return presets, rows.Err()
}

func (r *CatalogRepository) GetFilterByID(ctx context.Context, filterID, userID int) (*dto.CatalogFilterPreset, error) {
	var preset dto.CatalogFilterPreset
	err := r.DB.QueryRowContext(ctx, `
		SELECT id, user_id, name, filters, created_at
		FROM user_catalog_filters
		WHERE id = $1 AND user_id = $2
	`, filterID, userID).Scan(
		&preset.ID, &preset.UserID, &preset.Name, &preset.Filters, &preset.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &preset, nil
}

func (r *CatalogRepository) DeleteFilter(ctx context.Context, filterID, userID int) error {
	result, err := r.DB.ExecContext(ctx, `
		DELETE FROM user_catalog_filters
		WHERE id = $1 AND user_id = $2
	`, filterID, userID)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *CatalogRepository) CountUserFilters(ctx context.Context, userID int) (int, error) {
    var count int

    err := r.DB.QueryRowContext(ctx, `
        SELECT COUNT(*)
        FROM user_catalog_filters
        WHERE user_id = $1
    `, userID).Scan(&count)

    return count, err
}