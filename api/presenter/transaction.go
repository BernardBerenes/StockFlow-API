package presenter

import (
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
)

type TransactionResponse struct {
	UUID           uuid.UUID                `json:"uuid"`
	Store          StoreResponse            `json:"store"`
	Type           entities.TransactionType `json:"type"`
	Date           string                   `json:"date"`
	PaymentStatus  entities.PaymentStatus   `json:"payment_status"`
	DeliveryStatus entities.DeliveryStatus  `json:"delivery_status"`
}

type CreateUpdateRequestTransaction struct {
	StoreID        uuid.UUID                `json:"store_id" validate:"required"`
	Type           entities.TransactionType `json:"type" validate:"required,oneof=IN OUT"`
	Date           string                   `json:"date" validate:"required"`
	PaymentStatus  entities.PaymentStatus   `json:"payment_status" validate:"required,oneof=UNPAID PAID"`
	DeliveryStatus entities.DeliveryStatus  `json:"delivery_status" validate:"required,oneof=ON_DELIVERY DELIVERED"`
}

func ToTransactionResponse(transaction entities.Transaction) TransactionResponse {
	return TransactionResponse{
		UUID:           transaction.UUID,
		Store:          ToStoreResponse(transaction.Store),
		Type:           transaction.Type,
		Date:           transaction.Date.Format("2006-01-02"),
		PaymentStatus:  transaction.PaymentStatus,
		DeliveryStatus: transaction.DeliveryStatus,
	}
}
