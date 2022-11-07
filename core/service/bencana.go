package service

import (
	"context"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/repository"
)

type BencanaService interface {
	FindAll(ctx context.Context) []domain.Bencana
	FindById(ctx context.Context, id string) (*domain.Bencana, error)
	Create(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error)
	Update(ctx context.Context, id string, bencana domain.Bencana) (*domain.Bencana, error)
}
type bencanaService struct {
	repo repository.BencanaRepository
}

func NewBencanaService(repo repository.BencanaRepository) *bencanaService {
	return &bencanaService{repo: repo}
}
func (s *bencanaService) FindAll(ctx context.Context) []domain.Bencana {
	bencanaResult := s.repo.FindAll(ctx)

	return bencanaResult
}

func (s *bencanaService) FindById(ctx context.Context, id string) (*domain.Bencana, error) {
	bencana, err := s.repo.FindById(ctx, id)

	return bencana, err
}

func (s *bencanaService) Create(ctx context.Context, bencana domain.Bencana) (*domain.Bencana, error) {
	newBencana, err := s.repo.Create(ctx, bencana)

	return newBencana, err
}

func (s *bencanaService) Update(ctx context.Context, id string, bencana domain.Bencana) (*domain.Bencana, error) {
	existingBencana, err := s.repo.FindById(ctx, id)
	if err != nil {
		return existingBencana, err
	}

	existingBencana.Nama = bencana.Nama
	existingBencana.Lokasi = bencana.Lokasi
	existingBencana.TanggalKejadian = bencana.TanggalKejadian

	newBencana, err := s.repo.Update(ctx, *existingBencana)

	return newBencana, err
}
