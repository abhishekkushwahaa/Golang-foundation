package main

import "fmt"

// Public
const LoginToken string = "abcd$1234"

func main() {
	var name string = "Abhishek"
	fmt.Println(name)
	fmt.Printf("Variable is type: %T \n", name)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is type: %T \n", smallVal)

	var smallFloat float32 = 255.2345676543221345
	fmt.Println(smallFloat)
	fmt.Printf("Variable is type: %T \n", smallFloat)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is type: %T \n", anotherVariable)

	// implicit type
	var name1 = "abhi"
	fmt.Println(name1)
	fmt.Printf("Variable is type: %T \n", name1)

	// no var style := It is can only inside the method not allowed in global scope
	numberOfUser := 300.0
	fmt.Println(numberOfUser)
	fmt.Printf("Variable is type: %T \n", numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is type: %T \n", LoginToken)

}