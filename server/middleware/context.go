package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

type contextKey string

func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		key := contextKey("request_id")

		ctx := context.WithValue(r.Context(), key, uuid.New().String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
