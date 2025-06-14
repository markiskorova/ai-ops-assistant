package summarizer

import (
	"context"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

var useOpenAI = os.Getenv("USE_OPENAI") == "true"

func Summarize(raw string) (string, error) {
	if !useOpenAI {
		return fallbackSummary(raw), nil
	}

	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT4o, // or GPT3Dot5Turbo if needed
			Messages: []openai.ChatCompletionMessage{
				{
					Role:    openai.ChatMessageRoleSystem,
					Content: "You are a helpful assistant summarizing log files. Focus on critical errors and the overall flow.",
				},
				{
					Role:    openai.ChatMessageRoleUser,
					Content: "Please summarize the following log:\n\n" + raw,
				},
			},
			Temperature: 0.4,
		},
	)

	if err != nil {
		return "", err
	}

	return strings.TrimSpace(resp.Choices[0].Message.Content), nil
}

func fallbackSummary(raw string) string {
	lines := strings.Split(raw, "\n")
	var highlights []string

	for _, line := range lines {
		if strings.Contains(line, "ERROR") || strings.Contains(line, "FATAL") {
			highlights = append(highlights, line)
		}
	}

	summary := []string{}
	if len(lines) > 0 {
		summary = append(summary, "Start: "+lines[0])
	}
	if len(highlights) > 0 {
		summary = append(summary, "Errors: "+strings.Join(highlights, " | "))
	}
	if len(lines) > 1 {
		summary = append(summary, "End: "+lines[len(lines)-1])
	}

	return strings.Join(summary, "\n")
}
