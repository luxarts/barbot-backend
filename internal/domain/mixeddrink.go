package domain

type MixedDrink struct {
	ID     int64               `json:"id"`
	Name   string              `json:"name"`
	Drinks []DrinksPercentages `json:"drinks"`
}

type DrinksPercentages struct {
	ID         int64  `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

type DrinksPercentagesMap map[int64]int

func (md MixedDrink) ToDTO() MixedDrinkDTO {
	var dto MixedDrinkDTO

	dto.ID = md.ID
	dto.Name = md.Name
	dto.DrinksPercentages = make(DrinksPercentagesMap, len(md.Drinks))
	for _, d := range md.Drinks {
		dto.DrinksPercentages[d.ID] = d.Percentage
	}

	return dto
}

type MixedDrinkDTO struct {
	ID                int64                `json:"id"`
	Name              string               `json:"name"`
	DrinksPercentages DrinksPercentagesMap `json:"drinks_percentages"`
}
