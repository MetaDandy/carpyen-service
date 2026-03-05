package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MaterialDetails struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`

	BatchMaterialID uuid.UUID     `gorm:"type:uuid;index;"`
	BatchMaterial   BatchMaterial `gorm:"foreignKey:BatchMaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Purchase []Purchase `gorm:"foreignKey:MaterialDetailsID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (MaterialDetails) TableName() string {
	return "material_details"
}
