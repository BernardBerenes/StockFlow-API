package product

import (
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository struct {
	*pkg.Repository[entities.Product]
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Repository: pkg.NewRepository[entities.Product](db),
		db:         db,
	}
}
