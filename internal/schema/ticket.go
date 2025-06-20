package schema

import (
	"ai-ops-assistant/internal/models"
	"ai-ops-assistant/internal/triage"
	"errors"

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
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		idStr := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}
		var ticket models.Ticket
		if err := GetDB(p.Context).First(&ticket, "id = ?", id).Error; err != nil {
			return nil, err
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
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		idStr := p.Args["id"].(string)
		id, err := uuid.Parse(idStr)
		if err != nil {
			return nil, err
		}

		var ticket models.Ticket
		if err := GetDB(p.Context).First(&ticket, "id = ?", id).Error; err != nil {
			return nil, err
		}

		classifier := triage.NewClassifierFromEnv()
		classification, err := classifier.Classify(triage.Ticket{ID: ticket.ID.String(), Text: ticket.Description})
		if err != nil {
			return nil, err
		}

		ticket.Status = "triaged"
		ticket.Priority = classification.Severity
		ticket.Category = classification.Type

		if err := GetDB(p.Context).Save(&ticket).Error; err != nil {
			return nil, err
		}

		return ticket, nil
	},
}

var TicketsQueryField = &graphql.Field{
	Type: graphql.NewList(TicketType),
	Args: graphql.FieldConfigArgument{
		"status": &graphql.ArgumentConfig{Type: graphql.String},
	},
	Resolve: func(p graphql.ResolveParams) (interface{}, error) {
		_, ok := p.Context.Value("userID").(string)
		if !ok {
			return nil, errors.New("unauthorized")
		}

		var tickets []models.Ticket
		query := GetDB(p.Context).Model(&models.Ticket{})

		if status, ok := p.Args["status"].(string); ok && status != "" {
			query = query.Where("status = ?", status)
		}

		if err := query.Order("created_at DESC").Find(&tickets).Error; err != nil {
			return nil, err
		}

		return tickets, nil
	},
}
