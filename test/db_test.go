package test

import (
	"ai-ops-assistant/internal/db"
	"testing"
)

func TestDatabaseConnection(t *testing.T) {
	_, err := db.Connect()
	if err != nil {
		t.Fatalf("Failed to connect to database: %v", err)
	}
}
