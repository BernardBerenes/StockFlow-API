package handlers

import (
	"errors"
	"net/http"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/product"
	"github.com/gofiber/fiber/v3"
	"github.com/valyala/fasthttp"
)

func CreateProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestProduct

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return presenter.ErrorResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		}

		requestBody.Photo, err = ctx.FormFile("photo")
		if err != nil {
			if errors.Is(err, fasthttp.ErrMissingFile) {
				requestBody.Photo = nil
			} else {
				return err
			}
		}

		err = service.CreateProduct(&requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully create data", nil)
	}
}
