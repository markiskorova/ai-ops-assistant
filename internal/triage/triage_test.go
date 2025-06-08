package triage

import (
    "testing"
)

func TestClassifyTicket(t *testing.T) {
    ticket := Ticket{ID: "1", Text: "Service is down"}
    result, err := ClassifyTicket(ticket)
    if err != nil {
        t.Errorf("unexpected error: %v", err)
    }

    if result.Severity == "" {
        t.Errorf("expected severity to be classified")
    }
}
