package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectBatchProductMaterial struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   float64
	UnitPrice  float64
	TotalPrice float64

	ProjectID uuid.UUID `gorm:"type:uuid;index;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BatchProductMaterialID uuid.UUID            `gorm:"type:uuid;index;"`
	BatchProductMaterial   BatchProductMaterial `gorm:"foreignKey:BatchProductMaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;index;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProjectBatchProductMaterial) TableName() string {
	return "project_batch_product_material"
}
