package service

import (
	"context"
	"fmt"
	"mime/multipart"

	"github.com/lanxre/kyokusulib/internal/models/db"
	"github.com/lanxre/kyokusulib/internal/models/dto"
	"github.com/lanxre/kyokusulib/internal/repository"
	"github.com/lanxre/kyokusulib/internal/utils/files"
)

type TeamService struct {
	Repo *repository.TeamRepository
}

func NewTeamService(repo *repository.TeamRepository) *TeamService {
	return &TeamService{Repo: repo}
}

func (s *TeamService) GetTeams(ctx context.Context, search string, limit int, offset int) ([]*db.PublisherTeam, error) {
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	return s.Repo.GetTeams(ctx, search, limit, offset)
}

func (s *TeamService) Create(ctx context.Context, userID int, input dto.CreateTeamDTO) (*db.PublisherTeam, error) {
	existing, err := s.Repo.GetBySlug(ctx, input.Slug)
	if err != nil {
		return nil, err
	}
	if existing != nil {
		return nil, fmt.Errorf("Команда с таким Slug уже существует")
	}

	team := &db.PublisherTeam{
		OwnerID:     userID,
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
	}

	if err := s.Repo.Create(ctx, team); err != nil {
		return nil, fmt.Errorf("Команда с таким названием уже существует")
	}
	return team, nil
}

func (s *TeamService) Get(ctx context.Context, slug string) (*db.PublisherTeam, error) {
	return s.Repo.GetBySlug(ctx, slug)
}

func (s *TeamService) Update(ctx context.Context, userID int, slug string, input dto.UpdateTeamDTO) (*db.PublisherTeam, error) {
	team, err := s.Repo.GetBySlug(ctx, slug)
	if err != nil || team == nil {
		return nil, fmt.Errorf("team not found")
	}

	if team.OwnerID != userID {
		return nil, fmt.Errorf("forbidden")
	}

	if input.Description != nil {
		team.Description = *input.Description
	}
	if input.AvatarURL != nil {
		team.AvatarURL = *input.AvatarURL
	}

	if err := s.Repo.Update(ctx, team); err != nil {
		return nil, err
	}
	return team, nil
}

func (s *TeamService) Join(ctx context.Context, userID int, slug string) error {
	team, err := s.Repo.GetBySlug(ctx, slug)
	if err != nil || team == nil {
		return fmt.Errorf("team not found")
	}
	return s.Repo.AddMember(ctx, team.ID, userID)
}

func (s *TeamService) Leave(ctx context.Context, userID int, slug string) error {
	team, err := s.Repo.GetBySlug(ctx, slug)
	if err != nil || team == nil {
		return fmt.Errorf("team not found")
	}
	if team.OwnerID == userID {
		return fmt.Errorf("owner cannot leave")
	}
	return s.Repo.RemoveMember(ctx, team.ID, userID)
}

func (s *TeamService) UploadAvatar(ctx context.Context, file multipart.File, header *multipart.FileHeader, slug string) (string, error) {
	return files.UploadImage(ctx, file, header, fmt.Sprintf("teams/avatars/%s", slug), 400, 400)
}