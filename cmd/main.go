package main

import (
	"catalog-products/internal/controllers"
	"catalog-products/internal/database/connection"
	"catalog-products/internal/database/repository"
	"catalog-products/internal/logger"
	"catalog-products/internal/services"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

const idRule = "<int;min(1);max(2147483647)>"
const dbPortEnvKey = "POSTGRES_PRODUCTS_API_PORT"

func main() {
	app := fiber.New()
	conn := connection.Db()
	productRepo := repository.NewProductRepo(conn)
	ps := services.NewProductService(&productRepo)
	pc := controllers.NewProductController(&ps)

	pg := app.Group("/api/v1/products")
	pg.Get("/", pc.GetProducts)                       // список товарів
	pg.Get(idRoute(), pc.GetProduct)                  // інформація про товар
	pg.Post("/", pc.AddProduct)                       // створення нового товара
	pg.Put(idRoute(), pc.UpdateProduct)               // оновлення товара
	pg.Patch(idRoute()+"/trash", pc.TrashProduct)     // м'яке видалення
	pg.Patch(idRoute()+"/recover", pc.RecoverProduct) // відновлення
	pg.Delete(idRoute(), pc.RemoveProduct)            // остаточне видалення товара

	logger.Log().Fatal(app.Listen(":" + os.Getenv(dbPortEnvKey)))
}

func idRoute() string {
	return "/:id" + idRule
}
