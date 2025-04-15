package repositories

import (
	"grocery-management/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (r *UserRepository) Create(user *models.User) (*models.User, error) {
	// Hash password before saving
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	user.PasswordHash = string(hashedPassword)

	// Using generated code for create operation
	if err := r.q.User.Create(user); err != nil {
		return nil, err
	}
	// Return the created user
	return user, nil
}
