package user

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

	db.AutoMigrate(&User{})
	return db, nil
}

func TestRepository_Insert(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := User{
		Username: "User",
		Password: "Password",
	}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)
}

func TestRepository_Fetch(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	req := User{
		ID:       1,
		Username: "User",
		Password: "Password",
	}

	err = repo.Insert(ctx, req)
	assert.NoError(t, err)

	result, err := repo.Fetch(ctx, "User")
	assert.NoError(t, err)
	assert.Equal(t, req.Username, result.Username)
	assert.Equal(t, req.Password, result.Password)
}

func TestRepository_FetchError(t *testing.T) {
	db, err := setupTestDB()
	assert.NoError(t, err)

	repo := NewRepository(db)
	ctx := context.TODO()
	
	_, err = repo.Fetch(ctx, "Invalid")
	assert.Error(t, err)
}
