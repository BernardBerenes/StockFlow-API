package transaction_detail

import (
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	*pkg.Repository[entities.TransactionDetail]
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Repository: pkg.NewRepository[entities.TransactionDetail](db),
		db:         db,
	}
}

func FindByTransaction(transactionUuid uuid.UUID) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("transaction_id = ?", transactionUuid)
	}
}
