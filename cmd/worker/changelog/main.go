package main

import (
	"encoding/json"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ai-ops-assistant/internal/changelog"
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/observability/workermetrics"

	"gorm.io/gorm"
)

func main() {
	log.Println("🧠 Starting changelog worker...")
	dbConn := db.InitDB()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Metrics server for changelog worker
	go workermetrics.StartServer(":9103")

	go runChangelogLoop(dbConn)

	<-stop
	log.Println("🛑 Changelog worker stopped.")
}

func runChangelogLoop(db *gorm.DB) {
	for {
		// Optional: expose queue depth (unprocessed jobs)
		var pending int64
		if err := db.Model(&models.ChangelogJob{}).
			Where("processed = ?", false).
			Count(&pending).Error; err == nil {
			workermetrics.SetQueueDepth(int(pending))
		}

		var job models.ChangelogJob
		err := db.
			Where("processed = ?", false).
			Order("created_at ASC").
			First(&job).Error

		if err != nil {
			log.Printf("⏳ %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("📖 Processing changelog job: %s", job.ID)

		// Metrics: one unit of work
		workermetrics.IncStarted()
		timer := workermetrics.NewTimer()

		// Step 1: decode commit messages (stored as JSON []string)
		var commitMsgs []string
		if err := json.Unmarshal(job.CommitMessages, &commitMsgs); err != nil {
			log.Printf("❌ Failed to parse commit messages: %v", err)
			job.Error = err.Error()
			_ = db.Save(&job).Error
			timer.ObserveDuration()
			workermetrics.IncFailed()
			continue
		}

		// Step 2: map to []GitCommit
		var commits []changelog.GitCommit
		for _, msg := range commitMsgs {
			commits = append(commits, changelog.GitCommit{
				Message: msg,
				Author:  "unknown", // extend schema later if needed
				Date:    time.Now().Format(time.RFC3339),
			})
		}

		// Step 3: parse changelog entries
		entries, err := changelog.ParseChangelog(commits)
		if err != nil {
			log.Printf("❌ ParseChangelog error: %v", err)
			job.Error = err.Error()
			_ = db.Save(&job).Error
			timer.ObserveDuration()
			workermetrics.IncFailed()
			continue
		}

		// Step 4: save result into job.Result (JSON) and mark processed
		resultJSON, _ := json.Marshal(entries)
		job.Result = resultJSON
		job.Processed = true
		job.Error = ""

		if err := db.Save(&job).Error; err != nil {
			log.Printf("❌ Failed to save job result: %v", err)
			timer.ObserveDuration()
			workermetrics.IncFailed()
		} else {
			log.Printf("✅ Changelog job processed: %s", job.ID)
			timer.ObserveDuration()
			workermetrics.IncSucceeded()
		}
	}
}
