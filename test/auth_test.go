package test

import (
	"os"
	"testing"

	"ai-ops-assistant/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestHashAndComparePassword(t *testing.T) {
	password := "supersecure"
	hash, err := auth.HashPassword(password)
	assert.NoError(t, err)
	assert.True(t, auth.ComparePasswords(password, hash))
}

func TestGenerateAndValidateJWT(t *testing.T) {
	_ = os.Setenv("JWT_SECRET", "testsecret")

	token, err := auth.GenerateJWT("user-123")
	assert.NoError(t, err)
	assert.NotEmpty(t, token)

	userID, err := auth.ValidateJWT(token)
	assert.NoError(t, err)
	assert.Equal(t, "user-123", userID)
}
