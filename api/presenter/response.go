package presenter

import (
	"errors"
	"math"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
)

type PaginateResponse[T any] struct {
	Data             []T
	PaginateMetadata *PaginateMetadata
}

type PaginateMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	Total     int64 `json:"total"`
	TotalPage int   `json:"total_page"`
}

type Success[T any] struct {
	Message          string            `json:"message"`
	Data             T                 `json:"data,omitempty"`
	PaginateMetadata *PaginateMetadata `json:"metadata,omitempty"`
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

func SuccessResponsePaginate[T any](ctx fiber.Ctx, status int, message string, data T, paginateMetadata *PaginateMetadata) error {
	return ctx.Status(status).JSON(Success[T]{
		Message:          message,
		Data:             data,
		PaginateMetadata: paginateMetadata,
	})
}

func ErrorResponse(ctx fiber.Ctx, status int, message string, errors []ErrorItem) error {
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
			message = field + " is invalid"
		}

		result = append(result, ErrorItem{
			Field:   strings.ToLower(field),
			Message: message,
		})
	}

	return result
}

func MapToResponseListPaginate[T any, R any](items []T, total int64, page, size int, mapper func(T) R) ([]R, *PaginateMetadata) {
	data := MapToResponseList(items, mapper)

	totalPage := int(math.Ceil(float64(total) / float64(size)))

	metadata := &PaginateMetadata{
		Page:      page,
		Size:      size,
		Total:     total,
		TotalPage: totalPage,
	}

	return data, metadata
}

func MapToResponseList[T any, R any](items []T, mapper func(T) R) []R {
	result := make([]R, 0, len(items))

	for _, item := range items {
		result = append(result, mapper(item))
	}

	return result
}
