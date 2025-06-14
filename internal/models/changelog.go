package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type Changelog struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Commits     datatypes.JSON `gorm:"type:jsonb" json:"commits"`
	GeneratedAt time.Time      `json:"generatedAt"`
}

type ChangelogEntry struct {
	Scope   string `json:"scope"`
	Summary string `json:"summary"`
}

type ChangelogJob struct {
	ID             uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	CommitMessages datatypes.JSON `gorm:"type:jsonb" json:"commitMessages"` // stored as []string
	Result         datatypes.JSON `gorm:"type:jsonb" json:"result"`         // stored as []ChangelogEntry
	Processed      bool           `json:"processed"`
	Error          string         `json:"error"`
	CreatedAt      time.Time      `json:"createdAt"`
}
