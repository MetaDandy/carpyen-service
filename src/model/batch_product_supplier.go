package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BatchProductSupplier struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   uint
	UnitPrice  float64
	TotalPrice float64

	ProductID uuid.UUID `gorm:"type:uuid;"`
	Product   Product   `gorm:"foreignKey:ProductID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	SupplierID uuid.UUID `gorm:"type:uuid;"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectBatchProductSuppliers []ProjectBatchProductSupplier `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BatchProductSupplier) TableName() string {
	return "batch_product_supplier"
}
