package repository

import (
	"context"
	"database/sql"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserActivityRepository struct {
	DB *sql.DB
}

func NewUserActivityRepository(db *sql.DB) *UserActivityRepository {
	return &UserActivityRepository{DB: db}
}

func (r *UserActivityRepository) GetUserActivies(ctx context.Context, userID int) ([] *db.UserActivity, error) {
	var userActivity []*db.UserActivity
	query := `
		SELECT id, activity_type, target_id, metadata, timestamp
		FROM user_activities
		WHERE user_id = $1`
	
	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	
	for rows.Next() {
		var activity db.UserActivity
		err := rows.Scan(
			&activity.ID,
			&activity.ActivityType,
			&activity.TargetID,
			&activity.Metadata,
			&activity.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		userActivity = append(userActivity, &activity)
	}
	
	return userActivity, nil
}

func (r *UserActivityRepository) CreateUserActivity(ctx context.Context, userActivity *db.UserActivity) error {
	query := `
		INSERT INTO user_activities (user_id, activity_type, target_id, metadata, timestamp)
		VALUES ($1, $2, $3, $4, NOW())`
	_, err := r.DB.ExecContext(ctx, query, userActivity.UserID, userActivity.ActivityType, userActivity.TargetID, userActivity.Metadata)
	if err != nil {
		return err
	}
	
	return nil
}
