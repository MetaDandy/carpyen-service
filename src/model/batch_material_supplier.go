package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BatchMaterialSupplier struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity  uint
	UnitPrice decimal.Decimal
	TotalCost decimal.Decimal
	Stock     decimal.Decimal

	MaterialID uuid.UUID `gorm:"type:uuid;"`
	Material   Material  `gorm:"foreignKey:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	SupplierID uuid.UUID `gorm:"type:uuid;"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectBatchMaterialSuppliers []ProjectBatchMaterialSupplier `gorm:"foreignKey:BatchMaterialSupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BatchMaterialSupplier) TableName() string {
	return "batch_material_supplier"
}
