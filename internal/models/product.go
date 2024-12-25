package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title        string  `json:"title" validate:"required,max=255"`
	Price        float64 `json:"price" validate:"gte=0,lte=1000000000,numeric"`
	Color        string  `json:"color" validate:"iscolor"`
	Currency     string  `json:"currency" validate:"iso4217"`
	BestBefore   string  `json:"bestbefore" validate:"datetime=2006-01-02T15:04:05Z"`
	Manufacturer string  `json:"manufacturer" validate:"max=255"`
}

func (p *Product) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		ID           uint    `json:"id"`
		Title        string  `json:"title"`
		Price        float64 `json:"price"`
		Color        string  `json:"color"`
		Currency     string  `json:"currency"`
		BestBefore   string  `json:"bestbefore"`
		Manufacturer string  `json:"manufacturer"`
	}{
		ID:           p.ID,
		Title:        p.Title,
		Price:        p.Price,
		Color:        p.Color,
		Currency:     p.Currency,
		BestBefore:   p.BestBefore,
		Manufacturer: p.Manufacturer,
	})
}
