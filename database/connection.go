package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func ConnectDatabase() *pgxpool.Pool {
	godotenv.Load()
	ctx := context.Background()

	database_url := os.Getenv("DATABASE_URL")
	conn, err := pgxpool.New(ctx, database_url)

	if err != nil {
		panic(err)
	}

	return conn
}
