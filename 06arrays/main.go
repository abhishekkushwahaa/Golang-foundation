package main

import "fmt"

func main() {
	fmt.Println("Welcome to array")

	var frnds [5]string

	frnds[3] = "Harsh"
	frnds[1] = "Abhi"
	frnds[2] = "Lakshya"
	frnds[0] = "Balji"

	fmt.Println("List of friends is :", frnds)
	fmt.Println("List of friends is :", len(frnds)) // always give array length

	var gfList = [3]int{0,0,0}
	fmt.Println("GF list is:",gfList)
	fmt.Println("GF list is:",len(gfList))
}