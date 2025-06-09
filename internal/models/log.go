package models

import "time"

type LogEntry struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Raw       string    `json:"raw"`
	Summary   string    `json:"summary"`
	CreatedAt time.Time `json:"created_at"`
}
