package service

import (
	"barbot/internal/domain"
	"barbot/internal/repository"
)

type MixedDrinksService interface {
	GetAll() []domain.MixedDrinkDTO
}

type mixedDrinksService struct {
	repo repository.MixedDrinksRepository
}

func NewMixedDrinksService(repo repository.MixedDrinksRepository) MixedDrinksService {
	return &mixedDrinksService{repo: repo}
}

func (s *mixedDrinksService) GetAll() []domain.MixedDrinkDTO {
	mds := s.repo.GetAll()

	var mdDTOs []domain.MixedDrinkDTO
	for _, md := range mds {
		mdDTOs = append(mdDTOs, md.ToDTO())
	}

	return mdDTOs
}
