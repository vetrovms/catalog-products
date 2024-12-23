package repository

import (
	"catalog-products/internal/helpers"
	"catalog-products/internal/models"
	"time"

	"gorm.io/gorm"
)

type ProductRepo struct {
	DB *gorm.DB
}

func (rep *ProductRepo) List(products *[]models.Product, params map[string]string) {
	cond, bindParams := helpers.SearchQuery(params)
	orderQuery := helpers.OrderQuery(params)
	rep.DB.Where(cond, bindParams...).Order(orderQuery).Find(&products)
}

func (rep *ProductRepo) One(product *models.Product, id int) {
	rep.DB.First(&product, id)
}

func (rep *ProductRepo) OneUnscoped(product *models.Product, id int) {
	rep.DB.Unscoped().First(&product, id)
}

func (rep *ProductRepo) Create(product *models.Product) {
	rep.DB.Create(&product)
}

func (rep *ProductRepo) Save(product *models.Product) {
	rep.DB.Save(&product)
}

func (rep *ProductRepo) SoftDelete(product *models.Product) {
	rep.DB.Model(&product).Select("DeletedAt").Update("DeletedAt", time.Now())
}

func (rep *ProductRepo) Recover(product *models.Product) {
	rep.DB.Unscoped().Model(&product).Select("DeletedAt").Update("DeletedAt", nil)
}

func (rep *ProductRepo) Delete(product *models.Product) {
	rep.DB.Unscoped().Delete(&product)
}
