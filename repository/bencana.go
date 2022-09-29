package repository

import (
	"context"
	"fmt"
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/gorm"
)

type BencanaRepository interface {
	FindAll(ctx context.Context) []domain.Bencana
	FindById(ctx context.Context, id string) (*domain.Bencana, error)
	Create(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error)
	Update(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error)
}

type bencanaRepository struct {
	db *gorm.DB
}

func NewBencanaRepository(db *gorm.DB) *bencanaRepository {
	return &bencanaRepository{db: db}
}

func (r *bencanaRepository) FindAll(ctx context.Context) []domain.Bencana {
	var bencanaSlice []domain.Bencana

	r.db.WithContext(ctx).Find(&bencanaSlice)

	return bencanaSlice
}

func (r *bencanaRepository) FindById(ctx context.Context, id string) (*domain.Bencana, error) {
	var bencana domain.Bencana

	result := r.db.WithContext(ctx).Find(&bencana, id)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &bencana, nil
}

func (r *bencanaRepository) Create(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error) {
	err := r.db.WithContext(ctx).Create(&bencana).Error

	return &bencana, err
}

func (r *bencanaRepository) Update(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error) {
	err := r.db.Save(bencana).Error

	return &bencana, err
}
