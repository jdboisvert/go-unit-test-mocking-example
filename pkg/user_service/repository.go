package user_service

import (
	models "github.com/jdboisvert/go-unit-test-mocking/pkg/user_service/models"
)

type UserRepository struct{}

func (r UserRepository) GetAllUsers() ([]models.User, error) {
	users := []models.User{
		{FirstName: "Bob", LastName: "Tester 1"},
		{FirstName: "Bob", LastName: "Tester 2"},
	}
	return users, nil
}
