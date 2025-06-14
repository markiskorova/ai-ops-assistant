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
	db.InitDB()

	for {
		var job models.ChangelogJob
		err := db.DB.
			Where("processed = false").
			Order("created_at ASC").
			First(&job).Error

		if err != nil {
			if err == gorm.ErrRecordNotFound {
				log.Println("⏳ No pending changelog jobs. Retrying in 5s...")
			} else {
				log.Printf("❌ DB error: %v", err)
			}
			time.Sleep(5 * time.Second)
			continue
		}

		log.Printf("🧩 Generating changelog for job ID %s", job.ID)

		var messages []string
		if err := json.Unmarshal(job.CommitMessages, &messages); err != nil {
			log.Printf("❌ Failed to unmarshal commit messages: %v", err)
			job.Error = err.Error()
			db.DB.Save(&job)
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
			db.DB.Save(&job)
			continue
		}

		data, err := json.Marshal(entries)
		if err != nil {
			log.Printf("❌ Failed to marshal changelog result: %v", err)
			continue
		}

		job.Processed = true
		job.Result = datatypes.JSON(data)

		if err := db.DB.Save(&job).Error; err != nil {
			log.Printf("❌ Failed to save job result: %v", err)
			continue
		}

		log.Printf("✅ Changelog job %s completed", job.ID)
	}
}
