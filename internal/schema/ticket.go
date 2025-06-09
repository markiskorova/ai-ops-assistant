package schema

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/models"

	"github.com/graphql-go/graphql"
)

var ticketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ticket",
	Fields: graphql.Fields{
		"id":      &graphql.Field{Type: graphql.String},
		"status":  &graphql.Field{Type: graphql.String},
		"message": &graphql.Field{Type: graphql.String},
	},
})

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "Query",
	Fields: graphql.Fields{
		"ticket": &graphql.Field{
			Type: ticketType,
			Args: graphql.FieldConfigArgument{
				"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				id := p.Args["id"].(string)
				var ticket models.Ticket
				result := db.DB.First(&ticket, "id = ?", id)
				if result.Error != nil {
					return nil, result.Error
				}
				return ticket, nil
			},
		},
	},
})

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"triageTicket": &graphql.Field{
			Type: ticketType,
			Args: graphql.FieldConfigArgument{
				"id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"message": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				ticket := models.Ticket{
					ID:      p.Args["id"].(string),
					Status:  "triaged",
					Message: p.Args["message"].(string),
				}
				if err := db.DB.Create(&ticket).Error; err != nil {
					return nil, err
				}
				return ticket, nil
			},
		},
	},
})
