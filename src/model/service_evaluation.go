package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ServiceEvaluation struct {
	ID                  uuid.UUID `gorm:"type:uuid;primaryKey;"`
	DesignQualification float32
	FabricationQuality  float32
	InstallationQuality float32
	OverallSatisfaction float32
	Comments            string

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ServiceEvaluation) TableName() string {
	return "service_evaluation"
}
