package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

type SubQuote struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Ambient      string
	UnitCost     decimal.Decimal
	UnitQuantity decimal.Decimal
	UnitType     enum.Unit
	TotalCost    decimal.Decimal
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
