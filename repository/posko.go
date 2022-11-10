package repository

import (
	"context"
	"fmt"
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/gorm"
)

type PoskoRepository interface {
	FindAll(ctx context.Context, idBencana string) []domain.Posko
	FindById(ctx context.Context, idBencana string, idPosko string) (*domain.Posko, error)
	FindAllId(ctx context.Context, idBencana string) []string
	Create(ctx context.Context, posko domain.Posko) (*domain.Posko, error)
	Update(ctx context.Context, posko domain.Posko) (*domain.Posko, error)
	Delete(ctx context.Context, idPosko string) error
}

type poskoRepository struct {
	db *gorm.DB
}

func NewPoskoRepository(db *gorm.DB) *poskoRepository {
	return &poskoRepository{db: db}
}

func (r *poskoRepository) FindAll(ctx context.Context, idBencana string) []domain.Posko {
	var poskoSlice []domain.Posko

	r.db.WithContext(ctx).Where("bencana_id = ?", idBencana).Find(&poskoSlice)

	return poskoSlice
}

func (r *poskoRepository) FindById(ctx context.Context, idBencana string, idPosko string) (*domain.Posko, error) {
	var posko domain.Posko

	result := r.db.WithContext(ctx).Where("bencana_id = ?", idBencana).Find(&posko, idPosko)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &posko, nil
}

func (r *poskoRepository) FindAllId(ctx context.Context, idBencana string) []string {
	var poskoIds []string
	r.db.Model(&domain.Posko{}).Where("bencana_id = ?", idBencana).Pluck("id", &poskoIds)

	return poskoIds
}

func (r *poskoRepository) Create(ctx context.Context, posko domain.Posko) (*domain.Posko, error) {
	err := r.db.WithContext(ctx).Create(&posko).Error

	return &posko, err
}

func (r *poskoRepository) Update(ctx context.Context, posko domain.Posko) (*domain.Posko, error) {
	err := r.db.WithContext(ctx).Save(posko).Error

	return &posko, err
}

func (r *poskoRepository) Delete(ctx context.Context, idPosko string) error {
	err := r.db.WithContext(ctx).Delete(domain.Posko{}, idPosko).Error
	if err != nil {
		return fmt.Errorf("%v", "Cannot delete (violates Foreign Key constraint)")
	}

	return err
}
