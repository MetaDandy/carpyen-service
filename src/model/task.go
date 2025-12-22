package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Task struct {
	ID          uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Title       string
	Description string
	Status      enum.Status
	InitialHour time.Time
	FinalHour   time.Time

	ScheduleID uuid.UUID `gorm:"type:uuid;"`
	Schedule   Schedule  `gorm:"foreignKey:ScheduleID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserID uuid.UUID `gorm:"type:uuid;"`
	User   User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	UserAssignerID uuid.UUID `gorm:"type:uuid;"`
	UserAssigner   User      `gorm:"foreignKey:UserAssignerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (Task) TableName() string {
	return "task"
}
