package main

import "fmt"

func main() {
	defer fmt.Println("This is printed first")
	defer fmt.Println("This is printed second")
	defer fmt.Println("\nThis is printed third")
	fmt.Println("Welcome to defer in Golang!!")
	myDefer();
}

func myDefer(){
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}