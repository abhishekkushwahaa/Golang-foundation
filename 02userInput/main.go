package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the your name:")

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	fmt.Println("Nice name!" , input)
	fmt.Printf("Variable is type: %T \n", input)
}