package client

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	Create(c *fiber.Ctx) error
	Login(c *fiber.Ctx) error
	Edit(c *fiber.Ctx) error
	GetProfile(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	GetByID(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
	RegisterRoutes(router fiber.Router)
}

type handler struct {
	service Service
}

func NewClientHandler(service Service) Handler {
	return &handler{
		service: service,
	}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	users := router.Group("/clients")
	users.Post("/login", h.Login)

	users.Use(middleware.Jwt())

	users.Get("/me", h.GetProfile)
	users.Patch("/me", h.EditProfile)
	users.Get("/", h.FindAll)
	users.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleSeller}), h.Create)
	users.Get("/:id", h.GetByID)
	users.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleSeller}), h.Edit)
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
		return fiber.NewError(fiber.StatusNotFound, "Client not found")
	}

	return c.JSON(user)
}

func (h *handler) GetProfile(c *fiber.Ctx) error {
	clientID := c.Locals("user_id").(string)

	user, err := h.service.GetByID(clientID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Client not found")
	}

	return c.JSON(user)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	userID := c.Locals("user_id").(string)

	err := h.service.Create(input, userID)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create client")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var input Update
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}
	role := c.Locals("role").(string)

	if role == enum.RoleSeller.String() {
		if err := h.service.ValidateSeller(c.Locals("user_id").(string), id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}
	err := h.service.Update(id, input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) EditProfile(c *fiber.Ctx) error {
	clientID := c.Locals("user_id").(string)
	var input UpdateProfile
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.UpdateProfile(clientID, input)
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
