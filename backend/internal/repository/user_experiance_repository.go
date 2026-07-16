package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)


type UserExperienceRepository struct {
	DB *sql.DB
}

func NewUserExperianceRepository(db *sql.DB) *UserExperienceRepository {
	return &UserExperienceRepository{DB: db}
}

func (r *UserExperienceRepository) GetLevelDefinitions(ctx context.Context) ([]db.UserExperienceDefinitions, error) {
	query := "select level, title, total_xp_required from level_definitions"

	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var definitions []db.UserExperienceDefinitions
	for rows.Next() {
		var level db.UserExperienceDefinitions
		if err := rows.Scan(&level.Level, &level.Title, &level.Total_XP_Required); err != nil {
			return nil, err
		}
		definitions = append(definitions, level)
	}
	
	return definitions, nil
}

func (r *UserExperienceRepository) UpdateUserLevel(ctx context.Context, userID, level int, exp int64) error {
	_, err := r.DB.ExecContext(
		ctx,
		`UPDATE user_profile SET level = $1, experience = $2 WHERE user_id = $3`,
		level, exp, userID,
	)
	return err
}