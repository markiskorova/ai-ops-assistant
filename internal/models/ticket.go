package models

type Ticket struct {
	BaseModel
	Title       string `json:"title"`
	Description string `json:"description"`
	Category    string `json:"category"`
	Priority    string `json:"priority"`
	Status      string `json:"status"`
}
