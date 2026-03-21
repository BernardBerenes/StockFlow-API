package transaction

import (
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository struct {
	*pkg.Repository[entities.Transaction]
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Repository: pkg.NewRepository[entities.Transaction](db),
		db:         db,
	}
}
