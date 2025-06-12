package test

import (
	"ai-ops-assistant/internal/summarizer"
	"testing"
)

func TestSummarizer(t *testing.T) {
	rawLog := "Error: Disk full on node 3"
	result, err := summarizer.Summarize(rawLog)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if result == "" {
		t.Fatal("Summarizer returned empty result")
	}
}
