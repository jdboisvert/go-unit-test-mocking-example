package robot_service

import (
	models "github.com/jdboisvert/go-unit-test-mocking/pkg/robot_service/models"
)

type FetcherArgs struct {
	Fetcher  func() ([]models.Robot, error)
}

func GetAllRobots() ([]models.Robot, error) {
	robots := []models.Robot{
		{Name: "Optimus"},
		{Name: "Megatron"},
		{Name: "R.O.B."},
	}
	return robots, nil
}

func GetRobots(args FetcherArgs) ([]models.Robot, error) {
	return args.Fetcher() // Here we are calling the function passed into the method 
}
