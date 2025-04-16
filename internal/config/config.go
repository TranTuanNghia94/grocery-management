package config

import (
	"errors"
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
	AppPort    string
}

func LoadConfig(log *zap.Logger) (*Config, error) {
	_ = godotenv.Load()

	if os.Getenv("POSTGRES_USER") == "" {
		log.Error("Can not read file env")
		return nil, errors.New("DB_USER is not set")
	}

	log.Info("DATABASE HOST", zap.String("host", os.Getenv("POSTGRES_HOST")))
	log.Info("DATABASE PORT", zap.String("port", os.Getenv("POSTGRES_PORT")))

	log.Info("Loaded environment variables successfully")

	return &Config{
		DBHost:     os.Getenv("POSTGRES_HOST"),
		DBPort:     os.Getenv("POSTGRES_PORT"),
		DBUser:     os.Getenv("POSTGRES_USER"),
		DBPassword: os.Getenv("POSTGRES_PASSWORD"),
		DBName:     os.Getenv("POSTGRES_DB"),
		SSLMode:    os.Getenv("POSTGRES_SSLMODE"),
		AppPort:    os.Getenv("APP_PORT"),
	}, nil
}
