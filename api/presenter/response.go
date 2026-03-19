package presenter

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type ErrorItem struct {
	Key     string `json:"key"`
	Message string `json:"message"`
}

type Success[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type Error struct {
	Message string      `json:"message,omitempty"`
	Errors  []ErrorItem `json:"errors"`
}

func SuccessResponse[T any](ctx fiber.Ctx, status int, message string, data T) error {
	return ctx.Status(status).JSON(Success[T]{
		Message: message,
		Data:    data,
	})
}

func ErrorResponse(ctx fiber.Ctx, status int, message string, errors []ErrorItem) error {
	if status == http.StatusInternalServerError {
		message = "Internal server error"
	}

	return ctx.Status(status).JSON(Error{
		Message: message,
		Errors:  errors,
	})
}

func FormatValidationError(err error) []ErrorItem {
	var result []ErrorItem

	var validationErrors validator.ValidationErrors
	ok := errors.As(err, &validationErrors)
	if !ok {
		return result
	}

	for _, e := range validationErrors {
		var message string

		switch e.Tag() {
		case "required":
			message = "is required"
		case "min":
			message = "minimum " + e.Param() + " characters"
		default:
			message = "is invalid"
		}

		result = append(result, ErrorItem{
			Key:     strings.ToLower(e.Field()),
			Message: message,
		})
	}

	return result
}
