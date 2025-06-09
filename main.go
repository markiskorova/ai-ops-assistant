package main

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	// Initialize database and schema
	database := db.InitDB()

	err := schema.Init(database)
	if err != nil {
		log.Fatalf("❌ Failed to initialize GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true, // Enables browser UI at /query
	})

	http.Handle("/query", h)
	log.Println("🚀 GraphQL server running at http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
