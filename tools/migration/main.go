package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println("Warning: .env file not found or could not be loaded")
	}

	// Connect to the database
	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("DB connection error: %v", err)
	}
	defer db.Close()

	// Create an instance of the migration driver
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatalf("Migration driver error: %v", err)
	}

	// Create the migrate instance
	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations", // Folder with .sql files
		"postgres",          // Database name
		driver,
	)
	if err != nil {
		log.Fatalf("Migration instance error: %v", err)
	}

	// Run migrations
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("Migration up error: %v", err)
	}

	fmt.Println("Migrations applied successfully ✅")
}
