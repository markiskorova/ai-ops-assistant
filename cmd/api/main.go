package main

import (
	"ai-ops-assistant/internal/auth"
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/schema"
	"context"
	"log"
	"net/http"
	"strings"

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
		GraphiQL: true, // Enables browser UI at /query
	})

	http.HandleFunc("/graphql", func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "" {
			token = strings.TrimPrefix(token, "Bearer ")
			userID, err := auth.ValidateJWT(token)
			if err == nil {
				ctx := context.WithValue(r.Context(), "userID", userID)
				r = r.WithContext(ctx)
			}
		}
		h.ServeHTTP(w, r)
	})

	log.Println("🚀 GraphQL server running at http://localhost:8080/graphql")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
