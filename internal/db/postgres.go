package db

import (
	"fmt"
	"grocery-management/internal/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
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
		return nil, fmt.Errorf("unable to connect to database: %w", err)
	}

	// Configure connection pool settings
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("unable to get underlying database: %w", err)
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
