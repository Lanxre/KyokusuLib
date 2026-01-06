package repository

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/lanxre/kyokusulib/internal/models/db"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetByEmail(email string) (*db.User, error) {
	return r.findOne("email", email)
}

func (r *UserRepository) GetByID(id int) (*db.User, error) {
	return r.findOne("id", id)
}

func (r *UserRepository) GetByVerificationToken(token string) (*db.User, error) {
	return r.findOne("verification_token", token)
}

func (r *UserRepository) GetByResetToken(token string) (*db.User, error) {
	return r.findOne("reset_token", token)
}

func (r *UserRepository) findOne(field string, value interface{}) (*db.User, error) {
	u := &db.User{}

	var (
		passwordHash                 sql.NullString
		verificationToken            sql.NullString
		verificationTokenExpiresAt   sql.NullTime
		resetToken                   sql.NullString
		resetTokenExpiresAt          sql.NullTime
		name                         sql.NullString
		picture                      sql.NullString
		banner                       sql.NullString
		about                        sql.NullString
		birthday                     sql.NullTime
		gender                       sql.NullString
	)

	query := `
		SELECT 
			u.id, u.email, u.password_hash, u.role, u.status, u.last_login,
			u.is_verified, u.isPublic,
			u.verification_token, u.verification_token_expires_at,
			u.reset_token, u.reset_token_expires_at,
			u.create_at,
			p.name, p.picture, p.banner, p.about, p.birthday, p.gender,
			t.tag,
			ups.is_show_tag
		FROM users u
		LEFT JOIN user_profiles p ON p.user_id = u.id
		LEFT JOIN user_tags t ON t.id = p.tag_id
		LEFT JOIN user_profile_settings ups ON ups.user_id = u.id
		WHERE u.` + field + ` = $1`

	err := r.DB.QueryRow(query, value).Scan(
		&u.ID,
		&u.Email,
		&passwordHash,
		&u.Role,
		&u.Status,
		&u.LastLogin,
		&u.IsVerified,
		&u.IsPublic,
		&verificationToken,
		&verificationTokenExpiresAt,
		&resetToken,
		&resetTokenExpiresAt,
		&u.CreateAt,
		&name,
		&picture,
		&banner,
		&about,
		&birthday,
		&gender,
		&u.Tag,
		&u.IsShowTag,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if passwordHash.Valid {
		u.PasswordHash = passwordHash.String
	}
	if verificationToken.Valid {
		u.VerificationToken = verificationToken.String
	}
	if verificationTokenExpiresAt.Valid {
		u.VerificationTokenExpiresAt = &verificationTokenExpiresAt.Time
	}
	if resetToken.Valid {
		u.ResetToken = resetToken.String
	}
	if resetTokenExpiresAt.Valid {
		u.ResetTokenExpiresAt = &resetTokenExpiresAt.Time
	}
	if name.Valid {
		u.Name = name.String
	}
	if picture.Valid {
		u.Picture = picture.String
	}
	if about.Valid {
		u.About = about.String
	}
	if birthday.Valid {
		u.Birthday = &birthday.Time
	}
	if gender.Valid {
		u.Gender = db.UserGenere(gender.String)
	}
	if banner.Valid {
		u.Banner = banner.String
	}
    
	return u, nil
}

func (r *UserRepository) Create(u *db.User) error {
	if u.Role == "" {
		u.Role = "user"
	}

	var passwordHash interface{} = nil
	if u.PasswordHash != "" {
		passwordHash = u.PasswordHash
	}

	var verificationToken interface{} = nil
	if u.VerificationToken != "" {
		verificationToken = u.VerificationToken
	}

	var verificationTokenExpiresAt interface{} = nil
	if u.VerificationTokenExpiresAt != nil {
		verificationTokenExpiresAt = u.VerificationTokenExpiresAt
	}

	query := `
		INSERT INTO users 
			(email, password_hash, role, status, last_login, is_verified, verification_token, verification_token_expires_at) 
		VALUES ($1, $2, $3, 'online', $4, $5, $6, $7) 
		RETURNING id`

	err := r.DB.QueryRow(query,
		u.Email,
		passwordHash,
		u.Role,
		time.Now(),
		u.IsVerified,
		verificationToken,
		verificationTokenExpiresAt,
	).Scan(&u.ID)
	
	if err != nil {
		return err
	}

	return r.createOrUpdateProfile(u)
}

func (r *UserRepository) Update(u *db.User) error {
	var resetToken interface{} = nil
	if u.ResetToken != "" {
		resetToken = u.ResetToken
	}

	var resetTokenExpiresAt interface{} = nil
	if u.ResetTokenExpiresAt != nil {
		resetTokenExpiresAt = u.ResetTokenExpiresAt
	}

	query := `
		UPDATE users 
		SET password_hash = $1, 
		    reset_token = $2, 
		    reset_token_expires_at = $3,
		    status = $4,
		    last_login = $5,
		    is_verified = $6
		WHERE id = $7`

	_, err := r.DB.Exec(query,
		u.PasswordHash,
		resetToken,
		resetTokenExpiresAt,
		u.Status,
		u.LastLogin,
		u.IsVerified,
		u.ID,
	)
	return err
}

func (r *UserRepository) UpdateProfile(u *db.User) error {
	return r.createOrUpdateProfile(u)
}

func (r *UserRepository) createOrUpdateProfile(u *db.User) error {
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var tagID int
	if u.Tag != "" {
		err := tx.QueryRow("SELECT id FROM user_tags WHERE tag = $1", u.Tag).Scan(&tagID)	
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				tagID = 1
			} else {
				return err
			}
		}
	} else {
		tagID = 1 
	}

	query := `
		INSERT INTO user_profiles (user_id, name, picture, about, birthday, gender, tag_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		ON CONFLICT (user_id) DO UPDATE
		SET name = EXCLUDED.name,
		    picture = EXCLUDED.picture,
		    about = EXCLUDED.about,
		    birthday = EXCLUDED.birthday,
		    gender = EXCLUDED.gender,
		    tag_id = EXCLUDED.tag_id`

	_, err = tx.Exec(query,
		u.ID,
		u.Name,
		u.Picture,
		u.About,
		u.Birthday,
		u.Gender,
		tagID,
	)
	if err != nil {
		return err
	}

	query_public := `UPDATE users SET ispublic = $1 WHERE id = $2`
	_, err = tx.Exec(query_public, u.IsPublic, u.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (r *UserRepository) Delete(id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}

func (r *UserRepository) DeleteExpiredUnverified() error {
	query := `
		DELETE FROM users 
		WHERE is_verified = FALSE 
		  AND verification_token_expires_at < $1`
	_, err := r.DB.Exec(query, time.Now())
	return err
}

func (r *UserRepository) UpdateStatus(userID int, status string) error {
	query := `UPDATE users SET status = $1, last_login = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, status, time.Now(), userID)
	return err
}

func (r *UserRepository) UpdateDtoStatus(ctx context.Context, userId int, status string, lastActive time.Time) error {
	query := `UPDATE users SET status = $1, last_login = $2 WHERE id = $3`
	_, err := r.DB.Exec(query, status, lastActive, userId)
	return err
}

func (r *UserRepository) UpdateAvatar(userID int, avatar string) error {
	query := `UPDATE user_profiles SET picture = $1 WHERE user_id = $2`
	_, err := r.DB.Exec(query, avatar, userID)
	return err
}

func (r *UserRepository) UpdateBanner(userID int, banner string) error {
	query := `UPDATE user_profiles SET banner = $1 WHERE user_id = $2`
	_, err := r.DB.Exec(query, banner, userID)
	return err
}


func (r *UserRepository) GetAvatarUrl(userID int) (string, error) {
	var picture sql.NullString

	query := `
		SELECT p.picture
		FROM user_profiles p
		WHERE p.user_id = $1`

	err := r.DB.QueryRow(query, userID).Scan(&picture)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	if picture.Valid {
		return picture.String, nil
	}
	return "", nil
}

func (r *UserRepository) GetBannerUrl(userID int) (string, error) {
	var picture sql.NullString

	query := `
		SELECT p.banner
		FROM user_profiles p
		WHERE p.user_id = $1`

	err := r.DB.QueryRow(query, userID).Scan(&picture)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return "", nil
		}
		return "", err
	}

	if picture.Valid {
		return picture.String, nil
	}
	return "", nil
}

func (r *UserRepository) MarkUserVerified(userID int) error {
	query := `
		UPDATE users 
		SET is_verified = TRUE, 
		    verification_token = NULL, 
		    verification_token_expires_at = NULL 
		WHERE id = $1`
	_, err := r.DB.Exec(query, userID)
	return err
}

func (r *UserRepository) GetUserTags(ctx context.Context, userID int) ([]*db.UserTag, error) {
    const query = `
        SELECT ut.id, ut.tag
        FROM user_tags ut
        JOIN users_user_tags uut ON ut.id = uut.tag_id
        WHERE uut.user_id = $1
        ORDER BY ut.tag
    `

    rows, err := r.DB.QueryContext(ctx, query, userID)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var tags []*db.UserTag
    for rows.Next() {
        var t db.UserTag
        if err := rows.Scan(&t.TagID, &t.Tag); err != nil {
            return nil, err
        }
        tags = append(tags, &t)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return tags, nil
}

func (r *UserRepository) UpdateUserTag(ctx context.Context, userID int, tagID int) error {
	query := `UPDATE user_profiles
		SET tag_id = $2
		WHERE user_id = $1`
		
	_, err := r.DB.Exec(query, tagID, userID)
	return err
}
