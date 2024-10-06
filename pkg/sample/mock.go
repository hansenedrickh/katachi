package sample

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) Insert(ctx context.Context, sample Sample) error {
	args := m.Called(ctx, sample)
	return args.Error(0)
}

func (m *MockRepository) Fetch(ctx context.Context, id int) (Sample, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(Sample), args.Error(1)
}

func (m *MockRepository) Get(ctx context.Context, page, limit int) ([]Sample, error) {
	args := m.Called(ctx, page, limit)
	return args.Get(0).([]Sample), args.Error(1)
}

func (m *MockRepository) Update(ctx context.Context, sample Sample) error {
	args := m.Called(ctx, sample)
	return args.Error(0)
}

func (m *MockRepository) Delete(ctx context.Context, id int) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}
