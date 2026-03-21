package product

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/BernardBerenes/stockflow-api/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	CreateProduct(request *presenter.CreateUpdateRequestProduct) error
}

type Service struct {
	repository *Repository
	helpers    *helper.Helper
	validator  *validator.Validate
}

func NewService(repository *Repository, helpers *helper.Helper, validator *validator.Validate) IService {
	return &Service{
		repository: repository,
		helpers:    helpers,
		validator:  validator,
	}
}

func (s *Service) CreateProduct(request *presenter.CreateUpdateRequestProduct) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	product := &entities.Product{
		Name: request.Name,
	}

	product.UUID, err = uuid.NewV7()
	if err != nil {
		return err
	}

	var photoPath string
	if request.Photo != nil {
		photoPath, err = s.helpers.Minio.Insert(request.Photo, product.UUID.String())
		if err != nil {
			return err
		}

		product.Photo = &photoPath
	}

	return s.repository.Create(product)
}
