package request

import (
	"catalog-products/internal/models"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Title        string  `json:"title" validate:"omitempty,max=255"`
	Price        float64 `json:"price" validate:"gte=0,lte=1000000000,numeric"`
	Color        string  `json:"color" validate:"omitempty,iscolor"`
	Currency     string  `json:"currency" validate:"omitempty,iso4217"`
	BestBefore   string  `json:"bestbefore" validate:"omitempty,datetime=2006-01-02T15:04:05Z"`
	Manufacturer string  `json:"manufacturer" validate:"omitempty,max=255"`
}

func (r ProductRequest) Fill(m *models.Product) {
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
	var res []string
	validate := validator.New()
	errs := validate.Struct(r)

	if errs != nil {
		errMsgs := make([]string, 0)
		errMap := map[string]string{
			"max":      ": довжина має бути не більше %s символів",
			"gte":      ": значення має бути більше %s",
			"lte":      ": значення має бути менше %s",
			"iscolor":  ": значення має відповідати формату кольора%s",
			"iso4217":  ": значення має відповідати формату iso4217%s",
			"datetime": ": невірний формати дати %s",
		}
		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, err.StructField()+fmt.Sprintf(errMap[err.Tag()], err.Param()))
		}
		return errMsgs
	}

	return res
}
