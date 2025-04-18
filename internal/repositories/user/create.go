package repositories

import (
	"errors"
	"grocery-management/internal/models"
)

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	if user == nil {
		return nil, errors.New("user cannot be nil")
	}

	// Using generated code for create operation
	if err := r.q.User.Create(user); err != nil {
		return nil, err
	}
	// Return the created user
	return user, nil
}
