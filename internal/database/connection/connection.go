package connection

import (
	"catalog-products/internal/logger"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const dsn = "POSTGRES_PRODUCTS_API_DSN"

// Db. Повертає з'єднання з базою даних
func Db() *gorm.DB {
	cs := os.Getenv(dsn)
	db, err := gorm.Open(postgres.Open(cs))
	if err != nil {
		logger.Log().Fatal("failed to connect database")
	}
	return db
}
