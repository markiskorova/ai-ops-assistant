package test

import (
	"ai-ops-assistant/internal/triage"
	"testing"
)

func TestClassifyTicket(t *testing.T) {
	tests := []struct {
		text     string
		expected string
	}{
		{"System error occurred during login", "Bug"},
		{"Add support for export feature", "Feature"},
		{"Infrastructure update with Terraform", "Infra"},
		{"Some random feedback", "General"},
	}

	for _, tt := range tests {
		tc := triage.Ticket{ID: "test-123", Text: tt.text}
		result, _ := triage.ClassifyTicket(tc)
		if result.Type != tt.expected {
			t.Errorf("Expected type %s, got %s", tt.expected, result.Type)
		}
	}
}
