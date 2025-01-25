package main

import (
	"log"

	"github.com/TommyN987/guess_the_oscars/backend/internal/api"
	"github.com/TommyN987/guess_the_oscars/backend/internal/config"
	"github.com/TommyN987/guess_the_oscars/backend/internal/db"
	"github.com/TommyN987/guess_the_oscars/backend/internal/repository"
	"github.com/TommyN987/guess_the_oscars/backend/internal/service"
)

func main() {
	cfg := config.LoadConfig()

	pool := db.ConnectDB(cfg)
	defer pool.Close()

	repo := repository.NewPgxRepository(pool)

	svc := service.NewDefaultService(repo)

	app := api.NewRouter(svc)

	addr := ":8080"
	log.Printf("Starting server on %s... \n", addr)

	if err := app.Listen(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
