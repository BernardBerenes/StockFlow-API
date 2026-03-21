package presenter

import (
	"mime/multipart"

	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
)

type ProductResponse struct {
	UUID  uuid.UUID `json:"uuid"`
	Name  string    `json:"name"`
	Photo *string   `json:"photo"`
}

type CreateUpdateRequestProduct struct {
	Name  string                `form:"name" validate:"required,min=3"`
	Photo *multipart.FileHeader `form:"photo"`
}

func ToProductResponse(product entities.Product) ProductResponse {
	return ProductResponse{
		UUID:  product.UUID,
		Name:  product.Name,
		Photo: product.Photo,
	}
}
