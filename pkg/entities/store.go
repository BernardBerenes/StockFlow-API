package entities

import (
	"time"

	"github.com/google/uuid"
)

type Store struct {
	UUID      uuid.UUID  `gorm:"column:uuid;type:uuid;primaryKey"`
	Name      string     `gorm:"column:name;type:varchar(255);not null"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt *time.Time `gorm:"column:deleted_at;type:timestamp;index"`
}

func (Store) TableName() string {
	return "stores"
}
