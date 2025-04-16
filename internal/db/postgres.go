package db

import (
	"errors"
	"fmt"
	"grocery-management/internal/config"
	"time"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	// Instance is the global database instance
	Instance *PostgresDB
)

// PostgresDB implements the Database interface with PostgreSQL
type PostgresDB struct {
	db *gorm.DB
}

// NewPostgresDB creates a new PostgreSQL database connection using GORM
func NewPostgresDB(cfg config.Config) (*PostgresDB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.SSLMode,
	)

	// Configure GORM
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}

	// Create connection
	db, err := gorm.Open(postgres.Open(dsn), gormConfig)
	if err != nil {
		return nil, errors.New("unable to connect to database: " + err.Error())
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, errors.New("unable to get underlying database: " + err.Error())
	}

	// Set connection pool settings
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &PostgresDB{db: db}, nil
}

// Close closes the database connection pool
func (p *PostgresDB) Close() {
	if p.db != nil {
		sqlDB, err := p.db.DB()
		if err == nil {
			sqlDB.Close()
		}
	}
}

// GetDB returns the GORM DB instance for use in database operations
func (p *PostgresDB) GetDB() *gorm.DB {
	return p.db
}

func Initialize(cfg config.Config, log *zap.Logger) error {
	db, err := NewPostgresDB(cfg)
	if err != nil {
		return err
	}
	Instance = db
	log.Info("PostgreSQL database connection established", zap.String("host", cfg.DBHost), zap.String("port", cfg.DBPort))
	return nil
}
