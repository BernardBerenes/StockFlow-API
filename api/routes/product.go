package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/product"
	"github.com/gofiber/fiber/v3"
)

func ProductRouter(router fiber.Router, productService product.IService) {
	router.Post("create", handlers.CreateProduct(productService))
}
