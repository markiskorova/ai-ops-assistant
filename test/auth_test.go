package test

import (
	"ai-ops-assistant/internal/auth"
	"testing"
)

func TestHashAndVerifyPassword(t *testing.T) {
	password := "secure123"
	hashed, err := auth.HashPassword(password)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}
	if !auth.CheckPasswordHash(password, hashed) {
		t.Fatal("Password verification failed")
	}
}
