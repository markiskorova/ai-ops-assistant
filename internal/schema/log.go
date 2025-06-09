package schema

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/summarizer"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var logEntryType = graphql.NewObject(graphql.ObjectConfig{
	Name: "LogEntry",
	Fields: graphql.Fields{
		"id":         &graphql.Field{Type: graphql.String},
		"raw":        &graphql.Field{Type: graphql.String},
		"summary":    &graphql.Field{Type: graphql.String},
		"created_at": &graphql.Field{Type: graphql.String},
	},
})

func init() {
	// Mutation: summarizeLog
	rootMutation.AddFieldConfig("summarizeLog", &graphql.Field{
		Type: logEntryType,
		Args: graphql.FieldConfigArgument{
			"raw": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			raw := p.Args["raw"].(string)
			summary, err := summarizer.Summarize(raw)
			if err != nil {
				return nil, err
			}

			entry := models.LogEntry{
				ID:      uuid.New().String(),
				Raw:     raw,
				Summary: summary,
			}

			if err := db.DB.Create(&entry).Error; err != nil {
				return nil, err
			}

			return entry, nil
		},
	})

	// Query: logEntry(id)
	rootQuery.AddFieldConfig("logEntry", &graphql.Field{
		Type: logEntryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			id := p.Args["id"].(string)
			var entry models.LogEntry
			if err := db.DB.First(&entry, "id = ?", id).Error; err != nil {
				return nil, err
			}
			return entry, nil
		},
	})

	// Query: logEntries(limit)
	rootQuery.AddFieldConfig("logEntries", &graphql.Field{
		Type: graphql.NewList(logEntryType),
		Args: graphql.FieldConfigArgument{
			"limit": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			limit, ok := p.Args["limit"].(int)
			if !ok || limit <= 0 {
				limit = 10
			}

			var entries []models.LogEntry
			if err := db.DB.Order("id DESC").Limit(limit).Find(&entries).Error; err != nil {
				return nil, err
			}
			return entries, nil
		},
	})
}
