package handler

import (
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
	views "github.com/hansenedrickh/katachi/views/user"
)

type RegisterFormHandler struct{}

func NewRegisterFormHandler() *RegisterFormHandler {
	return &RegisterFormHandler{}
}

func (h *RegisterFormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	err := views.RegisterForm().Render(ctx, w)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}
}
