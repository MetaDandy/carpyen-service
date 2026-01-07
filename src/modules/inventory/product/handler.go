package product

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

func NewProductHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	product := router.Group("/products")

	product.Use(middleware.Jwt())

	product.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.create)
	product.Get("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller, enum.RoleInstaller}), h.findAll)
	product.Get("/:id", h.findByID)
	product.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.update)
	product.Delete("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.softDelete)
}

func (h *handler) create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	user_id := c.Locals("user_id").(string)

	if err := h.service.Create(input, user_id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create product")
	}

	return c.SendStatus(fiber.StatusCreated)

}

func (h *handler) findByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.service.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Product not found")
	}

	return c.JSON(product)
}

func (h *handler) findAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to retrieve products")
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
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update product")
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
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete product")
	}

	return c.SendStatus(fiber.StatusOK)
}
