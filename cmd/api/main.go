package main

import (
	"log"
	"net/http"

	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/observability/httpmetrics"
	"ai-ops-assistant/internal/schema"

	"github.com/graphql-go/handler"
)

func main() {
	database := db.InitDB()
	if err := schema.Init(database); err != nil {
		log.Fatalf("❌ Failed to initialize GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true, // Enables browser UI at /graphql
	})

	mux := http.NewServeMux()

	// Expose metrics
	mux.Handle("/metrics", httpmetrics.Handler())

	// GraphQL with auth + metrics
	graphql := httpmetrics.AuthMiddleware(h)
	mux.Handle("/graphql", httpmetrics.Instrument("/graphql", graphql))

	log.Println("🚀 GraphQL server running at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
