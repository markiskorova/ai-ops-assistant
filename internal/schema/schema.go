package schema

import (
	"github.com/graphql-go/graphql"
	"gorm.io/gorm"
)

var Schema graphql.Schema
var DB *gorm.DB

func Init(db *gorm.DB) error {
	DB = db

	rootQuery := graphql.NewObject(graphql.ObjectConfig{
		Name: "Query",
		Fields: mergeFields(
			graphql.Fields{
				"ticket":     TicketQueryField,
				"tickets":    TicketsQueryField,
				"logEntry":   LogEntryQueryFields["logEntry"],
				"logEntries": LogEntryQueryFields["logEntries"],
				"changelog":  ChangelogByIDField,
				"changelogs": ChangelogListField,
			},
		),
	})

	rootMutation := graphql.NewObject(graphql.ObjectConfig{
		Name: "Mutation",
		Fields: mergeFields(
			graphql.Fields{
				"triageTicket":      TicketMutationField,
				"summarizeLog":      LogEntryMutationFields["summarizeLog"],
				"generateChangelog": GenerateChangelogField,
				"login":             LoginField,
			},
		),
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	if err != nil {
		return err
	}
	Schema = schema
	return nil
}

// mergeFields combines multiple graphql.Fields maps
func mergeFields(fieldSets ...graphql.Fields) graphql.Fields {
	merged := graphql.Fields{}
	for _, fs := range fieldSets {
		for k, v := range fs {
			merged[k] = v
		}
	}
	return merged
}
