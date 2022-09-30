package main

import (
	"fmt"

	user_service "github.com/jdboisvert/go-unit-test-mocking/pkg/user_service"
)

func main() {
	// First create the main struct to be used which is the implementation of the Interface
	repository := user_service.UserRepository{}

	// Then create the service to be used which uses the interface
	// (duck typing so since the above implements all functions of interface it IS A UserRepository)
	service := user_service.UserService{UserRepositoryInterface: repository}

	// Should print out users correctly.
	users, _ := service.GetUser()
	fmt.Println(users)
}
