package seed

import (
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Materials(db *gorm.DB) error {

	userEmails := map[string]uuid.UUID{}

	emails := []string{
		"admin@carpyen.com",
		"juan.chief@carpyen.com",
		"maria.installer@carpyen.com",
	}

	for _, email := range emails {
		var user model.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			return err // si no existe
		}
		userEmails[email] = user.ID
	}

	materials := []model.Material{
		//materials created by admin
		{
			ID:          uuid.New(),
			Name:        "Madera Roble",
			Type:        enum.Wood,
			UnitMeasure: enum.SquareMeter,
			UnitPrice:   decimal.NewFromFloat(85.50),
			UserID:      userEmails["admin@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Madera Pino",
			Type:        enum.Wood,
			UnitMeasure: enum.SquareMeter,
			UnitPrice:   decimal.NewFromFloat(65.00),
			UserID:      userEmails["admin@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Metal Aluminio",
			Type:        enum.Metal,
			UnitMeasure: enum.Kilogram,
			UnitPrice:   decimal.NewFromFloat(40.00),
			UserID:      userEmails["admin@carpyen.com"],
		},

		//materials created by chief
		{
			ID:          uuid.New(),
			Name:        "Plancha Metal",
			Type:        enum.Metal,
			UnitMeasure: enum.Meter,
			UnitPrice:   decimal.NewFromFloat(120.00),
			UserID:      userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Vidrio Templado",
			Type:        enum.Glass,
			UnitMeasure: enum.SquareMeter,
			UnitPrice:   decimal.NewFromFloat(150.75),
			UserID:      userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Pintura Blanca",
			Type:        enum.Paint,
			UnitMeasure: enum.Liter,
			UnitPrice:   decimal.NewFromFloat(30.00),
			UserID:      userEmails["juan.chief@carpyen.com"],
		},

		//materials created by installer
		{
			ID:          uuid.New(),
			Name:        "Plástico ABS",
			Type:        enum.Plastic,
			UnitMeasure: enum.Kilogram,
			UnitPrice:   decimal.NewFromFloat(22.40),
			UserID:      userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Mueble Prefabricado",
			Type:        enum.Forniture,
			UnitMeasure: enum.MeasureUnit,
			UnitPrice:   decimal.NewFromFloat(950.00),
			UserID:      userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Barniz",
			Type:        enum.Paint,
			UnitMeasure: enum.Liter,
			UnitPrice:   decimal.NewFromFloat(45.30),
			UserID:      userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:          uuid.New(),
			Name:        "Resina",
			Type:        enum.Other,
			UnitMeasure: enum.CubicLiter,
			UnitPrice:   decimal.NewFromFloat(60.00),
			UserID:      userEmails["maria.installer@carpyen.com"],
		},
	}

	var count int64
	db.Model(&model.Material{}).
		Where("name = ?", "Madera Roble").
		Count(&count)

	if count > 0 {
		return nil
	}

	// BULK INSERT
	return db.Create(&materials).Error
}
