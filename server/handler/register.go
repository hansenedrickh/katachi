package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
)

type RegisterUsecase interface {
	Register(ctx context.Context, username, password string) error
}

type RegisterHandler struct {
	registerUsecase RegisterUsecase
}

func NewRegisterHandler(registerUsecase RegisterUsecase) *RegisterHandler {
	return &RegisterHandler{
		registerUsecase: registerUsecase,
	}
}

func (h *RegisterHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	username := r.FormValue("username")
	if username == "" {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "missing username parameter"}`)
		return
	}

	password := r.FormValue("password")
	if password == "" {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "missing password parameter"}`)
		return
	}

	err := h.registerUsecase.Register(ctx, username, password)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
