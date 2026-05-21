package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type TeamRepository struct {
	DB *sql.DB
}

func NewTeamRepository(db *sql.DB) *TeamRepository {
	return &TeamRepository{DB: db}
}

func (r *TeamRepository) GetTeams(ctx context.Context, search string, limit int, offset int, userID int) ([]*db.PublisherTeam, error) {
	var query string
	var rows *sql.Rows
	var err error

	if userID > 0 {
		query = `
			SELECT pt.id, pt.name, pt.slug, pt.description, pt.avatar_url, pt.banner_url, pt.owner_role_name, pt.moderator_role_name, pt.member_role_name, pt.owner_id, pt.created_at,
				(SELECT COUNT(*) FROM team_members WHERE team_id = pt.id) AS member_count,
				(SELECT COUNT(*) FROM team_subscribers WHERE team_id = pt.id) AS subscribers_count,
				CASE WHEN EXISTS (SELECT 1 FROM team_members WHERE team_id = pt.id AND user_id = $1) THEN true ELSE false END AS is_member
			FROM publisher_teams pt
			WHERE EXISTS (SELECT 1 FROM team_members WHERE team_id = pt.id AND user_id = $1)
			   OR EXISTS (SELECT 1 FROM team_subscribers WHERE team_id = pt.id AND user_id = $1)
			ORDER BY pt.id DESC LIMIT $2 OFFSET $3
			`
		rows, err = r.DB.QueryContext(ctx, query, userID, limit, offset)
	} else if search != "" {
		query = `
			SELECT id, name, slug, description, avatar_url, banner_url, owner_role_name, moderator_role_name, member_role_name, owner_id, created_at, 
			(SELECT COUNT(*) FROM team_members WHERE team_id = publisher_teams.id) AS member_count, 
			(SELECT COUNT(*) FROM team_subscribers WHERE team_id = publisher_teams.id) AS subscribers_count 
			FROM publisher_teams 
			WHERE name ILIKE $1::text OR slug ILIKE $1::text 
			ORDER BY id DESC LIMIT $2 OFFSET $3`
		rows, err = r.DB.QueryContext(ctx, query, "%"+search+"%", limit, offset)
	} else {
		query = `
			SELECT id, name, slug, description, avatar_url, banner_url, owner_role_name, moderator_role_name, member_role_name, owner_id, created_at, 
			(SELECT COUNT(*) FROM team_members WHERE team_id = publisher_teams.id) AS member_count, 
			(SELECT COUNT(*) FROM team_subscribers WHERE team_id = publisher_teams.id) AS subscribers_count 
			FROM publisher_teams 
			ORDER BY id DESC LIMIT $1 OFFSET $2`
		rows, err = r.DB.QueryContext(ctx, query, limit, offset)
	}

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var teams []*db.PublisherTeam
	for rows.Next() {
		var team db.PublisherTeam
		var description sql.NullString
		var avatarURL sql.NullString
		var bannerURL sql.NullString
		var memberCount int
		var subscribersCount int
		
		if err := rows.Scan(
			&team.ID, &team.Name, &team.Slug, 
			&description, &avatarURL, &bannerURL,
			&team.OwnerRoleName, &team.ModeratorRoleName, &team.MemberRoleName,
			&team.OwnerID, &team.CreatedAt,
			&memberCount, &subscribersCount,
			&team.IsMember,
		); err != nil {
			return nil, err
		}
		
		if description.Valid {
			team.Description = description.String
		}
		if avatarURL.Valid {
			team.AvatarURL = avatarURL.String
		}
		if bannerURL.Valid {
			team.BannerURL = bannerURL.String
		}
		team.MemberCount = memberCount
		team.SubscribersCount = subscribersCount

		teams = append(teams, &team)
	}
	
	if err := rows.Err(); err != nil {
		return nil, err
	}
	
	return teams, nil
}

var ErrTeamLimitReached = errors.New("team limit reached")

func (r *TeamRepository) Create(ctx context.Context, team *db.PublisherTeam) error {
	tx, err := r.DB.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var isAllowed bool
	checkQuery := `
		SELECT (
			u.role IN ('admin', 'moderator') OR 
			(SELECT COUNT(*) FROM publisher_teams WHERE owner_id = $1) < 5
		)
		FROM users u WHERE u.id = $1`

	err = tx.QueryRowContext(ctx, checkQuery, team.OwnerID).Scan(&isAllowed)
	if err != nil {
		return fmt.Errorf("failed to verify limits: %w", err)
	}

	if !isAllowed {
		return ErrTeamLimitReached
	}


	query := `
		INSERT INTO publisher_teams (owner_id, name, slug, description)
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at, updated_at`

	err = tx.QueryRowContext(ctx, query,
		team.OwnerID,
		team.Name,
		team.Slug,
		team.Description,
	).Scan(&team.ID, &team.CreatedAt, &team.UpdatedAt)
	if err != nil {
		return err
	}

	memberQuery := `INSERT INTO team_members (team_id, user_id) VALUES ($1, $2)`
	if _, err := tx.ExecContext(ctx, memberQuery, team.ID, team.OwnerID); err != nil {
		return err
	}

	return tx.Commit()
}

func (r *TeamRepository) GetBySlug(ctx context.Context, slug string, userID int) (*db.PublisherTeam, error) {
	team := &db.PublisherTeam{}
	var description sql.NullString
	var avatarURL sql.NullString
	var bannerURL sql.NullString

	query := `
		SELECT id, owner_id, name, slug, description, avatar_url, banner_url, owner_role_name, moderator_role_name, member_role_name, created_at, updated_at,
		(SELECT COUNT(*) FROM team_members WHERE team_id = publisher_teams.id) AS member_count,
		(SELECT COUNT(*) FROM team_subscribers WHERE team_id = publisher_teams.id) AS subscribers_count,
		EXISTS(SELECT 1 FROM team_members WHERE team_id = publisher_teams.id AND user_id = $2) AS is_member
		FROM publisher_teams WHERE slug = $1`

	err := r.DB.QueryRowContext(ctx, query, slug, userID).Scan(
		&team.ID, &team.OwnerID, &team.Name, &team.Slug,
		&description, &avatarURL, &bannerURL, 
		&team.OwnerRoleName, &team.ModeratorRoleName, &team.MemberRoleName,
		&team.CreatedAt, &team.UpdatedAt,
		&team.MemberCount, &team.SubscribersCount,
		&team.IsMember,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if description.Valid {
		team.Description = description.String
	}
	if avatarURL.Valid {
		team.AvatarURL = avatarURL.String
	}
	if bannerURL.Valid {
		team.BannerURL = bannerURL.String
	}

	if description.Valid {
		team.Description = description.String
	}
	if avatarURL.Valid {
		team.AvatarURL = avatarURL.String
	}
	
	return team, nil
}

func (r *TeamRepository) Update(ctx context.Context, team *db.PublisherTeam) error {
	query := `
		UPDATE publisher_teams 
		SET description = $1, avatar_url = $2, banner_url = $3, 
		    owner_role_name = $4, moderator_role_name = $5, member_role_name = $6,
		    updated_at = NOW()
		WHERE id = $7`
	_, err := r.DB.ExecContext(ctx, query, 
		team.Description, team.AvatarURL, team.BannerURL, 
		team.OwnerRoleName, team.ModeratorRoleName, team.MemberRoleName, 
		team.ID)
	return err
}

func (r *TeamRepository) GetMembers(ctx context.Context, slug string, limit int, offset int) ([]*db.TeamMemberUser, error) {
	query := `
		SELECT 
			u.id, up.name, up.picture, up.name as tag, 
			up.level, up.experience, ld.title, COALESCE(next_ld.total_xp_required, ld.total_xp_required) AS xp_for_next_level,
			tm.role, tm.custom_role_name, tm.joined_at,
			pt.owner_role_name, pt.moderator_role_name, pt.member_role_name
		FROM team_members tm
		JOIN publisher_teams pt ON tm.team_id = pt.id
		JOIN users u ON tm.user_id = u.id
		LEFT JOIN user_profiles up ON u.id = up.user_id
		LEFT JOIN level_definitions ld ON up.level = ld.level
		LEFT JOIN level_definitions next_ld ON next_ld.level = up.level + 1
		LEFT JOIN user_tags ut ON up.tag_id = ut.id
		WHERE pt.slug = $1
		ORDER BY 
			CASE tm.role 
				WHEN 'owner' THEN 1 
				WHEN 'moderator' THEN 2 
				ELSE 3 
			END, 
			tm.joined_at ASC
		LIMIT $2 OFFSET $3`

	rows, err := r.DB.QueryContext(ctx, query, slug, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []*db.TeamMemberUser
	for rows.Next() {
		var m db.TeamMemberUser
		var tag sql.NullString
		var picture sql.NullString
		var level, experience, xpForNextLevel sql.NullInt64
		var levelTitle sql.NullString
		var customRoleName, ownerRoleName, moderatorRoleName, memberRoleName sql.NullString

		if err := rows.Scan(
			&m.User.ID, &m.User.Name, &picture, &tag,
			&level, &experience, &levelTitle, &xpForNextLevel,
			&m.Role, &customRoleName, &m.JoinedAt,
			&ownerRoleName, &moderatorRoleName, &memberRoleName,
		); err != nil {
			return nil, err
		}

		if picture.Valid {
			m.User.Picture = picture.String
		}
		if tag.Valid {
			m.User.Tag = tag.String
		}

		if level.Valid {
			m.User.UserLevel.Level = int(level.Int64)
		}
		if experience.Valid {
			m.User.UserLevel.Experience = experience.Int64
		}
		if levelTitle.Valid {
			m.User.UserLevel.LevelTitle = levelTitle.String
		}
		if xpForNextLevel.Valid && experience.Valid {
			needed := max(xpForNextLevel.Int64 - experience.Int64, 0)
			m.User.UserLevel.XPForNext = needed
    	}

		if customRoleName.Valid && customRoleName.String != "" {
			m.CustomRole = customRoleName.String
		} else {
			switch m.Role {
			case "owner":
				m.CustomRole = ownerRoleName.String
				if m.CustomRole == "" { m.CustomRole = "Владелец" }
			case "moderator":
				m.CustomRole = moderatorRoleName.String
				if m.CustomRole == "" { m.CustomRole = "Модератор" }
			default:
				m.CustomRole = memberRoleName.String
				if m.CustomRole == "" { m.CustomRole = "Участник" }
			}
		}

		members = append(members, &m)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return members, nil
}

func (r *TeamRepository) Delete(ctx context.Context, id int) error {
	query := `DELETE FROM publisher_teams WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}

func (r *TeamRepository) AddMember(ctx context.Context, teamID, userID int) error {
	query := `INSERT INTO team_members (team_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.DB.ExecContext(ctx, query, teamID, userID)
	return err
}

func (r *TeamRepository) RemoveMember(ctx context.Context, teamID, userID int) error {
	query := `DELETE FROM team_members WHERE team_id = $1 AND user_id = $2`
	_, err := r.DB.ExecContext(ctx, query, teamID, userID)
	return err
}

func (r *TeamRepository) Subscribe(ctx context.Context, teamID, userID int) error {
	query := `INSERT INTO team_subscribers (team_id, user_id) VALUES ($1, $2) ON CONFLICT DO NOTHING`
	_, err := r.DB.ExecContext(ctx, query, teamID, userID)
	return err
}

func (r *TeamRepository) Unsubscribe(ctx context.Context, teamID, userID int) error {
	query := `DELETE FROM team_subscribers WHERE team_id = $1 AND user_id = $2`
	_, err := r.DB.ExecContext(ctx, query, teamID, userID)
	return err
}
