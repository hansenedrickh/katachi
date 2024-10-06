package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"
)

const ContextKeyRequestID = "request_id"

func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), ContextKeyRequestID, uuid.New().String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
