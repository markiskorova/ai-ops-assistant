package graph

import (
    "context"

    "ai-ops-assistant/internal/changelog"
    "ai-ops-assistant/internal/summarizer"
    "ai-ops-assistant/internal/triage"

    "ai-ops-assistant/graph/generated"
    "ai-ops-assistant/graph/model"
)

type Resolver struct{}

func (r *Resolver) Query() generated.QueryResolver {
    return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) HealthCheck(ctx context.Context) (string, error) {
    return "OK", nil
}

func (r *queryResolver) SummarizeLog(ctx context.Context, input model.LogInput) (*model.LogSummary, error) {
    result, err := summarizer.SummarizeLog(summarizer.LogInput{
        Text: input.Text,
    })
    if err != nil {
        return nil, err
    }
    return &model.LogSummary{
        Summary:   result.Summary,
        Category:  result.Category,
        Timestamp: result.Timestamp,
    }, nil
}

func (r *queryResolver) ClassifyTicket(ctx context.Context, input model.TicketInput) (*model.TicketClassification, error) {
    result, err := triage.ClassifyTicket(triage.Ticket{
        ID:   input.ID,
        Text: input.Text,
    })
    if err != nil {
        return nil, err
    }
    return &model.TicketClassification{
        Severity: result.Severity,
        Type:     result.Type,
        Owner:    result.Owner,
    }, nil
}

func (r *queryResolver) ParseChangelog(ctx context.Context, commits []*model.GitCommitInput) ([]*model.ChangelogEntry, error) {
    var commitList []changelog.GitCommit
    for _, c := range commits {
        commitList = append(commitList, changelog.GitCommit{
            Message: c.Message,
            Author:  c.Author,
            Date:    c.Date,
        })
    }

    entries, err := changelog.ParseChangelog(commitList)
    if err != nil {
        return nil, err
    }

    var result []*model.ChangelogEntry
    for _, e := range entries {
        result = append(result, &model.ChangelogEntry{
            Scope:   e.Scope,
            Summary: e.Summary,
        })
    }

    return result, nil
}
