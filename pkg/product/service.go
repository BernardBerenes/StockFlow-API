package product

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/BernardBerenes/stockflow-api/pkg/helper"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	ListPaginateProduct(paginateRequest *presenter.PaginateRequest) (*presenter.PaginateResponse[presenter.ProductResponse], error)
	ListProduct() ([]presenter.ProductResponse, error)
	CreateProduct(request *presenter.CreateUpdateRequestProduct) error
	UpdateProduct(uuid uuid.UUID, request *presenter.CreateUpdateRequestProduct) error
	DeleteProduct(uuid uuid.UUID) error
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

func (s *Service) ListPaginateProduct(paginateRequest *presenter.PaginateRequest) (*presenter.PaginateResponse[presenter.ProductResponse], error) {
	var products []entities.Product

	err := s.validator.Struct(paginateRequest)
	if err != nil {
		return nil, err
	}

	err = s.repository.ListPaginate(&products, paginateRequest)
	if err != nil {
		return nil, err
	}

	var total int64

	total, err = s.repository.Count()
	if err != nil {
		return nil, err
	}

	data, metadata := presenter.MapToResponseListPaginate(products, total, paginateRequest.Page, paginateRequest.Size, presenter.ToProductResponse)

	return &presenter.PaginateResponse[presenter.ProductResponse]{
		Data:             data,
		PaginateMetadata: metadata,
	}, nil
}

func (s *Service) ListProduct() ([]presenter.ProductResponse, error) {
	var products []entities.Product

	err := s.repository.List(&products)
	if err != nil {
		return nil, err
	}

	return presenter.MapToResponseList(products, presenter.ToProductResponse), nil
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
		photoPath, err = s.helpers.Minio.Insert(request.Photo, "products", product.UUID.String())
		if err != nil {
			return err
		}

		product.Photo = &photoPath
	}

	return s.repository.Create(product)
}

func (s *Service) UpdateProduct(uuid uuid.UUID, request *presenter.CreateUpdateRequestProduct) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	var product entities.Product

	err = s.repository.FindByUUID(&product, uuid)
	if err != nil {
		return err
	}

	product.Name = request.Name

	var photoPath string
	if request.Photo != nil {
		if product.Photo != nil {
			err = s.helpers.Minio.Delete(*product.Photo)
			if err != nil {
				return err
			}
		}

		photoPath, err = s.helpers.Minio.Insert(request.Photo, "products", product.UUID.String())
		if err != nil {
			return err
		}

		product.Photo = &photoPath
	}

	return s.repository.Update(&product)
}

func (s *Service) DeleteProduct(uuid uuid.UUID) error {
	var product entities.Product

	err := s.repository.FindByUUID(&product, uuid)
	if err != nil {
		return err
	}

	return s.repository.Delete(&product)
}
