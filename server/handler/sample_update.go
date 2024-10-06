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

type SampleUpdateUsecase interface {
	Update(ctx context.Context, id int, name string) error
}

type SampleUpdateHandler struct {
	sampleUsecase SampleUpdateUsecase
}

func NewSampleUpdateHandler(sampleUsecase SampleUpdateUsecase) *SampleUpdateHandler {
	return &SampleUpdateHandler{
		sampleUsecase: sampleUsecase,
	}
}

func (h *SampleUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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

	name := r.FormValue("name")
	if name == "" {
		utils.Response(w, http.StatusBadRequest, `{"success": false, "error": "missing name parameter"}`)
		return
	}

	err = h.sampleUsecase.Update(ctx, id, name)
	if err != nil {
		utils.Response(w, http.StatusInternalServerError, fmt.Sprintf(`{"success": false, "error": "%s"}`, err.Error()))
		return
	}

	w.Header().Set("HX-Location", "/samples?page=1&limit=10")
	w.WriteHeader(http.StatusNoContent)
}
