package repository

import (
	"context"
	"fmt"
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/gorm"
)

type PoskoRepository interface {
	FindAll(ctx context.Context, id_bencana string) []domain.Posko
	FindById(ctx context.Context, id_bencana string, id_posko string) (*domain.Posko, error)
	Create(ctx context.Context, id_bencana string, posko domain.Posko) (*domain.Posko, error)
}

type poskoRepository struct {
	db *gorm.DB
}

func NewPoskoRepository(db *gorm.DB) *poskoRepository {
	return &poskoRepository{db: db}
}

func (r *poskoRepository) FindAll(ctx context.Context, id_bencana string) []domain.Posko {
	var poskoSlice []domain.Posko

	r.db.WithContext(ctx).Where("bencana_id = ?", id_bencana).Find(&poskoSlice)

	return poskoSlice
}

func (r *poskoRepository) FindById(ctx context.Context, id_bencana string, id_posko string) (*domain.Posko, error) {
	var posko domain.Posko

	result := r.db.WithContext(ctx).Where("bencana_id = ?", id_bencana).Find(&posko, id_posko)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &posko, nil
}

func (r *poskoRepository) Create(ctx context.Context, id_bencana string, posko domain.Posko) (*domain.Posko, error) {
	err := r.db.WithContext(ctx).Where("bencana_id = ?", id_bencana).Create(&posko).Error

	return &posko, err
}
