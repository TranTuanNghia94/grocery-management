package repositories

import (
	"grocery-management/internal/models"
	"grocery-management/internal/query"

	"gorm.io/gorm"
)

var _ IUserRepository = (*UserRepository)(nil)

type IUserRepository interface {
	Create(user *models.User) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	// GetByID(id string) (*models.User, error)
	// GetByEmail(email string) (*models.User, error)
	// Update(user *models.User) (*models.User, error)
	// Delete(id string) error
	// GetAll() ([]*models.User, error)
}

type UserRepository struct {
	DB *gorm.DB
	q  *query.Query // The generated query interface
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		DB: db,
		q:  query.Use(db),
	}
}
