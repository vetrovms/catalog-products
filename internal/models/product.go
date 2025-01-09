package models

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title        string  `json:"title"`
	Price        float64 `json:"price"`
	Color        string  `json:"color"`
	Currency     string  `json:"currency"`
	BestBefore   string  `json:"bestbefore"`
	Manufacturer string  `json:"manufacturer"`
}

func (p *Product) DTO() ProductDTO {
	return ProductDTO{
		ID:           int(p.ID),
		Title:        p.Title,
		Price:        p.Price,
		Color:        p.Color,
		Currency:     p.Currency,
		BestBefore:   p.BestBefore,
		Manufacturer: p.Manufacturer,
	}
}
