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

const (
	MemberRoleModerator = "moderator"
	MemberRoleOwner     = "owner"
	MemberRoleMember    = "member"

	TeamTypePublic  = "open"
	TeamTypePrivate = "private"
)

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

	if input.TeamType == "" {
		input.TeamType = TeamTypePublic
	}

	team := &db.PublisherTeam{
		OwnerID:     userID,
		Name:        input.Name,
		Slug:        input.Slug,
		Description: input.Description,
		TeamType:    input.TeamType,
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

func (s *TeamService) Update(ctx context.Context, userID int, slug string, input dto.UpdateTeamDTO) (*dto.TeamDTO, error) {
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
	if input.TeamType != nil {
		team.TeamType = *input.TeamType
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
	return s.mapTeamToDTO(team), nil
}

func (s *TeamService) Join(ctx context.Context, userID int, slug string) (string, error) {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
	if err != nil || team == nil {
		return "", fmt.Errorf("Команда не найдена")
	}
	if team.TeamType == TeamTypePrivate {
		err := s.Repo.CreateJoinRequest(ctx, team.ID, userID)
		if err != nil {
			return "", err
		}
		return "request_sent", nil
	}
	err = s.Repo.AddMember(ctx, team.ID, userID)
	if err != nil {
		return "", err
	}
	return "joined", nil
}

func (s *TeamService) checkModPermission(ctx context.Context, team *db.PublisherTeam, modID int) (bool, error) {
	if team.OwnerID == modID {
		return true, nil
	}

	role, err := s.Repo.GetMemberRole(ctx, team.ID, modID)
	if err != nil {
		return false, err
	}
	
	return role == MemberRoleModerator || role == MemberRoleOwner, nil
}

func (s *TeamService) AddMemberByMod(ctx context.Context, modID int, slug string, targetUserID int) error {
	team, err := s.Repo.GetBySlug(ctx, slug, modID)
	if err != nil || team == nil {
		return fmt.Errorf("Команда не найдена")
	}

	isMod, err := s.checkModPermission(ctx, team, modID)
	if err != nil {
		return err
	}
	
	if !isMod {
		return fmt.Errorf("Доступ запрещен: только модератор или создатель могут добавлять участников")
	}

	return s.Repo.AddMember(ctx, team.ID, targetUserID)
}

func (s *TeamService) UpdateMember(ctx context.Context, modID int, slug string, targetUserID int, input dto.UpdateTeamMemberDTO) error {
	team, err := s.Repo.GetBySlug(ctx, slug, modID)
	if err != nil || team == nil {
		return fmt.Errorf("Команда не найдена")
	}

	isMod, err := s.checkModPermission(ctx, team, modID)
	if err != nil {
		return err
	}
	if !isMod {
		return fmt.Errorf("Доступ запрещен: только модератор или создатель могут обновлять участников")
	}
	
	targetMember, err := s.Repo.GetTeamMember(ctx, team.ID, targetUserID)
	
	if err != nil {
		return err
	}
	if targetMember == nil {
		return fmt.Errorf("Пользователь не является участником этой команды")
	}
	
	if targetMember.Role == MemberRoleOwner && modID != team.OwnerID {
		return fmt.Errorf("Доступ запрещен: только владелец может изменять другого владельца")
	}
	
	if input.Role != nil && *input.Role == MemberRoleOwner {
		return fmt.Errorf("Доступ запрещен: нельзя назначить роль владельца")
	}

	roleToSave := targetMember.Role
	if input.Role != nil {
		roleToSave = *input.Role
	}
	
	var customRoleToSave *string = input.CustomRoleName

	return s.Repo.UpdateMember(ctx, team.ID, targetUserID, roleToSave, customRoleToSave)
}

func (s *TeamService) AcceptJoinRequest(ctx context.Context, modID int, slug string, targetUserID int) error {
	team, err := s.Repo.GetBySlug(ctx, slug, modID)
	if err != nil || team == nil {
		return fmt.Errorf("Команда не найдена")
	}

	isMod, err := s.checkModPermission(ctx, team, modID)
	if err != nil {
		return err
	}
	if !isMod {
		return fmt.Errorf("Доступ запрещен: только модератор или создатель могут добавлять участников")
	}

	if err := s.Repo.AddMember(ctx, team.ID, targetUserID); err != nil {
		return err
	}
	return s.Repo.DeleteJoinRequest(ctx, team.ID, targetUserID)
}

func (s *TeamService) RejectJoinRequest(ctx context.Context, modID int, slug string, targetUserID int) error {
	team, err := s.Repo.GetBySlug(ctx, slug, modID)
	if err != nil || team == nil {
		return fmt.Errorf("Команда не найдена")
	}

	isMod, err := s.checkModPermission(ctx, team, modID)
	if err != nil {
		return err
	}
	if !isMod {
		return fmt.Errorf("Доступ запрещен: только модератор или создатель могут отклонять запросы")
	}

	return s.Repo.DeleteJoinRequest(ctx, team.ID, targetUserID)
}

func (s *TeamService) GetJoinRequests(ctx context.Context, modID int, slug string, limit, offset int) ([]*dto.TeamJoinRequestDTO, error) {
	team, err := s.Repo.GetBySlug(ctx, slug, modID)
	if err != nil || team == nil {
		return nil, fmt.Errorf("team not found")
	}

	isMod, err := s.checkModPermission(ctx, team, modID)
	if err != nil {
		return nil, err
	}
	if !isMod {
		return nil, fmt.Errorf("forbidden: only moderator or creator can view requests")
	}

	requests, err := s.Repo.GetJoinRequests(ctx, slug, limit, offset)
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.TeamJoinRequestDTO, len(requests))
	for i, req := range requests {
		dtos[i] = &dto.TeamJoinRequestDTO{
			User: dto.TeamMemberUserDTO{
				ID:        req.User.ID,
				Name:      req.User.Name,
				Picture:   req.User.Picture,
				ActiveTag: req.User.Tag,
				UserLevel: dto.UserLevelDTO{
					Level:      req.User.UserLevel.Level,
					Experience: req.User.UserLevel.Experience,
					LevelTitle: req.User.UserLevel.LevelTitle,
					XPForNext:  req.User.UserLevel.XPForNext,
				},
			},
			CreatedAt: req.CreatedAt.Format(time.RFC3339),
		}
	}
	return dtos, nil
}

func (s *TeamService) Leave(ctx context.Context, userID int, slug string) error {
	team, err := s.Repo.GetBySlug(ctx, slug, userID)
	if err != nil || team == nil {
		return fmt.Errorf("Команда не найдена")
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

func (s *TeamService) SubscribeToTeam(ctx context.Context, teamID int, userID int) error {

	isSubscribed, err := s.Repo.IsSubscribed(ctx, teamID, userID)
	if err != nil {
		return err
	}
	
	if isSubscribed {
		return fmt.Errorf("Вы уже подписаны")
	}
	
	return s.Repo.Subscribe(ctx, teamID, userID)
}

func (s *TeamService) UnsubscribeFromTeam(ctx context.Context, teamID int, userID int) error {
	isSubscribed, err := s.Repo.IsSubscribed(ctx, teamID, userID)
	if err != nil {
		return err
	}
	
	if !isSubscribed {
		return fmt.Errorf("Вы уже не подписаны")
	}
	return s.Repo.Unsubscribe(ctx, teamID, userID)
}

func (s *TeamService) GetSubscribers(ctx context.Context, slug string, limit, offset int) ([]*dto.TeamSubscriberDTO, error) {
	subscribers, err := s.Repo.GetSubscribers(ctx, slug, limit, offset)
	if err != nil {
		return nil, err
	}

	dtos := make([]*dto.TeamSubscriberDTO, len(subscribers))
	for i, sub := range subscribers {
		dtos[i] = s.mapSubscriberToDTO(sub)
	}

	return dtos, nil
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
		TeamType:    team.TeamType,
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
		IsSubscriber: team.IsSubscriber,
		HasRequested: team.HasRequested,
	}
}

func (s *TeamService) mapSubscriberToDTO(sub *db.TeamSubscriberUser) *dto.TeamSubscriberDTO {
	return &dto.TeamSubscriberDTO{
		User: dto.TeamMemberUserDTO{
			ID:        sub.User.ID,
			Name:      sub.User.Name,
			Picture:   sub.User.Picture,
			ActiveTag: sub.User.Tag,
			UserLevel: dto.UserLevelDTO{
				Level:      sub.User.UserLevel.Level,
				Experience: sub.User.UserLevel.Experience,
				LevelTitle: sub.User.UserLevel.LevelTitle,
				XPForNext:  sub.User.UserLevel.XPForNext,
			},
		},
		SubscribedAt: sub.SubscribedAt.Format(time.RFC3339),
	}
}