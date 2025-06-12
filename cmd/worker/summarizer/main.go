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
	dbConn := db.InitDB()
	summarizer := summarizer.NewSummarizerFromEnv()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	go runSummarizerLoop(dbConn, summarizer)

	<-stop
	log.Println("🛑 Summarizer worker stopped.")
}

func runSummarizerLoop(db *gorm.DB, s summarizer.Summarizer) {
	for {
		var entry models.LogEntry

		err := db.
			Where("summary = ''").
			Order("created_at ASC").
			First(&entry).Error

		if err != nil {
			log.Printf("⏳ %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("📝 Summarizing log ID: %s", entry.ID)
		summary, err := s.Summarize(entry.Raw)
		if err != nil {
			log.Printf("❌ Failed to summarize: %v", err)
			continue
		}

		entry.Summary = summary
		if err := db.Save(&entry).Error; err != nil {
			log.Printf("❌ Failed to save summary: %v", err)
		} else {
			log.Printf("✅ Summary saved for ID: %s", entry.ID)
		}
	}
}
