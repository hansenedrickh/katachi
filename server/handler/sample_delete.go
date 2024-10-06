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

type SampleDeleteUsecase interface {
	Delete(ctx context.Context, id int) error
}

type SampleDeleteHandler struct {
	sampleUsecase SampleDeleteUsecase
}

func NewSampleDeleteHandler(sampleUsecase SampleDeleteUsecase) *SampleDeleteHandler {
	return &SampleDeleteHandler{
		sampleUsecase: sampleUsecase,
	}
}

func (h *SampleDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	err = h.sampleUsecase.Delete(ctx, id)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	w.Header().Set("HX-Location", "/samples?page=1&limit=10")
	w.WriteHeader(http.StatusNoContent)
}
