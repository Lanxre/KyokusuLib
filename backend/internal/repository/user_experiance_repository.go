package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)


type UserExperianceRepository struct {
	DB *sql.DB
}

func NewUserExperianceRepository(db *sql.DB) *UserExperianceRepository {
	return &UserExperianceRepository{DB: db}
}

func (r *UserExperianceRepository) GetLevelDefinitions(ctx context.Context) ([]db.UserExperianceDefinitions, error) {
	query := "select level, title, total_xp_required from level_definitions"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var definitions []db.UserExperianceDefinitions
	for rows.Next() {
		var level db.UserExperianceDefinitions
		if err := rows.Scan(&level.Level, &level.Title, &level.Total_XP_Required); err != nil {
			return nil, err
		}
		definitions = append(definitions, level)
	}
	
	return definitions, nil
}