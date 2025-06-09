package main

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	database := db.InitDB()
	err := schema.Init(database)
	if err != nil {
		log.Fatalf("❌ Failed to initialize GraphQL schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true, // enable for testing
	})

	http.Handle("/graphql", h)

	log.Println("🚀 API server running on http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
