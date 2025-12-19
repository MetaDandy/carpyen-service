package user

import "github.com/gofiber/fiber/v2"

type UserHandler interface {
	RegisterRoutes(router fiber.Router)
	Login(c *fiber.Ctx) error
	Create(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	EditProfile(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
}
