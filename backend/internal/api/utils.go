package api

import (
	"github.com/gofiber/fiber/v2"
)

func setTokenAsCookie(c *fiber.Ctx, token string) {
	c.Cookie(&fiber.Cookie{
		Name:     "auth_token",
		Value:    token,
		HTTPOnly: true,
		Secure:   false,    // TODO: Set to `true` in production (requires HTTPS)
		SameSite: "Strict", // Prevents cross-site usage
	})
}
