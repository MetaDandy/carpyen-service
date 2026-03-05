package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BatchProduct struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   decimal.Decimal
	UnitPrice  decimal.Decimal
	TotalPrice decimal.Decimal
	Stock      decimal.Decimal

	ProductID uuid.UUID `gorm:"type:uuid;"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	WarehouseID uuid.UUID `gorm:"type:uuid;"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectBatchProducts []ProjectBatchProduct `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductDetails       []ProductDetails      `gorm:"foreignKey:BatchProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BatchProduct) TableName() string {
	return "batch_product_supplier"
}
