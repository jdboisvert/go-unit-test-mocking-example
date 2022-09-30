package food_service

import (
	models "github.com/jdboisvert/go-unit-test-mocking/pkg/food_service/models"
)

// Here we are simply using a function which will be later mocked.
func GetFood() ([]models.Food, error) {
	food_available, _ := GetAllFoodAvailable()
	return food_available, nil
}
