package models

type ProductDTO struct {
	ID           int     `json:"id"`
	Title        string  `json:"title"`
	Price        float64 `json:"price"`
	Color        string  `json:"color"`
	Currency     string  `json:"currency"`
	BestBefore   string  `json:"bestbefore"`
	Manufacturer string  `json:"manufacturer"`
}

func (dto *ProductDTO) FillModel(p *Product) {
	if dto.Title != "" {
		p.Title = dto.Title
	}
	if dto.Price != 0 {
		p.Price = dto.Price
	}
	if dto.Color != "" {
		p.Color = dto.Color
	}
	if dto.Currency != "" {
		p.Currency = dto.Currency
	}
	if dto.BestBefore != "" {
		p.BestBefore = dto.BestBefore
	}
	if dto.Manufacturer != "" {
		p.Manufacturer = dto.Manufacturer
	}
}
