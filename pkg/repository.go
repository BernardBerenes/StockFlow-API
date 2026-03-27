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

func applyScopes(db *gorm.DB, scopes ...func(*gorm.DB) *gorm.DB) *gorm.DB {
	for _, scope := range scopes {
		db = db.Scopes(scope)
	}
	return db
}

func WithRelations(relations ...string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, relation := range relations {
			db = db.Preload(relation)
		}
		return db
	}
}

func (r *Repository[T]) List(entity *[]T, scopes ...func(db *gorm.DB) *gorm.DB) error {
	db := applyScopes(r.db, scopes...)

	return db.Find(entity).Error
}

func (r *Repository[T]) FindByUUID(entity *T, uuid uuid.UUID, scopes ...func(db *gorm.DB) *gorm.DB) error {
	db := applyScopes(r.db, scopes...)

	return db.Where("uuid = ?", uuid).Take(entity).Error
}

func (r *Repository[T]) Create(entity *T) error {
	return r.db.Create(entity).Error
}

func (r *Repository[T]) CreateBulk(entity []T) error {
	return r.db.Create(&entity).Error
}

func (r *Repository[T]) Update(entity *T) error {
	return r.db.Save(entity).Error
}

func (r *Repository[T]) Delete(entity *T) error {
	return r.db.Delete(entity).Error
}
