package model

import (
	"time"

	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;primaryKey;"`
	Name     string
	Email    string `gorm:"uniqueIndex"`
	Phone    string
	Address  string
	Password string
	Role     enum.Role

	Projects  []Project  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Quotes    []Quote    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Designs   []Design   `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Schedules []Schedule `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Task      []Task     `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	DesignAssigneds   []Design           `gorm:"foreignKey:UserAssignerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MaterialProjects  []MaterialProject  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	MaterialSuppliers []MaterialSupplier `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ScheduleAssigneds []Schedule         `gorm:"foreignKey:UserAssignerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TaskAssigneds     []Task             `gorm:"foreignKey:UserAssignerID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Clients           []Client           `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (User) TableName() string {
	return "users"
}
