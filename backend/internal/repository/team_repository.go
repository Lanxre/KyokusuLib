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

func (r *TeamRepository) GetTeams(ctx context.Context, search string, limit int) ([]*db.PublisherTeam, error) {
	var query string
	var rows *sql.Rows
	var err error

	if search != "" {
		query = "SELECT id, name, slug, description, avatar_url, owner_id, created_at FROM publisher_teams WHERE name ILIKE $1::text OR slug ILIKE $1::text ORDER BY id DESC LIMIT $2"
		rows, err = r.DB.QueryContext(ctx, query, "%"+search+"%", limit)
	} else {
		query = "SELECT id, name, slug, description, avatar_url, owner_id, created_at FROM publisher_teams ORDER BY id DESC LIMIT $1"
		rows, err = r.DB.QueryContext(ctx, query, limit)
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
		
		if err := rows.Scan(
			&team.ID, &team.Name, &team.Slug, 
			&description, &avatarURL, 
			&team.OwnerID, &team.CreatedAt,
		); err != nil {
			return nil, err
		}
		
		if description.Valid {
			team.Description = description.String
		}
		if avatarURL.Valid {
			team.AvatarURL = avatarURL.String
		}
		
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

func (r *TeamRepository) GetBySlug(ctx context.Context, slug string) (*db.PublisherTeam, error) {
	team := &db.PublisherTeam{}
	query := `
		SELECT id, owner_id, name, slug, description, avatar_url, created_at, updated_at
		FROM publisher_teams WHERE slug = $1`

	err := r.DB.QueryRowContext(ctx, query, slug).Scan(
		&team.ID, &team.OwnerID, &team.Name, &team.Slug,
		&team.Description, &team.AvatarURL, &team.CreatedAt, &team.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return team, nil
}

func (r *TeamRepository) Update(ctx context.Context, team *db.PublisherTeam) error {
	query := `
		UPDATE publisher_teams 
		SET description = $1, avatar_url = $2, updated_at = NOW()
		WHERE id = $3`
	_, err := r.DB.ExecContext(ctx, query, team.Description, team.AvatarURL, team.ID)
	return err
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
