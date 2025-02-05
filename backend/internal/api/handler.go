package api

import (
	"net/http"
	"strconv"
	"time"

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

func getUserInfo() fiber.Handler {
	return func(c *fiber.Ctx) error {
		userID := c.Locals("userID")
		email := c.Locals("email")
		name := c.Locals("name")

		if userID == nil || email == nil || name == nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized",
			})
		}

		return c.JSON(fiber.Map{
			"id":    userID,
			"email": email,
			"name":  name,
		})
	}
}

func registerUser(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var user UserDTO
		if err := c.BodyParser(&user); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		err := svc.RegisterUser(c.Context(), toDomainUser(user))
		if err != nil {
			return c.Status(http.StatusConflict).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "Registration successful. Please check your email to confirm your account.",
		})
	}
}

func validateEmail(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		token := c.Query("token")
		if token == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid or missing token.",
			})
		}

		user, err := svc.ValidateToken(c.Context(), token)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		jwtToken, err := service.GenerateJWT(user)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to log in after validation.",
			})
		}

		setTokenAsCookie(c, jwtToken)

		return c.JSON(fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		})
	}

}

func loginUser(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var credentials struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&credentials); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body",
			})
		}

		token, user, err := svc.LoginUser(c.Context(), credentials.Email, credentials.Password)
		if err != nil {
			return c.Status(http.StatusUnauthorized).JSON(fiber.Map{
				"error": err.Error(),
			})
		}

		setTokenAsCookie(c, token)

		return c.JSON(fiber.Map{
			"id":    user.ID,
			"email": user.Email,
			"name":  user.Name,
		})
	}
}

func logoutUser() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// Clear the auth_token cookie
		c.Cookie(&fiber.Cookie{
			Name:     "auth_token",
			Value:    "",
			Expires:  time.Now().Add(-time.Hour),
			HTTPOnly: true,
			Secure:   false, // Set to true in production for HTTPS
			SameSite: "Strict",
		})

		return c.JSON(fiber.Map{
			"message": "Logged out successfully.",
		})
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

		userIDFloat, ok := c.Locals("userID").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Invalid user ID type",
			})
		}

		userID := int(userIDFloat)

		nominations, guess, err := svc.GetNominationsByCategory(c.Context(), userID, categoryID)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to retrieve nominations",
			})
		}

		nominationsResponse := make([]NominationResponse, len(nominations))
		for i, n := range nominations {
			nominationsResponse[i] = toResponseNomination(n, guess)
		}

		return c.JSON(fiber.Map{
			"nominations": nominationsResponse,
		})
	}
}

func submitGuess(svc service.Service) fiber.Handler {
	return func(c *fiber.Ctx) error {
		userIDFloat, ok := c.Locals("userID").(float64)
		if !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Unauthorized request.",
			})
		}
		userID := int(userIDFloat)

		var guess GuessDTO
		if err := c.BodyParser(&guess); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{
				"error": "Invalid request body.",
			})
		}

		domainGuess := toDomainGuess(guess, userID)

		err := svc.SubmitGuess(c.Context(), domainGuess)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
				"message": "Failed to submit guess.",
			})
		}

		return c.Status(http.StatusCreated).JSON(fiber.Map{
			"message": "Guess submitted successfully.",
		})
	}
}
