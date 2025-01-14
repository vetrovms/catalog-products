package controllers

import (
	"catalog-products/internal/logger"
	"catalog-products/internal/models"
	"catalog-products/internal/request"
	"catalog-products/internal/services"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// ProductController. Контролер товарів
type ProductController struct {
	ps *services.ProductService
}

// NewProductController. Конструктор контролера товарів
func NewProductController(service *services.ProductService) ProductController {
	return ProductController{
		ps: service,
	}
}

// GetProducts. Обробник список товарів
func (pc *ProductController) GetProducts(c *fiber.Ctx) error {
	products, err := pc.ps.List(c.Queries())
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	return c.JSON(products)
}

// GetProduct. Обробник інформація про товар
func (pc *ProductController) GetProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	product, err := pc.ps.One(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	if product.ID == 0 {
		return c.JSON(map[string]string{
			"error": productNotFound().Error(),
		})
	}

	return c.JSON(product)
}

// AddProduct. Обробник створення нового товара
func (pc *ProductController) AddProduct(c *fiber.Ctx) error {
	var product models.ProductDTO
	var productRequest request.ProductRequest

	if err := c.BodyParser(&productRequest); err != nil {
		logger.Log().Info(err)
		return err
	}
	if err := productRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	productRequest.Fill(&product)
	product, err := pc.ps.Create(product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(product)
}

// UpdateProduct. Обробник оновлення товара
func (pc *ProductController) UpdateProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	var productRequest request.ProductRequest

	if err := c.BodyParser(&productRequest); err != nil {
		logger.Log().Info(err)
		return err
	}

	if err := productRequest.Validate(); err != nil {
		logger.Log().Info(strings.Join(err, ";"))
		return c.Status(http.StatusBadRequest).JSON(err)
	}

	exists, err := pc.ps.Exists(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]string{
			"error": productNotFound().Error(),
		})
	}

	var product = models.ProductDTO{}
	productRequest.Fill(&product)
	product.ID = id
	product, err = pc.ps.Update(product)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(product)
}

// TrashProduct. Обробник м'яке видалення товара
func (pc *ProductController) TrashProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	exists, err := pc.ps.Exists(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]string{
			"error": productNotFound().Error(),
		})
	}

	dto, err := pc.ps.SoftDelete(models.ProductDTO{ID: id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(dto)
}

// RecoverProduct. Обробник відновлення товара
func (pc *ProductController) RecoverProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	exists, err := pc.ps.ExistsUnscoped(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	if !exists {
		err = productNotFound()
		return c.Status(http.StatusNotFound).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	dto, err := pc.ps.Recover(models.ProductDTO{ID: id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(dto)
}

// RemoveProduct. Обробник остаточне видалення товара
func (pc *ProductController) RemoveProduct(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	exists, err := pc.ps.ExistsUnscoped(id)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}
	if !exists {
		return c.Status(http.StatusNotFound).JSON(map[string]string{
			"error": productNotFound().Error(),
		})
	}

	dto, err := pc.ps.Delete(models.ProductDTO{ID: id})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(map[string]string{
			"error": err.Error(),
		})
	}

	return c.JSON(dto)
}

// productNotFound. Логування помилки про запит на неіснуючий товар
func productNotFound() error {
	msg := "product not found"
	logger.Log().Info(msg)
	return errors.New(msg)
}
