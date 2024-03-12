package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else in Golang!!")

	loginCount := 23
    var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10{
		result = "Special User"
	} else {
		result = "Exactly 10 logins!"
	}
	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("Number is even")
	} else {
		fmt.Println("Number is odd")
	}

	if num := 3; num < 10 {
		fmt.Println(num, "is less than 10")
	} else {
		fmt.Println(num, "is greater than 10")
	}
}