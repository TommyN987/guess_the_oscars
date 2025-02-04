package api

import (
	"net/http"

	"github.com/TommyN987/guess_the_oscars/backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func AuthMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Get the JWT from the cookie
		tokenStr := c.Cookies("auth_token")
		if tokenStr == "" {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: No token provided.",
			})
		}

		// Validate the token
		token, err := service.ValidateJWT(tokenStr)
		if err != nil || !token.Valid {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized: Invalid token.",
			})
		}

		// Extract claims and add to request context
		claims := token.Claims.(jwt.MapClaims)
		c.Locals("userID", claims["id"])
		c.Locals("email", claims["email"])
		c.Locals("name", claims["name"])

		return c.Next()
	}
}
