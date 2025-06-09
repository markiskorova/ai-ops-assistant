package models

import (
    "time"

    "github.com/google/uuid"
    "gorm.io/datatypes"
)

type Changelog struct {
    ID          uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
    Commits     datatypes.JSON
    GeneratedAt time.Time
}