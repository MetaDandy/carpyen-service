package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProductDetails struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`

	BatchProductID uuid.UUID    `gorm:"type:uuid;index;"`
	BatchProduct   BatchProduct `gorm:"foreignKey:BatchProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	PurchaseID uuid.UUID `gorm:"type:uuid;index;"`
	Purchase   Purchase  `gorm:"foreignKey:PurchaseID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProductDetails) TableName() string {
	return "product_details"
}
