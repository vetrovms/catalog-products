package main

import (
	"catalog-products/internal/controllers"
	"catalog-products/internal/database/driver"
	"catalog-products/internal/database/repository"
	"catalog-products/internal/helpers"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

var (
	app *fiber.App
	pc  controllers.ProductController
)

func init() {
	app = fiber.New()
	gotenv.Load()
	pc = controllers.ProductController{
		Repo: &repository.ProductRepo{
			DB: driver.DB,
		},
	}
}

func main() {
	app.Get("/api/v1/products", pc.GetProducts())                  // список товарів
	app.Get("/api/v1/products/:id", pc.GetProduct())               // інформація про товар
	app.Post("/api/v1/products", pc.AddProduct())                  // створення нового товара
	app.Put("/api/v1/products/:id", pc.UpdateProduct())            // оновлення товара
	app.Patch("/api/v1/products/:id/trash", pc.TrashProduct())     // м'яке видалення
	app.Patch("/api/v1/products/:id/recover", pc.RecoverProduct()) // відновлення
	app.Delete("/api/v1/products/:id", pc.RemoveProduct())         // остаточне видалення товара
	helpers.Log.Fatal(app.Listen(":8080"))
}
