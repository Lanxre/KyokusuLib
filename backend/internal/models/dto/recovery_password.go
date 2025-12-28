package dto

import (
	"errors"
	"github.com/lanxre/kyokusulib/internal/utils/validation"
)

type ForgotPasswordDTO struct {
	Email string `json:"email"`
}

type ResetPasswordDTO struct {
	Token    string `json:"token"`
	Password string `json:"password"`
}

func (d *ResetPasswordDTO) Validate() error {
	if d.Token == "" {
		return errors.New("Неправильный токен")
	}

	err := validation.PasswordValidate(d.Password)
	return err
}

func (d *ForgotPasswordDTO) Validate() error {
	if err := validation.EmailValidate(d.Email); err != nil {
		return err
	}
	
	return nil
}