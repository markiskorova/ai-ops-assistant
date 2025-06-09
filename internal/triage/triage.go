package triage

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
	// Simulate ticket classification with mock logic
	return TicketClassification{
		Severity: "High",
		Type:     "Outage",
		Owner:    "ops-team@example.com",
	}, nil
}
