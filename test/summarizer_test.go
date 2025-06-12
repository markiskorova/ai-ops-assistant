package test

import (
	"testing"

	"ai-ops-assistant/internal/summarizer"
	"github.com/stretchr/testify/assert"
)

type MockSummarizer struct{}

func (m *MockSummarizer) Summarize(raw string) (string, error) {
	return "Mock summary result", nil
}

func TestMockSummarizer(t *testing.T) {
	var s summarizer.Summarizer = &MockSummarizer{}

	result, err := s.Summarize("This is a raw log line with ERROR and FATAL")
	assert.NoError(t, err)
	assert.Equal(t, "Mock summary result", result)
}
