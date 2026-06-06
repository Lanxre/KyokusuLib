package dto

import (
	"errors"
	"time"
	"unicode/utf8"

	"github.com/lanxre/kyokusulib/internal/utils/validation"
)

type UpdateProfileDTO struct {
	Nickname string `json:"name"`
	About    string `json:"about"`
	Gender   string `json:"gender"`
	Birthday string `json:"birthday"`
	IsPublic bool   `json:"is_public"`
}

type UpdateAvatarDTO struct {
	Avatar string `json:"picture"`
}

type ChangePasswordDTO struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

type ChangeEmaildDTO struct {
	Email 		string `json:"email"`
	OldPassword string `json:"old_password"`
}

type ProfileSettingsDTO struct {
	Theme 	string `json:"theme"`
}

type NotifySettingsPatchDTO struct {
	IsAppNotify 		 *bool `json:"is_app_notify"`
	IsNewPublishedNotify *bool `json:"is_new_published_notify"`
}

type UserInterfacePatchDTO struct {
	IsShowTag *bool `json:"is_show_tag"`
	IsShowBookmark *bool `json:"is_show_bookmark"`
}

type ReaderSettingsPatchDTO struct {
	FontSize   *int     `json:"font_size"`
	LineWeight *float64 `json:"line_weight"`
	FontFamily *string  `json:"font_family"`
	AutoScroll *bool    `json:"auto_scroll"`
}

func (d *UpdateProfileDTO) Validate() error {
	if err := validation.LoginValidate(d.Nickname); err != nil {
		return err
	}

	if utf8.RuneCountInString(d.About) > 500 {
		return errors.New("about section too long")
	}

	switch d.Gender {
	case "male", "female", "hidden":
	default:
		return errors.New("invalid gender")
	}

	if d.Birthday != "" {
		_, err := time.Parse("2006-01-02", d.Birthday)
		if err != nil {
			return errors.New("invalid birthday format")
		}
	}

	return nil
}

func (d *ChangePasswordDTO) Validate() error {
	if d.OldPassword == "" {
		return errors.New("Остутсвуют данные")
	}

	if err := validation.PasswordValidate(d.OldPassword); err != nil {
		return err
	}

	return nil
}

func (d *ChangeEmaildDTO) Validate() error {
	if err := validation.EmailValidate(d.Email); err != nil {
		return err
	}

	if err := validation.PasswordValidate(d.OldPassword); err != nil {
		return err
	}

	return nil
}

