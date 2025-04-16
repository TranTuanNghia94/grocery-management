package repositories

import "grocery-management/internal/models"

func (g *GroceryRepository) GetAll() ([]models.Grocery, error) {
	var groceries []models.Grocery
	err := g.DB.Find(&groceries).Error
	if err != nil {
		return nil, err
	}
	return groceries, nil
}
