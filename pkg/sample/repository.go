package sample

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

func (repo *Repository) Insert(ctx context.Context, sample Sample) error {
	result := repo.db.WithContext(ctx).Create(&sample)
	if result.Error != nil {
		return fmt.Errorf("[SampleRepository] Insert Error for %v: %w", sample, result.Error)
	}

	return nil
}

func (repo *Repository) Fetch(ctx context.Context, id int) (Sample, error) {
	var sample Sample
	result := repo.db.WithContext(ctx).First(&sample, id)
	if result.Error != nil {
		return sample, fmt.Errorf("[SampleRepository] Fetch Error for %v: %w", id, result.Error)
	}

	return sample, nil
}

func (repo *Repository) Get(ctx context.Context, page, limit int) ([]Sample, error) {
	var samples []Sample
	result := repo.db.WithContext(ctx).
		Limit(limit).
		Offset((page - 1) * limit).
		Order("id").
		Find(&samples)
	if result.Error != nil {
		return samples, fmt.Errorf("[SampleRepository] Get Error for %v: %w", page, result.Error)
	}

	return samples, nil
}

func (repo *Repository) Update(ctx context.Context, sample Sample) error {
	result := repo.db.WithContext(ctx).Save(&sample)
	if result.Error != nil {
		return fmt.Errorf("[SampleRepository] Update Error for %v: %w", sample, result.Error)
	}

	return nil
}

func (repo *Repository) Delete(ctx context.Context, id int) error {
	result := repo.db.WithContext(ctx).Delete(&Sample{}, id)
	if result.Error != nil {
		return fmt.Errorf("[SampleRepository] Delete Error for %v: %w", id, result.Error)
	}

	return nil
}
