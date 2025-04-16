package services

import "grocery-management/internal/models"

func (s *GroceryService) GetAllGroceries() ([]models.Grocery, error) {
	groceries, err := s.repo.GetAll()
	if err != nil {
		return nil, err
	}

	return groceries, nil
}
