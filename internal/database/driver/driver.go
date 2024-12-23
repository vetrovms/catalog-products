package driver

import (
	"catalog-products/internal/helpers"
	"os"

	"github.com/subosito/gotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	gotenv.Load()
	if DB != nil {
		return
	}
	cs := os.Getenv("POSTGRES_PRODUCTS_API_DSN")
	DB, err = gorm.Open(postgres.Open(cs))
	if err != nil {
		helpers.Log.Fatal("failed to connect database")
	}
}
