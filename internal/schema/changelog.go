package schema

import (
	"ai-ops-assistant/internal/changelog"
	"ai-ops-assistant/internal/models"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var ChangelogType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Changelog",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.String},
		"commits":     &graphql.Field{Type: graphql.String},
		"generatedAt": &graphql.Field{Type: graphql.String},
	},
})

var GenerateChangelogField = &graphql.Field{
	Type: ChangelogType,
	Args: graphql.FieldConfigArgument{
		"commits": &graphql.ArgumentConfig{Type: graphql.NewList(graphql.NewNonNull(graphql.String))},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		rawCommits := p.Args["commits"].([]interface{})
		var commits []changelog.GitCommit
		for _, c := range rawCommits {
			commits = append(commits, changelog.GitCommit{Message: c.(string)})
		}
		entries, err := changelog.ParseChangelog(commits)
		if err != nil {
			return nil, err
		}

		grouped := map[string][]string{
			"features": {},
			"fixes":    {},
			"other":    {},
		}

		for _, entry := range entries {
			line := entry.Summary
			switch entry.Scope {
			case "feature":
				grouped["features"] = append(grouped["features"], line)
			case "bugfix":
				grouped["fixes"] = append(grouped["fixes"], line)
			default:
				grouped["other"] = append(grouped["other"], line)
			}
		}

		jsonData, _ := json.Marshal(grouped)

		entry := models.Changelog{
    Commits:     jsonData,
    GeneratedAt: time.Now(),
}
entry.ID = uuid.New()

		if err := GetDB(p.Context).Create(&entry).Error; err != nil {
			return nil, err
		}

		return entry, nil
	},
}

var ChangelogByIDField = &graphql.Field{
	Type: ChangelogType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		idStr := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		var entry models.Changelog
		if err := GetDB(p.Context).First(&entry, "id = ?", id).Error; err != nil {
			return nil, err
		}
		return entry, nil
	},
}

var ChangelogListField = &graphql.Field{
	Type: graphql.NewList(ChangelogType),
	Args: graphql.FieldConfigArgument{
		"limit": &graphql.ArgumentConfig{Type: graphql.Int},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		limit, ok := p.Args["limit"].(int)
		if !ok || limit <= 0 {
			limit = 10
		}

		var entries []models.Changelog
		if err := GetDB(p.Context).Order("generated_at DESC").Limit(limit).Find(&entries).Error; err != nil {
			return nil, err
		}
		return entries, nil
	},
}

func contains(text, keyword string) bool {
	return strings.Contains(strings.ToLower(text), keyword)
}
