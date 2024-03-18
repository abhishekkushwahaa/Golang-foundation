package main

import (
	"fmt"
	"time"
)

func main() {
	// goroutine means a function that runs concurrently with other functions
	go greeter("Hello G")

	greeter("Sir")
}

func greeter(s string) {
	for i := 0; i < 5; i++ {
		// It is not better approach please use sync package for better approach
		time.Sleep(3 * time.Millisecond)
		fmt.Println(s)
	}
}