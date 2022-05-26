package repository

import "barbot/internal/domain"

type MixedDrinksRepository interface {
	GetAll() []domain.MixedDrink
}

type mixedDrinksRepository struct {
}

func NewMixedDrinksRepository() MixedDrinksRepository {
	return &mixedDrinksRepository{}
}

func (r *mixedDrinksRepository) GetAll() []domain.MixedDrink {
	return []domain.MixedDrink{
		{
			ID: 0, Name: "Fernet Cola", Drinks: []domain.DrinksPercentages{
				{ID: 0, Name: "Fernet", Percentage: 30},
				{ID: 1, Name: "Coca Cola", Percentage: 70},
			},
		},
		{
			ID: 1, Name: "Mojito", Drinks: []domain.DrinksPercentages{
				{ID: 2, Name: "Agua con gas", Percentage: 57},
				{ID: 3, Name: "Jugo de Limón", Percentage: 14},
				{ID: 4, Name: "Ron", Percentage: 29}},
		},
		{
			ID: 2, Name: "Gin Tonic", Drinks: []domain.DrinksPercentages{
				{ID: 5, Name: "Ginebra", Percentage: 20},
				{ID: 6, Name: "Agua tónica", Percentage: 80},
			},
		},
		{
			ID: 3, Name: "Martini", Drinks: []domain.DrinksPercentages{
				{ID: 5, Name: "Ginebra", Percentage: 83},
				{ID: 7, Name: "Vermut", Percentage: 17},
			},
		},
	}
}
