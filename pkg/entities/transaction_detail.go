package entities

import (
	"time"

	"github.com/google/uuid"
)

type TransactionDetail struct {
	UUID          uuid.UUID  `gorm:"column:uuid;type:uuid;primaryKey"`
	TransactionID uuid.UUID  `gorm:"column:transaction_id;type:uuid;not null"`
	ProductID     uuid.UUID  `gorm:"column:product_id;type:uuid;not null"`
	Quantity      int64      `gorm:"column:quantity;type:int;not null"`
	Unit          string     `gorm:"column:unit;type:varchar(10);not null"`
	Price         float64    `gorm:"column:price;type:numeric(10,2);not null"`
	CreatedAt     *time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt     *time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	DeletedAt     *time.Time `gorm:"column:deleted_at;index"`

	Transaction Transaction `gorm:"foreignKey:TransactionID;references:UUID"`
	Product     Product     `gorm:"foreignKey:ProductID;references:UUID"`
}
