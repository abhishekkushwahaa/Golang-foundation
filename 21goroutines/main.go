package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// // goroutine means a function that runs concurrently with other functions
	// go greeter("Hello G")

	// greeter("Sir")

	websiteList := []string{
		"https://golang.org",
		"https://go.dev",
		"https://lco.dev",
		"https://fb.com",
	}

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println("All go routines are done")
}

// func greeter(s string) {
// 	for i := 0; i < 5; i++ {
// 		// It is not better approach please use sync package for better approach
// 		time.Sleep(3 * time.Millisecond)
// 		fmt.Println(s)
// 	}
// }

func getStatusCode(endpoint string)  {
	defer wg.Done()
	res, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("OOPs in endpoint")
	}
	fmt.Println("Status code is", res.StatusCode)
}