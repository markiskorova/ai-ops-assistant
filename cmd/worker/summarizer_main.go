package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/summarizer"

	"gorm.io/gorm"
)

func main() {
	log.Println("🔁 Starting summarization worker...")
	db.InitDB()

	// Handle Ctrl+C graceful shutdown
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go func() {
		for {
			var entry models.LogEntry

			err := db.DB.
				Where("summary = ''").
				Order("created_at ASC").
				First(&entry).Error

			if err != nil {
				if err == gorm.ErrRecordNotFound {
					log.Println("⏳ No logs to summarize. Retrying in 5s...")
				} else {
					log.Printf("❌ Database error: %v", err)
				}
				time.Sleep(5 * time.Second)
				continue
			}

			log.Printf("📝 Summarizing log ID: %s", entry.ID)
			summary, err := summarizer.Summarize(entry.Raw)
			if err != nil {
				log.Printf("❌ Failed to summarize: %v", err)
				continue
			}

			entry.Summary = summary
			if err := db.DB.Save(&entry).Error; err != nil {
				log.Printf("❌ Failed to save summary: %v", err)
			} else {
				log.Printf("✅ Summary saved for ID: %s", entry.ID)
			}
		}
	}()

	<-stop
	log.Println("🛑 Worker shutting down...")
}
