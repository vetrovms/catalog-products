package repository

import (
	"catalog-products/internal/models"
	"catalog-products/internal/query"
	"context"

	"gorm.io/gorm"
)

// ProductRepo. Репозиторій товарів
type ProductRepo struct {
	db *gorm.DB
}

// NewProductRepo. Конструктор репозиторія товарів
func NewProductRepo(conn *gorm.DB) ProductRepo {
	return ProductRepo{
		db: conn,
	}
}

// List. Повертає список товарів
func (rep *ProductRepo) List(ctx context.Context, params map[string]string) []models.Product {
	var products []models.Product
	cond, bindParams := query.SearchQuery(params)
	orderQuery := query.OrderQuery(params)
	rep.db.WithContext(ctx).Where(cond, bindParams...).Order(orderQuery).Find(&products)
	return products
}

// One. Повертає товар за ідентифікатором
func (rep *ProductRepo) One(ctx context.Context, id int) models.Product {
	var product models.Product
	rep.db.WithContext(ctx).First(&product, id)
	// rep.db.WithContext(ctx).Exec("select pg_sleep(10);") // @debug
	return product
}

// OneUnscoped. Повертає м'яко видалений товар за ідентифікатором
func (rep *ProductRepo) OneUnscoped(ctx context.Context, id int) models.Product {
	var product models.Product
	rep.db.WithContext(ctx).Unscoped().First(&product, id)
	return product
}

// Create. Створює новий товар
func (rep *ProductRepo) Create(ctx context.Context, product *models.Product) {
	rep.db.WithContext(ctx).Create(&product)
}

// Save. Зберігає товар
func (rep *ProductRepo) Save(ctx context.Context, product *models.Product) {
	rep.db.WithContext(ctx).Save(&product)
}

// SoftDelete. М'яке видалення товара
func (rep *ProductRepo) SoftDelete(ctx context.Context, product *models.Product) {
	rep.db.WithContext(ctx).Model(&product).Delete(&product)
}

// Recover. Відновлення м'яко видаленного товара
func (rep *ProductRepo) Recover(ctx context.Context, product *models.Product) {
	rep.db.WithContext(ctx).Unscoped().Model(&product).Update("DeletedAt", nil)
}

// Delete. Остаточне видалення товара
func (rep *ProductRepo) Delete(ctx context.Context, product *models.Product) {
	rep.db.WithContext(ctx).Unscoped().Delete(&product)
}

// Exists. Перевірка існування товара за ідентифікатором
func (rep *ProductRepo) Exists(ctx context.Context, id int) bool {
	var exists bool
	rep.db.WithContext(ctx).Model(models.Product{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists)
	return exists
}

// ExistsUnscoped. Перевірка існування м'яко видаленого товара за ідентифікатаром
func (rep *ProductRepo) ExistsUnscoped(ctx context.Context, id int) bool {
	var exists bool
	rep.db.WithContext(ctx).Unscoped().Model(models.Product{}).Select("count(*) > 0").Where("id = ?", id).Find(&exists)
	return exists
}
