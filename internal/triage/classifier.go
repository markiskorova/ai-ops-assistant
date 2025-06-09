package triage

import "ai-ops-assistant/internal/models"

// Classify simulates categorizing a ticket with mock logic
func Classify(t *models.Ticket) {
    t.Status = "triaged"
    t.Priority = "medium"
    t.Category = "bug"
}