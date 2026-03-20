package config

import (
	"errors"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"gorm.io/gorm"
)

func NewFiber() *fiber.App {
	return fiber.New(fiber.Config{
		AppName:       "StockFlow",
		CaseSensitive: true,
		ErrorHandler:  ErrorHandler(),
		StrictRouting: true,
	})
}

func ErrorHandler() fiber.ErrorHandler {
	return func(ctx fiber.Ctx, err error) error {
		status := fiber.StatusInternalServerError
		message := "Something went wrong"

		if validationErrors, ok := errors.AsType[validator.ValidationErrors](err); ok {
			status = fiber.StatusBadRequest
			errs := presenter.FormatValidationError(validationErrors)

			return presenter.ErrorResponse(ctx, status, "Validation error", errs)
		}

		if errors.Is(err, gorm.ErrRecordNotFound) {
			status = fiber.StatusNotFound
			message = "Record not found"

			return presenter.ErrorResponse(ctx, status, message, nil)
		}

		if fiberErr, ok := errors.AsType[*fiber.Error](err); ok {
			status = fiberErr.Code
			message = fiberErr.Message

			return presenter.ErrorResponse(ctx, status, message, nil)
		}

		return presenter.ErrorResponse(ctx, status, message, nil)
	}
}
