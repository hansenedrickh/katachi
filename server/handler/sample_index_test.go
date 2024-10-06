package handler

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hansenedrickh/katachi/pkg/sample"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockSampleIndexUsecase struct {
	mock.Mock
}

func (m *MockSampleIndexUsecase) Get(ctx context.Context, page, limit int) ([]sample.Sample, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]sample.Sample), args.Error(1)
}

func TestSampleIndexHandler_ServeHTTP_Success(t *testing.T) {
	mockUsecase := new(MockSampleIndexUsecase)

	h := NewSampleIndexHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	mockUsecase.On("Get", mock.Anything, 1, 10).Return([]sample.Sample{}, nil)

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusOK, w.Code)
	mockUsecase.AssertExpectations(t)
}

func TestSampleIndexHandler_ServeHTTP_InvalidPageParameter(t *testing.T) {
	mockUsecase := new(MockSampleIndexUsecase)
	h := NewSampleIndexHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/?page=invalid&limit=10", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSampleIndexHandler_ServeHTTP_InvalidLimitParameter(t *testing.T) {
	mockUsecase := new(MockSampleIndexUsecase)
	h := NewSampleIndexHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/?page=1&limit=invalid", nil)
	w := httptest.NewRecorder()

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusBadRequest, w.Code)
}

func TestSampleIndexHandler_ServeHTTP_UsecaseError(t *testing.T) {
	mockUsecase := new(MockSampleIndexUsecase)
	h := NewSampleIndexHandler(mockUsecase)
	req := httptest.NewRequest("GET", "/?page=1&limit=10", nil)
	w := httptest.NewRecorder()

	mockUsecase.On("Get", mock.Anything, 1, 10).Return([]sample.Sample{}, errors.New("usecase error"))

	h.ServeHTTP(w, req)

	require.Equal(t, http.StatusInternalServerError, w.Code)
	mockUsecase.AssertExpectations(t)
}
