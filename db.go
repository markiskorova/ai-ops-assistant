package main

import (
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"ai-ops-assistant/internal/models"
)

var DB *gorm.DB

func initDB() {
	// Load from env or use default
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=localhost user=postgres password=postgres dbname=aiops port=5432 sslmode=disable"
	}

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	// Auto-migrate schema
	if err := DB.AutoMigrate(
		&models.Ticket{},
		&models.LogEntry{}, // 👈 this is the important line
	); err != nil {
		log.Fatal("❌ Failed to migrate:", err)
	}

	log.Println("✅ Database connected and migrated")
}
