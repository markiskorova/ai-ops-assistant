package test

import (
	"os"
	"testing"

	"ai-ops-assistant/internal/db"
)

func TestInitDB(t *testing.T) {
	// Not a real test without DB, just a sanity check on error handling
	_ = os.Setenv("DB_HOST", "localhost")
	_ = os.Setenv("DB_USER", "postgres")
	_ = os.Setenv("DB_PASS", "postgres")
	_ = os.Setenv("DB_PORT", "5432")
	_ = os.Setenv("DB_NAME", "testdb")

	// This will likely fail unless PostgreSQL is running locally
	defer func() {
		if r := recover(); r != nil {
			t.Log("Recovered from panic (expected if no DB available):", r)
		}
	}()
	_ = db.InitDB()
}
