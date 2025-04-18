package repositories

import (
	"errors"
	"grocery-management/internal/models"
	"grocery-management/pkg/logger"

	"go.uber.org/zap"
)

func (r *UserRepository) GetByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	user := &models.User{}
	err := r.DB.Model(user).Where("username = ?", username).Error

	if err != nil {
		logger.Log.Error("failed to query user by username", zap.Error(err))
		return nil, errors.New("failed to query user by username")
	}

	return user, nil
}
