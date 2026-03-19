package handlers

import (
	"errors"
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

func CreateStore(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateRequest

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			ctx.Status(http.StatusBadRequest)
			return ctx.JSON(&fiber.Map{
				"message": err.Error(),
			})
		}

		err = service.CreateStore(&requestBody)
		if err != nil {
			if validationErrors, ok := errors.AsType[validator.ValidationErrors](err); ok {
				errs := presenter.FormatValidationError(validationErrors)
				return presenter.ErrorResponse(ctx, http.StatusBadRequest, "", errs)
			}

			return presenter.ErrorResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}
