package user_service

import (
	"fmt"
	"testing"

	models "github.com/jdboisvert/go-unit-test-mocking/pkg/user_service/models"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Create Mocked repository struct with a mock member
type UserRepositoryMock struct {
	mock.Mock
}

// Create a function that follows the same pattern as
/*
Create a function that follows the same pattern as

	type UserRepositoryInterface interface {
		GetAllUsers() ([]models.User, error)
	}
This is the interface created in service.go. This is how once can mock the functionality for tests.
*/
func (r UserRepositoryMock) GetAllUsers() ([]models.User, error) {
	args := r.Called()
	users := []models.User{
		{FirstName: "Mock", LastName: "This"},
	}
	return users, args.Error(1)
}

func TestService_GetUserMocked(t *testing.T) {
	repository := UserRepositoryMock{}
	repository.On("GetAllUsers").Return([]models.User{}, nil)

	service := UserService{UserRepositoryInterface: repository}
	users, _ := service.GetUser()
	for i := range users {
		assert.Equal(t, users[i].FirstName, "Mock", "First name is not correct")
	}

	// Here this will show that this is actually the mocked method above ^^
	fmt.Println(users)
}

func TestService_GetUser(t *testing.T) {
	repository := UserRepository{}
	service := UserService{UserRepositoryInterface: repository}
	users, _ := service.GetUser()

	assert.Equal(t, users[0].FirstName, "Bob", "First name is not correct")
	assert.Equal(t, users[0].LastName, "Tester 1", "First name is not correct")

	// Here this will show that this is actually the NOT the mocked method above but the original
	fmt.Println(users)
}
