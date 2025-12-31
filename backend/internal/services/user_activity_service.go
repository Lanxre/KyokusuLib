package service

import (
	"context"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type UserActivityService struct {
	Repo *repository.UserActivityRepository
}

func NewUserActivityService(repo *repository.UserActivityRepository) *UserActivityService {
	return &UserActivityService{Repo: repo}
}

func (s *UserActivityService) GetUserActivity(ctx context.Context, userID int) ([]*dto.GetUserActivity, error) {
	var userActivity []*dto.GetUserActivity
	activities, err := s.Repo.GetUserActivies(ctx, userID)
	if err != nil {
		return nil, err
	}
	
	for _, activity := range activities {
		userActivity = append(userActivity, &dto.GetUserActivity{
			ID: activity.ID,
			ActivityType: activity.ActivityType,
			Timestamp: activity.Timestamp,
			Metadata: activity.Metadata,
		})
	}
	return userActivity, nil
}

func (s *UserActivityService) CreateUserActivity(ctx context.Context, userID int, userActivity *dto.CreateUserActivity) error {
	model := &db.UserActivity{
		UserID: userID,
		ActivityType: userActivity.ActivityType,
		TargetID: userActivity.TargetID,
		Metadata: userActivity.Metadata,
	}
	return s.Repo.CreateUserActivity(ctx, model)
}
