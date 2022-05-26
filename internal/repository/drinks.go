package repository

import "barbot/internal/domain"

type DrinksRepository interface {
	GetAll() []domain.Drink
}

type drinksRepository struct {
}

func NewDrinksRepository() DrinksRepository {
	return &drinksRepository{}
}

func (r *drinksRepository) GetAll() []domain.Drink {
	return []domain.Drink{
		{ID: 0, Name: "Fernet"},
		{ID: 1, Name: "Coca Cola"},
		{ID: 2, Name: "Agua con gas"},
		{ID: 3, Name: "Jugo de Limón"},
		{ID: 4, Name: "Ron"},
		{ID: 5, Name: "Ginebra"},
		{ID: 6, Name: "Agua tónica"},
		{ID: 7, Name: "Vermut"},
	}
}
