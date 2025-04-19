package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB(databaseURL string) {
	var err error

	DB, err = pgxpool.New(context.Background(), databaseURL)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")
}

func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}
