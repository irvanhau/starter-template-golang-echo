package helper

import (
	"github.com/go-playground/validator/v10"
	"unicode"
)

var validate *validator.Validate

func ValidateJSON(data interface{}) (bool, map[string]string) {
	validate = validator.New()

	if err := validate.Struct(data); err != nil {
		var errors = map[string]string{}

		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = err.Tag()
		}
		return false, errors
	}
	return true, nil
}

func ValidatePassword(password string) bool {
	hasLetter := false
	hasDigit := false
	hasSymbol := false

	for _, char := range password {
		if unicode.IsLetter(char) {
			hasLetter = true
		} else if unicode.IsDigit(char) {
			hasDigit = true
		} else if unicode.IsPunct(char) || unicode.IsSymbol(char) {
			hasSymbol = true
		}
	}

	result := hasLetter && hasDigit && hasSymbol

	return result
}
