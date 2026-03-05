package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Purchase struct {
	ID            uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Date          time.Time
	ReceiptNumber string
	//gloss

	MaterialDetailsID uuid.UUID       `gorm:"type:uuid;index;"`
	MaterialDetails   MaterialDetails `gorm:"foreignKey:MaterialDetailsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProductDetailsID uuid.UUID      `gorm:"type:uuid;index;"`
	ProductDetails   ProductDetails `gorm:"foreignKey:ProductDetailsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	SupplierID uuid.UUID `gorm:"type:uuid;index;"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Purchase) TableName() string {
	return "purchase"
}
