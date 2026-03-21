package presenter

import (
	"errors"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type Success[T any] struct {
	Message string `json:"message"`
	Data    T      `json:"data,omitempty"`
}

type Error struct {
	Message string      `json:"message,omitempty"`
	Errors  []ErrorItem `json:"errors,omitempty"`
}

type ErrorItem struct {
	Field   string `json:"field"`
	Message string `json:"message"`
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

		field := e.Field()

		switch e.Tag() {
		case "required":
			message = field + " is required"
		case "min":
			message = field + " minimum is " + e.Param() + " characters"
		default:
			message = field + "is invalid"
		}

		result = append(result, ErrorItem{
			Field:   strings.ToLower(field),
			Message: message,
		})
	}

	return result
}

func MapToResponseList[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, 0, len(items))

	for _, item := range items {
		result = append(result, mapper(item))
	}

	return result
}
