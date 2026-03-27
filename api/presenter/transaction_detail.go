package presenter

import (
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
)

type TransactionDetailResponse struct {
	UUID     uuid.UUID                      `json:"uuid"`
	Product  ProductResponse                `json:"product"`
	Quantity int                            `json:"quantity"`
	Unit     entities.TransactionDetailUnit `json:"unit"`
	Price    float64                        `json:"price"`
}

type CreateUpdateTransactionDetail struct {
	ProductID uuid.UUID                      `json:"product_id" validate:"required"`
	Quantity  int                            `json:"quantity" validate:"required"`
	Unit      entities.TransactionDetailUnit `json:"unit" validate:"required"`
	Price     float64                        `json:"price" validate:"required"`
}

type CreateUpdateRequestTransactionDetail struct {
	TransactionDetails []CreateUpdateTransactionDetail `json:"transaction_details" validate:"required,min=1,check_duplicate_products,dive"`
}

func ToTransactionDetailResponse(transactionDetail entities.TransactionDetail) TransactionDetailResponse {
	return TransactionDetailResponse{
		UUID:     transactionDetail.UUID,
		Product:  ToProductResponse(transactionDetail.Product),
		Quantity: transactionDetail.Quantity,
		Unit:     transactionDetail.Unit,
		Price:    transactionDetail.Price,
	}
}
