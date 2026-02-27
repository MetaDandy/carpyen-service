package middleware

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func Jwt() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Leer access token de la cookie
		accessToken := c.Cookies("accessToken")
		if accessToken == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Access token missing",
				"code":  "NO_TOKEN",
			})
		}

		token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil {
			// Access token expiró, enviar código específico
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Access token expired or invalid",
				"code":  "TOKEN_EXPIRED",
			})
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Verificar que sea un access token
			if tokenType, exists := claims["type"].(string); exists && tokenType != "access" {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Invalid token type",
				})
			}

			c.Locals("user_id", claims["sub"])
			c.Locals("email", claims["email"])
			if role, ok := claims["role"].(string); ok {
				c.Locals("role", role)
			} else {
				return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
					"error": "Role claim missing in token",
				})
			}
		} else {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid token claims",
			})
		}

		return c.Next()
	}
}
