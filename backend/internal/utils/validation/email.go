package validation

import (
	"errors"
	"regexp"
)

func EmailValidate(email string) error {
	emailRegex := regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)
	if !emailRegex.MatchString(email) {
		return errors.New("Некорректный формат email")
	}

	return nil
}