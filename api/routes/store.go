package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/gofiber/fiber/v3"
)

func StoreRouter(router fiber.Router, storeService store.IService) {
	router.Get("list", handlers.ListStore(storeService))
	router.Post("create", handlers.CreateStore(storeService))
	router.Patch("update/:uuid", handlers.UpdateStore(storeService))
	router.Delete("delete/:uuid", handlers.DeleteStore(storeService))
}
