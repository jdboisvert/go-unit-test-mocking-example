package food_service

import (
	"testing"

	models "github.com/jdboisvert/go-unit-test-mocking/pkg/food_service/models"
)

// Here we want to test GetFood but we will mock the GetAllFoodAvailable
// To return Different Food items than it normally does.
func Test_GetFoodGetCatalogMocked(t *testing.T) {
	// preserve the original function
	originalGetAllFoodAvailableFunction := GetAllFoodAvailable

	// Here we declare our new function. In theory this can be anything however so be careful when patching.
	GetAllFoodAvailable = func() ([]models.Food, error) {
		food := []models.Food{
			{Name: "Chocolate"},
		}
		return food, nil
	}

	food, _ := GetFood()

	if food[0].Name != "Chocolate" {
		t.Errorf("Error it seems the patch or function failed. Expected Chocolate got %v instead", food[0].Name)

	}

	// Acts as a clean up to re-assign the mock back.
	// As you can imagine this is not ideal and can lead to fragile tests.
	GetAllFoodAvailable = originalGetAllFoodAvailableFunction
}
