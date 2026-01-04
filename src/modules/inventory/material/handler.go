package material

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

func NewMaterialHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	material := router.Group("/materials")

	material.Use(middleware.Jwt())

	material.Post("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.create)
	material.Get("/", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller, enum.RoleInstaller}), h.findAll)
	material.Get("/:id", h.findByID)
	material.Patch("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}), h.update)
	material.Delete("/:id", middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleChiefInstaller}), h.softDelete)
}

func (h *handler) create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	if err := h.service.Create(input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create material")
	}

	return c.SendStatus(fiber.StatusCreated)

}

func (h *handler) findByID(c *fiber.Ctx) error {
	id := c.Params("id")

	material, err := h.service.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Material not found")
	}

	return c.JSON(material)
}

func (h *handler) findAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
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
		if err := h.service.ValidateChiefInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.Update(id, input); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) softDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	role := c.Locals("role").(string)
	user_id := c.Locals("user_id").(string)

	if role == enum.RoleChiefInstaller.String() {
		if err := h.service.ValidateChiefInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.SoftDelete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete supplier")
	}

	return c.SendStatus(fiber.StatusOK)
}
