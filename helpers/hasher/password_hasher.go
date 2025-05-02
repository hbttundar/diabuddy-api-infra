package hasher

import (
	diabuddyErrors "github.com/hbttundar/diabuddy-errors"
	"golang.org/x/crypto/bcrypt"
	"strings"
)

// HashPassword takes a plain password and returns the bcrypt hash of the password
func HashPassword(password string) (string, diabuddyErrors.ApiErrors) {
	if strings.Trim(password, " ") == "" {
		return "", diabuddyErrors.NewApiError(diabuddyErrors.BadRequestErrorType, "empty password is invalid")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", diabuddyErrors.NewApiError(diabuddyErrors.InternalServerErrorType, "error hashing password", diabuddyErrors.WithInternalError(err))
	}
	return string(hashedPassword), nil
}

// CheckPassword compares the hashed password and the plain password to see if they match
func CheckPassword(hashedPassword, plainPassword string) error {
	if strings.Trim(plainPassword, " ") == "" {
		return diabuddyErrors.NewApiError(diabuddyErrors.BadRequestErrorType, "empty password is invalid for checking")
	}
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return diabuddyErrors.NewApiError(diabuddyErrors.BadRequestErrorType, "invalid password", diabuddyErrors.WithInternalError(err))
	}
	return nil
}
