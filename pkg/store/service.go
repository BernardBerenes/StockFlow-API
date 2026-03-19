package store

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	List() ([]presenter.StoreResponse, error)
	Create(request *presenter.CreateRequest) error
}

type Service struct {
	repository *Repository
	validator  *validator.Validate
}

func NewService(repository *Repository, validator *validator.Validate) IService {
	return &Service{
		repository: repository,
		validator:  validator,
	}
}

func (s *Service) List() ([]presenter.StoreResponse, error) {
	var stores []entities.Store

	err := s.repository.List(&stores, nil)
	if err != nil {
		return nil, err
	}

	return presenter.ToStoreResponseList(stores), nil
}

func (s *Service) Create(request *presenter.CreateRequest) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	store := &entities.Store{
		Name: request.Name,
	}

	store.UUID, err = uuid.NewV7()
	if err != nil {
		return err
	}

	return s.repository.Create(store)
}
