package services

import "grocery-management/internal/models"

func (s *GroceryService) CreateGrocery(input models.Grocery) (*models.Grocery, error) {
	grocery, err := s.repo.Create(&input)
	if err != nil {
		return nil, err
	}
	return grocery, nil
}
