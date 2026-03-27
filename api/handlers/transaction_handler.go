package handlers

import (
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ListTransaction(service transaction.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		transactions, err := service.ListTransaction()

		if err != nil {
			return err
		}

		return presenter.SuccessResponse(ctx, 200, "Successfully get data", transactions)
	}
}

func CreateTransaction(service transaction.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestTransaction

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return presenter.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		}

		err = service.CreateTransaction(&requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}

func UpdateTransaction(service transaction.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestTransaction

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return err
		}

		var parsedUuid uuid.UUID

		uuidParam := ctx.Params("uuid")
		parsedUuid, err = uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.UpdateTransaction(parsedUuid, &requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully update data", nil)
	}
}

func DeleteTransaction(service transaction.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.DeleteTransaction(parsedUuid)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully delete data", nil)
	}
}
