package database

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"CRUD_echo_ent/ent"
	_ "github.com/lib/pq"
)

// ConnectDatabase establishes a connection to the database and returns an ent.Client
func ConnectDatabase() (*ent.Client, error) {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("Warning: Error loading .env file")
	}

	// Construct database connection string
	dbURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASSWORD"),
	)

	// Create ent client
	client, err := ent.Open("postgres", dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	return client, nil
}