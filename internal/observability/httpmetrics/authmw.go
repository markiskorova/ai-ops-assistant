package httpmetrics

import (
	"context"
	"net/http"
	"strings"

	"ai-ops-assistant/internal/auth"
)

// AuthMiddleware adds userID from a validated JWT (if present) into the request context.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token != "" {
			token = strings.TrimPrefix(token, "Bearer ")
			if userID, err := auth.ValidateJWT(token); err == nil {
				ctx := context.WithValue(r.Context(), "userID", userID)
				r = r.WithContext(ctx)
			}
		}
		next.ServeHTTP(w, r)
	})
}
