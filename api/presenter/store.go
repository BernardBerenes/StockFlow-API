package presenter

import (
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
)

type StoreResponse struct {
	UUID uuid.UUID `json:"uuid"`
	Name string    `json:"name"`
}

type CreateUpdateRequestStore struct {
	Name string `json:"name" validate:"required,min=3"`
}

func ToStoreResponse(store entities.Store) StoreResponse {
	return StoreResponse{
		UUID: store.UUID,
		Name: store.Name,
	}
}

func ToStoreResponseList(stores []entities.Store) []StoreResponse {
	result := make([]StoreResponse, 0, len(stores))

	for _, s := range stores {
		result = append(result, ToStoreResponse(s))
	}

	return result
}
