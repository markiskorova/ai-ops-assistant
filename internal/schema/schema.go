package schema

import "github.com/graphql-go/graphql"

var Schema graphql.Schema

func init() {
	Schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
}