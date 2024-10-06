package user

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (repo *Repository) Insert(ctx context.Context, user User) error {
	result := repo.db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return fmt.Errorf("[UserRepository] Insert Error for %v: %w", user, result.Error)
	}

	return nil
}

func (repo *Repository) Fetch(ctx context.Context, username string) (User, error) {
	var user User
	result := repo.db.WithContext(ctx).Model(User{Username: username}).First(&user)
	if result.Error != nil {
		return user, fmt.Errorf("[UserRepository] Fetch Error for %v: %w", username, result.Error)
	}

	return user, nil
}
