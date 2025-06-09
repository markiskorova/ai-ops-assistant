package summarizer

// Summarize returns a standard mock summary response
func Summarize(raw string) (string, error) {
	return "Mock summary: This is a placeholder summary for the given log.", nil
}