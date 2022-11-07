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
	newKorban, err := s.repoKorban.Create(ctx, id_bencana, korban)

	return newKorban, err
}
