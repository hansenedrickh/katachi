package handler

import (
	"context"
	"errors"
	"github.com/gorilla/mux"
	"github.com/hansenedrickh/katachi/pkg/sample"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockSampleFormUsecase struct {
	mock.Mock
}

func (m *MockSampleFormUsecase) Fetch(ctx context.Context, id int) (sample.Sample, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(sample.Sample), args.Error(1)
}

func TestSampleFormHandler_ServeHTTP_Success(t *testing.T) {
	mockUsecase := new(MockSampleFormUsecase)
	h := NewSampleFormHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/samples/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	sampleData := sample.Sample{}
	mockUsecase.On("Fetch", mock.Anything, 1).Return(sampleData, nil)

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestSampleFormHandler_ServeHTTP_InvalidIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleFormUsecase)
	h := NewSampleFormHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/samples/invalid", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSampleFormHandler_ServeHTTP_ZeroIDParameter(t *testing.T) {
	mockUsecase := new(MockSampleFormUsecase)
	h := NewSampleFormHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/samples/0", nil)
	vars := map[string]string{
		"id": "0",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
}

func TestSampleFormHandler_ServeHTTP_UsecaseError(t *testing.T) {
	mockUsecase := new(MockSampleFormUsecase)
	h := NewSampleFormHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/samples/1", nil)

	vars := map[string]string{
		"id": "1",
	}
	req = mux.SetURLVars(req, vars)
	w := httptest.NewRecorder()

	mockUsecase.On("Fetch", mock.Anything, 1).Return(sample.Sample{}, errors.New("usecase error"))

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockUsecase.AssertExpectations(t)
}
