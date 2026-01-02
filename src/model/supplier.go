package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Supplier struct {
	ID      uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name    string
	Contact string
	Phone   string
	Email   string `gorm:"uniqueIndex"`
	Address string

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BatchMaterialSuppliers []BatchMaterialSupplier `gorm:"foreignKey:SupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	BatchProductSuppliers  []BatchProductSupplier  `gorm:"foreignKey:SupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Supplier) TableName() string {
	return "supplier"
}
