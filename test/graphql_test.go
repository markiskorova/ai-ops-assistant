package test

import (
	"testing"

	"github.com/graphql-go/graphql"
	"github.com/stretchr/testify/assert"
)

func TestSimpleSchema(t *testing.T) {
	testType := graphql.NewObject(graphql.ObjectConfig{
		Name: "Test",
		Fields: graphql.Fields{
			"message": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "Hello", nil
				},
			},
		},
	})

	schema, err := graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"test": &graphql.Field{Type: testType},
			},
		}),
	})
	assert.NoError(t, err)

	query := `{ test { message } }`
	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)

	assert.Empty(t, result.Errors)
	assert.Equal(t, "Hello", result.Data.(map[string]interface{})["test"].(map[string]interface{})["message"])
}
