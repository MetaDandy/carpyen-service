package batchproductsupplier

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterRoutes(router fiber.Router)
	create(c *fiber.Ctx) error
	findByID(c *fiber.Ctx) error
	findAll(c *fiber.Ctx) error
	update(c *fiber.Ctx) error
	softDelete(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewBatchProductSupplierHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	batch_product_suppliers := router.Group("/batch_product_suppliers")

	batch_product_suppliers.Use(middleware.Jwt())

	batch_product_suppliers.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.create)
	batch_product_suppliers.Get("/", h.findAll)
	batch_product_suppliers.Get("/:id", h.findByID)
	batch_product_suppliers.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.update)
	batch_product_suppliers.Delete("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.softDelete)
}

func (h *handler) create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	user_id := c.Locals("user_id").(string)

	if err := h.service.Create(input, user_id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create the batch product supplier")
	}

	return c.SendStatus(fiber.StatusCreated)

}

func (h *handler) findByID(c *fiber.Ctx) error {
	id := c.Params("id")

	batchproductsupplier, err := h.service.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Batch product supplier not found")
	}

	return c.JSON(batchproductsupplier)
}

func (h *handler) findAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve batch product suppliers")
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

	if role == enum.RoleInstaller.String() {
		if err := h.service.ValidateInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.Update(id, input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update batch product supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) softDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	role := c.Locals("role").(string)
	user_id := c.Locals("user_id").(string)

	if role == enum.RoleInstaller.String() {
		if err := h.service.ValidateInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.SoftDelete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete the batch product supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}
