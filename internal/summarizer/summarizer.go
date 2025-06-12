package summarizer

import (
	"context"
	"errors"
	"os"
	"strings"

	openai "github.com/sashabaranov/go-openai"
)

type Summarizer interface {
	Summarize(raw string) (string, error)
}

type OpenAISummarizer struct {
	client *openai.Client
	model  string
}

type FallbackSummarizer struct{}

func NewSummarizerFromEnv() Summarizer {
	if os.Getenv("USE_OPENAI") == "true" && os.Getenv("OPENAI_API_KEY") != "" {
		return &OpenAISummarizer{
			client: openai.NewClient(os.Getenv("OPENAI_API_KEY")),
			model:  openai.GPT4o, // or openai.GPT3Dot5Turbo
		}
	}
	return &FallbackSummarizer{}
}

func (s *OpenAISummarizer) Summarize(raw string) (string, error) {
	if s.client == nil {
		return "", errors.New("OpenAI client not initialized")
	}

	resp, err := s.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: s.model,
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

func (f *FallbackSummarizer) Summarize(raw string) (string, error) {
	lines := strings.Split(raw, "\n")
	var highlights []string

	for _, line := range lines {
		if strings.Contains(line, "ERROR") || strings.Contains(line, "FATAL") {
			highlights = append(highlights, line)
		}
	}

	var summary []string
	if len(lines) > 0 {
		summary = append(summary, "Start: "+lines[0])
	}
	if len(highlights) > 0 {
		summary = append(summary, "Errors: "+strings.Join(highlights, " | "))
	}
	if len(lines) > 1 {
		summary = append(summary, "End: "+lines[len(lines)-1])
	}

	return strings.Join(summary, "\n"), nil
}
