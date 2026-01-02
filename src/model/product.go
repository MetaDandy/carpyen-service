package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID        uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name      string
	Type      enum.Product
	UnitPrice float64

	BatchProductSuppliers []BatchProductSupplier `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BatchProductMaterials []BatchProductMaterial `gorm:"foreignKey:ProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Product) TableName() string {
	return "product"
}
