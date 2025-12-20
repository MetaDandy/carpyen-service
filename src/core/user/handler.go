package user

import (
	"github.com/MetaDandy/go-fiber-skeleton/helper"
	"github.com/MetaDandy/go-fiber-skeleton/middleware"
	"github.com/gofiber/fiber/v2"
)

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

type Handler struct {
	service UserService
}

func NewUserHandler(service UserService) UserHandler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	users := router.Group("/users")
	users.Post("/login", h.Login)
	users.Get("/", h.FindAll)

	users.Use(middleware.Jwt())

	users.Post("/", h.Create)
	users.Get("/me", h.GetProfile)
	users.Get("/:id", h.GetByID)
	users.Put("/:id", h.Edit)
	users.Put("/me/edit", h.EditProfile)
	users.Delete("/:id", h.SoftDelete)
}

func (h *Handler) Login(c *fiber.Ctx) error {
	var input LoginDTO
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	token, err := h.service.Login(input)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	return c.JSON(fiber.Map{
		"token": token,
	})
}

func (h *Handler) GetByID(c *fiber.Ctx) error {
	id := c.Params("id")

	user, err := h.service.GetByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

func (h *Handler) GetProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)

	user, err := h.service.GetByID(userID)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	return c.JSON(user)
}

func (h *Handler) Create(c *fiber.Ctx) error {
	var input CreateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.Create(input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not create user")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *Handler) Edit(c *fiber.Ctx) error {
	id := c.Params("id")
	var input UpdateUserDTO
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.Update(id, input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update user")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) EditProfile(c *fiber.Ctx) error {
	userID := c.Locals("user_id").(string)
	var input UpdateUserProfileDTO
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	err := h.service.UpdateProfile(userID, input)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not update profile")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *Handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not retrieve users")
	}

	return c.JSON(finded)
}

func (h *Handler) SoftDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := h.service.SoftDelete(id)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not delete user")
	}

	return c.SendStatus(fiber.StatusOK)
}
