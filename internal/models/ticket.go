package models

import (
	"time"

	"github.com/google/uuid"
)

type Ticket struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Title       string
	Description string
	Category    string
	Priority    string
	Status      string
	CreatedAt   time.Time
}
