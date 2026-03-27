package handlers

import (
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction_detail"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ListTransactionDetail(service transaction_detail.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		transactionUuidParam := ctx.Params("transaction_uuid")
		parsedTransactionUuid, err := uuid.Parse(transactionUuidParam)
		if err != nil {
			return err
		}

		transactionDetails, err := service.ListTransactionDetail(parsedTransactionUuid)
		if err != nil {
			return err
		}
		
		return presenter.SuccessResponse(ctx, 200, "Successfully get data", transactionDetails)
	}
}

func CreateTransactionDetail(service transaction_detail.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestTransactionDetail

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return presenter.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		}

		var parsedTransactionUuid uuid.UUID

		transactionUuidParam := ctx.Params("transaction_uuid")
		parsedTransactionUuid, err = uuid.Parse(transactionUuidParam)
		if err != nil {
			return err
		}

		err = service.CreateTransactionDetail(parsedTransactionUuid, requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}
