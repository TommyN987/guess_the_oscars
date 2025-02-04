package api

import (
	"github.com/TommyN987/guess_the_oscars/backend/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func NewRouter(svc service.Service) *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:5173",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, OPTIONS",
		AllowCredentials: true,
	}))

	app.Get("/validate", validateEmail(svc))
	app.Post("/api/register", registerUser(svc))
	app.Post("/api/login", loginUser(svc))
	app.Get("/api/categories", getAllCategories(svc))
	app.Get("/api/categories/:id/nominations", getNominationsByCategory(svc))
	app.Get("/health", checkHealth(svc))
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Guess the Oscars!")
	})

	protected := app.Group("/api/p", AuthMiddleware())
	protected.Get("/me", getUserInfo())
	protected.Post("/logout", logoutUser())
	protected.Post("/guesses", submitGuesses())

	return app
}
