package pkg

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any](db *gorm.DB) *Repository[T] {
	return &Repository[T]{
		db: db,
	}
}

func (r *Repository[T]) List(entity *[]T, scope func(db *gorm.DB) *gorm.DB) error {
	db := r.db
	if scope != nil {
		db = db.Scopes(scope)
	}

	return db.Find(entity).Error
}

func (r *Repository[T]) FindByUUID(entity *T, uuid uuid.UUID) error {
	return r.db.Where("uuid = ?", uuid).Take(entity).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}
