package handler

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/hansenedrickh/katachi/pkg/sample"
	"github.com/hansenedrickh/katachi/server/utils"
	views "github.com/hansenedrickh/katachi/views/sample"

	"github.com/gorilla/mux"
)

type SampleFormUsecase interface {
	Fetch(ctx context.Context, id int) (sample.Sample, error)
}

type SampleFormHandler struct {
	sampleUsecase SampleFormUsecase
}

func NewSampleFormHandler(sampleUsecase SampleFormUsecase) *SampleFormHandler {
	return &SampleFormHandler{
		sampleUsecase: sampleUsecase,
	}
}

func (h *SampleFormHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "Invalid id parameter"}`)
		return
	}

	if id == 0 {
		err = views.Form(sample.Sample{}).Render(ctx, w)
		if err != nil {
			utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
			return
		}
		return
	}

	res, err := h.sampleUsecase.Fetch(ctx, id)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	err = views.Form(res).Render(ctx, w)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}
}
