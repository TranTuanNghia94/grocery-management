package repositories

import "grocery-management/internal/models"

func (r *GroceryRepository) Create(grocery *models.Grocery) (*models.Grocery, error) {
	// Use the generated query interface to create a grocery
	err := r.q.Grocery.Create(grocery)
	if err != nil {
		return nil, err
	}
	return grocery, nil
}
