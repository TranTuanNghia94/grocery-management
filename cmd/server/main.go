package main

import (
	"grocery-management/internal/config"
	"grocery-management/internal/db"
	"log"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}
	log.Print("Loaded configuration database successfully")

	// Initialize database connection
	db, err := db.NewPostgresDB(*cfg)
	if err != nil {
		log.Fatalf("could not connect to database: %v", err)
	}
	defer db.Close()
	log.Print("Connected to database successfully")

}
