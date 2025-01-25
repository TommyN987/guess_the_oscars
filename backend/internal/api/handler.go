package api

import (
	"net/http"
	"strconv"

	"github.com/TommyN987/guess_the_oscars/backend/internal/service"
	"github.com/gofiber/fiber/v2"
)

func checkHealth(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		health, err := svc.CheckHealth(c.Context())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Database connection failed",
			})
		}
		return c.JSON(health)
	}
}

func getAllCategories(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categories, err := svc.GetAllCategories(c.Context())
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve categories",
			})
		}
		return c.JSON(categories)
	}
}

func getNominationsByCategory(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		categoryID, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid category ID",
			})
		}

		category, nominations, err := svc.GetNominationsByCategory(c.Context(), categoryID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve nominations",
			})
		}

		return c.JSON(fiber.Map{
			"category":    category,
			"nominations": nominations,
		})
	}
}
