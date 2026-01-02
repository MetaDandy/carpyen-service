package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProjectBatchProductSupplier struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   uint
	UnitPrice  float64
	TotalPrice float64

	ProjectID uuid.UUID `gorm:"type:uuid;index;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BatchProductSupplierID uuid.UUID            `gorm:"type:uuid;index;"`
	BatchProductSupplier   BatchProductSupplier `gorm:"foreignKey:BatchProductSupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;index;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProjectBatchProductSupplier) TableName() string {
	return "project_batch_product_supplier"
}
