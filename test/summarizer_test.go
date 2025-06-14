package test

import (
	"os"
	"strings"
	"testing"

	"ai-ops-assistant/internal/summarizer"
)

func TestSummarizerBasic(t *testing.T) {
	// Force fallback mode (safe default for local/dev testing)
	_ = os.Setenv("USE_OPENAI", "false")

	log := `
2025-06-14 12:01 INFO Service starting
2025-06-14 12:02 ERROR Could not connect
2025-06-14 12:03 INFO Shutting down
`

	summary, err := summarizer.Summarize(log)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if summary == "" {
		t.Fatal("Expected non-empty summary")
	}

	if !strings.Contains(summary, "Start:") || !strings.Contains(summary, "End:") {
		t.Error("Expected summary to contain 'Start:' and 'End:' sections")
	}
}
