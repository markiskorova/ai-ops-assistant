package db

import (
	"fmt"
	"log"
	"os"
	"time"

	"ai-ops-assistant/internal/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

	var db *gorm.DB
	var err error
	for attempts := 1; attempts <= 5; attempts++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("DB not ready, retrying in %d seconds...", attempts*2)
		time.Sleep(time.Duration(attempts*2) * time.Second)
	}
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&models.Ticket{}, &models.LogEntry{}, &models.ChangelogJob{}); err != nil {
		log.Fatal("❌ Failed to migrate:", err)
	}

	log.Println("✅ Database initialized and models migrated")
	return db
}
