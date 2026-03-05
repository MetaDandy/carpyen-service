package model

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type ProjectBatchProduct struct {
	ID         uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Quantity   uint
	UnitPrice  decimal.Decimal
	TotalPrice decimal.Decimal

	ProjectID uuid.UUID `gorm:"type:uuid;index;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	BatchProductID uuid.UUID    `gorm:"type:uuid;index;"`
	BatchProduct   BatchProduct `gorm:"foreignKey:BatchProductID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;index;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ProjectBatchProduct) TableName() string {
	return "project_batch_product"
}
