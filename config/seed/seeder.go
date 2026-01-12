package seed

import (
	"log"

	"gorm.io/gorm"
)

func Seeder(db *gorm.DB) {

	if err := Users(db); err != nil {
		log.Fatalf("Error al seedear usuarios: %v", err)
	}

	if err := Materials(db); err != nil {
		log.Fatalf("Error al seedear materials: %v", err)
	}

	if err := Products(db); err != nil {
		log.Fatalf("Error al seedear products: %v", err)
	}

	if err := Suppliers(db); err != nil {
		log.Fatalf("Error al seedear suppliers: %v", err)
	}
}
