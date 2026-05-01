package handlers

import (
	"errors"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/product"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/valyala/fasthttp"
)

func ListPaginateProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var paginateRequest presenter.PaginateRequest

		err := ctx.Bind().Query(&paginateRequest)
		if err != nil {
			return err
		}

		name := ctx.Query("name")

		var paginateProducts *presenter.PaginateResponse[presenter.ProductResponse]

		paginateProducts, err = service.ListPaginateProduct(&paginateRequest, name)
		if err != nil {
			return err
		}

		return presenter.SuccessResponsePaginate(ctx, 200, "Successfully get data", paginateProducts.Data, paginateProducts.PaginateMetadata)
	}
}

func ListProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		name := ctx.Query("name")

		products, err := service.ListProduct(name)

		if err != nil {
			return err
		}

		return presenter.SuccessResponse(ctx, 200, "Successfully get data", products)
	}
}

func CreateProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestProduct

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return err
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

func UpdateProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		var requestBody presenter.CreateUpdateRequestProduct

		err := ctx.Bind().Body(&requestBody)
		if err != nil {
			return err
		}

		requestBody.Photo, err = ctx.FormFile("photo")
		if err != nil {
			if errors.Is(err, fasthttp.ErrMissingFile) {
				requestBody.Photo = nil
			} else {
				return err
			}
		}

		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.UpdateProduct(parsedUuid, &requestBody)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully update data", nil)
	}
}

func DeleteProduct(service product.IService) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		uuidParam := ctx.Params("uuid")
		parsedUuid, err := uuid.Parse(uuidParam)
		if err != nil {
			return err
		}

		err = service.DeleteProduct(parsedUuid)
		if err != nil {
			return err
		}

		return presenter.SuccessResponse[any](ctx, 200, "Successfully delete data", nil)
	}
}
