package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type SubQuote struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Ambient      string
	UnitCost     float64
	UnitQuantity float64
	UnitType     enum.Unit
	TotalCost    float64
	Status       enum.Status
	Description  string

	QuoteID uuid.UUID `gorm:"type:uuid;"`
	Quote   Quote     `gorm:"foreignKey:QuoteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (SubQuote) TableName() string {
	return "sub_quote"
}
