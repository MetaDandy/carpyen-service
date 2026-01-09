package config

import (
	"log"
	"os"
	"time"

	"github.com/MetaDandy/carpyen-service/config/seed"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	Port string
)

func Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	Port = os.Getenv("PORT")
	if Port == "" {
		Port = "8001"
	}

	maxRetries := 10
	for i := range maxRetries {
		dns := os.Getenv("DATABASE_URL")
		if dns == "" {
			log.Fatal("DATABASE_URL not set in .env file")
		}

		log.Println("Running migrations...")
		Migrate(dns)
		log.Println("Migrations completed")

		log.Println("Connecting to database...")
		// Agregar parámetros para Neon: timeouts y connection pooling
		dsn := dns + "?sslmode=require&connect_timeout=30"
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			// Configurar connection pool
			sqlDB, _ := DB.DB()
			sqlDB.SetMaxIdleConns(5)
			sqlDB.SetMaxOpenConns(20)
			sqlDB.SetConnMaxLifetime(time.Hour)
			log.Printf("Database connected successfully after %d attempt(s)", i+1)
			log.Println("Starting database seeding...")
			seed.Seeder(DB)
			log.Println("Database seeding completed")
			return
		}

		log.Printf("Failed to connect to database, retrying (%d/%d): %v", i+1, maxRetries, err)
		time.Sleep(2 * time.Second)
	}

	log.Fatalf("Error connecting to database after %d retries", maxRetries)

	log.Printf("Database connected")
}
