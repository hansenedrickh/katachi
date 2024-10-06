package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hansenedrickh/katachi/pkg/sample"
	"github.com/hansenedrickh/katachi/server/utils"
	views "github.com/hansenedrickh/katachi/views/sample"
)

type SampleIndexUsecase interface {
	Get(ctx context.Context, page, limit int) ([]sample.Sample, error)
}

type SampleIndexHandler struct {
	sampleUsecase SampleIndexUsecase
}

func NewSampleIndexHandler(sampleUsecase SampleIndexUsecase) *SampleIndexHandler {
	return &SampleIndexHandler{
		sampleUsecase: sampleUsecase,
	}
}

func (h *SampleIndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	pageStr := r.URL.Query().Get("page")
	limitStr := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "Invalid page parameter"}`)
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "Invalid limit parameter"}`)
		return
	}

	samples, err := h.sampleUsecase.Get(ctx, page, limit)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	err = views.Index(samples).Render(ctx, w)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}
}
