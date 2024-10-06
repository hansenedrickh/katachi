package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hansenedrickh/katachi/server/utils"
)

type SampleInsertUsecase interface {
	Insert(ctx context.Context, name string) error
}

type SampleInsertHandler struct {
	sampleUsecase SampleInsertUsecase
}

func NewSampleInsertHandler(sampleUsecase SampleInsertUsecase) *SampleInsertHandler {
	return &SampleInsertHandler{
		sampleUsecase: sampleUsecase,
	}
}

func (h *SampleInsertHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	name := r.FormValue("name")
	if name == "" {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "missing name parameter"}`)
		return
	}

	err := h.sampleUsecase.Insert(ctx, name)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	w.Header().Set("HX-Location", "/samples?page=1&limit=10")
	w.WriteHeader(http.StatusNoContent)
}
