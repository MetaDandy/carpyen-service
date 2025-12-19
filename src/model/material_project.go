package model

import "github.com/google/uuid"

type MaterialProject struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;"`

	Quantity uint

	MaterialID uuid.UUID `gorm:"type:uuid;"`
	Material   Material  `gorm:"foreignKey:MaterialID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (MaterialProject) TableName() string {
	return "material_project"
}
