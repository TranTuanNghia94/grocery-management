package services

import (
	"errors"
	"grocery-management/internal/models"
	"grocery-management/pkg/security"
)

func (s *UserService) GetUserByUsername(username string) (*models.User, error) {
	if username == "" {
		return nil, errors.New("username cannot be empty")
	}

	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	return user, nil
}

func (s *UserService) Login(username, password string) (*string, error) {
	if username == "" || password == "" {
		return nil, errors.New("username and password cannot be empty")
	}

	user, err := s.repo.GetByUsername(username)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	if user.IsActive == nil || !*user.IsActive {
		return nil, errors.New("user is not active")
	}

	if !security.CheckPasswordHash(password, user.PasswordHash) {
		return nil, errors.New("invalid credentials")
	}

	jwt, err := security.GenerateJWT(user)
	if err != nil {
		return nil, err
	}

	return &jwt, nil

}
