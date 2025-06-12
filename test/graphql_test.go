package test

import (
	"ai-ops-assistant/internal/db"
	"ai-ops-assistant/internal/schema"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/graphql-go/handler"
)

func SetupGraphQLHandler() http.Handler {
	os.Setenv("DB_HOST", "db")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "aiops")
	os.Setenv("DB_PASS", "secret")
	os.Setenv("DB_NAME", "aiops_db")

	dbConn := db.InitDB()
	_ = schema.Init(dbConn)

	return handler.New(&handler.Config{
		Schema:   &schema.Schema,
		Pretty:   true,
		GraphiQL: false,
	})
}

func TestGraphQL_TicketsQuery(t *testing.T) {
	query := `{"query":"{ tickets { id status } }"}`

	req := httptest.NewRequest("POST", "/graphql", bytes.NewBuffer([]byte(query)))
	req.Header.Set("Content-Type", "application/json")

	rr := httptest.NewRecorder()
	handler := SetupGraphQLHandler()
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %v", rr.Code)
	}

	var respBody map[string]interface{}
	err := json.Unmarshal(rr.Body.Bytes(), &respBody)
	if err != nil {
		t.Fatalf("Invalid JSON response: %v", err)
	}

	data, ok := respBody["data"]
	if !ok {
		t.Fatalf("No data returned in GraphQL response: %s", rr.Body.String())
	}

	t.Logf("GraphQL response data: %+v", data)
}
