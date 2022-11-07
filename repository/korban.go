package repository

import (
	"context"
	"fmt"
	"github.com/sikoba-tm/api/core/domain"
	"gorm.io/gorm"
)

type KorbanRepository interface {
	FindAllByPosko(ctx context.Context, id_posko string) []domain.Korban
	//FindByKondisi(ctx context.Context, kondisi string) []domain.Korban
	FindById(ctx context.Context, id_korban string) (*domain.Korban, error)
	Create(ctx context.Context, id_bencana string, korban domain.Korban) (*domain.Korban, error)
	//Update(ctx context.Context, korban domain.Korban) (*domain.Korban, error)
}

type korbanRepository struct {
	db *gorm.DB
}

func NewKorbanRepository(db *gorm.DB) *korbanRepository {
	return &korbanRepository{db: db}
}

func (r *korbanRepository) FindAllByPosko(ctx context.Context, id_posko string) []domain.Korban {
	var korbanSlice []domain.Korban

	r.db.WithContext(ctx).Where("posko_id = ?", id_posko).Find(&korbanSlice)

	return korbanSlice
}

func (r *korbanRepository) FindById(ctx context.Context, id_korban string) (*domain.Korban, error) {
	var korban domain.Korban

	result := r.db.WithContext(ctx).Joins("Posko").Find(&korban, id_korban)

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

func (r *korbanRepository) Create(ctx context.Context, id_bencana string, korban domain.Korban) (*domain.Korban, error) {
	fmt.Println(korban)
	err := r.db.WithContext(ctx).Create(&korban).Error

	return &korban, err
}
