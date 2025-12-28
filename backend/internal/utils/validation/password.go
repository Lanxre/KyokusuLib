package validation

import (
	"errors"
	"regexp"
	"unicode"
)

func PasswordValidate(password string) error {
	if password == "" {
		return errors.New("Остутсвуют данные")
	}

	passwordBaseRegex := regexp.MustCompile(`^[a-zA-Z0-9]{6,50}$`)
	if !passwordBaseRegex.MatchString(password) {
		return errors.New("Пароль содержит недопустимые символы или слишком короткий")
	}

	var hasUpper, hasLower, hasDigit bool
	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasDigit = true
		}
	}

	if !hasUpper || !hasLower || !hasDigit {
		return errors.New("Пароль должен содержать латиницу, цифру, заглавную и строчную букву")
	}

	return nil
}