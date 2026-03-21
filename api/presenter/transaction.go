package presenter

import (
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
)

type CreateUpdateRequestTransaction struct {
	StoreID        uuid.UUID                `json:"store_id" validate:"required"`
	Type           entities.TransactionType `json:"type" validate:"required"`
	Date           string                   `json:"date" validate:"required"`
	PaymentStatus  entities.PaymentStatus   `json:"payment_status" validate:"required"`
	DeliveryStatus entities.DeliveryStatus  `json:"delivery_status" validate:"required"`
}
