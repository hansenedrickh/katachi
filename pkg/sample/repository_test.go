package sample

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&Sample{})
	return db, nil
}

func TestRepository_Insert(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := Sample{
		Name: "Sample",
	}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)
}

func TestRepository_Fetch(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := Sample{
		ID:   1,
		Name: "Sample",
	}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)

	result, err := repo.Fetch(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, req.Name, result.Name)

	// Error case: Fetch non-existent ID
	_, err = repo.Fetch(ctx, 999)
	assert.Error(t, err)
}

func TestRepository_Update(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := Sample{
		ID:   1,
		Name: "Sample",
	}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)

	req.Name = "UpdatedValue"
	err = repo.Update(ctx, req)
	assert.NoError(t, err)

	updatedSample, err := repo.Fetch(ctx, 1)
	assert.NoError(t, err)
	assert.Equal(t, "UpdatedValue", updatedSample.Name)
}

func TestRepository_Delete(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := Sample{ID: 1}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)

	err = repo.Delete(ctx, 1)
	assert.NoError(t, err)
}

func TestRepository_Get_ReturnsCorrectSamples(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	samples := []Sample{
		{Name: "Sample1"},
		{Name: "Sample2"},
		{Name: "Sample3"},
	}

	for _, sample := range samples {
		err = repo.Insert(ctx, sample)
		assert.NoError(t, err)
	}

	result, err := repo.Get(ctx, 1, 2)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Sample1", result[0].Name)
	assert.Equal(t, "Sample2", result[1].Name)
}

func TestRepository_Get_ReturnsEmptySliceForNoSamples(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	result, err := repo.Get(ctx, 1, 2)
	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestRepository_Get_ReturnsRemainingSamplesIfLessThanLimit(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	samples := []Sample{
		{Name: "Sample1"},
		{Name: "Sample2"},
	}

	for _, sample := range samples {
		err = repo.Insert(ctx, sample)
		assert.NoError(t, err)
	}

	result, err := repo.Get(ctx, 1, 3)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Sample1", result[0].Name)
	assert.Equal(t, "Sample2", result[1].Name)
}

func TestRepository_Get_ReturnsCorrectPage(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()

	samples := []Sample{
		{Name: "Sample1"},
		{Name: "Sample2"},
		{Name: "Sample3"},
		{Name: "Sample4"},
	}

	for _, sample := range samples {
		err = repo.Insert(ctx, sample)
		assert.NoError(t, err)
	}

	result, err := repo.Get(ctx, 2, 2)
	assert.NoError(t, err)
	assert.Len(t, result, 2)
	assert.Equal(t, "Sample3", result[0].Name)
	assert.Equal(t, "Sample4", result[1].Name)
}
