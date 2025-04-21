package repositories

import (
	"grocery-management/internal/models"
	"grocery-management/internal/query"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var _ IUserRepository = (*UserRepository)(nil)

type IUserRepository interface {
	Create(ctx *gin.Context, user *models.User) (*models.User, error)
	GetByUsername(ctx *gin.Context, username string) (*models.User, error)
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
