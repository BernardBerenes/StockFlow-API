package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/product"
	"github.com/gofiber/fiber/v3"
)

func ProductRouter(router fiber.Router, productService product.IService) {
	router.Get("list-paginate", handlers.ListPaginateProduct(productService))
	router.Get("list", handlers.ListProduct(productService))
	router.Post("create", handlers.CreateProduct(productService))
	router.Patch("update/:uuid", handlers.UpdateProduct(productService))
	router.Delete("delete/:uuid", handlers.DeleteProduct(productService))
}
