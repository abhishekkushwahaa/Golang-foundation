package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is", ptr)

	myNumber := 23
	var ptr = &myNumber
	// fmt.Println("Value of number is", myNumber)
	fmt.Println("Value of pointer is", ptr) // Address memory
	fmt.Println("Value of pointer is", *ptr) // Actual value

	*ptr = *ptr + 2
	fmt.Println("New vlaue is:", myNumber)

}