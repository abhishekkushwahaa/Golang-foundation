package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loop, Break and Continue in Golang!!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// The above for loop can be written as below
	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("Index: %v, Day: %v\n", index, day)
	// }

	for _, day := range days {
		fmt.Printf("Day: %v\n", day)
	}

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 3 {
			goto lco
		}

		// if rougueValue == 5 {
		// 	break
		// } 

		if rougueValue == 5 {
			rougueValue++
			continue
		}

		fmt.Println("Value of rogueValue is: ", rougueValue)
		rougueValue++
	}

	lco:
	    fmt.Println("Jumping to LCO")

}