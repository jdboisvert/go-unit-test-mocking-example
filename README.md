# Go Unit Test Mocking via Interfaces 

A simply project based on several online examples to showcase how one can write their Go code in a way to make it easily testable and possible to mock when needed. There will be times mocks are not needed but when dealing with dependencies this can be crucial. 

### How To Start

```
# Download all dependencies 
go mod download

# To run app 
go run main.go 

# To run the tests showcasing the results (runs all tests in project)
go test ./...
```

### Features

- `food_service` Shows how one can mock functions in tests making it so a function is just a variable and in the test you can reassign this function to any function you choose. 
- `robot_service` Shows how one can mock functions in tests using dependency injection in the form of passing the function to be used as a function parameter. 
- `user_service` Shows how one can mock functions using interfaces. Similar to to the dependency injection example however makes use of interfaces for the function itself versus simply passing the function into another function.

Simply review the code as well for comments explaining each part more. 