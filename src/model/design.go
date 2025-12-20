package model

import (
	"time"

	"github.com/MetaDandy/go-fiber-skeleton/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Design struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey;"`
	URLRender       string
	IluminatedPlane string
	State           enum.Status
	Comments        string

	QuoteID uuid.UUID `gorm:"type:uuid;"`
	Quote   Quote     `gorm:"foreignKey:QuoteID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserAssignerID uuid.UUID `gorm:"type:uuid;"`
	UserAssigner   User      `gorm:"foreignKey:UserAssignerID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Design) TableName() string {
	return "design"
}
