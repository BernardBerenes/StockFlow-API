package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	UUID      uuid.UUID      `gorm:"column:uuid;type:uuid;primaryKey"`
	Name      string         `gorm:"column:name;type:varchar(255);not null"`
	Photo     string         `gorm:"column:photo;type:text"`
	CreatedAt time.Time      `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time      `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;type:timestamp;index"`
}

func (Product) TableName() string {
	return "products"
}
