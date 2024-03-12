package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs")
	// no inheritance in golang; No super or parent

	abhishek := User{"abhi", "abhi@gmail.com", true, 19}
	fmt.Println(abhishek)
	fmt.Printf("abhishek details are: %+v\n", abhishek)
	fmt.Printf("Name is %v and Email is %v", abhishek.Name, abhishek.Email)
}

type User struct{
	Name string 
	Email string
	Status bool
	Age int
}