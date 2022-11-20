package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/core/service/external"
	"github.com/sikoba-tm/api/repository"
)

type KorbanService interface {
	FindAll(ctx context.Context, idBencana string) []domain.Korban
	FindById(ctx context.Context, idKorban uuid.UUID) (*domain.Korban, error)
	Create(ctx context.Context, idBencana string, idPosko string, korban domain.Korban) (*domain.Korban, error)
	Update(ctx context.Context, idKorban uuid.UUID, korban domain.Korban) (*domain.Korban, error)
	Delete(ctx context.Context, idKorban uuid.UUID) error
	SearchByFoto(ctx context.Context, idBencana string, reference []byte) ([]domain.Korban, error)
}

type korbanService struct {
	repoKorban repository.KorbanRepository
	repoPosko  repository.PoskoRepository
	gcs        external.CloudStorageService
	rekog      external.RekognitionService
}

func NewKorbanService(repoKorban repository.KorbanRepository, repoPosko repository.PoskoRepository, gcs external.CloudStorageService,
	rekog external.RekognitionService) *korbanService {
	return &korbanService{repoKorban: repoKorban, repoPosko: repoPosko, gcs: gcs, rekog: rekog}
}

func (s *korbanService) FindAll(ctx context.Context, idBencana string) []domain.Korban {
	allPosko := s.repoPosko.FindAllId(ctx, idBencana)

	var results = make([]domain.Korban, 0)
	for _, e := range allPosko {
		korbans := s.repoKorban.FindAllByPosko(ctx, e)
		for _, k := range korbans {
			results = append(results, k)
		}
	}

	return results
}

func (s *korbanService) FindById(ctx context.Context, idKorban uuid.UUID) (*domain.Korban, error) {
	korban, err := s.repoKorban.FindById(ctx, idKorban)

	return korban, err
}

func (s *korbanService) Create(ctx context.Context, idBencana string, idPosko string, korban domain.Korban) (*domain.Korban, error) {
	posko, err := s.repoPosko.FindById(ctx, idBencana, idPosko)
	korban.Posko = *posko
	newKorban, err := s.repoKorban.Create(ctx, korban)

	return newKorban, err
}

func (s *korbanService) Update(ctx context.Context, idKorban uuid.UUID, korban domain.Korban) (*domain.Korban, error) {
	existingKorban, err := s.repoKorban.FindById(ctx, idKorban)
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

func (s *korbanService) Delete(ctx context.Context, idKorban uuid.UUID) error {
	_, err := s.repoKorban.FindById(ctx, idKorban)
	if err != nil {
		return err
	}

	err = s.repoKorban.Delete(ctx, idKorban)

	return err
}

func (s *korbanService) SearchByFoto(ctx context.Context, idBencana string, reference []byte) ([]domain.Korban, error) {
	const ThresholdLimit = 80.000
	var searchResults = make([]domain.Korban, 0)
	var idOfResults = make([]uuid.UUID, 0)
	var korbans = s.FindAll(ctx, idBencana)
	for _, k := range korbans {
		comparison, err := s.gcs.DownloadFile(ctx, "korban/", k.ID.String())

		if err != nil {
			return searchResults, err
		}
		comparisonRes, err := s.rekog.CompareImages(ctx, reference, comparison, ThresholdLimit)
		if err != nil {
			return searchResults, err
		}
		matches := comparisonRes.FaceMatches
		if len(matches) > 0 {
			idOfResults = append(idOfResults, k.ID)
		}

	}

	if len(idOfResults) != 0 {
		searchResults = s.repoKorban.FindAllKorbanByIdBulk(ctx, idOfResults)
	}

	return searchResults, nil
}
