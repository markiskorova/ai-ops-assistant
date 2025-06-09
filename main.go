package main

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/schema"
	"log"
	"net/http"

	"github.com/graphql-go/handler"
)

func main() {
	db.Init()

	h := handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: true, // Enables browser UI at /query
	})

	http.Handle("/query", h)
	log.Println("🚀 GraphQL server running at http://localhost:8080/query")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
