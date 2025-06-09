package models

type Ticket struct {
	ID      string `gorm:"primaryKey" json:"id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}