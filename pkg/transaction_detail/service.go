package transaction_detail

import (
	"github.com/BernardBerenes/stockflow-api/api/presenter"
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/BernardBerenes/stockflow-api/pkg/transaction"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type IService interface {
	ListTransactionDetail(transactionUuid uuid.UUID) ([]presenter.TransactionDetailResponse, error)
	DetailTransactionDetail(transactionDetailUuid uuid.UUID) (*presenter.TransactionDetailResponse, error)
	CreateTransactionDetail(transactionUuid uuid.UUID, request presenter.CreateUpdateRequestTransactionDetail) error
}

type Service struct {
	repository            *Repository
	transactionRepository *transaction.Repository
	validator             *validator.Validate
}

func NewService(repository *Repository, transactionRepository *transaction.Repository, validator *validator.Validate) IService {
	return &Service{
		repository:            repository,
		transactionRepository: transactionRepository,
		validator:             validator,
	}
}

func (s *Service) ListTransactionDetail(transactionUuid uuid.UUID) ([]presenter.TransactionDetailResponse, error) {
	var transactionDetails []entities.TransactionDetail

	err := s.repository.List(&transactionDetails, FindByTransaction(transactionUuid), pkg.WithRelations("Product"))
	if err != nil {
		return nil, err
	}

	return presenter.MapToResponseList(transactionDetails, presenter.ToTransactionDetailResponse), nil
}

func (s *Service) DetailTransactionDetail(transactionDetailUuid uuid.UUID) (*presenter.TransactionDetailResponse, error) {
	var transactionDetail entities.TransactionDetail

	err := s.repository.FindByUUID(&transactionDetail, transactionDetailUuid, pkg.WithRelations("Product"))
	if err != nil {
		return nil, err
	}

	return new(presenter.ToTransactionDetailResponse(transactionDetail)), nil
}

func (s *Service) CreateTransactionDetail(transactionUuid uuid.UUID, request presenter.CreateUpdateRequestTransactionDetail) error {
	err := s.validator.Struct(request)
	if err != nil {
		return err
	}

	var existingTransaction entities.Transaction

	err = s.transactionRepository.FindByUUID(&existingTransaction, transactionUuid)
	if err != nil {
		return err
	}

	var transactionDetails []entities.TransactionDetail

	for _, req := range request.TransactionDetails {
		transactionDetail := entities.TransactionDetail{
			TransactionID: existingTransaction.UUID,
			ProductID:     req.ProductID,
			Quantity:      req.Quantity,
			Unit:          req.Unit,
			Price:         req.Price,
		}

		transactionDetail.UUID, err = uuid.NewV7()
		if err != nil {
			return err
		}

		transactionDetails = append(transactionDetails, transactionDetail)
	}

	return s.repository.CreateBulk(transactionDetails)
}
