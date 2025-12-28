package validation

import (
	"errors"
	"regexp"
	"unicode"
)

func LoginValidate(login string) error {
	usernameBaseRegex := regexp.MustCompile(`^[a-zA-Z]{6,50}$`)
	if !usernameBaseRegex.MatchString(login) {
		return errors.New("Никнейм должен содержать только латиницу (6-50 символов)")
	}

	var hasUserUpper, hasUserLower bool
	for _, char := range login {
		switch {
		case unicode.IsUpper(char):
			hasUserUpper = true
		case unicode.IsLower(char):
			hasUserLower = true
		}
	}

	if !hasUserUpper || !hasUserLower {
		return errors.New("Никнейм должен содержать минимум 1 большую и 1 маленькую букву")
	}

	return nil
}