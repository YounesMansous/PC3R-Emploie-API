package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

var DB *pgx.Conn

func ConnectDB() {
	var err error
	DB, err = pgx.Connect(context.Background(), "DATABASE_URL")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Database connected")
}

func CloseDB() {
	if DB != nil {
		DB.Close(context.Background())
	}
}
