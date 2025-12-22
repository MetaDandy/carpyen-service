package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Material struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name         string
	Type         enum.Material
	UniteMessure string
	UnitPrice    float64
	Stock        uint

	MaterialProjects  []MaterialProject  `gorm:"foreignKey:MaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MaterialSuppliers []MaterialSupplier `gorm:"foreignKey:MaterialID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Material) TableName() string {
	return "material"
}
