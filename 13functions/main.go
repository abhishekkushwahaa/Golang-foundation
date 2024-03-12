package main

import "fmt"

func main() {
	greet()

	fmt.Println("Welcome to Functions in Golang!!")

	result := add(2, 3)
	fmt.Println("Result of 2 + 3 is: ", result)

	proResult := proAdd(2, 3, 4, 5, 6)
	fmt.Println("Result of 2 + 3 + 4 + 5 + 6 is: ", proResult)
}

func add(a int, b int) int {
	return a + b
}

// Pro Function
func proAdd(values... int) int{
	result := 0
	for _, value := range values {
		result += value
	}
	return result
}

func greet()  {
	fmt.Println("Good Morning!!")
}