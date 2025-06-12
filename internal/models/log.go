package models

type LogEntry struct {
	ID        string `gorm:"primaryKey" json:"id"`
	Raw       string `json:"raw"`
	Summary   string `json:"summary"`
	CreatedAt string `json:"createdAt"`
}
