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
	//los details cambiar los id al reves con purchase
	SupplierID uuid.UUID `gorm:"type:uuid;index;"`
	Supplier   Supplier  `gorm:"foreignKey:SupplierID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	MaterialDetails []MaterialDetails `gorm:"foreignKey:PurchaseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProductDetails  []ProductDetails  `gorm:"foreignKey:PurchaseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Purchase) TableName() string {
	return "purchase"
}
