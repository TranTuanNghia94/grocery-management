package main

import (
	"grocery-management/internal/config"
	"grocery-management/internal/db"
	"log"

	"go.uber.org/zap"
)

func main() {

	// Initialize logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	// Initialize DB
	cfg, err := config.LoadConfig(logger)
	if err != nil {
		log.Fatalf("could not load config: %v", err)
	}

	if err := db.Initialize(*cfg, logger); err != nil {
		logger.Fatal("failed to initialize database", zap.Error(err))
	}

	defer db.Instance.Close()
	// Load configuration

}
