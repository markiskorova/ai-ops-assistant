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
	"ai-ops-assistant/internal/triage"

	"gorm.io/gorm"
)

func main() {
	log.Println("🧠 Starting ticket triage worker...")
	dbConn := db.InitDB()
	classifier := triage.NewClassifierFromEnv()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Metrics server for triage worker
	go workermetrics.StartServer(":9101")

	go runTriageLoop(dbConn, classifier)

	<-stop
	log.Println("🛑 Triage worker stopped.")
}

func runTriageLoop(db *gorm.DB, c triage.Classifier) {
	for {
		// Optional: expose queue depth
		var pending int64
		if err := db.Model(&models.Ticket{}).
			Where("status = ?", "untriaged").
			Count(&pending).Error; err == nil {
			workermetrics.SetQueueDepth(int(pending))
		}

		var ticket models.Ticket
		err := db.
			Where("status = ?", "untriaged").
			Order("created_at ASC").
			First(&ticket).Error

		if err != nil {
			log.Printf("⏳ %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("📌 Triage ticket: %s", ticket.ID)

		// Metrics: one unit of work
		workermetrics.IncStarted()
		timer := workermetrics.NewTimer()

		classification, err := c.Classify(triage.Ticket{
			ID:   ticket.ID.String(),
			Text: ticket.Description,
		})
		if err != nil {
			log.Printf("❌ Classification error: %v", err)
			timer.ObserveDuration()
			workermetrics.IncFailed()
			continue
		}

		ticket.Status = "triaged"
		ticket.Priority = classification.Severity
		ticket.Category = classification.Type

		if err := db.Save(&ticket).Error; err != nil {
			log.Printf("❌ Failed to save triaged ticket: %v", err)
			timer.ObserveDuration()
			workermetrics.IncFailed()
		} else {
			log.Printf("✅ Ticket triaged: %s", ticket.ID)
			timer.ObserveDuration()
			workermetrics.IncSucceeded()
		}
	}
}
