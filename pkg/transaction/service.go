package transaction

import (
	"time"

	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/BernardBerenes/stockflow-api/pkg/store"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	CreateTransaction(request *presenter.CreateUpdateRequestTransaction) error
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
