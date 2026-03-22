package handlers

import (
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
)

func ListStore(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		stores, err := service.ListStore()

		if err != nil {
			return err
		}

		return presenter.SuccessResponse(ctx, 200, "Successfully get data", stores)
	}
}

func CreateStore(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestStore

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return presenter.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		}

		err = service.CreateStore(&requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}

func UpdateStore(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestStore

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

		err = service.UpdateStore(parsedUuid, &requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully update data", nil)
	}
}

func DeleteStore(service store.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.DeleteStore(parsedUuid)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully delete data", nil)
	}
}
