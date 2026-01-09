package main

import (
	"log"
	"os"

	"github.com/MetaDandy/carpyen-service/cmd/api"
	"github.com/MetaDandy/carpyen-service/config"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	config.Load()

	app := fiber.New()
	app.Use(middleware.Logger())

	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("ALLOW_ORIGINS"),
		AllowMethods: "GET,POST,PATCH,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	log.Println("Setting up container...")
	c := src.SetupContainer()
	log.Println("Setting up API routes...")
	api.SetupApi(app, c)

	log.Println("Starting server on port " + config.Port)
	app.Listen("0.0.0.0:" + config.Port)
}
