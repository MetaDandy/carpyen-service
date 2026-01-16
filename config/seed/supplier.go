package seed

import (
	"github.com/MetaDandy/carpyen-service/src/model"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func Suppliers(db *gorm.DB) error {

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

	suppliers := []model.Supplier{
		//suppliers created by admin
		{
			ID:      uuid.New(),
			Name:    "Juanito Miles",
			Contact: "cosas",
			Phone:   "123456789",
			Email:   "juanitosupplier@gmail.com",
			Address: "Calle Falsa 123",
			UserID:  userEmails["admin@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Mario Lopez",
			Contact: "cosas de mario",
			Phone:   "77789456",
			Email:   "marioosupplier@gmail.com",
			Address: "Calle Falsa 124",
			UserID:  userEmails["admin@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Narancio Lorax",
			Contact: "cosas de Lorax",
			Phone:   "55467890",
			Email:   "Elloraxsupplier@gmail.com",
			Address: "EL bosque",
			UserID:  userEmails["admin@carpyen.com"],
		},

		//suppliers created by chief
		{
			ID:      uuid.New(),
			Name:    "Kurt Cobain",
			Contact: "cosas de kurt",
			Phone:   "489775413",
			Email:   "kurtcobainsupplier@gmail.com",
			Address: "Nirvana St",
			UserID:  userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Karla Smith",
			Contact: "cosas de karla",
			Phone:   "64547180",
			Email:   "karlasupplier@gmail.com",
			Address: "Calle prueba 456",
			UserID:  userEmails["juan.chief@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "William Mamani",
			Contact: "cosas de William",
			Phone:   "745896412",
			Email:   "williammamanisupplier@gmail.com",
			Address: "Calle prueba 789",
			UserID:  userEmails["juan.chief@carpyen.com"],
		},

		//suppliers created by installer
		{
			ID:      uuid.New(),
			Name:    "Lisa Ono",
			Contact: "cosas de Lisa",
			Phone:   "745896412",
			Email:   "lisagarciasupplier@gmail.com",
			Address: "Calle cariñito 789",
			UserID:  userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Charly García",
			Contact: "cosas de Charly",
			Phone:   "33124589",
			Email:   "charlygarciasupplier@gmail.com",
			Address: "calle acanconcagua",
			UserID:  userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Axel Rosas",
			Contact: "cosas de Axel",
			Phone:   "15454987",
			Email:   "axelrosasupplier@gmail.com",
			Address: "calle de las rosas 123",
			UserID:  userEmails["maria.installer@carpyen.com"],
		},
		{
			ID:      uuid.New(),
			Name:    "Juan Contreras",
			Contact: "cosas de Juan",
			Phone:   "15454987",
			Email:   "juancontrerasupplier@gmail.com",
			Address: "calle perico de los palotes 123",
			UserID:  userEmails["maria.installer@carpyen.com"],
		},
	}

	var count int64
	db.Model(&model.Supplier{}).
		Where("name = ?", "Juanito Miles").
		Count(&count)

	if count > 0 {
		return nil
	}

	// BULK INSERT
	return db.Create(&suppliers).Error
}
