package main

import (
	"encoding/json"
	"log"
	"time"

	"ai-ops-assistant/internal/changelog"
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

func main() {
	log.Println("📜 Starting changelog generation worker...")
	dbConn := db.InitDB()

	runChangelogLoop(dbConn)
}

func runChangelogLoop(db *gorm.DB) {
	for {
		var job models.ChangelogJob
		err := db.
			Where("processed = false").
			Order("created_at ASC").
			First(&job).Error

		if err != nil {
			log.Printf("⏳ %v", err)
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("🧩 Generating changelog for job ID %s", job.ID)

		var messages []string
		if err := json.Unmarshal(job.CommitMessages, &messages); err != nil {
			log.Printf("❌ Failed to unmarshal: %v", err)
			job.Error = err.Error()
			db.Save(&job)
			continue
		}

		var commits []changelog.GitCommit
		for _, msg := range messages {
			commits = append(commits, changelog.GitCommit{Message: msg})
		}

		entries, err := changelog.ParseChangelog(commits)
		if err != nil {
			log.Printf("❌ Parse error: %v", err)
			job.Error = err.Error()
			db.Save(&job)
			continue
		}

		data, err := json.Marshal(entries)
		if err != nil {
			log.Printf("❌ Failed to marshal result: %v", err)
			continue
		}

		job.Processed = true
		job.Result = datatypes.JSON(data)

		if err := db.Save(&job).Error; err != nil {
			log.Printf("❌ Failed to save result: %v", err)
		} else {
			log.Printf("✅ Changelog job %s completed", job.ID)
		}
	}
}
