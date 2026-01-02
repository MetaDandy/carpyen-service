package supplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	registerRoutes(router fiber.Router)
	create(c *fiber.Ctx) error
	findByID(c *fiber.Ctx) error
	findAll(c *fiber.Ctx) error
	update(c *fiber.Ctx) error
	softDelete(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) registerRoutes(router fiber.Router) {
	supplier := router.Group("/suppliers")

	supplier.Use(middleware.Jwt())

	supplier.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.create)
	supplier.Get("/", h.findAll)
	supplier.Get("/:id", h.findByID)
	supplier.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.update)
	supplier.Delete("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.softDelete)
}

func (h *handler) create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := h.service.create(input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create supplier")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) findByID(c *fiber.Ctx) error {
	id := c.Params("id")

	supplier, err := h.service.findByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Supplier not found")
	}

	return c.JSON(supplier)
}

func (h *handler) findAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.findAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve suppliers")
	}

	return c.JSON(finded)
}

func (h *handler) update(c *fiber.Ctx) error {
	id := c.Params("id")
	var input Update
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	role := c.Locals("role").(string)
	user_id := c.Locals("user_id").(string)

	if role == enum.RoleChiefInstaller.String() {
		if err := h.service.validateChiefInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.update(id, input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) softDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	role := c.Locals("role").(string)
	user_id := c.Locals("user_id").(string)

	if role == enum.RoleChiefInstaller.String() {
		if err := h.service.validateChiefInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.softDelete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}
