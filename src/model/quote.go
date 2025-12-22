package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quote struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	MeterType       enum.Meter
	MeterCost       float64
	MeterQuantity   float64
	FurnitureNumber uint
	FurnitureCost   float64
	TotalCost       float64
	State           enum.Status
	Comments        string

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Designs []Design `gorm:"foreignKey:QuoteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Quote) TableName() string {
	return "quote"
}
