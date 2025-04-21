package main

import (
	"grocery-management/internal/api"
	"grocery-management/internal/config"
	"grocery-management/internal/db"
	"grocery-management/internal/di"
	"grocery-management/pkg/logger"
	"grocery-management/pkg/security"

	"go.uber.org/zap"
)

func main() {
	// Initialize logger
	logger.InitWithRotation(".logs/app.log", "info")
	defer logger.Sync()

	logger.Info("Starting Grocery Management")

	// Initialize DB
	cfg, err := config.LoadConfig(logger.Log)
	if err != nil {
		logger.Fatal("could not load config: %v", zap.Error(err))
	}

	if err := db.Initialize(*cfg, logger.Log); err != nil {
		logger.Fatal("failed to initialize database", zap.Error(err))
	}

	defer db.Instance.Close()

	// Setup dependency injection
	container := di.NewContainer(db.Instance.GetDB(), cfg)

	// Setup routes
	r := api.SetupRouter(container)

	security.Init(logger.Log)

	// Start server
	logger.Info("Starting server", zap.String("port", cfg.AppPort))
	if err := r.Run(":" + cfg.AppPort); err != nil {
		logger.Fatal("failed to start server", zap.Error(err))
	}
}
