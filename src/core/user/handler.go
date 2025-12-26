package user

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
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

type handler struct {
	service Service
}

func NewUserHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Post("/login", h.Login)

	users.Use(middleware.Jwt())

	users.Get("/me", h.GetProfile)
	users.Patch("/me", h.EditProfile)
	users.Get("/", h.FindAll)
	users.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin}), h.Create)
	users.Get("/:id", h.GetByID)
	users.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin}), h.Edit)
	users.Delete("/:id", h.SoftDelete)
}

func (h *handler) Login(c *fiber.Ctx) error {
	var input Login
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	token, err := h.service.Login(input)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"token": token,
	})
}

func (h *handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.service.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

func (h *handler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.service.GetByID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.Create(input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var input Update
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.Update(id, input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) EditProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var input UpdateProfile
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.UpdateProfile(userID, input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update profile")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not retrieve users")
	}

	return c.JSON(finded)
}

func (h *handler) SoftDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.SoftDelete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete user")
	}

	return c.SendStatus(fiber.StatusOK)
}
