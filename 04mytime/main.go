package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time of golang")

	presentTime := time.Now()
	fmt.Println(presentTime);

	// it's format standard
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createdDate := time.Date(2024, time.March, 10, 8, 24, 0, 0, time.UTC)
	fmt.Println(createdDate);
	fmt.Println(createdDate.Format("01-02-2006 Monday"))
} 