package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Config for database connection
type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func main() {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		fmt.Printf("Warning: .env file not found or couldn't be loaded: %v\n", err)
		// Continue execution even if .env file is not found
	}

	// Configure database connection
	dbConfig := Config{
		Host:     getEnvOrDefault("POSTGRES_HOST", "localhost"),
		Port:     getEnvOrDefault("POSTGRES_PORT", "5432"),
		User:     getEnvOrDefault("POSTGRES_USER", "postgres"),
		Password: getEnvOrDefault("POSTGRES_PASSWORD", "postgres"),
		DBName:   getEnvOrDefault("POSTGRES_DB", "grocery_db"),
		SSLMode:  getEnvOrDefault("POSTGRES_SSLMODE", "disable"),
	}

	// Connect to the database
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		dbConfig.Host, dbConfig.Port, dbConfig.User, dbConfig.Password, dbConfig.DBName, dbConfig.SSLMode,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		os.Exit(1)
	}

	// Initialize generator
	g := gen.NewGenerator(gen.Config{
		OutPath:      "internal/query",
		ModelPkgPath: "models",
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	// Use the connected db
	g.UseDB(db)

	// Generate basic models for all tables
	// Get all tables from the database
	allTables := g.GenerateAllTable()

	// Apply to generator
	g.ApplyBasic(allTables...)

	// Generate the code
	g.Execute()
}

// getEnvOrDefault returns environment variable value or default if not set
func getEnvOrDefault(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
