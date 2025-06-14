package triage

import (
	"strings"
)

type Ticket struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

type TicketClassification struct {
	Severity string `json:"severity"`
	Type     string `json:"type"`
	Owner    string `json:"owner"`
}

func ClassifyTicket(t Ticket) (TicketClassification, error) {
	text := strings.ToLower(t.Text)

	// Default values
	classification := TicketClassification{
		Severity: "Low",
		Type:     "General",
		Owner:    "support@example.com",
	}

	if strings.Contains(text, "error") || strings.Contains(text, "fail") || strings.Contains(text, "exception") {
		classification.Severity = "High"
		classification.Type = "Bug"
		classification.Owner = "ops-team@example.com"
	} else if strings.Contains(text, "deploy") || strings.Contains(text, "infrastructure") || strings.Contains(text, "terraform") {
		classification.Severity = "Medium"
		classification.Type = "Infra"
		classification.Owner = "devops@example.com"
	} else if strings.Contains(text, "add") || strings.Contains(text, "feature") || strings.Contains(text, "support") {
		classification.Severity = "Medium"
		classification.Type = "Feature"
		classification.Owner = "product@example.com"
	}

	return classification, nil
}
