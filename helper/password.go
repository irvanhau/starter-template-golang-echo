package helper

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashed, input string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashed), []byte(input))
}

func HashPassword(password string) (string, error) {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return "", err
	}
	return string(hashPass), nil
}
