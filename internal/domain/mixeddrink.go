package domain

type MixedDrink struct {
	ID     int64                   `json:"id"`
	Name   string                  `json:"name"`
	Img    string                  `json:"img"`
	Drinks []DrinksWithPercentages `json:"drinks"`
}

type DrinksWithPercentages struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

type DrinksPercentagesMap map[int64]int

func (md MixedDrink) ToDTO() MixedDrinkDTO {
	var dto MixedDrinkDTO

	dto.ID = md.ID
	dto.Name = md.Name
	dto.Img = md.Img
	dto.Drinks = md.Drinks

	return dto
}

type MixedDrinkDTO struct {
	ID     int64                   `json:"id"`
	Name   string                  `json:"name"`
	Img    string                  `json:"img"`
	Drinks []DrinksWithPercentages `json:"drinks"`
}
