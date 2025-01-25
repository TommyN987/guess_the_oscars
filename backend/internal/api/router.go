package api

import (
	"github.com/TommyN987/guess_the_oscars/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func NewRouter(svc service.Service) *fiber.App {
	app := fiber.New()

	app.Get("/categories", getAllCategories(svc))
	app.Get("/categories/:id/nominations", getNominationsByCategory(svc))
	app.Get("/health", checkHealth(svc))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Guess the Oscars!")
	})

	return app
}
