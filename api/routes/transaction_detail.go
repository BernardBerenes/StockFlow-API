package routes

import (
	"github.com/BernardBerenes/stockflow-api/api/handlers"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction_detail"
	"github.com/gofiber/fiber/v3"
)

func TransactionDetailRouter(router fiber.Router, transactionDetailService transaction_detail.IService) {
	router.Get("list/:transaction_uuid", handlers.ListTransactionDetail(transactionDetailService))
	router.Post("create/:transaction_uuid", handlers.CreateTransactionDetail(transactionDetailService))
}
