package model

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ClientObservation struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Commment string

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (ClientObservation) TableName() string {
	return "client_observation"
}
