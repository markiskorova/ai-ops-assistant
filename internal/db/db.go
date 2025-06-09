package db

import (
	"fmt"
	"log"
	"os"

	"ai-ops-assistant/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	_ = godotenv.Load() // Load .env silently

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, pass, name, port,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	if err := DB.AutoMigrate(&models.Ticket{}, &models.LogEntry{}); err != nil {
		log.Fatal("❌ Failed to migrate:", err)
	}

	log.Println("✅ Database initialized and models migrated")
	return DB
}
