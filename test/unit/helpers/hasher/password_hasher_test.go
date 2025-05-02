package hasher_test

import (
	"github.com/hbttundar/diabuddy-api-infra/helpers/hasher"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashPassword(t *testing.T) {
	tests := []struct {
		name        string
		password    string
		expectError bool
	}{
		{"Hash Password Successfully", "ValidPassword123!", false},
		{"Invalid password to hash", "       ", true},
		{"too long to hash", " ValidPassword123!ValidPassword123!ValidPassword123!ValidPassword123!ValidPassword123!", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, err := hasher.HashPassword(tt.password)
			if tt.expectError {
				assert.Empty(t, hashedPassword)
				assert.Error(t, err)
			} else {
				assert.NotEmpty(t, hashedPassword)
				assert.NoError(t, err)
			}
		})
	}
}

func TestCheckPasswordHash(t *testing.T) {
	tests := []struct {
		name             string
		password         string
		checkingPassword string
		expectError      bool
	}{
		{"Check Matching Password", "ValidPassword123!", "ValidPassword123!", false},
		{"Check Non-Matching Password", "ValidPassword123!", "DifferentPassword456!", true},
		{"Check empty Password", "ValidPassword123!", "    ", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hashedPassword, _ := hasher.HashPassword(tt.password)
			err := hasher.CheckPassword(hashedPassword, tt.checkingPassword)
			if tt.expectError {
				assert.NotNil(t, err, "password should not match the hash")

			} else {
				assert.Nil(t, err, "password should match the hash")
			}
		})
	}
}

func TestHashPassword_Error(t *testing.T) {
	// This test is a bit artificial because bcrypt does not usually fail unless there is an internal error.
	// However, we will keep it for completion and future purposes where hashing can introduce errors.
	t.Run("Hash Password With Error Handling", func(t *testing.T) {
		password := ""
		_, err := hasher.HashPassword(password)
		assert.Error(t, err, "hashing an empty password should result in an error")
	})
}
