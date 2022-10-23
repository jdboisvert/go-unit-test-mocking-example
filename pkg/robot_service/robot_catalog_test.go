package robot_service

import (
	"fmt"
	"testing"

	models "github.com/jdboisvert/go-unit-test-mocking/pkg/robot_service/models"

	"github.com/stretchr/testify/assert"
)

/*
Create a mock function to then pass in the function you wish to test. It simply
needs to have the same function pattern as declared in the strict so in this case

	func() ([]models.Robot, error)

This is how one can mock the functionality for tests.
*/
func TestRobotCatalog_GetRobotsMocked(t *testing.T) {
	var mockedFetcher = func() ([]models.Robot, error) {
		robots := []models.Robot{
			{Name: "Mock"},
		}
		return robots, nil
	}

	fetcherArgs := FetcherArgs{
		Fetcher: mockedFetcher,
	}

	robots, _ := GetRobots(fetcherArgs)
	for i := range robots {
		assert.Equal(t, robots[i].Name, "Mock", "First name is not correct")
	}

	// Here this will show that this is actually the mocked method above ^^
	fmt.Println(robots)
}

func TestService_GetRobotsNotMocked(t *testing.T) {
	fetcherArgs := FetcherArgs{
		Fetcher: GetAllRobots,
	}

	robots, _ := GetRobots(fetcherArgs)

	assert.Equal(t, robots[0].Name, "Optimus", "First robot's name is not correct")
	assert.Equal(t, robots[1].Name, "Megatron", "Second robot's name is not correct")
	assert.Equal(t, robots[2].Name, "R.O.B.", "Third robot's name is not correct")

	// Here this will show that this is actually the NOT the mocked method above but a regular function passed.
	fmt.Println(robots)
}
