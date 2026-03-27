package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionDetailUnit string

const (
	Piece  TransactionDetailUnit = "PIECE"
	Dozen  TransactionDetailUnit = "DOZEN"
	Box    TransactionDetailUnit = "BOX"
	Carton TransactionDetailUnit = "CARTON"
)

type TransactionDetail struct {
	UUID          uuid.UUID             `gorm:"column:uuid;type:uuid;primaryKey"`
	TransactionID uuid.UUID             `gorm:"column:transaction_id;type:uuid;not null"`
	ProductID     uuid.UUID             `gorm:"column:product_id;type:uuid;not null"`
	Quantity      int                   `gorm:"column:quantity;type:int;not null"`
	Unit          TransactionDetailUnit `gorm:"column:unit;type:transaction_detail_unit;not null"`
	Price         float64               `gorm:"column:price;type:numeric(10,2);not null"`
	CreatedAt     time.Time             `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt     time.Time             `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt     gorm.DeletedAt        `gorm:"column:deleted_at;type:timestamp;index"`

	Transaction Transaction `gorm:"foreignKey:TransactionID;references:UUID"`
	Product     Product     `gorm:"foreignKey:ProductID;references:UUID"`
}

func (TransactionDetail) TableName() string {
	return "transaction_details"
}
