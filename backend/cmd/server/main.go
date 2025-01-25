package main

import (
	"log"

	"github.com/TommyN987/guess_the_oscars/backend/internal/config"
	"github.com/TommyN987/guess_the_oscars/backend/internal/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	cfg := config.LoadConfig()

	pool := db.ConnectDB(cfg)
	defer pool.Close()

	app := fiber.New()

	app.Get("/health", func(c *fiber.Ctx) error {
		if err := pool.Ping(c.Context()); err != nil {
			return c.Status(500).SendString("Database connection error")
		}
		return c.SendString("Database connection healthy")
	})

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to Guess the Oscars!")
	})

	addr := ":8080"
	log.Printf("Starting server on %s... \n", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
