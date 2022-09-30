package food_service

import (
	models "github.com/jdboisvert/go-unit-test-mocking/pkg/food_service/models"
)

// Here the function is declared as a variable. Several articles showed this was done but the community was split on it.
var GetAllFoodAvailable = func() ([]models.Food, error) {
	food := []models.Food{
		{Name: "Carrot"},
		{Name: "Cheese"},
		{Name: "Lettuce"},
	}
	return food, nil
}
