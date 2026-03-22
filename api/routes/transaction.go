package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/gofiber/fiber/v3"
)

func TransactionRouter(router fiber.Router, transactionService transaction.IService) {
	router.Get("list", handlers.ListTransaction(transactionService))
	router.Post("create", handlers.CreateTransaction(transactionService))
	router.Patch("update/:uuid", handlers.UpdateTransaction(transactionService))
	router.Delete("delete/:uuid", handlers.DeleteTransaction(transactionService))
}
