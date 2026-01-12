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

func (s *UserActivityService) GetUserActivity(ctx context.Context, userID int, params dto.QueryParams) ([]*dto.GetUserActivity, error) {

	if params.Limit <= 0 || params.Limit > 100 {
		params.Limit = 20
	}

	activities, err := s.Repo.GetUserActivies(ctx, userID, params.Limit)
	if err != nil {
		return nil, err
	}
	
	userActivity := make([]*dto.GetUserActivity, 0, len(activities))
	for _, activity := range activities {
		userActivity = append(userActivity, &dto.GetUserActivity{
			ID:           activity.ID,
			ActivityType: activity.ActivityType,
			TargetID:     activity.TargetID,
			Timestamp:    activity.Timestamp,
			Metadata:     activity.Metadata,
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
