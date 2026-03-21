package store

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	ListStore() ([]presenter.StoreResponse, error)
	CreateStore(request *presenter.CreateUpdateRequestStore) error
	UpdateStore(uuid uuid.UUID, request *presenter.CreateUpdateRequestStore) error
	DeleteStore(uuid uuid.UUID) error
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

func (s *Service) ListStore() ([]presenter.StoreResponse, error) {
	var stores []entities.Store

	err := s.repository.List(&stores, nil)
	if err != nil {
		return nil, err
	}

	return presenter.ToStoreResponseList(stores), nil
}

func (s *Service) CreateStore(request *presenter.CreateUpdateRequestStore) error {
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

func (s *Service) UpdateStore(uuid uuid.UUID, request *presenter.CreateUpdateRequestStore) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	var store entities.Store

	err = s.repository.FindByUUID(&store, uuid)
	if err != nil {
		return err
	}

	store.Name = request.Name

	return s.repository.Update(&store)
}

func (s *Service) DeleteStore(uuid uuid.UUID) error {
	var store entities.Store

	err := s.repository.FindByUUID(&store, uuid)
	if err != nil {
		return err
	}

	return s.repository.Delete(&store)
}
