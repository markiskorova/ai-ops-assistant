package models

import (
    "time"

    "github.com/google/uuid"
)

type User struct {
    ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Email     string    `gorm:"uniqueIndex"`
    Password  string
    CreatedAt time.Time
}