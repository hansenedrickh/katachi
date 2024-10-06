package handler

import (
	"context"
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
)

type MockSampleInsertUsecase struct {
	mock.Mock
}

func (m *MockSampleInsertUsecase) Insert(ctx context.Context, name string) error {
	args := m.Called(ctx, name)
	return args.Error(0)
}

func TestSampleInsertHandler_ServeHTTP_Success(t *testing.T) {
	mockUsecase := new(MockSampleInsertUsecase)
	h := NewSampleInsertHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples", nil)
	req.Form = url.Values{}
	req.Form.Add("name", "sample name")
	w := httptest.NewRecorder()

	mockUsecase.On("Insert", mock.Anything, "sample name").Return(nil)

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
	require.Equal(t, "/samples?page=1&limit=10", w.Header().Get("HX-Location"))
	mockUsecase.AssertExpectations(t)
}

func TestSampleInsertHandler_ServeHTTP_MissingNameParameter(t *testing.T) {
	mockUsecase := new(MockSampleInsertUsecase)
	h := NewSampleInsertHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
	require.JSONEq(t, `{"success": false, "error": "missing name parameter"}`, w.Body.String())
}

func TestSampleInsertHandler_ServeHTTP_UsecaseError(t *testing.T) {
	mockUsecase := new(MockSampleInsertUsecase)
	h := NewSampleInsertHandler(mockUsecase)
	req := httptest.NewRequest("POST", "/samples", nil)
	req.Form = url.Values{}
	req.Form.Add("name", "sample name")
	w := httptest.NewRecorder()

	mockUsecase.On("Insert", mock.Anything, "sample name").Return(errors.New("usecase error"))

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	require.JSONEq(t, `{"success": false, "error": "usecase error"}`, w.Body.String())
}
