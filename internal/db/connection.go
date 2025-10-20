package db

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

var DB *pgxpool.Pool

func InitDB() {

	// load env
	err := godotenv.Load()
	if err != nil {
		log.Println(".env file not found")
	}

	// get data store name
	dsn := os.Getenv("DB_CONNECTION_STRING")

	// connect to DB, context.Background is root level context,
	// saying that no timeout is needed for DB connection
	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("Cannot connect to database: %v", err)
	}

	// test connection
	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatalf("Cannot ping database: %v", err)
	}

	log.Println("Connected to database")
	DB = pool
}
