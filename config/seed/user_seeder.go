package seed

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {

	users := []model.User{
		{
			ID:       uuid.New(),
			Name:     "Admin System",
			Email:    "admin@carpyen.com",
			Phone:    "+591 70123456",
			Address:  "Avenida Equipetrol 100, Santa Cruz, Bolivia",
			Password: "admin123",
			Role:     enum.RoleAdmin,
		},
		{
			ID:       uuid.New(),
			Name:     "Carlos Designer",
			Email:    "carlos.designer@carpyen.com",
			Phone:    "+591 71234567",
			Address:  "Calle Arenales 234, Santa Cruz, Bolivia",
			Password: "designer123",
			Role:     enum.RoleDesigner,
		},
		{
			ID:       uuid.New(),
			Name:     "Sofia Diseñadora",
			Email:    "sofia.designer@carpyen.com",
			Phone:    "+591 72345678",
			Address:  "Avenida San Martín 456, Santa Cruz, Bolivia",
			Password: "designer123",
			Role:     enum.RoleDesigner,
		},
		// SELLER - 2 usuarios
		{
			ID:       uuid.New(),
			Name:     "Miguel Vendedor",
			Email:    "miguel.seller@carpyen.com",
			Phone:    "+591 73456789",
			Address:  "Calle El Barrio 567, Santa Cruz, Bolivia",
			Password: "seller123",
			Role:     enum.RoleSeller,
		},
		{
			ID:       uuid.New(),
			Name:     "Andrea Ventas",
			Email:    "andrea.seller@carpyen.com",
			Phone:    "+591 74567890",
			Address:  "Avenida Casariego 678, Santa Cruz, Bolivia",
			Password: "seller123",
			Role:     enum.RoleSeller,
		},
		// CHIEF_INSTALLER - 2 usuarios
		{
			ID:       uuid.New(),
			Name:     "Juan Jefe Instalador",
			Email:    "juan.chief@carpyen.com",
			Phone:    "+591 75678901",
			Address:  "Calle Monseñor Rivero 789, Santa Cruz, Bolivia",
			Password: "chief123",
			Role:     enum.RoleChiefInstaller,
		},
		{
			ID:       uuid.New(),
			Name:     "Roberto Supervisor",
			Email:    "roberto.chief@carpyen.com",
			Phone:    "+591 76789012",
			Address:  "Avenida Germán Busch 890, Santa Cruz, Bolivia",
			Password: "chief123",
			Role:     enum.RoleChiefInstaller,
		},
		// INSTALLER - 2 usuarios
		{
			ID:       uuid.New(),
			Name:     "María Instaladora",
			Email:    "maria.installer@carpyen.com",
			Phone:    "+591 77890123",
			Address:  "Calle Beni 901, Santa Cruz, Bolivia",
			Password: "installer123",
			Role:     enum.RoleInstaller,
		},
		{
			ID:       uuid.New(),
			Name:     "Pedro Técnico",
			Email:    "pedro.installer@carpyen.com",
			Phone:    "+591 78901234",
			Address:  "Avenida Cristo Redentor 1012, Santa Cruz, Bolivia",
			Password: "installer123",
			Role:     enum.RoleInstaller,
		},
	}

	for _, user := range users {
		var existingUser model.User
		if err := db.Where("email = ?", user.Email).First(&existingUser).Error; err == nil {
			continue
		}

		hashed, err := helper.HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashed

		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}
