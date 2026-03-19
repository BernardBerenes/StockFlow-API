package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/gofiber/fiber/v3"
)

func StoreRouter(router fiber.Router, storeService store.IService) {
	router.Get("list", handlers.List(storeService))
	router.Post("create", handlers.Create(storeService))
	router.Patch("update/:uuid", handlers.Update(storeService))
}
