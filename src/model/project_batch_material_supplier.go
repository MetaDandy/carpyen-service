package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProjectBatchMaterialSupplier struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   decimal.Decimal
	UnitPrice  decimal.Decimal
	TotalPrice decimal.Decimal

	ProjectID uuid.UUID `gorm:"type:uuid;index;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BatchMaterialSupplierID uuid.UUID             `gorm:"type:uuid;index;"`
	BatchMaterialSupplier   BatchMaterialSupplier `gorm:"foreignKey:BatchMaterialSupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;index;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProjectBatchMaterialSupplier) TableName() string {
	return "project_batch_material_supplier"
}
