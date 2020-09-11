package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/hansenedrickh/katachi/constant"
)

func TestContext(t *testing.T) {
	w := httptest.NewRecorder()
	r, _ := http.NewRequest("GET", "/", nil)

	handler := http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.NotEmpty(t, r.Context().Value(constant.ContextKeyRequestID))
	}))

	contextHandler := Context(handler)
	contextHandler.ServeHTTP(w, r)
}
