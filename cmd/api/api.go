package api

import (
	"github.com/MetaDandy/carpyen-service/src"
	"github.com/gofiber/fiber/v2"
)

func SetupApi(app *fiber.App, c *src.Container) {
	v1 := app.Group("/api/v1")

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Aloha")
	})

	handlers := []func(fiber.Router){
		c.User.RegisterRoutes,
		c.Client.RegisterRoutes,
		c.Supplier.RegisterRoutes,
		c.Material.RegisterRoutes,
		c.Product.RegisterRoutes,
		c.BPM.RegisterRoutes,
		c.PM.RegisterRoutes,
	}

	for _, register := range handlers {
		register(v1)
	}
}
