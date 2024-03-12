package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to sliecs")

	var frndsList = []string{"A", "B", "C"}
	fmt.Printf("Type of frndsList is %T\n", frndsList)

	frndsList = append(frndsList, "A", "S")
	fmt.Println(frndsList)

	frndsList = append(frndsList[1:2])
	fmt.Println(frndsList)

	highScores := make([]int, 4)

	highScores[0] = 101
	highScores[2] = 202
	highScores[1] = 303
	highScores[3] = 404

	highScores = append(highScores, 505,606)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))


	// how to remove a value from slices based on index
	var name = []string{"abhi", "shek", "balji", "mauryaji", "shikha"}
	fmt.Println(name)

	var index int = 1
	name = append(name[:index], name[index+1:]...)
	fmt.Println(name)
}