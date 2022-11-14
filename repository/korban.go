package repository

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/gorm"
)

type KorbanRepository interface {
	FindAllByPosko(ctx context.Context, idPosko string) []domain.Korban
	//FindByKondisi(ctx context.Context, kondisi string) []domain.Korban
	FindById(ctx context.Context, idKorban uuid.UUID) (*domain.Korban, error)
	Create(ctx context.Context, korban domain.Korban) (*domain.Korban, error)
	Update(ctx context.Context, korban domain.Korban) (*domain.Korban, error)
	Delete(ctx context.Context, idKorban uuid.UUID) error
}

type korbanRepository struct {
	db *gorm.DB
}

func NewKorbanRepository(db *gorm.DB) *korbanRepository {
	return &korbanRepository{db: db}
}

func (r *korbanRepository) FindAllByPosko(ctx context.Context, idPosko string) []domain.Korban {
	var korbanSlice []domain.Korban

	r.db.WithContext(ctx).Where("posko_id = ?", idPosko).Find(&korbanSlice)

	return korbanSlice
}

func (r *korbanRepository) FindById(ctx context.Context, idKorban uuid.UUID) (*domain.Korban, error) {
	var korban domain.Korban

	result := r.db.WithContext(ctx).Joins("Posko").Find(&korban, idKorban)

	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}

	if result.RowsAffected == 0 {
		return nil, fmt.Errorf("%v", "no record found")
	}

	return &korban, nil
}

//func (r *korbanRepository) FindByKondisi(ctx context.Context, kondisi string) []domain.Korban {
//	var korbanSlice []domain.Korban
//
//	r.db.Where(ctx).Where(&domain.Korban{Kondisi: kondisi})
//
//	return korbanSlice
//}

func (r *korbanRepository) Create(ctx context.Context, korban domain.Korban) (*domain.Korban, error) {
	err := r.db.WithContext(ctx).Create(&korban).Error

	return &korban, err
}

func (r *korbanRepository) Update(ctx context.Context, korban domain.Korban) (*domain.Korban, error) {
	err := r.db.WithContext(ctx).Save(&korban).Error

	return &korban, err
}

func (r *korbanRepository) Delete(ctx context.Context, idKorban uuid.UUID) error {
	err := r.db.WithContext(ctx).Delete(domain.Korban{}, idKorban).Error
	if err != nil {
		return fmt.Errorf("%v", "Cannot delete (violates Foreign Key constraint)")
	}

	return err
}
