package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type BatchMaterial struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity  decimal.Decimal
	UnitPrice decimal.Decimal
	TotalCost decimal.Decimal
	Stock     decimal.Decimal

	MaterialID uuid.UUID `gorm:"type:uuid;"`
	Material   Material  `gorm:"foreignKey:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	WarehouseID uuid.UUID `gorm:"type:uuid;"`
	Warehouse   Warehouse `gorm:"foreignKey:WarehouseID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectBatchMaterials []ProjectBatchMaterial `gorm:"foreignKey:BatchMaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MaterialDetails       []MaterialDetails      `gorm:"foreignKey:BatchMaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (BatchMaterial) TableName() string {
	return "batch_material_supplier"
}
