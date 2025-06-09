package main

import (
    "log"
    "os"
    "os/signal"
    "syscall"
    "time"

    "ai-ops-assistant/internal/db"
    "ai-ops-assistant/internal/models"
    "ai-ops-assistant/internal/triage"

    "gorm.io/gorm"
)

func main() {
    log.Println("🧠 Starting ticket triage worker...")
    db.InitDB()

    stop := make(chan os.Signal, 1)
    signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

    go func() {
        for {
            var ticket models.Ticket
            err := db.DB.
                Where("status = ?", "untriaged").
                Order("created_at ASC").
                First(&ticket).Error

            if err != nil {
                if err == gorm.ErrRecordNotFound {
                    log.Println("⏳ No untriaged tickets. Retrying in 5s...")
                } else {
                    log.Printf("❌ DB error: %v", err)
                }
                time.Sleep(5 * time.Second)
                continue
            }

            log.Printf("📌 Triage ticket: %s\n", ticket.ID)
            triage.Classify(&ticket)

            if err := db.DB.Save(&ticket).Error; err != nil {
                log.Printf("❌ Failed to save triaged ticket: %v", err)
            } else {
                log.Printf("✅ Ticket triaged: %s\n", ticket.ID)
            }
        }
    }()

    <-stop
    log.Println("🛑 Triage worker stopped.")
}