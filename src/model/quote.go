package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Quote struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	TotalCost    float64
	Status       enum.Status
	Comments     string
	ValidDays    int
	DeliveryDays int

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Designs   []Design   `gorm:"foreignKey:QuoteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	SubQuotes []SubQuote `gorm:"foreignKey:QuoteID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Quote) TableName() string {
	return "quote"
}
