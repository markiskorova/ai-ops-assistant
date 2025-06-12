package test

import (
	"testing"

	"ai-ops-assistant/internal/triage"
)

func TestClassifyTicket(t *testing.T) {
	input := triage.Ticket{
		ID:   "123",
		Text: "User reports a crash when opening the dashboard.",
	}

	result, err := triage.ClassifyTicket(input)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if result.Type == "" || result.Owner == "" || result.Severity == "" {
		t.Fatal("Classification result is incomplete")
	}
}
