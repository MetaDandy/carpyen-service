package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialSupplier struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity  uint
	UnitPrice float64
	TotalCost float64

	MaterialID uuid.UUID `gorm:"type:uuid;"`
	Material   Material  `gorm:"foreignKey:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	SupplierID uuid.UUID `gorm:"type:uuid;"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MaterialSupplier) TableName() string {
	return "material_supplier"
}
