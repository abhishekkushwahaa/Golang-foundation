package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Maths in GoLang")

	// Addition
	var a int = 10
	var b float64 = 20
	fmt.Println("Addition of a and b is: ", a + int(b))

	// Random number for math/rand
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println("Random number: ", rand.Intn(10) + 1)

	// Random number for crypto/rand
	RandomNumber, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println("Random number: ", RandomNumber)
}