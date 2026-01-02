package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductMaterial struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity  uint
	UnitPrice float64
	TotalCost float64

	BatchProductMaterialID uuid.UUID            `gorm:"type:uuid;"`
	BatchProductMaterial   BatchProductMaterial `gorm:"foreignKey:BatchProductMaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	MaterialID uuid.UUID `gorm:"type:uuid;"`
	Material   Material  `gorm:"foreignKey:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductMaterial) TableName() string {
	return "product_material"
}
