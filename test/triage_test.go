package test

import (
	"testing"

	"ai-ops-assistant/internal/triage"
	"github.com/stretchr/testify/assert"
)

type MockClassifier struct{}

func (m *MockClassifier) Classify(t triage.Ticket) (triage.TicketClassification, error) {
	return triage.TicketClassification{
		Severity: "High",
		Type:     "Bug",
		Owner:    "mock@example.com",
	}, nil
}

func TestMockClassifier(t *testing.T) {
	var c triage.Classifier = &MockClassifier{}

	ticket := triage.Ticket{
		ID:   "abc-123",
		Text: "Something failed in the deploy script.",
	}

	classification, err := c.Classify(ticket)
	assert.NoError(t, err)
	assert.Equal(t, "High", classification.Severity)
	assert.Equal(t, "Bug", classification.Type)
	assert.Equal(t, "mock@example.com", classification.Owner)
}
