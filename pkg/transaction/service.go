package transaction

import (
	"time"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	ListPaginateTransaction(paginateRequest *presenter.PaginateRequest) (*presenter.PaginateResponse[presenter.TransactionResponse], error)
	ListTransaction() ([]presenter.TransactionResponse, error)
	CreateTransaction(request *presenter.CreateUpdateRequestTransaction) error
	UpdateTransaction(uuid uuid.UUID, request *presenter.CreateUpdateRequestTransaction) error
	DeleteTransaction(uuid uuid.UUID) error
}

type Service struct {
	repository      *Repository
	storeRepository *store.Repository
	validator       *validator.Validate
}

func NewService(repository *Repository, storeRepository *store.Repository, validator *validator.Validate) IService {
	return &Service{
		repository:      repository,
		storeRepository: storeRepository,
		validator:       validator,
	}
}

func (s *Service) ListPaginateTransaction(paginateRequest *presenter.PaginateRequest) (*presenter.PaginateResponse[presenter.TransactionResponse], error) {
	var transactions []entities.Transaction

	err := s.validator.Struct(paginateRequest)
	if err != nil {
		return nil, err
	}

	err = s.repository.ListPaginate(&transactions, paginateRequest, pkg.WithRelations("Store"))
	if err != nil {
		return nil, err
	}

	var total int64

	total, err = s.repository.Count()
	if err != nil {
		return nil, err
	}

	data, metadata := presenter.MapToResponseListPaginate(transactions, total, paginateRequest.Page, paginateRequest.Size, presenter.ToTransactionResponse)

	return &presenter.PaginateResponse[presenter.TransactionResponse]{
		Data:             data,
		PaginateMetadata: metadata,
	}, nil
}

func (s *Service) ListTransaction() ([]presenter.TransactionResponse, error) {
	var transactions []entities.Transaction

	err := s.repository.List(&transactions, pkg.WithRelations("Store"))
	if err != nil {
		return nil, err
	}

	return presenter.MapToResponseList(transactions, presenter.ToTransactionResponse), nil
}

func (s *Service) CreateTransaction(request *presenter.CreateUpdateRequestTransaction) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	var existingStore entities.Store

	err = s.storeRepository.FindByUUID(&existingStore, request.StoreID)
	if err != nil {
		return err
	}

	var parsedDate time.Time

	parsedDate, err = time.Parse("2006-01-02", request.Date)
	if err != nil {
		return err
	}

	transaction := &entities.Transaction{
		StoreID:        existingStore.UUID,
		Type:           request.Type,
		Date:           parsedDate,
		PaymentStatus:  request.PaymentStatus,
		DeliveryStatus: request.DeliveryStatus,
	}

	transaction.UUID, err = uuid.NewV7()
	if err != nil {
		return err
	}

	return s.repository.Create(transaction)
}

func (s *Service) UpdateTransaction(uuid uuid.UUID, request *presenter.CreateUpdateRequestTransaction) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	var transaction entities.Transaction

	err = s.repository.FindByUUID(&transaction, uuid)
	if err != nil {
		return err
	}

	var existingStore entities.Store

	err = s.storeRepository.FindByUUID(&existingStore, request.StoreID)
	if err != nil {
		return err
	}

	var parsedDate time.Time

	parsedDate, err = time.Parse("2006-01-02", request.Date)
	if err != nil {
		return err
	}

	transaction.StoreID = existingStore.UUID
	transaction.Type = request.Type
	transaction.Date = parsedDate
	transaction.PaymentStatus = request.PaymentStatus
	transaction.DeliveryStatus = request.DeliveryStatus

	return s.repository.Update(&transaction)
}

func (s *Service) DeleteTransaction(uuid uuid.UUID) error {
	var transaction entities.Transaction

	err := s.repository.FindByUUID(&transaction, uuid)
	if err != nil {
		return err
	}

	return s.repository.Delete(&transaction)
}
