package service

import (
	"context"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/repository"
)

type SocialsService struct {
	Repo *repository.UserSocialsRepository
}

func NewSocialsService(repo *repository.UserSocialsRepository) *SocialsService {
	return &SocialsService{Repo: repo}
}

func (s *SocialsService) GetUserSocials(ctx context.Context, userID int) (*db.UserSocials, error) {
	return s.Repo.GetUserSocials(ctx, userID)
}

func (s *SocialsService) LinkAccount(ctx context.Context, userID int, provider db.SocialProvider, externalID, refresh_token string) error {
	return s.Repo.LinkSocial(ctx, userID, provider, externalID, refresh_token)
}

func (s *SocialsService) UnlinkAccount(ctx context.Context, userID int, provider db.SocialProvider) error {
	return s.Repo.UnlinkSocial(ctx, userID, provider)
}