package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Schedule struct {
	ID           uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title        string
	Description  string
	InitialDate  time.Time
	FinalDate    time.Time
	EstimateDays uint8
	State        enum.Status

	ProjectID uuid.UUID `gorm:"type:uuid;"`
	Project   Project   `gorm:"foreignKey:ProjectID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserAssignerID uuid.UUID `gorm:"type:uuid;"`
	UserAssigner   User      `gorm:"foreignKey:UserAssignerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	Tasks []Task `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Schedule) TableName() string {
	return "schedule"
}
