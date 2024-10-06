package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
)

type LoginUsecase interface {
	Login(ctx context.Context, username, password string) (string, error)
}

type LoginHandler struct {
	loginUsecase LoginUsecase
}

func NewLoginHandler(loginUsecase LoginUsecase) *LoginHandler {
	return &LoginHandler{
		loginUsecase: loginUsecase,
	}
}

func (h *LoginHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	token, err := h.loginUsecase.Login(ctx, username, password)
	if err != nil && err.Error() == "wrong password" {
		utils.Response(w, http.StatusUnauthorized, `{"success": false, "error": "unauthorized"}`)
		return
	} else if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		Domain:   "localhost",
		MaxAge:   3600,
		Secure:   false,
		HttpOnly: true,
	})
	http.Redirect(w, r, "/samples?page=1&limit=10", http.StatusSeeOther)
}
