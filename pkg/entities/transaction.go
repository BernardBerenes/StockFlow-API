package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type TransactionType string
type PaymentStatus string
type DeliveryStatus string

const (
	In  TransactionType = "IN"
	Out TransactionType = "OUT"

	Unpaid PaymentStatus = "UNPAID"
	Paid   PaymentStatus = "PAID"

	OnDelivery DeliveryStatus = "ON_DELIVERY"
	Delivered  DeliveryStatus = "DELIVERED"
)

type Transaction struct {
	UUID           uuid.UUID       `gorm:"column:uuid;type:uuid;primaryKey"`
	StoreID        uuid.UUID       `gorm:"column:store_id;type:uuid;not null"`
	Type           TransactionType `gorm:"column:type;type:transaction_type;not null"`
	Date           time.Time       `gorm:"column:date;type:date;not null"`
	PaymentStatus  PaymentStatus   `gorm:"column:payment_status;type:payment_status;not null"`
	DeliveryStatus DeliveryStatus  `gorm:"column:delivery_status;type:delivery_status;not null"`
	CreatedAt      time.Time       `gorm:"column:created_at;type:timestamp;autoCreateTime"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;type:timestamp;autoUpdateTime"`
	DeletedAt      gorm.DeletedAt  `gorm:"column:deleted_at;type:timestamp;index"`

	Store Store `gorm:"foreignKey:StoreID;references:UUID"`
}

func (Transaction) TableName() string {
	return "transactions"
}
