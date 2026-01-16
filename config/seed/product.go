package seed

import (
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func Products(db *gorm.DB) error {

	userEmails := map[string]uuid.UUID{}

	emails := []string{
		"admin@carpyen.com",
		"juan.chief@carpyen.com",
		"maria.installer@carpyen.com",
	}

	for _, email := range emails {
		var user model.User
		if err := db.Where("email = ?", email).First(&user).Error; err != nil {
			return err
		}
		userEmails[email] = user.ID
	}

	products := []model.Product{
		// productos creados por admin
		{
			ID:        uuid.New(),
			Name:      "Silla Moderna",
			Type:      enum.Chair,
			UnitPrice: decimal.NewFromFloat(320.00),
			UserID:    userEmails["admin@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Mesa Comedor",
			Type:      enum.Table,
			UnitPrice: decimal.NewFromFloat(950.00),
			UserID:    userEmails["admin@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Sofá 3 Plazas",
			Type:      enum.Sofa,
			UnitPrice: decimal.NewFromFloat(2100.00),
			UserID:    userEmails["admin@carpyen.com"],
		},

		// productos creados por chief installer
		{
			ID:        uuid.New(),
			Name:      "Cama Queen",
			Type:      enum.Bed,
			UnitPrice: decimal.NewFromFloat(1800.00),
			UserID:    userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Gabinete Cocina",
			Type:      enum.Cabinet,
			UnitPrice: decimal.NewFromFloat(1250.00),
			UserID:    userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Escritorio Oficina",
			Type:      enum.Desk,
			UnitPrice: decimal.NewFromFloat(780.00),
			UserID:    userEmails["juan.chief@carpyen.com"],
		},

		// productos creados por installer
		{
			ID:        uuid.New(),
			Name:      "Estante Madera",
			Type:      enum.Shelf,
			UnitPrice: decimal.NewFromFloat(450.00),
			UserID:    userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Lámpara de Pie",
			Type:      enum.Lamp,
			UnitPrice: decimal.NewFromFloat(260.00),
			UserID:    userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Alfombra Decorativa",
			Type:      enum.Rug,
			UnitPrice: decimal.NewFromFloat(390.00),
			UserID:    userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:        uuid.New(),
			Name:      "Cortina Blackout",
			Type:      enum.Curtain,
			UnitPrice: decimal.NewFromFloat(310.00),
			UserID:    userEmails["maria.installer@carpyen.com"],
		},
	}

	var count int64
	db.Model(&model.Product{}).
		Where("name = ?", "Silla Moderna").
		Count(&count)

	if count > 0 {
		return nil
	}

	// Bulk insert
	return db.Create(&products).Error
}
