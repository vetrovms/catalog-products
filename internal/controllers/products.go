package controllers

import (
	"catalog-products/internal/database/repository"
	"catalog-products/internal/helpers"
	"catalog-products/internal/models"
	"catalog-products/internal/request"
	"errors"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type ProductController struct {
	Repo *repository.ProductRepo
}

// список товарів
func (pc *ProductController) GetProducts() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var products []models.Product
		pc.Repo.List(&products, c.Queries())
		return c.JSON(products)
	}
}

// інформація про товар
func (pc *ProductController) GetProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := checkInvalidId(c)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		var product models.Product
		pc.Repo.One(&product, id)
		if err := productNotFound(product); err != nil {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		//return c.JSON(product) // чомусь не спрацьовує MarshalJSON для об'єкта, але спрацьовує для слайса (getProducts)
		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

// створення нового товара
func (pc *ProductController) AddProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var product models.Product
		var productRequest request.ProductRequest

		if err := c.BodyParser(&productRequest); err != nil {
			helpers.Log.Info(err)
			return err
		}
		if err := productRequest.Validate(); err != nil {
			helpers.Log.Info(strings.Join(err, ";"))
			return c.Status(http.StatusBadRequest).JSON(err)
		}

		productRequest.Fill(&product)
		pc.Repo.Save(&product)

		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

// оновлення товара
func (pc *ProductController) UpdateProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := checkInvalidId(c)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		var product models.Product
		pc.Repo.One(&product, id)
		if err := productNotFound(product); err != nil {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		var productRequest request.ProductRequest

		if err := c.BodyParser(&productRequest); err != nil {
			helpers.Log.Info(err)
			return err
		}
		if err := productRequest.Validate(); err != nil {
			helpers.Log.Info(strings.Join(err, ";"))
			return c.Status(http.StatusBadRequest).JSON(err)
		}

		productRequest.Fill(&product)
		pc.Repo.Save(&product)

		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

// м'яке видалення
func (pc *ProductController) TrashProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := checkInvalidId(c)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		var product models.Product
		pc.Repo.One(&product, id)
		if err := productNotFound(product); err != nil {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		pc.Repo.SoftDelete(&product)
		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

// відновлення
func (pc *ProductController) RecoverProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := checkInvalidId(c)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		var product models.Product
		pc.Repo.OneUnscoped(&product, id)
		if err := productNotFound(product); err != nil {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		pc.Repo.Recover(&product)

		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

// остаточне видалення товара
func (pc *ProductController) RemoveProduct() func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		id, err := checkInvalidId(c)
		if err != nil {
			return c.Status(http.StatusBadRequest).SendString(err.Error())
		}

		var product models.Product
		pc.Repo.One(&product, id)
		if err := productNotFound(product); err != nil {
			return c.Status(http.StatusNotFound).SendString(err.Error())
		}

		pc.Repo.Delete(&product)

		json, _ := product.MarshalJSON()
		c.Set("Content-Type", "application/json")
		return c.SendString(string(json))
	}
}

func checkInvalidId(c *fiber.Ctx) (int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		msg := "invalid param id"
		helpers.Log.Info(msg)
		return 0, errors.New(msg)
	}
	return id, nil
}

func productNotFound(product models.Product) error {
	if product.ID == 0 {
		msg := "product not found"
		helpers.Log.Info(msg)
		return errors.New(msg)
	}
	return nil
}
