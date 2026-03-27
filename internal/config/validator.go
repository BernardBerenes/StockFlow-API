package config

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func NewValidator() *validator.Validate {
	return validator.New()
}

func CheckDuplicateProduct(fl validator.FieldLevel) bool {
	items, ok := fl.Field().Interface().([]presenter.CreateUpdateTransactionDetail)
	if !ok {
		return false
	}

	seen := make(map[uuid.UUID]bool)
	for _, item := range items {
		if seen[item.ProductID] {
			return false
		}

		seen[item.ProductID] = true
	}

	return true
}
