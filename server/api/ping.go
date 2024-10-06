package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
)

type PingUsecase interface {
	Ping(ctx context.Context) (string, error)
}

type PingHandler struct {
	pingUsecase PingUsecase
}

func NewPingHandler(pingUsecase PingUsecase) *PingHandler {
	return &PingHandler{
		pingUsecase: pingUsecase,
	}
}

func (h *PingHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	res, err := h.pingUsecase.Ping(ctx)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	utils.Response(w, http.StatusOK, fmt.Sprintf(`{"success": true, "data": "%s"}`, res))
}
