package dto

import (
	"github.com/lanxre/kyokusulib/internal/utils/validation"
)

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (d *LoginDTO) Validate() error {
	if err := validation.EmailValidate(d.Email); err != nil {
		return err
	}

	if err := validation.PasswordValidate(d.Password); err != nil {
		return err
	}

	return nil
}