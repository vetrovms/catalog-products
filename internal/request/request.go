package request

import (
	"catalog-products/internal/models"
	"catalog-products/internal/validator"
)

type ProductRequest struct {
	Title        string  `json:"title" validate:"omitempty,max=255"`
	Price        float64 `json:"price" validate:"gte=0,lte=1000000000,numeric"`
	Color        string  `json:"color" validate:"omitempty,iscolor"`
	Currency     string  `json:"currency" validate:"omitempty,iso4217"`
	BestBefore   string  `json:"bestbefore" validate:"omitempty,datetime=2006-01-02T15:04:05Z"`
	Manufacturer string  `json:"manufacturer" validate:"omitempty,max=255"`
}

func (r ProductRequest) Fill(m *models.ProductDTO) {
	if r.Title != "" {
		m.Title = r.Title
	}
	if r.Price != 0 {
		m.Price = r.Price
	}
	if r.Color != "" {
		m.Color = r.Color
	}
	if r.Currency != "" {
		m.Currency = r.Currency
	}
	if r.Manufacturer != "" {
		m.Manufacturer = r.Manufacturer
	}
	if r.BestBefore != "" {
		m.BestBefore = r.BestBefore
	}
}

func (r ProductRequest) Validate() []string {
	return validator.Validate(r)
}
