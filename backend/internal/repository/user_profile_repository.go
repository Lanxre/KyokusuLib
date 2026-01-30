package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserProfileRepository struct {
	DB *sql.DB
}

func NewUserProfileRepository(db *sql.DB) *UserProfileRepository {
	return &UserProfileRepository{DB: db}
}

func (r *UserProfileRepository) GetUserLevel(ctx context.Context, userID int) (*db.UserLevel, error) {
	u := &db.UserLevel{}
	
	var (
		level                int
        experience           int64
        levelTitle           sql.NullString
        xpForThisLevel       sql.NullInt64
        xpForNextLevel       sql.NullInt64
	)
	
	query := `
	SELECT 
		p.level, p.experience,
		ld.title AS level_title,
		ld.total_xp_required AS xp_for_this_level,
		COALESCE(next_ld.total_xp_required, ld.total_xp_required) AS xp_for_next_level
	FROM user_profiles p
	LEFT JOIN level_definitions ld ON ld.level = p.level
	LEFT JOIN level_definitions next_ld ON next_ld.level = p.level + 1
	WHERE p.user_id = $1
	`
	
	err := r.DB.QueryRowContext(ctx, query, userID).Scan(&level, &experience, &levelTitle, &xpForThisLevel, &xpForNextLevel)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}
		return nil, err
	}
	
	if levelTitle.Valid {
        u.LevelTitle = levelTitle.String
    }
    
    if xpForNextLevel.Valid && xpForThisLevel.Valid {
        needed := max(xpForNextLevel.Int64 - experience, 0)
        u.XPForNext = needed
    }

	u.Level = level
	u.Experience = experience
	
	return u, nil
}