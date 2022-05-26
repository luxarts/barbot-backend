package service

import (
	"barbot/internal/domain"
	"barbot/internal/repository"
)

type DrinksService interface {
	GetAll() []domain.Drink
}

type drinksService struct {
	repo repository.DrinksRepository
}

func NewDrinksService(repo repository.DrinksRepository) DrinksService {
	return &drinksService{repo: repo}
}

func (s *drinksService) GetAll() []domain.Drink {
	return s.repo.GetAll()
}
