package productmaterial

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterRoutes(routes fiber.Router)
	Create(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	Delete(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(routes fiber.Router) {
	productMaterial := routes.Group("/product-materials")

	roles := []enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}

	productMaterial.Use(middleware.Jwt())

	productMaterial.Post("/", middleware.RequireRole(roles), h.Create)
	productMaterial.Get("/:id", h.FindAll)
	productMaterial.Patch("/:id", middleware.RequireRole(roles), h.Update)
	productMaterial.Delete("/:id", middleware.RequireRole(roles), h.Delete)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := h.service.Create(input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create product material")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) FindAll(c *fiber.Ctx) error {
	id := c.Params("id")

	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(id, opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch product materials")
	}

	return c.Status(fiber.StatusOK).JSON(finded)
}

func (h *handler) Update(c *fiber.Ctx) error {
	id := c.Params("id")

	var input Update
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := h.service.Update(id, input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update product material")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	if err := h.service.Delete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete product material")
	}

	return c.SendStatus(fiber.StatusOK)
}
