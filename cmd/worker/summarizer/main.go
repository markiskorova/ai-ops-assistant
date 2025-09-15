package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/observability/workermetrics"
	"ai-ops-assistant/internal/summarizer"

	"gorm.io/gorm"
)

func main() {
	log.Println("🧠 Starting summarizer worker...")
	dbConn := db.InitDB()
	s := summarizer.NewSummarizerFromEnv()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Metrics server for summarizer worker
	go workermetrics.StartServer(":9102")

	go runSummarizerLoop(dbConn, s)

	<-stop
	log.Println("🛑 Summarizer worker stopped.")
}

func runSummarizerLoop(db *gorm.DB, s summarizer.Summarizer) {
	for {
		// Optional: expose queue depth (logs needing summaries)
		var pending int64
		if err := db.Model(&models.LogEntry{}).
			Where("summary = '' OR summary IS NULL").
			Count(&pending).Error; err == nil {
			workermetrics.SetQueueDepth(int(pending))
		}

		var logEntry models.LogEntry
		err := db.
			Where("summary = '' OR summary IS NULL").
			Order("created_at ASC").
			First(&logEntry).Error

		if err != nil {
			log.Printf("⏳ %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("📝 Summarizing log: %s", logEntry.ID)

		// Metrics: one unit of work
		workermetrics.IncStarted()
		timer := workermetrics.NewTimer()

		summary, err := s.Summarize(logEntry.Raw)
		if err != nil {
			log.Printf("❌ Summarization error: %v", err)
			timer.ObserveDuration()
			workermetrics.IncFailed()
			continue
		}

		logEntry.Summary = summary

		if err := db.Save(&logEntry).Error; err != nil {
			log.Printf("❌ Failed to save summary: %v", err)
			timer.ObserveDuration()
			workermetrics.IncFailed()
		} else {
			log.Printf("✅ Log summarized: %s", logEntry.ID)
			timer.ObserveDuration()
			workermetrics.IncSucceeded()
		}
	}
}
