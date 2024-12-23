package request

import (
	"catalog-products/internal/models"
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
)

type ProductRequest struct {
	Title        string    `json:"title" validate:"omitempty,max=255"`
	Price        float64   `json:"price" validate:"gte=0,lte=1000000000,numeric"`
	Color        string    `json:"color" validate:"omitempty,iscolor"`
	Currency     string    `json:"currency" validate:"omitempty,iso4217"`
	BestBefore   time.Time `json:"bestbefore"`
	Manufacturer string    `json:"manufacturer" validate:"omitempty,max=255"`
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
}

func (r ProductRequest) Validate() []string {
	var res []string
	validate := validator.New()
	errs := validate.Struct(r)

	if errs != nil {
		errMsgs := make([]string, 0)
		for _, err := range errs.(validator.ValidationErrors) {
			errMsgs = append(errMsgs, fmt.Sprintf(
				"[%s]: не відповідає правилу '%s' %s",
				err.StructField(),
				err.Tag(),
				err.Param(),
			))
		}
		return errMsgs
	}

	return res
}
