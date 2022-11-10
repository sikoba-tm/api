package service

import (
	"context"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/repository"
)

type KorbanService interface {
	FindAll(ctx context.Context, id_bencana string) []domain.Korban
	FindById(ctx context.Context, id_korban string) (*domain.Korban, error)
	Create(ctx context.Context, id_bencana string, id_posko string, korban domain.Korban) (*domain.Korban, error)
	Update(ctx context.Context, id_korban string, korban domain.Korban) (*domain.Korban, error)
	Delete(ctx context.Context, id_korban string) error
}

type korbanService struct {
	repoKorban repository.KorbanRepository
	repoPosko  repository.PoskoRepository
}

func NewKorbanService(repoKorban repository.KorbanRepository, repoPosko repository.PoskoRepository) *korbanService {
	return &korbanService{repoKorban: repoKorban, repoPosko: repoPosko}
}

func (s *korbanService) FindAll(ctx context.Context, id_bencana string) []domain.Korban {
	allPosko := s.repoPosko.FindAllId(ctx, id_bencana)

	var results = make([]domain.Korban, 0)
	for _, e := range allPosko {
		korbans := s.repoKorban.FindAllByPosko(ctx, e)
		for _, k := range korbans {
			results = append(results, k)
		}
	}

	return results
}

func (s *korbanService) FindById(ctx context.Context, id_korban string) (*domain.Korban, error) {
	korban, err := s.repoKorban.FindById(ctx, id_korban)

	return korban, err
}

func (s *korbanService) Create(ctx context.Context, id_bencana string, id_posko string, korban domain.Korban) (*domain.Korban, error) {
	posko, err := s.repoPosko.FindById(ctx, id_bencana, id_posko)
	korban.Posko = *posko
	newKorban, err := s.repoKorban.Create(ctx, korban)

	return newKorban, err
}

func (s *korbanService) Update(ctx context.Context, id_korban string, korban domain.Korban) (*domain.Korban, error) {
	existingKorban, err := s.repoKorban.FindById(ctx, id_korban)
	if err != nil {
		return existingKorban, err
	}

	existingKorban.Nama = korban.Nama
	existingKorban.Kondisi = korban.Kondisi
	existingKorban.Foto = korban.Foto
	existingKorban.NamaIbuKandung = korban.NamaIbuKandung
	existingKorban.RentangUsia = korban.RentangUsia
	existingKorban.TempatLahir = korban.TempatLahir
	existingKorban.TanggalLahir = korban.TanggalLahir

	updatedKorban, err := s.repoKorban.Update(ctx, *existingKorban)

	return updatedKorban, err
}

func (s *korbanService) Delete(ctx context.Context, id_korban string) error {
	_, err := s.repoKorban.FindById(ctx, id_korban)
	if err != nil {
		return err
	}

	err = s.repoKorban.Delete(ctx, id_korban)

	return err
}
