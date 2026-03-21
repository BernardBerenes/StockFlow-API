package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/gofiber/fiber/v3"
)

func TransactionRouter(router fiber.Router, transactionService transaction.IService) {
	router.Post("create", handlers.CreateTransaction(transactionService))
}
