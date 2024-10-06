package handler

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockSampleDeleteUsecase struct {
	mock.Mock
}

func (m *MockSampleDeleteUsecase) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

func TestSampleDeleteHandler_ServeHTTP_Success(t *testing.T) {
	mockUsecase := new(MockSampleDeleteUsecase)
	h := NewSampleDeleteHandler(mockUsecase)
	req := httptest.NewRequest("DELETE", "/samples/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	mockUsecase.On("Delete", mock.Anything, 1).Return(nil)

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusNoContent, w.Code)
	require.Equal(t, "/samples?page=1&limit=10", w.Header().Get("HX-Location"))
	mockUsecase.AssertExpectations(t)
}

func TestSampleDeleteHandler_ServeHTTP_InvalidIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleDeleteUsecase)
	h := NewSampleDeleteHandler(mockUsecase)
	req := httptest.NewRequest("DELETE", "/sample/invalid", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSampleDeleteHandler_ServeHTTP_ZeroIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleDeleteUsecase)
	h := NewSampleDeleteHandler(mockUsecase)
	req := httptest.NewRequest("DELETE", "/sample/0", nil)
	vars := map[string]string{
		"id": "0",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestSampleDeleteHandler_ServeHTTP_UsecaseError(t *testing.T) {
	mockUsecase := new(MockSampleDeleteUsecase)
	h := NewSampleDeleteHandler(mockUsecase)
	req := httptest.NewRequest("DELETE", "/sample/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	mockUsecase.On("Delete", mock.Anything, 1).Return(errors.New("usecase error"))

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockUsecase.AssertExpectations(t)
}
