package models

import (
	"time"

	"gorm.io/datatypes"
)

type Changelog struct {
	BaseModel
	Commits     datatypes.JSON `gorm:"type:jsonb" json:"commits"` // []GitCommit
	GeneratedAt time.Time      `json:"generatedAt"`
}

type ChangelogEntry struct {
	Scope   string `json:"scope"`
	Summary string `json:"summary"`
}

type ChangelogJob struct {
	BaseModel
	CommitMessages datatypes.JSON `gorm:"type:jsonb" json:"commitMessages"` // []string
	Result         datatypes.JSON `gorm:"type:jsonb" json:"result"`         // []ChangelogEntry
	Processed      bool           `json:"processed"`
	Error          string         `json:"error"`
}
