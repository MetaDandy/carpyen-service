package user

import (
	"fmt"

	"github.com/MetaDandy/carpyen-service/helper"
	"github.com/MetaDandy/carpyen-service/middleware"
	"github.com/MetaDandy/carpyen-service/src/enum"
	"github.com/gofiber/fiber/v2"
)

type Handler interface {
	RegisterRoutes(router fiber.Router)
	Login(c *fiber.Ctx) error
	Logout(c *fiber.Ctx) error
	RefreshToken(c *fiber.Ctx) error
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
	users.Post("/refresh", h.RefreshToken) // Renovar access token
	users.Post("/logout", h.Logout)

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

	user, err := h.service.Login(input)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid credentials")
	}

	// Generar AMBOS tokens (Access + Refresh)
	tokenPair, err := helper.GenerateTokenPair(user.ID, user.Email, user.Role)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not generate tokens")
	}

	// Access Token: httpOnly, 1 hora
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    tokenPair.AccessToken,
		HTTPOnly: true,
		Secure:   false, // Cambiar a true en producción con HTTPS
		SameSite: "Strict",
		MaxAge:   3600, // 1 hora
		Path:     "/",
	})

	// Refresh Token: httpOnly, 7 días, más seguro
	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    tokenPair.RefreshToken,
		HTTPOnly: true,
		Secure:   false, // Cambiar a true en producción con HTTPS
		SameSite: "Strict",
		MaxAge:   7 * 24 * 3600, // 7 días
		Path:     "/",
	})

	// Cookie con info de expiración (sin httpOnly para que sea legible desde JS)
	// Contiene timestamp Unix en segundos de cuando caduca el access token
	c.Cookie(&fiber.Cookie{
		Name:     "tokenExpiry",
		Value:    fmt.Sprintf("%d", tokenPair.AccessExpire),
		HTTPOnly: false, // Permitir lectura desde JavaScript
		Secure:   false, // Cambiar a true en producción con HTTPS
		SameSite: "Strict",
		MaxAge:   3600, // Mismo tiempo que el access token
		Path:     "/",
	})

	// Retornar solo datos del usuario (tokens están en cookies)
	return c.Status(fiber.StatusOK).JSON(user)
}

func (h *handler) Logout(c *fiber.Ctx) error {
	// Eliminar la cookie estableciendo MaxAge en -1
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    "",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   -1, // Esto elimina la cookie
		Path:     "/",
	})

	c.Cookie(&fiber.Cookie{
		Name:     "refreshToken",
		Value:    "",
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   -1,
		Path:     "/",
	})

	// Eliminar cookie de expiración
	c.Cookie(&fiber.Cookie{
		Name:     "tokenExpiry",
		Value:    "",
		HTTPOnly: false,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   -1,
		Path:     "/",
	})

	return c.SendStatus(fiber.StatusOK)
}

func (h *handler) RefreshToken(c *fiber.Ctx) error {
	// Leer refresh token de la cookie
	refreshTokenString := c.Cookies("refreshToken")
	if refreshTokenString == "" {
		return fiber.NewError(fiber.StatusUnauthorized, "Refresh token missing")
	}

	// Validar que sea un refresh token válido
	claims, err := helper.ValidateToken(refreshTokenString)
	if err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid refresh token")
	}

	// Verificar que sea realmente un refresh token
	if tokenType, ok := claims["type"].(string); !ok || tokenType != "refresh" {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid token type")
	}

	userID := claims["sub"].(string)
	email := claims["email"].(string)
	role := claims["role"].(string)

	// Generar NUEVO access token
	newAccessToken, expireTime, err := helper.GenerateAccessToken(userID, email, role)
	if err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Could not generate token")
	}

	// Establecer nuevo access token
	c.Cookie(&fiber.Cookie{
		Name:     "accessToken",
		Value:    newAccessToken,
		HTTPOnly: true,
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   3600, // 1 hora
		Path:     "/",
	})

	// Actualizar cookie de expiración
	c.Cookie(&fiber.Cookie{
		Name:     "tokenExpiry",
		Value:    fmt.Sprintf("%d", expireTime),
		HTTPOnly: false, // Permitir lectura desde JavaScript
		Secure:   false,
		SameSite: "Strict",
		MaxAge:   3600, // 1 hora
		Path:     "/",
	})

	return c.SendStatus(fiber.StatusOK)
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
