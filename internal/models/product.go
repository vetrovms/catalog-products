package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Title        string    `json:"title" validate:"required,max=255"`
	Price        float64   `json:"price" validate:"gte=0,lte=1000000000,numeric"`
	Color        string    `json:"color" validate:"iscolor"`
	Currency     string    `json:"currency" validate:"iso4217"`
	BestBefore   time.Time `json:"bestbefore"`
	Manufacturer string    `json:"manufacturer" validate:"max=255"`
}
