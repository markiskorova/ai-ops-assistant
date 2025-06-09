package main

import (
	"log"
	"time"

	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/summarizer"
)

func main() {
	log.Println("🔁 Starting summarization worker...")
	db.Init()

	for {
		var entry models.LogEntry

		err := db.DB.
			Where("summary = ''").
			Order("created_at ASC").
			First(&entry).Error

		if err != nil {
			log.Println("No logs to summarize. Retrying in 5s...")
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("📝 Summarizing log ID: %s\n", entry.ID)
		summary, err := summarizer.Summarize(entry.Raw)
		if err != nil {
			log.Printf("❌ Failed to summarize: %v", err)
			continue
		}

		entry.Summary = summary
		if err := db.DB.Save(&entry).Error; err != nil {
			log.Printf("❌ Failed to save summary: %v", err)
		} else {
			log.Printf("✅ Summary saved for ID: %s\n", entry.ID)
		}
	}
}
