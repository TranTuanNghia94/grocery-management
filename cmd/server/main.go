package main

import (
	"grocery-management/internal/api"
	"grocery-management/internal/config"
	"grocery-management/internal/db"
	"grocery-management/internal/di"
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

	// Setup dependency injection
	container := di.NewContainer(db.Instance.GetDB(), cfg)

	// Setup routes
	r := api.SetupRouter(container)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.AppPort))
	if err := r.Run(":" + cfg.AppPort); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
