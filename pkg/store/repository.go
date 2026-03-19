package store

import (
	"github.com/BernardBerenes/stockflow-api/pkg"
	"github.com/BernardBerenes/stockflow-api/pkg/entities"
	"gorm.io/gorm"
)

type Repository struct {
	*pkg.Repository[entities.Store]
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Repository: pkg.NewRepository[entities.Store](db),
		db:         db,
	}
}
