package handler

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockSampleUpdateUsecase struct {
	mock.Mock
}

func (m *MockSampleUpdateUsecase) Update(ctx context.Context, id int, name string) error {
	args := m.Called(ctx, id, name)
	return args.Error(0)
}

func TestSampleUpdateHandler_ServeHTTP_Success(t *testing.T) {
	mockUsecase := new(MockSampleUpdateUsecase)
	h := NewSampleUpdateHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples/1", nil)
	req.Form = url.Values{}
	req.Form.Add("name", "sample name")
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	mockUsecase.On("Update", mock.Anything, 1, "sample name").Return(nil)

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
	require.Equal(t, "/samples?page=1&limit=10", w.Header().Get("HX-Location"))
	mockUsecase.AssertExpectations(t)
}

func TestSampleUpdateHandler_ServeHTTP_InvalidIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleUpdateUsecase)
	h := NewSampleUpdateHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples/invalid", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.JSONEq(t, `{"success": false, "error": "Invalid id parameter"}`, w.Body.String())
}

func TestSampleUpdateHandler_ServeHTTP_ZeroIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleUpdateUsecase)
	h := NewSampleUpdateHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples/0", nil)
	vars := map[string]string{
		"id": "0",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestSampleUpdateHandler_ServeHTTP_MissingNameParameter(t *testing.T) {
	mockUsecase := new(MockSampleUpdateUsecase)
	h := NewSampleUpdateHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples/1", nil)
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.JSONEq(t, `{"success": false, "error": "missing name parameter"}`, w.Body.String())
}

func TestSampleUpdateHandler_ServeHTTP_UsecaseError(t *testing.T) {
	mockUsecase := new(MockSampleUpdateUsecase)
	h := NewSampleUpdateHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples/1", nil)
	req.Form = url.Values{}
	req.Form.Add("name", "sample name")
	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	mockUsecase.On("Update", mock.Anything, 1, "sample name").Return(errors.New("usecase error"))

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.JSONEq(t, `{"success": false, "error": "usecase error"}`, w.Body.String())
}
