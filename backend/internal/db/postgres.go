package db

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/TommyN987/guess_the_oscars/backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDB(cfg config.Config) *pgxpool.Pool {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s",
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
	)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.New(ctx, dsn)

	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}

	log.Println("Database connection established")

	return pool
}
