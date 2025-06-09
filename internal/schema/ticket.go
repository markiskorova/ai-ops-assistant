package schema

import (
	"ai-ops-assistant/internal/models"

	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

var TicketType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Ticket",
	Fields: graphql.Fields{
		"id":          &graphql.Field{Type: graphql.ID},
		"title":       &graphql.Field{Type: graphql.String},
		"description": &graphql.Field{Type: graphql.String},
		"category":    &graphql.Field{Type: graphql.String},
		"priority":    &graphql.Field{Type: graphql.String},
		"status":      &graphql.Field{Type: graphql.String},
		"message":     &graphql.Field{Type: graphql.String},
		"createdAt":   &graphql.Field{Type: graphql.String},
	},
})

var TicketQueryField = &graphql.Field{
	Type: TicketType,
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.ID)},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idStr := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		var ticket models.Ticket
		result := DB.First(&ticket, "id = ?", id)
		if result.Error != nil {
			return nil, result.Error
		}
		return ticket, nil
	},
}

var TicketMutationField = &graphql.Field{
	Type: TicketType,
	Args: graphql.FieldConfigArgument{
		"id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		"message": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		idStr := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		ticket := models.Ticket{
			ID:     id,
			Status: "triaged",
		}

		if err := DB.Create(&ticket).Error; err != nil {
			return nil, err
		}

		return ticket, nil
	},
}
