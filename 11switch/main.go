package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switch case in Golang!!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("You rolled: 1")
	case 2:
		fmt.Println("You rolled: 2")
	case 3:
		fmt.Println("You rolled: 3")
		// fallthrough means that the next case will also be executed
		fallthrough
	case 4:
		fmt.Println("You rolled: 4")
	default:
		fmt.Println("You rolled: Default")
	}
}