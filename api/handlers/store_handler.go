package handlers

import (
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func List(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		stores, err := service.List()

		if err != nil {
			return err
		}

		return presenter.SuccessResponse(ctx, 200, "Successfully get data", stores)
	}
}

func Create(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequest

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return presenter.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		}

		err = service.Create(&requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}

func Update(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequest

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return err
		}

		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.Update(parsedUuid, &requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully update data", nil)
	}
}

func Delete(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.Delete(parsedUuid)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully delete data", nil)
	}
}
