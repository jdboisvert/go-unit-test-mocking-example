package user_service

import (
	models "github.com/jdboisvert/go-unit-test-mocking/pkg/user_service/models"
)

type UserRepositoryInterface interface {
	GetAllUsers() ([]models.User, error)
}

type UserService struct {
	UserRepositoryInterface
}

// Here we use the interface given to the struct to then use (has-a)
func (s UserService) GetUser() ([]models.User, error) {
	users, _ := s.UserRepositoryInterface.GetAllUsers()
	return users, nil
}
