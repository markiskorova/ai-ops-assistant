package summarizer

type LogInput struct {
    Text string `json:"text"`
}

type LogSummary struct {
    Summary   string `json:"summary"`
    Category  string `json:"category"`
    Timestamp string `json:"timestamp"`
}

func SummarizeLog(input LogInput) (LogSummary, error) {
    // Simulate summarization with mock logic
    return LogSummary{
        Summary:   "Detected timeout error in Service X",
        Category:  "TimeoutError",
        Timestamp: "2025-06-06T17:00:00Z",
    }, nil
}
