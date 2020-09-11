package middleware

import (
	"context"
	"net/http"

	"github.com/google/uuid"

	"github.com/hansenedrickh/katachi/constant"
)

func Context(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), constant.ContextKeyRequestID, uuid.New().String())

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
