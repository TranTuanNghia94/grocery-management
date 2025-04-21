package repositories

import (
	"errors"
	"grocery-management/internal/models"
	"grocery-management/pkg/logger"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func (r *UserRepository) GetByUsername(ctx *gin.Context, username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	u := r.q.User

	user, err := u.WithContext(ctx).Where(u.Username.Eq(username)).First()
	if err != nil {
		logger.Log.Error("failed to query user by username", zap.Error(err))
		return nil, errors.New("failed to query user by username")
	}

	logger.Log.Info("query user by username", zap.String("username", username))

	if err != nil {
		logger.Log.Error("failed to query user by username", zap.Error(err))
		return nil, errors.New("failed to query user by username")
	}

	return user, nil
}
