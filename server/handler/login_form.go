package handler

import (
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
	views "github.com/hansenedrickh/katachi/views/user"
)

type LoginFormHandler struct{}

func NewLoginFormHandler() *LoginFormHandler {
	return &LoginFormHandler{}
}

func (h *LoginFormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := views.LoginForm().Render(ctx, w)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}
}
