package batchproductmaterial

import (
	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterRoutes(router fiber.Router)
	Create(c *fiber.Ctx) error
	FindByID(c *fiber.Ctx) error
	FindAll(c *fiber.Ctx) error
	Update(c *fiber.Ctx) error
	SoftDelete(c *fiber.Ctx) error
}

type handler struct {
	service Service
}

func NewBatchProductMaterialHandler(service Service) Handler {
	return &handler{service: service}
}

func (h *handler) RegisterRoutes(router fiber.Router) {
	batchProductMaterial := router.Group("/batch-product-materials")

	batchProductMaterial.Use(middleware.Jwt())

	batchProductMaterial.Post(
		"/",
		middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}),
		h.Create,
	)
	batchProductMaterial.Get("/", h.FindAll)
	batchProductMaterial.Get("/:id", h.FindByID)
	batchProductMaterial.Patch(
		"/:id",
		middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}),
		h.Update,
	)
	batchProductMaterial.Delete(
		"/:id",
		middleware.RequireRole([]enum.Role{enum.RoleAdmin, enum.RoleInstaller, enum.RoleChiefInstaller}),
		h.SoftDelete,
	)
}

func (h *handler) Create(c *fiber.Ctx) error {
	var input Create
	if err := c.BodyParser(&input); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid input")
	}

	user_id := c.Locals("user_id").(string)

	if err := h.service.Create(input, user_id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to create batch product material")
	}

	return c.SendStatus(fiber.StatusCreated)
}

func (h *handler) FindByID(c *fiber.Ctx) error {
	id := c.Params("id")

	batchProductMaterial, err := h.service.FindByID(id)
	if err != nil {
		return fiber.NewError(fiber.StatusNotFound, "Batch product material not found")
	}

	return c.JSON(batchProductMaterial)
}

func (h *handler) FindAll(c *fiber.Ctx) error {
	opts := helper.NewFindAllOptionsFromQuery(c)

	finded, err := h.service.FindAll(opts)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to fetch batch product materials")
	}

	return c.JSON(finded)
}

func (h *handler) Update(c *fiber.Ctx) error {
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
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update batch product material")
	}

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) SoftDelete(c *fiber.Ctx) error {
	id := c.Params("id")

	role := c.Locals("role").(string)
	user_id := c.Locals("user_id").(string)

	if role == enum.RoleInstaller.String() {
		if err := h.service.ValidateInstaller(id, user_id); err != nil {
			return fiber.NewError(fiber.StatusUnauthorized, "Unauthorized")
		}
	}

	if err := h.service.SoftDelete(id); err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to delete batch product material")
	}

	return c.SendStatus(fiber.StatusOK)
}
