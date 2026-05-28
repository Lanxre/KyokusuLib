package service

import (
	"context"
	"fmt"
	"mime/multipart"
	"time"

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

func (s *TeamService) GetTeams(ctx context.Context, search string, limit int, offset int, userID int) ([]*dto.TeamDTO, error) {
	if limit <= 0 {
		limit = 10
	} else if limit > 100 {
		limit = 100
	}
	
	teams, err := s.Repo.GetTeams(ctx, search, limit, offset, userID)
	if err != nil {
		return nil, err
	}

	dtoTeams := make([]*dto.TeamDTO, len(teams))
	for i, team := range teams {
		dtoTeams[i] = s.mapTeamToDTO(team)
	}
	
	return dtoTeams, nil
}

func (s *TeamService) GetMembers(ctx context.Context, slug string, limit int, offset int) ([]*dto.TeamMemberDTO, error) {
	if limit <= 0 {
		limit = 20
	} else if limit > 100 {
		limit = 100
	}
	
	members, err := s.Repo.GetMembers(ctx, slug, limit, offset)
	if err != nil {
		return nil, err
	}

	dtoMembers := make([]*dto.TeamMemberDTO, len(members))
	for i, m := range members {
		dtoMembers[i] = &dto.TeamMemberDTO{
			User: dto.TeamMemberUserDTO{
				ID:      m.User.ID,
				Name:    m.User.Name,
				Picture: m.User.Picture,
				ActiveTag: m.User.Tag,
				UserLevel: dto.UserLevelDTO{
					Level:      m.User.UserLevel.Level,
					Experience: m.User.UserLevel.Experience,
					LevelTitle: m.User.UserLevel.LevelTitle,
					XPForNext:  m.User.UserLevel.XPForNext,
				},
			},
			Role:     m.Role,
			RoleName: m.CustomRole,
			JoinedAt: m.JoinedAt.Format(time.RFC3339),
		}
	}
	
	return dtoMembers, nil
}

func (s *TeamService) Create(ctx context.Context, userID int, input dto.CreateTeamDTO) (*db.PublisherTeam, error) {
	existing, err := s.Repo.GetBySlug(ctx, input.Slug, 0)
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

func (s *TeamService) Get(ctx context.Context, slug string, userID int) (*dto.TeamDTO, error) {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
	if err != nil {
		return nil, err
	}

	if team == nil {
		return nil, fmt.Errorf("Team not found")
	}
	
	return s.mapTeamToDTO(team), nil
}

func (s *TeamService) Update(ctx context.Context, userID int, slug string, input dto.UpdateTeamDTO) (*db.PublisherTeam, error) {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
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
	if input.BannerURL != nil {
		team.BannerURL = *input.BannerURL
	}
	if input.OwnerRoleName != nil {
		team.OwnerRoleName = *input.OwnerRoleName
	}
	if input.ModeratorRoleName != nil {
		team.ModeratorRoleName = *input.ModeratorRoleName
	}
	if input.MemberRoleName != nil {
		team.MemberRoleName = *input.MemberRoleName
	}

	if err := s.Repo.Update(ctx, team); err != nil {
		return nil, err
	}
	return team, nil
}

func (s *TeamService) Join(ctx context.Context, userID int, slug string) error {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
	if err != nil || team == nil {
		return fmt.Errorf("team not found")
	}
	return s.Repo.AddMember(ctx, team.ID, userID)
}

func (s *TeamService) Leave(ctx context.Context, userID int, slug string) error {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
	if err != nil || team == nil {
		return fmt.Errorf("team not found")
	}
	if team.OwnerID == userID {
		if err := s.Repo.DeleteTeam(ctx, team.ID); err != nil {
			return err
		}
		return nil
	}
	return s.Repo.RemoveMember(ctx, team.ID, userID)
}

func (s *TeamService) UploadAvatar(ctx context.Context, file multipart.File, header *multipart.FileHeader, slug string) (string, error) {
	return files.UploadImage(ctx, file, header, fmt.Sprintf("teams/avatars/%s", slug), 400, 400)
}

func (s *TeamService) UploadBanner(ctx context.Context, file multipart.File, header *multipart.FileHeader, slug string) (string, error) {
	return files.UploadImage(ctx, file, header, fmt.Sprintf("teams/banners/%s", slug), 1200, 400)
}

func (s TeamService) DeleteTeam(ctx context.Context, teamID int) error {
	return s.Repo.DeleteTeam(ctx, teamID)
}

func (s *TeamService) mapTeamToDTO(team *db.PublisherTeam) *dto.TeamDTO {
	return &dto.TeamDTO{
		ID:          team.ID,
		OwnerID:     team.OwnerID,
		Name:        team.Name,
		Slug:        team.Slug,
		Description: team.Description,
		AvatarURL:   team.AvatarURL,
		BannerURL:   team.BannerURL,
		RoleNames: dto.TeamRoleNames{
			Owner:     team.OwnerRoleName,
			Moderator: team.ModeratorRoleName,
			Member:    team.MemberRoleName,
		},
		CreatedAt:    team.CreatedAt.Format(time.RFC3339),
		Stats: dto.TeamStats{
			MemberCount:      team.MemberCount,
			SubscribersCount: team.SubscribersCount,
		},
		IsMember:     team.IsMember,
	}
}