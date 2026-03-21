package handlers

import (
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/gofiber/fiber/v3"
)

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
