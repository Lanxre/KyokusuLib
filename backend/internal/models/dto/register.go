package dto

import (
	"errors"

	"github.com/lanxre/kyokusulib/internal/utils/validation"
)

type RegisterDTO struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

func (d *RegisterDTO) Validate() error {
	if d.Email == "" || d.Username == "" || d.Password == "" {
		return errors.New("Заполните все поля")
	}

	if err := validation.EmailValidate(d.Email); err != nil {
		return err
	}
	
	if err := validation.LoginValidate(d.Username); err != nil {
		return err
	}
	
	if err := validation.PasswordValidate(d.Password); err != nil {
		return err
	}

	return nil
}