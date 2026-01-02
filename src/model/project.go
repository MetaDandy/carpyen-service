package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Project struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name        string
	Description string
	Location    string
	State       enum.Status

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ClientID uuid.UUID `gorm:"type:uuid;"`
	Client   Client    `gorm:"foreignKey:ClientID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ClientObservations []ClientObservation `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ServiceEvaluations []ServiceEvaluation `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quotes             []Quote             `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedules          []Schedule          `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	ProjectBatchMaterialSuppliers []ProjectBatchMaterialSupplier `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectBatchProductMaterials  []ProjectBatchProductMaterial  `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ProjectBatchProductSuppliers  []ProjectBatchProductSupplier  `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Project) TableName() string {
	return "project"
}
