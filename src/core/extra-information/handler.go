package extrainformation

import (
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler struct{}

func NewExtraInformationHandler() Handler {
	return Handler{}
}

func (h *Handler) RegisterRoutes(router fiber.Router) {
	extra := router.Group("/extra-information")

	extra.Get("/measures", h.GetMeasures)
	extra.Get("/material-types", h.GetMaterialTypes)
	extra.Get("/product-types", h.GetProductTypes)
	extra.Get("/roles", h.GetRoles)

}

func (h *Handler) GetMeasures(c *fiber.Ctx) error {
	return c.JSON(enum.MeasuretoArray())
}

func (h *Handler) GetMaterialTypes(c *fiber.Ctx) error {
	return c.JSON(enum.MaterialToArray())
}

func (h *Handler) GetProductTypes(c *fiber.Ctx) error {
	return c.JSON(enum.ProductToArray())
}

func (h *Handler) GetRoles(c *fiber.Ctx) error {
	return c.JSON(enum.RoleToArray())
}
