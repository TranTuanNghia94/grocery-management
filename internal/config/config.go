package config

import (
	"os"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	SSLMode    string
}

func LoadConfig(log *zap.Logger) (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	log.Info("Loaded environment variables successfully")
	log.Info("DB_HOST", zap.String("DB_HOST", os.Getenv("DB_HOST")))

	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		SSLMode:    os.Getenv("SSL_MODE"),
	}, nil
}
