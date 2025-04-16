package repositories

import (
	"grocery-management/internal/models"
	"grocery-management/internal/query"

	"gorm.io/gorm"
)

var _ IGroceryRepository = (*GroceryRepository)(nil)

type IGroceryRepository interface {
	Create(grocery *models.Grocery) (*models.Grocery, error)
	GetAll() ([]models.Grocery, error)
}

type GroceryRepository struct {
	DB *gorm.DB
	q  *query.Query // The generated query interface
}

func NewGroceryRepo(db *gorm.DB) IGroceryRepository {
	return &GroceryRepository{
		DB: db,
		q:  query.Use(db),
	}
}
