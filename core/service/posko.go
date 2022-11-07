package service

import (
	"context"
	"github.com/sikoba-tm/api/core/domain"
	"github.com/sikoba-tm/api/repository"
	"strconv"
)

type PoskoService interface {
	FindAll(ctx context.Context, id_bencana string) []domain.Posko
	FindById(ctx context.Context, id_bencana string, id_posko string) (*domain.Posko, error)
	Create(ctx context.Context, id_bencana string, posko domain.Posko) (*domain.Posko, error)
}
type poskoService struct {
	repo repository.PoskoRepository
}

func NewPoskoService(repo repository.PoskoRepository) *poskoService {
	return &poskoService{repo: repo}
}
func (s *poskoService) FindAll(ctx context.Context, id_bencana string) []domain.Posko {
	poskoResult := s.repo.FindAll(ctx, id_bencana)

	return poskoResult
}

func (s *poskoService) FindById(ctx context.Context, id_bencana string, id_posko string) (*domain.Posko, error) {
	posko, err := s.repo.FindById(ctx, id_bencana, id_posko)

	return posko, err
}

func (s *poskoService) Create(ctx context.Context, id_bencana string, posko domain.Posko) (*domain.Posko, error) {
	idBencana64, _ := strconv.ParseUint(id_bencana, 10, 64)
	posko.BencanaID = uint(idBencana64)
	newPosko, err := s.repo.Create(ctx, id_bencana, posko)

	return newPosko, err
}
