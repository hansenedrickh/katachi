package sample

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUsecase_Insert(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	name := "Sample"

	mockRepo.On("Insert", ctx, mock.Anything).Return(nil)

	err := usecase.Insert(ctx, name)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_InsertError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	name := "Sample"

	mockRepo.On("Insert", ctx, mock.Anything).Return(errors.New("some-error"))

	err := usecase.Insert(ctx, name)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Fetch(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1
	expectedSample := Sample{ID: uint(id), Name: "Sample"}

	mockRepo.On("Fetch", ctx, id).Return(expectedSample, nil)

	result, err := usecase.Fetch(ctx, id)
	assert.NoError(t, err)
	assert.Equal(t, expectedSample, result)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_FetchError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1

	mockRepo.On("Fetch", ctx, id).Return(Sample{}, errors.New("some-error"))

	result, err := usecase.Fetch(ctx, id)
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Update(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1
	name := "UpdatedSample"
	existingSample := Sample{ID: uint(id), Name: "Sample"}

	mockRepo.On("Fetch", ctx, id).Return(existingSample, nil)
	mockRepo.On("Update", ctx, mock.Anything).Return(nil)

	err := usecase.Update(ctx, id, name)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_UpdateWhenGetError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1
	name := "UpdatedSample"

	mockRepo.On("Fetch", ctx, id).Return(Sample{}, errors.New("some-error"))

	err := usecase.Update(ctx, id, name)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_UpdateError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1
	name := "UpdatedSample"
	existingSample := Sample{ID: uint(id), Name: "Sample"}

	mockRepo.On("Fetch", ctx, id).Return(existingSample, nil)
	mockRepo.On("Update", ctx, mock.Anything).Return(errors.New("some-error"))

	err := usecase.Update(ctx, id, name)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_Delete(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1

	mockRepo.On("Delete", ctx, id).Return(nil)

	err := usecase.Delete(ctx, id)
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_DeleteError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	id := 1

	mockRepo.On("Delete", ctx, id).Return(errors.New("some-error"))

	err := usecase.Delete(ctx, id)
	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_GetSuccess(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	page := 1
	limit := 10
	expectedSamples := []Sample{
		{ID: 1, Name: "Sample1"},
		{ID: 2, Name: "Sample2"},
	}

	mockRepo.On("Get", ctx, page, limit).Return(expectedSamples, nil)

	result, err := usecase.Get(ctx, page, limit)
	assert.NoError(t, err)
	assert.Equal(t, expectedSamples, result)
	mockRepo.AssertExpectations(t)
}

func TestUsecase_GetError(t *testing.T) {
	mockRepo := new(MockRepository)
	usecase := NewUsecase(mockRepo)
	ctx := context.TODO()
	page := 1
	limit := 10

	mockRepo.On("Get", ctx, page, limit).Return([]Sample{}, errors.New("some-error"))

	result, err := usecase.Get(ctx, page, limit)
	assert.Error(t, err)
	assert.Empty(t, result)
	mockRepo.AssertExpectations(t)
}
