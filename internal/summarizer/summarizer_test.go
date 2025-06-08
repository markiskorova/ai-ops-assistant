package summarizer

import (
    "testing"
)

func TestSummarizeLog(t *testing.T) {
    input := LogInput{Text: "Example log entry"}
    summary, err := SummarizeLog(input)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if summary.Summary == "" {
        t.Errorf("expected non-empty summary")
    }
}
